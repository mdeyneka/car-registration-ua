package database

import (
	"github.com/mdeineka/car-registration-ua/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.VehicleRegistration{})
}
