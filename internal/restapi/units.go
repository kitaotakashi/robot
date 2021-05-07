package db

import (
	//"fmt"
	"database/sql"
	"github.com/guregu/null"
	"net/http"
	"strconv"
	"time"
)

type geoLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
type unitProfile struct {
	UnitType null.String `json:"unit_type"`
	Purpose  null.String `json:"purpose"`
	Location geoLocation `json:"location"`
}
type unitStatus struct {
	IsCharging bool          `json:"is_charging"`
	IsWorking  bool          `json:"is_working"`
	Soc        int           `json:"soc"`
	Soh        sql.NullInt32 `json:"soh"`
	Capacity   sql.NullInt32 `json:"capacity"`
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
		UnitID         string      `json:"unit_id"`
		CustomerName   string      `json:"customer_name"`
		DepartmentName string      `json:"department_name"`
		ContractID int `json:"contract_id"`
		//ContractName   string      `json:"contract_name"`
		//Profile        unitProfile `json:"profile"`
		IsCharging     bool        `json:"is_charging"`
		IsWorking      bool        `json:"is_working"`
		Soc            int         `json:"soc"`
		RequiredAction string      `json:"required_action"`
		ErrorCode  	int   `json:"errorcode"`
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
			UnitID         string      `json:"unit_id"`
			CustomerName   string      `json:"customer_name"`
			DepartmentName string      `json:"department_name"`
			ContractID int `json:"contract_id"`
			//ContractName   string      `json:"contract_name"`
			//Profile        unitProfile `json:"profile"`
			IsCharging     bool        `json:"is_charging"`
			IsWorking      bool        `json:"is_working"`
			Soc            int         `json:"soc"`
			RequiredAction string      `json:"required_action"`
			ErrorCode  		int   `json:"errorcode"`
		}
		var unitElm unitElm
		Columns := columns(&unitElm)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		unit.UnitID = unitElm.UnitID

		if unitElm.ErrorCode.Valid == true {
			errorcode := int(unitElm.ErrorCode.Int32)
			unit.ErrorCode = errorcode
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

		//契約IDがnullでない場合、事業所名と企業名を取得
		//test　db実装まち
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
		if time.Now().Sub(unitElm.Time) > time.Minute*5 {
			unit.IsWorking = false
			//unit.UnitDetail.Status.IsWorking = "off"
		} else {
			unit.IsWorking = true
			//unit.UnitDetail.Status.IsWorking = "on"
		}
		unit.Soc = unitElm.Soc
		//unit profileはここで使わないのでoff
		/*
		unit.Profile.Purpose = unitElm.Purpose
		unit.Profile.UnitType = unitElm.UnitType
		//unit.UnitDetail.Profile.UnitType = unitElm.UnitType
		unit.Profile.Location.Latitude = unitElm.Latitude
		//unit.UnitDetail.Profile.Location.Latitude = unitElm.Latitude
		unit.Profile.Location.Longitude = unitElm.Longitude
		*/
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
		units = append(units, unit)
	}
	send(units, w)
}
