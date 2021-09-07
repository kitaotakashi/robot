package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
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
	send(batteryoptions, w)
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

func BatteryRequestView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT battery_option_id,option_name,department_id,contract_id FROM battery_options WHERE contract_id IS NOT NULL")
	if err != nil {
		panic(err.Error())
	}
	var batteryrequests []batteryRequestElm
	for results.Next() {
		var batteryrequest batteryRequestElm
		var department_id int
		var contract_id int
		err = results.Scan(&batteryrequest.BatteryOptionId,&batteryrequest.OptionName,&department_id,&contract_id)
		if err != nil {
			panic(err.Error())
		}
		//get department name
		results2, err := db.Query("SELECT department_id,department_name FROM departments WHERE department_id =" + strconv.Itoa(department_id))
		if err != nil {
			panic(err.Error())
		}
		for results2.Next(){
			err = results2.Scan(&batteryrequest.DepartmentID,&batteryrequest.DepartmentName)
			if err != nil {
				panic(err.Error())
			}
		}
		//get contract name
		results3, err := db.Query("SELECT contract_id, contract_name FROM contracts WHERE contract_id =" + strconv.Itoa(contract_id))
		if err != nil {
			panic(err.Error())
		}
		for results3.Next(){
			err = results3.Scan(&batteryrequest.ContractID,&batteryrequest.ContractName)
			if err != nil {
				panic(err.Error())
			}
		}

		batteryrequests = append(batteryrequests, batteryrequest) //各customerをcustomersに格納
	}
	fmt.Println(batteryrequests)
	send(batteryrequests, w)
}

//契約詳細時に、contractに紐づいたoptionsを全て取得（field以外）
func ContractBatteryOptionView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "contract_id")
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM battery_options WHERE contract_id =" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var batteryoptions []batteryOptionElm
	for results.Next() {
		var batteryoption batteryOptionElm
		Columns := columns(&batteryoption)
		err = results.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		
		batteryoptions = append(batteryoptions, batteryoption) //各customerをcustomersに格納
	}
	fmt.Println(batteryoptions)
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
	contract := keyVal["contract_id"]
	udate := keyVal["update_date"]
	info := keyVal["info_type"]
	name := keyVal["option_name"]
	forklift := keyVal["forklift"]
	otype:= keyVal["type"]
	serial := keyVal["serial_number"]
	voltage := keyVal["voltage"]
	capacity := keyVal["capacity"]
	weight := keyVal["weight"]
	vertical := keyVal["vertical"]
	horizontal := keyVal["horizontal"]
	height := keyVal["height"]
	how2 := keyVal["how2change"]
	dwm := keyVal["daily_working_minute"]
	dcm := keyVal["daily_charging_minute"]
	noc := keyVal["number_of_change"]
	environment := keyVal["forklift_environment"]
	els := keyVal["environment_else"] 
	input := keyVal["input_plug"]
	output := keyVal["output_plug"]
	help := keyVal["change_help"]
	comment := keyVal["comment"]
	request := keyVal["request"]
	pic1 :=	keyVal["pic_forklift"]
	pic2 := keyVal["pic_forklift_plate"]
	pic3 :=	keyVal["pic_battery"]
	pic4 := keyVal["pic_battery_plate"]
	pic5 := keyVal["pic_change_place"]
	pic6 := keyVal["pic_battery_plug"]
	pic7 := keyVal["pic_forklift_plug"]
	pic8 := keyVal["pic_change_equipment"]

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

	stmt, err = db.Prepare("UPDATE battery_options SET contract_id = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 

	fmt.Println(err,contract)
	if contract != "-1" {
		_, err = stmt.Exec(contract,id)
	}else{
		fmt.Println("field registration")
	}

  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET update_date = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(udate,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET info_type = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(info,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET option_name = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET forklift = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(forklift,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET type = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(otype,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET serial_number = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(serial,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET voltage = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(voltage,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET capacity = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(capacity,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET weight = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(weight,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET vertical = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(vertical,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET horizontal = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(horizontal,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET height = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(height,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET how2change = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(how2,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET daily_working_minute = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dwm,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET daily_charging_minute = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dcm,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET number_of_change = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(noc,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET forklift_environment = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(environment,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET environment_else = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(els,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET input_plug = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(input,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET output_plug = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(output,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET change_help = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(help,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET comment = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(comment,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET request = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(request,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET pic_forklift = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic1,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET pic_forklift_plate = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic2,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET pic_battery = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic3,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET pic_battery_plate = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic4,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET pic_change_place = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic5,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET pic_battery_plug = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic6,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET pic_forklift_plug = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic7,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET pic_change_equipment = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic8,id)
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

func UpdateBatteryOption(w http.ResponseWriter, r *http.Request) {

	idtmp := query(r, "battery_option_id")
	id := idtmp[0]

	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
	did := keyVal["department_id"]
	contract := keyVal["contract_id"]
	udate := keyVal["update_date"]
	info := keyVal["info_type"]
	name := keyVal["option_name"]
	forklift := keyVal["forklift"]
	otype:= keyVal["type"]
	serial := keyVal["serial_number"]
	voltage := keyVal["voltage"]
	capacity := keyVal["capacity"]
	weight := keyVal["weight"]
	vertical := keyVal["vertical"]
	horizontal := keyVal["horizonal"]
	height := keyVal["height"]
	how2 := keyVal["how2change"]
	dwm := keyVal["daily_working_minute"]
	dcm := keyVal["daily_charging_minute"]
	noc := keyVal["number_of_change"]
	environment := keyVal["forklift_environment"]
	els := keyVal["environment_else"] 
	input := keyVal["input_plug"]
	output := keyVal["output_plug"]
	help := keyVal["change_help"]
	comment := keyVal["comment"]

	db := open()
	defer db.Close()
	
	stmt, err := db.Prepare("UPDATE battery_options SET department_id = ? WHERE battery_option_id = ?")
	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(did,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	  stmt, err = db.Prepare("UPDATE battery_options SET contract_id = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(contract,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET update_date = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(udate,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET info_type = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(info,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET option_name = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET forklift = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(forklift,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET type = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(otype,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET serial_number = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(serial,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET voltage = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(voltage,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET capacity = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(capacity,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET weight = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(weight,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET vertical = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(vertical,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET horizontal = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(horizontal,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET height = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(height,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET how2change = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(how2,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET daily_working_minute = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dwm,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET daily_charging_minute = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dcm,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET number_of_change = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(noc,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET forklift_environment = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(environment,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE battery_options SET environment_else = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(els,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET input_plug = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(input,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET output_plug = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(output,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET change_help = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(help,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE battery_options SET comment = ? WHERE battery_option_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(comment,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("update!", w)
	
}