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

func InitiateDB(db *gorm.DB) {
	dbConnect = db
}

func GetVehicleRegistrations(c *gin.Context) {
	var vehicles []models.VehicleRegistration
	status := dbConnect.Find(&vehicles)
	if status.Error != nil {
		log.Printf("Error while getting all vehicle registration, Reason: %v\n", status.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Vehicle registrations",
		"data":    vehicles,
	})
}

func PostVehicleRegistrations(c *gin.Context) {
	/*
		var newVehicleRegistration models.VehicleRegistration

		if err := c.BindJSON(&newVehicleRegistration); err != nil {
			return
		}

		database.VehRegistrations = append(database.VehRegistrations, newVehicleRegistration)
		c.IndentedJSON(http.StatusCreated, newVehicleRegistration)
	*/
}

func GetVehicleRegistrationByID(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.VehicleRegistration
	status := dbConnect.Where("id = ?", id).First(&vehicle)

	if status.Error != nil {
		log.Printf("Error while getting vehicle registration with ID: %s - %v\n", id, status.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error while getting vehicle registration with ID: %s - %v", id, status.Error),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   vehicle,
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
