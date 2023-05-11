package mico2

import (
	"time"
	"database/sql"
	//"github.com/guregu/null"
)

type Token struct {
	Token	string		`json:"access_token"`
	RefreshToken string	`json:"refresh_token"`
	//Scope	string		`json:"scope"`
	Expires_in	int	`json:"expires_in"`
	TokenType	string 	`json:"token_type"`
}

type unitTimeStamp struct {
	RegisterdAt time.Time `json:"registerd_at"`
	Time        time.Time `json:"time"`
}

type pageElm struct{
	PageNow	int		`json:"page_now"`
	PageMax	int		`json:"page_max"` 
	DataNum int		`json:"data_num"`
}

type batteryPnt struct{
	Page	pageElm			`json:"page"`
	Data	[]batteryData	`json:"batteries"`
}

type batteryData struct{
	Data 		unitData 		`json:"data"`
	Management 	managementElm	`json:"management"`
}

type managementElm struct{
	IsRestration 	bool		`json:"is_registration"`
	SerialNumber 	string		`json:"serial_number"`
	Customer 		string		`json:"customer"`
	IsError			bool		`json:"is_error"`	
}

type unitData struct{
	UnitID     string      `json:"unit_id"`
	Time       time.Time   `json:"time"`
	BmsVersion string `json:"bms_version"`
	LastIOtime time.Time   `json:"last_io_time"`
	//LastChargerError int   `json:"last_charger_error"`
	//LastChargerErrorTime time.Time   `json:"last_charger_error_time"`
	Longitude  float32     `json:"longitude"`
	Latitude   float32     `json:"latitude"`
	ChargeMode string	   `json:"charge_mode"`
	BatteryCurrent float32 `json:"battery_current"`
	BatteryVoltage float32 `json:"battery_voltage"`
	BatteryError sql.NullInt32	   `json:"battery_error"`
	Soc        float32     `json:"soc"`
	OutputCurrent    float32     `json:"output_current"`
	OutputVoltage    float32     `json:"output_voltage"`
	IsCharging string      `json:"is_charging"`
	ChargerError int 	   `json:"charger_error"`
	UsageTime	float32	   `json:"usage_time"`
	NumberOfCharges int	   `json:"number_of_charges"`
	MaxCellVoltage float32 `json:"max_cell_voltage"`
	MinCellVoltage float32 `json:"min_cell_voltage"`
	MaxTemperature float32 `json:"max_temperature"`
	MinTemperature float32 `json:"min_temperature"`
}

type errorsElm struct {
	ErrorCode      int 			 `json:"error_code"`
	ErrorCategory	string 		 `json:"error_category"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
}