package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// BatteryView はbatteryに格納しているデータをDBから取得し、JSONで返す
func BatteryView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "serial_number")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM batteries WHERE serial_number=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var battery batteryElm
	for results1.Next() {
		Columns := columns(&battery)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(battery)
	send(battery, w)
}

// BatteryView はbatteryに格納しているデータをDBから取得し、JSONで返す
func ContractBatteryView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "contract_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM batteries WHERE contract_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var batteries []batteryElm
	for results1.Next() {
		var battery batteryElm
		Columns := columns(&battery)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		batteries = append(batteries, battery)
	}
	fmt.Println(batteries)
	send(batteries, w)
}

func BatteryOptionBatteryView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "battery_option_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM batteries WHERE battery_option_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var batteries []batteryElm
	for results1.Next() {
		var battery batteryElm
		Columns := columns(&battery)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		batteries = append(batteries, battery)
	}
	fmt.Println(batteries)
	send(batteries, w)
}

func CreateBattery(w http.ResponseWriter, r *http.Request) {
	//var battery batteryElm
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	id := keyVal["serial_number"]
	unit := keyVal["unit_id"]
	contract:= keyVal["contract_id"]
	option := keyVal["battery_option_id"]
	btype := keyVal["battery_type_id"]
	mf_date := keyVal["date_of_manufacture"]
	purpose := keyVal["purpose"]
	state := keyVal["unit_state"]
	fmt.Print(id,unit,contract,option,btype,mf_date,purpose,state)
	//fmt.Println("new customer name: ", name)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO batteries(serial_number) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET unit_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(unit,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE batteries SET contract_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(contract,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET battery_option_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(option,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET battery_type_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(btype,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET date_of_manufacture = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(mf_date,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET purpose = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(purpose,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET unit_state = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(state,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("create!", w)
}

func UpdateBattery(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "serial_number")
	id := idtmp[0]

	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
	unit := keyVal["unit_id"]
	contract:= keyVal["contract_id"]
	option := keyVal["battery_option_id"]
	btype := keyVal["battery_type_id"]
	mf_date := keyVal["date_of_manufacture"]
	purpose := keyVal["purpose"]
	state := keyVal["unit_state"]
	fmt.Print(id,unit,contract,option,btype,mf_date,purpose,state)
	//fmt.Println("new customer name: ", name)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE batteries SET unit_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(unit,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE batteries SET contract_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(contract,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET battery_option_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(option,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET battery_type_id = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(btype,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET date_of_manufacture = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(mf_date,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET purpose = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(purpose,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE batteries SET unit_state = ? WHERE serial_number = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(state,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("update!", w)
}