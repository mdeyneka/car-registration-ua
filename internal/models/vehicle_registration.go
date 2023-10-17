package models

import (
	"time"
)

type VehicleRegistration struct {
	ID                    string    `json:"id"`
	Person                string    `json:"person"`
	RegistrationAddress   string    `json:"reg_address"`
	OperationCode         uint16    `json:"op_code"`
	OperationName         string    `json:"op_name"`
	RegistrationDate      time.Time `json:"reg_date"`
	DepartmentCode        uint16    `json:"dep_code"`
	DepartmentName        string    `json:"dep_name"`
	Brand                 string    `json:"brand"`
	Model                 string    `json:"model"`
	VIN                   string    `json:"vin"`
	ProductionYear        uint16    `json:"prod_year"`
	Color                 string    `json:"color"`
	VehicleType           string    `json:"veh_type"`
	VehicleBody           string    `json:"veh_body"`
	Purpose               string    `json:"purpose"`
	FuelType              string    `json:"fuel_type"`
	EngineCapacity        uint16    `json:"eng_cap"`
	OwnWeight             uint16    `json:"own_weight"`
	TotalWeight           uint16    `json:"total_weight"`
	NewRegistrationNumber string    `json:"new_reg_num"`
}
