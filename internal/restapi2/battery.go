package mico2

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// UnitsView はunitページに必要なデータをDBから取得し、JSONで返す
func BatteriesView(w http.ResponseWriter, r *http.Request) {
	//id := query(r, "unit_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM units")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("chk")
	fmt.Println(results1)

	var units []unitPnt
	//unitsベージで必要な情報

	for results1.Next() {
		var unit unitPnt
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
				var errorElm errorsElm
				err = results2.Scan(&errorElm.RequiredAction)
				if err != nil {
					panic(err.Error())
				}
				unit.RequiredAction = errorElm.RequiredAction
			}
		}else{
			unit.RequiredAction = ""
		}

		//まず battery一覧を取得し、unitIDで問い合わせてcontractIDがあるかどうか
		//契約IDがnullでない場合、事業所名と企業名を取得
		var flg=0
		results3, err := db.Query("SELECT * FROM batteries WHERE unit_id=" + unit.UnitID)
		
		for results3.Next() {
			flg=1
			var batteryElm batteryElm
			Columns := columns(& batteryElm)
			err = results3.Scan(Columns...)

			if err != nil {
				panic(err.Error())
			}

			/*
			var battery_type_id = batteryElm.BatteryTypeID
			results40, err := db.Query("SELECT type FROM battery_types WHERE battery_type_id=" + strconv.Itoa(battery_type_id))
			for results40.Next() {
				var battery_type string
				
				err = results40.Scan(&battery_type)
				if err != nil {
					panic(err.Error())
				}

				//detail.Profile.UnitType = batteryElm.BatteryTypeID
				unit.Profile.UnitType = battery_type
			}

			unit.ContractID = batteryElm.ContractID
			unit.Profile.Purpose = batteryElm.Purpose//unitElm.Purpose //from battery DB
			//unit.Profile.UnitType = "test_RB-***"//<-battery_type_idから製品table取得

			//contractIDから取得
			//var accountId string
			results4, err := db.Query("SELECT department_id FROM contracts WHERE contract_id=" + strconv.Itoa(batteryElm.ContractID))
			for results4.Next() {
				var contractElm contractElm
				err = results4.Scan(&contractElm.DepartmentID)
				if err != nil {
					panic(err.Error())
				}

				//acountId = contractElm.AccountID
				//department_idからparent_id=account_idを取得
				var departmentElm departmentElm
				results7, err := db.Query("SELECT parent_id FROM departments WHERE department_id=" + strconv.Itoa(contractElm.DepartmentID))
				for results7.Next(){
					
					err = results7.Scan(&departmentElm.ParentID)
					if err != nil {
						panic(err.Error())
					}
				}

				//results5, err := db.Query("SELECT corporation_name FROM customers WHERE account_id=" + strconv.Itoa(contractElm.DepartmentID))
				results5, err := db.Query("SELECT corporation_name FROM customers WHERE account_id=" + strconv.Itoa(departmentElm.ParentID))
				for results5.Next() {
					var customerElm customerElm
					err = results5.Scan(&customerElm.CorporationName)
					if err != nil {
						panic(err.Error())
					}

					unit.CustomerName = customerElm.CorporationName.String	
				}
				results6, err := db.Query("SELECT department_name FROM departments WHERE parent_id=" + strconv.Itoa(contractElm.DepartmentID))
				for results6.Next() {
					//var departmentElm departmentElm
					err = results6.Scan(&departmentElm.DepartmentName)
					if err != nil {
						panic(err.Error())
					}
					unit.DepartmentName = departmentElm.DepartmentName.String
				}
				
			}
			*/
			
		}
		//契約がない＝customerがない場合
		if flg==0{
			unit.ContractID=-1
			unit.CustomerName = ""	
			unit.DepartmentName = ""
		}
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

		//unit.UnitDetail.Profile.UnitType = unitElm.UnitType
		unit.Profile.Location.Latitude = unitElm.Latitude
		//unit.UnitDetail.Profile.Location.Latitude = unitElm.Latitude
		unit.Profile.Location.Longitude = unitElm.Longitude

		fmt.Println(unit)
		units = append(units, unit)
	}
	send(units, w)
}