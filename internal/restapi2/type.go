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

type manageInfoPnt struct{
	Page	pageElm			`json:"page"`
	Data	[]manageInfoData	`json:"manage_info"`
}

type batteryData struct{
	Data 		unitData 		`json:"data"`
	Management 	managementMinElm	`json:"management"`
	//Data_BMU 	unitBMUData 		`json:"data_bmu"`
	//Error		[]errorsElm		`json:"error"`
}

type batteryDetailData struct{
	Data 		unitData 		`json:"data"`
	Management 	managementElm	`json:"management"`
	Error		[]errorsElm		`json:"error"`
}

type manageInfoData struct{
	SerialNumber	string		`json:"serial_number"`
	UnitID			string      `json:"unit_id"`
	//UnitID			int     `json:"unit_id"`
	BatteryType		string		`json:"battery_type"`
	CreateAt		time.Time	`json:"create_at"`
	Voltage			float32		`json:"voltage"`
	Current			float32		`json:"current"`
	OutputVoltage	float32		`json:"output_voltage"`
	OutputCurrent	float32		`json:"output_current"`
	SoC				float32     `json:"soc"`
	Customer 		string		`json:"customer"`
	CarModel		string		`json:"car_model"`
	Charger			string		`json:"charger"`
	Seller			string		`json:"seller"`
	Comment			string		`json:"comment"`
	IsError			bool		`json:"is_error"`
	State    		string		`json:"registration_state"`
}

type managementElm struct{
	IsError			bool		`json:"is_error"`
	IsRegistered 	bool		`json:"is_registered"`
	SerialNumber 	string		`json:"serial_number"`
	Customer 		string		`json:"customer"`
	Voltage			float32		`json:"voltage"`
	Current			float32		`json:"current"`
}

type managementMinElm struct{
	IsError			bool		`json:"is_error"`
	IsRegistered 	bool		`json:"is_registered"`
	SerialNumber 	string		`json:"serial_number"`
	//Customer 		string		`json:"customer"`
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

type unitBMUData struct {
    BmuID         uint64    `json:"bmu_id"`
    Time            time.Time `json:"time"`
    LastIOTime      time.Time `json:"last_io_time"`
    ErrCode         string    `json:"err_code"`         // char(1)
    VCell0          float32   `json:"vcell0"`
    VCell1          float32   `json:"vcell1"`
    VCell2          float32   `json:"vcell2"`
    VCell3          float32   `json:"vcell3"`
    VCell4          float32   `json:"vcell4"`
    VCell5          float32   `json:"vcell5"`
    VCell6          float32   `json:"vcell6"`
    VCell7          float32   `json:"vcell7"`
    VCell8          float32   `json:"vcell8"`
    VCell9          float32   `json:"vcell9"`
    VCell10         float32   `json:"vcell10"`
    VCell11         float32   `json:"vcell11"`
    VCell12         float32   `json:"vcell12"`
    VCell13         float32   `json:"vcell13"`
    WorkState       string    `json:"work_state"`       // char(16)
    ViSense50       float32   `json:"visense50"`        // float
    ViSense10       float32   `json:"visense10"`        // float
    Vim050          uint16    `json:"vim050"`          // smallint unsigned
    Vim010          uint16    `json:"vim010"`          // smallint unsigned
    TAdr0           int16     `json:"tadr0"`           // smallint
    TAdr1           int16     `json:"tadr1"`           // smallint
    VtPCB0          float32   `json:"vt_pcb0"`         // float
    VtPCB1          float32   `json:"vt_pcb1"`         // float
    BatteryLED      string    `json:"battery_led"`     // enum('01','02','04','08','10')
    FetState        string    `json:"fet_state"`       // enum('00','01','02','03')
    BalanceNow      string    `json:"balance_now"`     // char(16)
    ChargeNum       uint32    `json:"charge_num"`      // int unsigned
    DischargeNum    uint32    `json:"discharge_num"`   // int unsigned
    OvChNum         uint32    `json:"ovch_num"`        // int unsigned
    OvDisNum        uint32    `json:"ovdis_num"`       // int unsigned
    OvChargeNum     uint32    `json:"ovcharge_num"`    // int unsigned
    OvDischargeNum  uint32    `json:"ovdischarge_num"` // int unsigned
    ShortNum        uint32    `json:"short_num"`       // int unsigned
    UDTEMPechNum    uint32    `json:"udtempech_num"`   // int unsigned
    UDTEMPedisNum   uint32    `json:"udtempedis_num"`  // int unsigned
    OVTEMPechNum    uint32    `json:"ovtempech_num"`   // int unsigned
    OVTEMPedisNum   uint32    `json:"ovtempedis_num"`  // int unsigned
    Energy          uint32    `json:"energy"`          // int unsigned
    Energy100       uint32    `json:"energy100"`       // int unsigned
    IntDetection    string    `json:"int_detection"`   // char(2)
}

type errorsElm struct {
	ErrorCode      int 			 `json:"error_code"`
	ErrorCategory	string 		 `json:"error_category"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
	ErrorTime		time.Time	 `json:"error_time"`
}

type errorState struct {
	UnitID			int 		`json:"unit_id"`
	ErrorCode      	int 		`json:"error_code"`
	ErrorTime		time.Time	`json:"error_time"`
}

type userElm struct {
	UserName	string			`json:"user_name"`
	UserRole	string			`json:"user_role"`
}

type errorsListElm struct{
	ErrorCode      int 			 `json:"error_code"`
	ErrorCategory	string 		 `json:"error_category"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
}

type carModelListElm struct{
	CarModelID      int 		 `json:"car_model_id"`
	CarModelName	string 		 `json:"car_model_name"`
	Comment   		string       `json:"comment"`
}

type customerListElm struct{
	CustomerID      int			`json:"customer_id"`
	CustomerName	string		`json:"customer_name"`
	Comment   		string       `json:"comment"`
}