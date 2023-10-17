package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdeineka/car-registration-ua/internal/database"
	"github.com/mdeineka/car-registration-ua/internal/models"
)

func GetVehicleRegistrations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.VehRegistrations)
}

func PostVehicleRegistrations(c *gin.Context) {
	var newVehicleRegistration models.VehicleRegistration

	if err := c.BindJSON(&newVehicleRegistration); err != nil {
		return
	}

	database.VehRegistrations = append(database.VehRegistrations, newVehicleRegistration)
	c.IndentedJSON(http.StatusCreated, newVehicleRegistration)
}

func GetVehicleRegistrationByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range database.VehRegistrations {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vehicle registration not found"})
}
