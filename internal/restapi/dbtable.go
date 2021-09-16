package db

import (
	"database/sql"
	"github.com/guregu/null"
	"github.com/lib/pq"
	"time"
	//"encoding/json"
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
	ParentID		int			`json:"parent_id"`
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

type batteryRequestElm struct {
	BatteryOptionId	int	`json:"battery_option_id"`
	OptionName		null.String `json:"option_name"`
	DepartmentID	int		`json:"department_id"`
	DepartmentName	null.String    `json:"department_name"`
	ContractID		int		`json:"contract_id"`
	ContractName	string   `json:"contract_name"`
}

// contractElm は契約情報を格納する
type contractElm struct {
	ContractID		int		`json:"contract_id"`
	DepartmentID      	int       `json:"department_id"`
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
	Horizontal		sql.NullInt32			`json:"horizontal"`
	Height			sql.NullInt32			`json:"height"`
	How2Change		null.String	`json:"how2change"`
	DailyWorkingMinute	sql.NullInt32		`json:"daily_working_minute"`
	DailyChargingMinute	sql.NullInt32		`json:"daily_charging_minute"`
	NumberOfChanges		sql.NullInt32		`json:"number_of_change"`
	ForkliftEnvironment	null.String	`json:"forklift_environment"`
	EnvironmentElse		null.String `json:"environment_else"`
	InputPlug			null.String `json:"input_plug"`
	OutputPlug			null.String `json:"output_plug"`
	ChangeHelp			null.String `json:"change_help"`
	Comment				null.String	`json:"comment"`
	Request				null.String     `json:"request"`
	PicForklift			null.String		`json:"pic_forklift"`
	PicForkliftPlate	null.String		`json:"pic_forklift_plate"`
	PicBattery			null.String		`json:"pic_battery"`
	PicBatteryPlate		null.String		`json:"pic_battery_plate"`
	PicChangePlace		null.String		`json:"pic_change_place"`
	PicBatteryPlug		null.String		`json:"pic_battery_plug"`
	PicForkliftPlug		null.String		`json:"pic_forklift_plug"`
	PicChangeEquipment	null.String		`json:"pic_change_equipment"`
}

type chargerElm  struct{
	ChargerId	int	`json:"charger_id"`
	DepartmentID	int `json:"department_id"`
	ContractID		sql.NullInt32  `json:"contract_id"`
	UpdateDate		pq.NullTime	`json:"update_date"`
	InfoType		null.String	`json:"info_type"`
	ChargerName		null.String `json:"charger_name"`
	
	ChargeEnvironment	null.String	`json:"charge_environment"`
	ChargeEnvironmentElse	null.String	`json:"charge_environment_else"`
	How2Supply		null.String		`json:"how2supply"`
	SupplyPlug		null.String		`json:"supply_plug"`
	SupplyElse		null.String		`json:"supply_else"`
	PowerSupplyAmpere 	sql.NullInt32  `json:"power_supply_ampere"`
	Stand			null.String		`json:"stand"`
	PowerSupply2ChargerCable	sql.NullFloat64	`json:"power_supply2charger_cable_length"`
	Charger2ForkliftCable	sql.NullFloat64	`json:"charger2forklift_cable_length"`
	ChargerSettingHelp	null.String		`json:"charger_setting_help"`
	Comment				null.String		`json:"comment"`
	Request				null.String     `json:"request"`
	PicChargerStand			null.String		`json:"pic_charger_stand"`
	PicPowerSupply			null.String		`json:"pic_power_supply"`
	PicSupplyPlug			null.String		`json:"pic_supply_plug"`
}

// errorElm はエラー情報を格納する
type errorStatesElm struct {
	ErrorCode      	int 		`json:"error_code"`
	ObjectType		string 		`json:"object_type"`
	ObjectId   		int        	`json:"object_id"`
	ErrorTime 		time.Time   `json:"error_time"`
}

type errorsElm struct {
	ErrorCode      int 			 `json:"error_code"`
	ErrorCategory	string 		 `json:"error_category"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
}
