package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdeineka/car-registration-ua/internal/models"
	"gorm.io/gorm"
)

var dbConnect *gorm.DB

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
	/*
		id := c.Param("id")
		for _, a := range database.VehRegistrations {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vehicle registration not found"})
	*/
}
