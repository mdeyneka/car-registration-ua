package models

import (
	"time"
)

type VehicleRegistration struct {
	ID                    uint32    `gorm:"type:serial;primaryKey"`
	Person                string    `gorm:"type:varchar(1)" json:"person"`
	RegistrationAddress   string    `gorm:"type:varchar(10);column:registrationaddress" json:"reg_address"`
	OperationCode         uint16    `gorm:"type:smallint;column:operationcode" json:"op_code"`
	OperationName         string    `gorm:"type:varchar(255);column:operationname" json:"op_name"`
	RegistrationDate      time.Time `gorm:"type:date;column:registrationdate" json:"reg_date"`
	DepartmentCode        uint32    `gorm:"type:int;column:departmentcode" json:"dep_code"`
	DepartmentName        string    `gorm:"type:varchar(256);column:departmentname" json:"dep_name"`
	VehicleBrand          string    `gorm:"type:varchar(64);column:brand" json:"veh_brand"`
	VehicleModel          string    `gorm:"type:varchar(64);column:model" json:"veh_model"`
	VIN                   string    `gorm:"type:varchar(64);index;column:vin" json:"vin"`
	ProductionYear        uint16    `gorm:"type:smallint;column:productionyear" json:"prod_year"`
	Color                 string    `gorm:"type:varchar(32);column:color" json:"color"`
	VehicleType           string    `gorm:"type:varchar(32);column:vehicletype" json:"veh_type"`
	VehicleBody           string    `gorm:"type:varchar(64);column:vehiclebody" json:"veh_body"`
	Purpose               string    `gorm:"type:varchar(32);column:purpose" json:"purpose"`
	FuelType              string    `gorm:"type:varchar(32);column:fueltype" json:"fuel_type"`
	EngineCapacity        uint32    `gorm:"type:int;column:enginecapacity" json:"eng_cap"`
	OwnWeight             float32   `gorm:"type:real;column:ownweight" json:"own_weight"`
	TotalWeight           float32   `gorm:"type:real;column:totalweight" json:"total_weight"`
	NewRegistrationNumber string    `gorm:"type:varchar(16);index;column:newregistrationnumber" json:"new_reg_num"`
}
