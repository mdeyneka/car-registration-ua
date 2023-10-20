package models

import (
	"time"
)

type VehicleRegistration struct {
	ID                    uint32    `gorm:"primaryKey"`
	Person                string    `gorm:"type:varchar(1)" json:"person"`
	RegistrationAddress   string    `gorm:"type:varchar(10);column:registrationaddress" json:"reg_address"`
	OperationCode         uint16    `gorm:"type:uint16;column:operationcode" json:"op_code"`
	OperationName         string    `gorm:"type:varchar(255);column:operationname" json:"op_name"`
	RegistrationDate      time.Time `gorm:"type:time;column:registrationdate" json:"reg_date"`
	DepartmentCode        uint16    `gorm:"type:uint16;column:departmentcode" json:"dep_code"`
	DepartmentName        string    `gorm:"type:varchar(64);column:departmentname" json:"dep_name"`
	VehicleBrand          string    `gorm:"type:varchar(32);column:brand" json:"veh_brand"`
	VehicleModel          string    `gorm:"type:varchar(32);column:model" json:"veh_model"`
	VIN                   string    `gorm:"type:varchar(10);column:vin" json:"vin"`
	ProductionYear        uint16    `gorm:"type:uint16;column:productionyear" json:"prod_year"`
	Color                 string    `gorm:"type:varchar(32);column:color" json:"color"`
	VehicleType           string    `gorm:"type:varchar(32);column:vehicletype" json:"veh_type"`
	VehicleBody           string    `gorm:"type:varchar(32);column:vehiclebody" json:"veh_body"`
	Purpose               string    `gorm:"type:varchar(32);column:purpose" json:"purpose"`
	FuelType              string    `gorm:"type:varchar(32);column:fueltype" json:"fuel_type"`
	EngineCapacity        uint16    `gorm:"type:uint16;column:enginecapacity" json:"eng_cap"`
	OwnWeight             uint16    `gorm:"type:uint16;column:ownweight" json:"own_weight"`
	TotalWeight           uint16    `gorm:"type:uint16;column:totalweight" json:"total_weight"`
	NewRegistrationNumber string    `gorm:"type:varchar(10);column:newregistrationnumber" json:"new_reg_num"`
}
