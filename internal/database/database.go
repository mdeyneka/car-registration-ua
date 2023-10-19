package database

import (
	"fmt"
	"log"
	"time"

	"github.com/mdeineka/car-registration-ua/internal/config"
	"github.com/mdeineka/car-registration-ua/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var VehRegistrations = []models.VehicleRegistration{
	{ID: "1", Person: "P", RegistrationAddress: "1225284101", OperationCode: 308, OperationName: "ПЕРЕРЕЄСТРАЦІЯ НА НОВОГО ВЛАСНИКА ЗА ДОГ. КУПIВЛI-ПРОДАЖУ (СГ)", RegistrationDate: time.Date(2022, 12, 21, 0, 0, 0, 0, time.Local), DepartmentCode: 12537, DepartmentName: "ТСЦ 1248", Brand: "KIA", Model: "SOUL", VIN: "KNAJT811AC7443261", ProductionYear: 2012, Color: "КОРИЧНЕВИЙ", VehicleType: "ЛЕГКОВИЙ", VehicleBody: "УНІВЕРСАЛ", Purpose: "ЗАГАЛЬНИЙ", FuelType: "БЕНЗИН АБО ГАЗ", EngineCapacity: 1591, OwnWeight: 1217, TotalWeight: 1690, NewRegistrationNumber: "АЕ7265ХА"},
	{ID: "2", Person: "P", RegistrationAddress: "3220880905", OperationCode: 254, OperationName: "НАЛЕЖНИЙ КОРИСТУВАЧ. РЕЄСТРАЦІЯ", RegistrationDate: time.Date(2022, 9, 7, 0, 0, 0, 0, time.Local), DepartmentCode: 10000, DepartmentName: "ДДАІ МВС УКРАЇНИ", Brand: "NISSAN", Model: "QASHQAI", VIN: "SJNFBNJ10U2471634", ProductionYear: 2012, Color: "БІЛИЙ", VehicleType: "ЛЕГКОВИЙ", VehicleBody: "ХЕТЧБЕК", Purpose: "ЗАГАЛЬНИЙ", FuelType: "БЕНЗИН АБО ГАЗ", EngineCapacity: 1997, OwnWeight: 1310, TotalWeight: 1960, NewRegistrationNumber: "АІ1954ЕМ"},
}

func Init(class interface{}) *gorm.DB {
	cfg := class.(config.Config)

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
