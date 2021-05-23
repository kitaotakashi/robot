package db

import (
	"fmt"
	//"database/sql"
	//"github.com/guregu/null"
	"net/http"
	"strconv"
	"time"
)

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
}
type unitTimeStamp struct {
	RegisterdAt time.Time `json:"registerd_at"`
	Time        time.Time `json:"time"`
}

/*
type unitSummary struct {
	UnitID         string      `json:"unit_id"`
	RequiredAction string      `json:"required_action"`
	Profile        unitProfile `json:"profile"`
	IsCharging     bool        `json:"is_charging"`
	IsWorking      bool        `json:"is_working"`
	Soc            int         `json:"soc"`
}
*/

type unitError struct {
	ErrorCode      int    `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	RequiredAction string `json:"required_action"`
}
type unitDetail struct {
	UnitID string    `json:"unit_id"`
	Error  unitError `json:"error"`
	//RequiredAction string        `json:"required_action"`
	Profile    unitProfile   `json:"profile"`
	Status     unitStatus    `json:"status"`
	TimeStamps unitTimeStamp `json:"time_stamps"`
	CustomerID string        `json:"customer_id"`
}

/*
type unit struct {
	UnitSummary unitSummary `json:"unitSummary"`
	//UnitDetail  unitDetail  `json:"unit_datail"`
}
*/

// UnitsView はunitページに必要なデータをDBから取得し、JSONで返す
func UnitsView(w http.ResponseWriter, r *http.Request) {
	//id := query(r, "unit_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM units")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("chk")
	fmt.Println(results1)
	/*
	var units []struct {
		UnitID         string      `json:"unit_id"`
		RequiredAction string      `json:"required_action"`
		Profile        unitProfile `json:"profile"`
		IsCharging     bool        `json:"is_charging"`
		IsWorking      bool        `json:"is_working"`
		Soc            int         `json:"soc"`
	}
	*/
	//unitsベージで必要な情報
	var units []struct {
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

	for results1.Next() {
		/*
		var unit struct {
			UnitID         string      `json:"unit_id"`
			RequiredAction string      `json:"required_action"`
			Profile        unitProfile `json:"profile"`
			IsCharging     bool        `json:"is_charging"`
			IsWorking      bool        `json:"is_working"`
			Soc            int         `json:"soc"`
		}
		*/
		var unit struct{
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
		var unitElm unitElm
		Columns := columns(&unitElm)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		unit.UnitID = unitElm.UnitID
		unit.LastIOtime = unitElm.LastIOtime

		if unitElm.BatteryError.Valid == true {
			errorcode := int(unitElm.BatteryError.Int32)
			unit.BatteryError = errorcode
			//error DBから取得
			results2, err := db.Query("SELECT required_action FROM errors WHERE error_code=" + strconv.Itoa(errorcode))
			if err != nil {
				panic(err.Error())
			}
			for results2.Next() {
				var errorElm errorElm
				err = results2.Scan(&errorElm.RequiredAction)
				if err != nil {
					panic(err.Error())
				}
				unit.RequiredAction = errorElm.RequiredAction
			}
		}

		//test　db実装まち
		//まず battery一覧を取得し、unitIDで問い合わせてcontractIDがあるかどうか
		//契約IDがnullでない場合、事業所名と企業名を取得
		unit.CustomerName = "test"
		unit.DepartmentName = "test"
		unit.ContractID = 1
		/*
		if unitElm.ContractID.Valid == true {
			contract_id := int(unitElm.ContractID.Int32)
			var departmentElm departmentElm
			results3, err := db.Query("SELECT * FROM departments WHERE department_id=(SELECT department_id FROM contracts WHERE contract_id=" + strconv.Itoa(contract_id) + ")")
			if err != nil {
				panic(err.Error())
			}
			for results3.Next() {
				Columns = columns(&departmentElm)
				err = results3.Scan(Columns...)
				if err != nil {
					panic(err.Error())
				}
				unit.DepartmentName = departmentElm.DepartmentName

				var customerElm customerElm
				department_id := departmentElm.ParentID
				results4, err := db.Query("SELECT * FROM customers WHERE customer_id = (SELECT parent_id FROM departments WHERE department_id="+ strconv.Itoa(department_id) +")")
				if err != nil {
					panic(err.Error())
				}
				for results4.Next() {
					err = results4.Scan(&customerElm.CorporationName)
					if err != nil {
						panic(err.Error())
					}
					unit.CustomerName = customerElm.CorporationName
				}
			}
		}
		*/


		//unit.UnitDetail.UnitID = unitElm.UnitID
		//unit.UnitDetail.Profile.UnitID = unitElm.UnitID
		/*
			if unitElm.ErrorCode.Valid == true {
				errorcode := int(unitElm.ErrorCode.Int32)
				unit.UnitDetail.Error.ErrorCode = errorcode
				fmt.Println(errorcode)
				results2, err := db.Query("SELECT * FROM errors WHERE error_code=" + strconv.Itoa(errorcode))
				if err != nil {
					panic(err.Error())
				}
				for results2.Next() {
					var errorElm errorElm

					err = results2.Scan(&errorElm.ErrorCode.Int32, &errorElm.ErrorMessage, &errorElm.RequiredAction)
					if err != nil {
						panic(err.Error())
					}
					unit.UnitDetail.Error.ErrorMessage = errorElm.ErrorMessage
					unit.UnitDetail.Error.RequiredAction = errorElm.RequiredAction
				}
			}
		*/
		if unitElm.IsCharging == "close" {
			unit.IsCharging = false
		} else {
			unit.IsCharging = true
		}
		if time.Now().Sub(unitElm.LastIOtime) > time.Minute*5 {
			unit.IsWorking = false
			//unit.UnitDetail.Status.IsWorking = "off"
		} else {
			unit.IsWorking = true
			//unit.UnitDetail.Status.IsWorking = "on"
		}
		unit.Soc = unitElm.Soc

		unit.Profile.Purpose = "test_forklift"//unitElm.Purpose //from battery DB
		unit.Profile.UnitType = "test_RB-***"//unitElm.UnitType //from battery DB
		//unit.UnitDetail.Profile.UnitType = unitElm.UnitType
		unit.Profile.Location.Latitude = unitElm.Latitude
		//unit.UnitDetail.Profile.Location.Latitude = unitElm.Latitude
		unit.Profile.Location.Longitude = unitElm.Longitude
		
		/*
			unit.UnitDetail.Profile.Location.Longitude = unitElm.Longitude
			unit.UnitDetail.Status.Soc = unitElm.Soc
			unit.UnitDetail.Status.Soh = unitElm.Soh
			unit.UnitDetail.Status.Voltage = unitElm.Voltage
			unit.UnitDetail.Status.IsCharging = unitElm.IsCharging
			unit.UnitDetail.Status.Current = unitElm.Current
			unit.UnitDetail.Status.Capacity = unitElm.Capacity
			unit.UnitDetail.TimeStamps.Time = unitElm.Time
		*/
		fmt.Println(unit)
		units = append(units, unit)
	}
	send(units, w)
}
