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
	var units []struct {
		UnitID         string      `json:"unit_id"`
		RequiredAction string      `json:"required_action"`
		Profile        unitProfile `json:"profile"`
		IsCharging     bool        `json:"is_charging"`
		IsWorking      bool        `json:"is_working"`
		Soc            int         `json:"soc"`
	}
	for results1.Next() {
		var unit struct {
			UnitID         string      `json:"unit_id"`
			RequiredAction string      `json:"required_action"`
			Profile        unitProfile `json:"profile"`
			IsCharging     bool        `json:"is_charging"`
			IsWorking      bool        `json:"is_working"`
			Soc            int         `json:"soc"`
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
		unit.Profile.Purpose = unitElm.Purpose
		unit.Profile.UnitType = unitElm.UnitType
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
		units = append(units, unit)
	}
	send(units, w)
}
