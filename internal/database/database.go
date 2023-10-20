package database

import (
	"fmt"
	"log"

	"github.com/mdeineka/car-registration-ua/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(class interface{}) *gorm.DB {
	cfg := class.(config.Config)

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
