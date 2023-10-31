package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdeineka/car-registration-ua/internal/config"
	"github.com/mdeineka/car-registration-ua/internal/controllers"
	"github.com/mdeineka/car-registration-ua/internal/database"
)

func main() {
	db := database.Init(config.Dbconfig)
	controllers.InitiateDB(db)
	router := gin.Default()
	v1_group := router.Group("/v1")
	{
		v1_group.GET("/vehicle_registrations", controllers.GetVehicleRegistrations)
		v1_group.GET("/vehicle_registrations/:id", controllers.GetVehicleRegistrationByID)
		v1_group.POST("/vehicle_registrations/batchload", controllers.PostBatchVehicleRegistrations)
	}

	router.Run("localhost:8080")
}
