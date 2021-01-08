package db

import (
	"database/sql"
	"time"
)

// customerElm は顧客情報を格納する
type customerElm struct {
	AccountID       string    `json:"account_id"`
	CorporationName string    `json:"corporation_name"`
	Position        string    `json:"position"`
	Sector          string    `json:"sector"`
	Name            string    `json:"name"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	PostalCode      string    `json:"postal_code"`
	Address         string    `json:"address"`
	Mail            string    `json:"mail"`
	Phone           string    `json:"phone"`
}

// contractElm は契約情報を格納する
type contractElm struct {
	UnitID         uint      `json:"unit_id"`
	AccountID      string    `json:"account_id"`
	ContractType   string    `json:"contract_type"`
	ExecutionDate  time.Time `json:"execution_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	BillingDate    time.Time `json:"billing_date"`
}

// unitElm はバッテリー情報を格納する
type unitElm struct { //要素名を変更
	UnitID     string    `json:"unit_id"`
	UnitType   string    `json:"unit_type"`
	Purpose    string    `json:"purpose"`
	BmsVersion string    `json:"bms_version"`
	UnitState  string    `json:"unit_state"`
	Time       time.Time `json:"time"`
	//IsWorking  string    `json:"is_working"`
	Soc        int           `json:"soc"`
	Soh        int           `json:"soh"`
	Capacity   int           `json:"capacity"`
	Current    float32       `json:"current"`
	Voltage    float32       `json:"voltage"`
	Latitude   float32       `json:"latitude"`
	Longitude  float32       `json:"longitude"`
	IsCharging string        `json:"is_charging"`
	ErrorCode  sql.NullInt32 `json:"error_code"`
}

// errorElm はエラー情報を格納する
type errorElm struct {
	ErrorCode      sql.NullInt32 `json:"error_code"`
	ErrorMessage   string        `json:"error_message"`
	RequiredAction string        `json:"required_action"`
}
