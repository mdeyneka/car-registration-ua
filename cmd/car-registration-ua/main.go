package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdeineka/car-registration-ua/internal/controllers"
)

func main() {
	router := gin.Default()
	v1_group := router.Group("/v1")
	{
		v1_group.GET("/vehicle_registrations", controllers.GetVehicleRegistrations)
		v1_group.GET("/vehicle_registrations/:id", controllers.GetVehicleRegistrationByID)
		v1_group.POST("/vehicle_registrations", controllers.PostVehicleRegistrations)
	}

	router.Run("localhost:8080")
}
