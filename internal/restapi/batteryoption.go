package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// customer は顧客の詳細情報を格納する
type BatteryOption struct {
	BatteryOption batteryOptionElm `json:"battery_option"`
}

// CustomerView はCustomerページに必要なデータをDBから取得し、JSONで返す
func BatteryOptionView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "battery_option_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM battery_options WHERE battery_option_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var batteryoption BatteryOption
	for results1.Next() {
		Columns := columns(&batteryoption.BatteryOption)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(batteryoption)
	send(batteryoption, w)
}

func CustomerBatteryOptionView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "department_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM battery_options WHERE department_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var batteryoptions []batteryOptionElm
	for results1.Next() {
		var batteryoption batteryOptionElm
		Columns := columns(&batteryoption)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		batteryoptions = append(batteryoptions, batteryoption)
	}
	fmt.Println(batteryoptions)
	/*
	len := len(batteryoption)
	i:= 0
	for i < len {
		if batteryoption[i][1]!= id[0] {
			batteryoption[i]=""
		}
	}
	*/
	send(batteryoptions, w)
}

func CreateBatteryOption(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	id := keyVal["battery_option_id"]
	did := keyVal["department_id"]

	fmt.Println(id,did)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO battery_options(battery_option_id) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET department_id = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(did,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("create!", w)
}

/*
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer customerElm
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	name := keyVal["name"]
	//fmt.Println("new customer name: ", name)

	//json.NewDecoder(r.Body).Decode(&customer)
    //fmt.Println("new customer: ", customer)
	//fmt.Println("new customer name: ", customer["name"])

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO customers(name) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 

	_, err = stmt.Exec(name)
  	if err != nil {
    	panic(err.Error())
  	}

	//test := Test{ID:1, FirstName:"kitao"} 
	send("update", w)
}
*/

func DeleteBatteryOption(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "battery_option_id")
	id := idtmp[0]

	db := open()
	defer db.Close()
	
	stmt, err := db.Prepare("DELETE FROM battery_options WHERE battery_option_id = ?")

	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "battery_option ID = %s was deleted",id)
}