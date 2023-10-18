package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mdeineka/car-registration-ua/internal/controllers"
)

type Config struct {
	DbHost     string `env:"DB_HOST,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbName     string `env:"DB_NAME,required"`
	DBbPort    int    `env:"DB_PORT,required"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env file: %e", err)
	}

	cfg := Config{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Unable to parse ennvironment variables: %e", err)
	}

	router := gin.Default()
	v1_group := router.Group("/v1")
	{
		v1_group.GET("/vehicle_registrations", controllers.GetVehicleRegistrations)
		v1_group.GET("/vehicle_registrations/:id", controllers.GetVehicleRegistrationByID)
		v1_group.POST("/vehicle_registrations", controllers.PostVehicleRegistrations)
	}

	router.Run("localhost:8080")
}
