package db

import (
	"database/sql"
	"github.com/guregu/null"
	"github.com/lib/pq"
	"time"
)

// customerElm は顧客情報を格納する
type customerElm struct {
	AccountID       int       `json:"account_id"`
	CorporationName null.String      `json:"corporation_name"`
	Sector          null.String    `json:"sector"`
	PostalCode      null.String    `json:"postal_code"`
	Address         null.String    `json:"address"`
	Name            null.String    `json:"name"`
	Position        null.String    `json:"position"`
	//DateOfBirth     time.Time `json:"date_of_birth"`
	Mail            null.String    `json:"mail"`
	Phone           null.String    `json:"phone"`
}

type departmentElm struct {
	DepartmentID	int		`json:"department_id"`
	DepartmentName	null.String    `json:"department_name"`
	//AccountID       int       `json:"account_id"`
	ParentID		int       `json:"parent_id"`
	PostalCode      null.String    `json:"postal_code"`
	Address         null.String    `json:"address"`
	DailyWorkingHour	sql.NullInt32 `json:"daily_working_hour"`
	WeeklyHoliday	sql.NullInt32 `json:"weekly_holiday"`
	Name			null.String		`json:"name"`
	Position        null.String		`json:"position"`
	Mail            null.String		`json:"mail"`
	Phone           null.String		`json:"phone"`
	//DailyWorkingHour	int `json:"daily_working_hour"`
	//WeeklyHoliday		int `json:"weekly_holiday"`
}

// contractElm は契約情報を格納する
type contractElm struct {
	ContractID		int		`json:"contract_id"`
	AccountID      	int       `json:"account_id"`
	ContractName	string   `json:"contract_name"`
	ContractType   	string    `json:"contract_type"`
	ExecutionDate  	time.Time `json:"execution_date"`
	ExpirationDate 	time.Time `json:"expiration_date"`
	//BillingDate    time.Time `json:"billing_date"`
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

// unitElm はバッテリー情報を格納する
/*
type unitElm struct { //要素名を変更
	UnitID     string      `json:"unit_id"`
	UnitType   null.String `json:"unit_type"`
	Purpose    null.String `json:"purpose"`
	BmsVersion null.String `json:"bms_version"`
	UnitState  null.String `json:"unit_state"`
	Time       time.Time   `json:"time"`
	//IsWorking  string    `json:"is_working"`
	Soc        int           `json:"soc"`
	Soh        sql.NullInt32 `json:"soh"`
	Capacity   sql.NullInt32 `json:"capacity"`
	Current    float32       `json:"current"`
	Voltage    float32       `json:"voltage"`
	Latitude   float32       `json:"latitude"`
	Longitude  float32       `json:"longitude"`
	IsCharging string        `json:"is_charging"`
	ErrorCode  sql.NullInt32 `json:"error_code"`
	//ContractID sql.NullInt32 `json:"contract_id"` -> db側で実装まち
}*/
//新DB用に変更する
type unitElm struct { //要素名を変更
	UnitID     string      `json:"unit_id"`
	Time       time.Time   `json:"time"`
	BmsVersion string `json:"bms_version"`
	LastIOtime time.Time   `json:"last_io_time"`
	//LastChargerError int   `json:"last_charger_error"`
	//LastChargerErrorTime time.Time   `json:"last_charger_error_time"`
	Latitude   float32     `json:"latitude"`
	Longitude  float32     `json:"longitude"`
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

type batteryOptionElm  struct{
	BatteryOptionId	int	`json:"battery_option_id"`
	DepartmentID	int `json:"department_id"`
	ContractID		sql.NullInt32  `json:"contract_id"`
	UpdateDate		pq.NullTime	`json:"update_date"`
	//UpdateDate		sql.NullInt32	`json:"update_date"`
	InfoType		null.String	`json:"info_type"`
	OptionName		null.String `json:"option_name"`
	Forklift		null.String `json:"forklift"`
	Type			null.String `json:"type"`
	SerialNumber	null.String `json:"serial_number"`
	Voltage			sql.NullFloat64		`json:"voltage"`
	Capacity		sql.NullFloat64		`json:"capacity"`
	Weight			sql.NullFloat64		`json:"weight"`
	Vertical		sql.NullInt32			`json:"vertical"`
	Horizonal		sql.NullInt32			`json:"horizonal"`
	Height			sql.NullInt32			`json:"height"`
	How2Charge		null.String	`json:"how2charge"`
	DailyWorkingMinute	sql.NullInt32		`json:"daily_working_minute"`
	DailyChargingMinute	sql.NullInt32		`json:"daily_charging_minute"`
	NumberOfCharges		sql.NullInt32		`json:"number_of_charge"`
	ForkliftEnvironment	null.String	`json:"forklift_environment"`
	EnvironmentElse		null.String `json:"environment_else"`
	InputPlug			null.String `json:"input_plug"`
	OutputPlug			null.String `json:"output_plug"`
	ChargeHelp			null.String `json:"charge_help"`
	Comment				null.String	`json:"comment"`
	PicForklift			[]uint8		`json:"pic_forklift"`
	PicForkliftPlate	[]uint8		`json:"pic_forklift_plate"`
	PicBattery			[]uint8		`json:"pic_battery"`
	PicBatteryPlate		[]uint8		`json:"pic_battery_plate"`
	PicChargerPlace		[]uint8		`json:"pic_charger_place"`
	PicBatteryPlug		[]uint8		`json:"pic_battery_plug"`
	PicForkliftPlug		[]uint8		`json:"pic_forklift_plug"`
	PicChangeEquipment	[]uint8		`json:"pic_change_equipment"`
}

// errorElm はエラー情報を格納する
type errorElm struct {
	ErrorCode      sql.NullInt32 `json:"error_code"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
}
