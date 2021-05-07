package db

import (
	"database/sql"
	"github.com/guregu/null"
	"time"
)

// customerElm は顧客情報を格納する
type customerElm struct {
	AccountID       int       `json:"account_id"`
	CorporationName null.String    `json:"corporation_name"`
	Sector          null.String    `json:"sector"`
	Name            null.String    `json:"name"`
	Position        null.String    `json:"position"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	PostalCode      null.String    `json:"postal_code"`
	Address         null.String    `json:"address"`
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
	Name		null.String    `json:"name"`
	Position        null.String    `json:"position"`
	Mail            null.String    `json:"mail"`
	Phone           null.String    `json:"phone"`
	DailyWorkingHour	sql.NullInt32 `json:"daily_working_hour"`
	WeeklyHoliday	sql.NullInt32 `json:"weekly_holiday"`
}

// contractElm は契約情報を格納する
type contractElm struct {
	UnitID         uint      `json:"unit_id"`
	AccountID      int       `json:"account_id"`
	//ContractID		int		`json:"contract_id"`
	//DepartmentID	int		`json:"department_id"`
	ContractType   string    `json:"contract_type"`
	ExecutionDate  time.Time `json:"execution_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	BillingDate    time.Time `json:"billing_date"`
}

// unitElm はバッテリー情報を格納する
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
}

// errorElm はエラー情報を格納する
type errorElm struct {
	ErrorCode      sql.NullInt32 `json:"error_code"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
}
