package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdeineka/car-registration-ua/internal/models"
	"gorm.io/gorm"
)

var dbConnect *gorm.DB

type Form struct {
	CSVFile *multipart.FileHeader `form:"file" binding:"required"`
}

type VehicleRegistrationRequest struct {
	Id                    int       `form:"id,default=0"`
	Vin                   string    `form:"vin"`
	Newregistrationnumber string    `form:"newRegNumber"`
	Registrationdate      time.Time `form:"registrationDate" time_format:"2006-01-02" time_utc:"2"`
}

func InitiateDB(db *gorm.DB) {
	dbConnect = db
}

func GetVehicleRegistrations(c *gin.Context) {
	var vehicleRegistrationsRequest VehicleRegistrationRequest
	var vehicles []models.VehicleRegistration

	if err := c.ShouldBindQuery(&vehicleRegistrationsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := dbConnect.Limit(100).Find(&vehicles, vehicleRegistrationsRequest).Error
	if err != nil {
		log.Printf("Can't find anything %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Can't find anything %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   vehicles,
	})
}

func PostBatchVehicleRegistrations(c *gin.Context) {
	var form Form

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if form.CSVFile == nil {
		c.JSON(http.StatusBadRequest, "File is missing")
		return
	}

	file, err := form.CSVFile.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error in opening file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var vehicleRegistrations []models.VehicleRegistration
	for line, record := range records[1:] {
		operationCode, err := strconv.ParseUint(record[2], 10, 16)
		if err != nil {
			text := fmt.Sprintf("At line %d in parameter operationCode error: %s", line, err.Error())
			c.JSON(http.StatusBadRequest, text)
			return
		}

		registrationDate, err := time.Parse("02.01.2006", record[4])
		if err != nil {
			text := fmt.Sprintf("At line %d in parameter registrationDate error: %s", line, err.Error())
			c.JSON(http.StatusBadRequest, text)
			return
		}

		departmentCode, err := strconv.ParseUint(record[5], 10, 32)
		if err != nil {
			text := fmt.Sprintf("At line %d in parameter departmentCode error: %s", line, err.Error())
			c.JSON(http.StatusBadRequest, text)
			return
		}

		productionYear, err := strconv.ParseUint(record[10], 10, 16)
		if err != nil {
			text := fmt.Sprintf("At line %d in parameter productionYear error: %s", line, err.Error())
			c.JSON(http.StatusBadRequest, text)
			return
		}

		var engineCapacity uint64 = 0
		if record[16] != "" {
			engineCapacity, err = strconv.ParseUint(record[16], 10, 32)
			if err != nil {
				text := fmt.Sprintf("At line %d in parameter engineCapacity error: %s", line, err.Error())
				c.JSON(http.StatusBadRequest, text)
				return
			}
		}

		var ownWeight float64 = 0
		if record[17] != "" {
			ownWeight, err = strconv.ParseFloat(record[17], 32)
			if err != nil {
				text := fmt.Sprintf("At line %d in parameter ownWeight error: %s", line, err.Error())
				c.JSON(http.StatusBadRequest, text)
				return
			}
		}

		var totalWeight float64 = 0
		if record[18] != "" {
			totalWeight, err = strconv.ParseFloat(record[18], 32)
			if err != nil {
				text := fmt.Sprintf("At line %d in parameter totalWeight error: %s", line, err.Error())
				c.JSON(http.StatusBadRequest, text)
				return
			}
		}
		vehreg := models.VehicleRegistration{
			Person:                record[0],
			RegistrationAddress:   record[1],
			OperationCode:         uint16(operationCode),
			OperationName:         record[3],
			RegistrationDate:      registrationDate,
			DepartmentCode:        uint32(departmentCode),
			DepartmentName:        record[6],
			VehicleBrand:          record[7],
			VehicleModel:          record[8],
			VIN:                   record[9],
			ProductionYear:        uint16(productionYear),
			Color:                 record[11],
			VehicleType:           record[12],
			VehicleBody:           record[13],
			Purpose:               record[14],
			FuelType:              record[15],
			EngineCapacity:        uint32(engineCapacity),
			OwnWeight:             float32(ownWeight),
			TotalWeight:           float32(totalWeight),
			NewRegistrationNumber: record[19],
		}

		vehicleRegistrations = append(vehicleRegistrations, vehreg)
	}
	result := dbConnect.CreateInBatches(&vehicleRegistrations, 1000)

	if result.Error != nil {
		c.JSON(http.StatusExpectationFailed, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("%d Rows affected", result.RowsAffected),
	})
}
