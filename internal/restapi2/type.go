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

type batteryElm  struct{
	SerialNumber	int		`json:"serial_number"`
	UnitID     		string      `json:"unit_id"`
	ContractID		int		`json:"contract_id"`
	BatteryOptionID	int		`json:"battery_option_id"`
	DateOfManufacture time.Time `json:"date_of_manufacture"`
	BatteryTypeID	int 	`json:"battery_type_id"`
	Purpose			string      `json:"purpose"`
	UnitState		string		`json:"unit_state"`
}

type geoLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
type unitProfile struct {
	//null.stringに変更すること
	UnitType string `json:"unit_type"`//battery_type_id
	Purpose  string `json:"purpose"`
	Location geoLocation `json:"location"`
}
type unitStatus struct {
	IsCharging bool          `json:"is_charging"`
	IsWorking  bool          `json:"is_working"`
	Soc        float32           `json:"soc"`
	//Soh        sql.NullInt32 `json:"soh"`
	//Capacity   sql.NullInt32 `json:"capacity"`
	Current    float32       `json:"current"`
	Voltage    float32       `json:"voltage"`
	OutputCurrent    float32     `json:"output_current"`
	OutputVoltage    float32     `json:"output_voltage"`
	UsageTime	float32	   `json:"usage_time"`
	NumberOfCharges int	   `json:"number_of_charges"`
	MaxCellVoltage float32 `json:"max_cell_voltage"`
	MinCellVoltage float32 `json:"min_cell_voltage"`
	MaxTemperature float32 `json:"max_temperature"`
	MinTemperature float32 `json:"min_temperature"`
	LastIOtime 	   time.Time   `json:"last_io_time"`
}
type unitTimeStamp struct {
	RegisterdAt time.Time `json:"registerd_at"`
	Time        time.Time `json:"time"`
}

type unitPnt struct{
	UnitID         string      `json:"unit_id"`//units_DBから取得
	CustomerName   string      `json:"customer_name"`//contract_DBから取得
	DepartmentName string      `json:"department_name"`//contract_DBから取得
	ContractID int `json:"contract_id"`//contract_DBから取得
	LastIOtime 	   time.Time   `json:"last_io_time"`
	//ContractName   string      `json:"contract_name"`
	Profile        unitProfile `json:"profile"`
	IsCharging     bool        `json:"is_charging"`//units_DBから取得
	IsWorking      bool        `json:"is_working"`//last_io_timeから別途計算
	Soc            float32     `json:"soc"`//units_DBから取得
	RequiredAction string      `json:"required_action"`//error_DBから取得
	BatteryError  	int   `json:"battery_error""`//units_DBから取得
}

type unitElm struct { //要素名を変更
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

	/*
	UnitType   null.String `json:"unit_type"`
	Purpose    null.String `json:"purpose"`
	UnitState  null.String `json:"unit_state"`
	IsWorking  string    `json:"is_working"`
	Soh        sql.NullInt32 `json:"soh"`
	Capacity   sql.NullInt32 `json:"capacity"`
	ErrorCode  sql.NullInt32 `json:"error_code"`
	//ContractID sql.NullInt32 `json:"contract_id"` -> db側で実装まち
	*/
}

type errorsElm struct {
	ErrorCode      int 			 `json:"error_code"`
	ErrorCategory	string 		 `json:"error_category"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
}