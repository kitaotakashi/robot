package db

import (
	//"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
	//"github.com/guregu/null"
)

type unitDetail struct {
	UnitID string    `json:"unit_id"`
	Error  unitError `json:"error"`
	//RequiredAction string        `json:"required_action"`
	Profile    unitProfile   `json:"profile"`
	Status     unitStatus    `json:"status"`
	TimeStamps unitTimeStamp `json:"time_stamps"`
	CustomerID string        `json:"customer_id"`
	CorporationName string `json:"corporation_name"`
}

// detailed はバッテリーの詳細情報を格納する
type detailed struct {
	// 契約情報
	Contract contractElm `json:"contract"`
	// バッテリー情報
	Unit unitElm `json:"unit"`
	//　顧客情報
	Customer customerElm `json:"customer"`
}

// DetailView はdetailedページに必要なデータをDBから取得し、JSONで返す
func DetailView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "unit_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM units WHERE unit_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	/*
		var details []struct {
			// unitInfo
			UnitID         string        `json:"unit_id"`
			ErrorCode      sql.NullInt32 `json:"error_code"`
			RequiredAction string        `json:"required_action"`
			Profile        unitProfile   `json:"profile"`
			Status         unitStatus    `json:"status"`
			TimeStamps     unitTimeStamp `json:"time_stamps"`
			// customerInfo
			CustomerID      string `json:"customer_id"`
			CorporationName string `json:"corporation_name"`
			Position        string `json:"position"`
			Name            string `json:"name"`
			Mail            string `json:"mail"`
			Phone           string `json:"phone"`
		}
	*/
	/*
	var details []struct {
		// unitInfo
		UnitID         string        `json:"unit_id"`
		ErrorCode      sql.NullInt32 `json:"error_code"`
		RequiredAction string        `json:"required_action"`
		Profile        unitProfile   `json:"profile"`
		Status         unitStatus    `json:"status"`
		TimeStamps     unitTimeStamp `json:"time_stamps"`
		LastIOtime 	   time.Time   `json:"last_io_time"`
		// customerInfo
		CustomerID      int    `json:"customer_id"` //string->int
		CorporationName null.String `json:"corporation_name"`
		Position        null.String `json:"position"`
		Name            null.String `json:"name"`
		Mail            null.String `json:"mail"`
		Phone           null.String `json:"phone"`
	}*/
	var details []unitDetail

	for results1.Next() {
		/*
		var detail struct {
			// unitInfo
			UnitID         string        `json:"unit_id"`
			ErrorCode      sql.NullInt32 `json:"error_code"`
			RequiredAction string        `json:"required_action"`
			Profile        unitProfile   `json:"profile"`
			Status         unitStatus    `json:"status"`
			TimeStamps     unitTimeStamp `json:"time_stamps"`
			LastIOtime 	   time.Time   `json:"last_io_time"`
			// customerInfo
			CustomerID      int    `json:"customer_id"` //string->int
			CorporationName null.String `json:"corporation_name"`
			Position        null.String `json:"position"`
			Name            null.String `json:"name"`
			Mail            null.String `json:"mail"`
			Phone           null.String `json:"phone"`
		}*/

		var detail unitDetail

		var unitElm unitElm
		Columns := columns(&unitElm)
		err = results1.Scan(Columns...)
		fmt.Println(unitElm)
		if err != nil {
			panic(err.Error())
		}
		detail.UnitID = unitElm.UnitID
		if unitElm.BatteryError.Valid == true {
			errorcode := int(unitElm.BatteryError.Int32)
			results2, err := db.Query("SELECT error_code,required_action FROM errors WHERE error_code=" + strconv.Itoa(errorcode))
			if err != nil {
				panic(err.Error())
			}
			for results2.Next() {
				var errorElm errorElm
				err = results2.Scan(&errorElm.ErrorCode, &errorElm.RequiredAction)
				if err != nil {
					panic(err.Error())
				}
				detail.Error.ErrorCode = errorElm.ErrorCode
				detail.Error.RequiredAction = errorElm.RequiredAction
			}
		}else{
			detail.Error.RequiredAction = ""
		}
		
		detail.Profile.Location.Latitude = unitElm.Latitude
		detail.Profile.Location.Longitude = unitElm.Longitude

		if unitElm.IsCharging == "close" {
			detail.Status.IsCharging = false
		} else {
			detail.Status.IsCharging = true
		}
		if time.Now().Sub(unitElm.Time) > time.Minute*5 {
			detail.Status.IsWorking = false
		} else {
			detail.Status.IsWorking = true
		}
		detail.Status.Soc = unitElm.Soc
		//detail.Status.Soh = unitElm.Soh
		//detail.Status.Capacity = unitElm.Capacity
		detail.Status.Current = unitElm.OutputCurrent
		detail.Status.Voltage = unitElm.OutputVoltage
		detail.TimeStamps.Time = unitElm.Time
		detail.Status.LastIOtime= unitElm.LastIOtime

		//batteryDBからとる
		var flg=0
		results3, err := db.Query("SELECT * FROM batteries WHERE unit_id=" + unitElm.UnitID)
		for results3.Next() {
			flg=1
			var batteryElm batteryElm
			Columns := columns(& batteryElm)
			err = results3.Scan(Columns...)

			if err != nil {
				panic(err.Error())
			}
			//detail.Profile.UnitType = unitElm.UnitType
			detail.Profile.UnitType = "test_RB_****"//battery_type_id from batteries DB
			//detail.Profile.Purpose = unitElm.Purpose//from batteries DB
			detail.Profile.Purpose = batteryElm.Purpose

			results4, err := db.Query("SELECT account_id FROM contracts WHERE contract_id=" + strconv.Itoa(batteryElm.ContractID))
			for results4.Next() {
				var contractElm contractElm
				err = results4.Scan(&contractElm.AccountID)
				if err != nil {
					panic(err.Error())
				}

				detail.CustomerID = strconv.Itoa(contractElm.AccountID)

				results5, err := db.Query("SELECT corporation_name FROM customers WHERE account_id=" + strconv.Itoa(contractElm.AccountID))
				for results5.Next() {
					var customerElm customerElm
					err = results5.Scan(&customerElm.CorporationName)
					if err != nil {
						panic(err.Error())
					}
					detail.CorporationName = customerElm.CorporationName.String
				}
			}
		}
		if flg==0{
			detail.Profile.UnitType = ""//battery_type_id from batteries DB
			detail.Profile.Purpose = ""
			detail.CustomerID = "-1"
			detail.CorporationName = ""
		}

		/*
		//cutomerInfo
		var customerElm customerElm
		//batteriesからと
		results2, err := db.Query("SELECT * FROM customers WHERE account_id=(SELECT account_id FROM contracts WHERE unit_id=" + id[0] + ")")
		flag := true
		if err != nil {
			flag = false
			fmt.Println("no customer")
		}

		if flag == true {
			for results2.Next() {
				Columns = columns(&customerElm)
				err = results2.Scan(Columns...)
				if err != nil {
					panic(err.Error())
				}
			}
			detail.CustomerID = string(customerElm.AccountID)
			detail.CorporationName = customerElm.CorporationName
			/*
			detail.CorporationName = customerElm.CorporationName
			detail.Position = customerElm.Position
			detail.Name = customerElm.Name
			detail.Mail = customerElm.Mail
			detail.Phone = customerElm.Phone
			*/
		/*
		} else {
			//未登録処理
			detail.CustomerID = "-1"
			detail.CorporationName = "no customer"
		}
		*/
		details = append(details, detail)
	}
	fmt.Println(details)
	send(details, w)
}
