package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

// contract は契約の詳細情報を格納する
/*
type customer struct {
	// 契約情報
	Contracts []contractElm `json:"contracts"`
	// バッテリー情報
	Units []unitElm `json:"units"`
	//　顧客情報
	Customer customerElm `json:"customer"`
}
*/

type contractDetail struct{
	Contract []contractElm `json:"contracts"`
	DepartmentName string `json:"department_name"`
}

// CustomerView はCustomerページに必要なデータをDBから取得し、JSONで返す
func ContractView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "contract_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM contracts WHERE contract_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var contracts []contractElm
	for results1.Next() {
		var contract contractElm
		Columns := columns(&contract)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		contracts = append(contracts, contract)
	}
	fmt.Println(contracts)
	send(contracts, w)
}

func ContractDefaultView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "contract_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM contracts WHERE contract_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var contracts []ContractHomeDefaultElm
	for results1.Next() {
		var contract ContractHomeDefaultElm
		Columns := columns(&contract.Contract)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}

		var department_id = contract.Contract.DepartmentID

		var account_id int
		results2, err := db.Query("SELECT parent_id FROM departments WHERE department_id=" + strconv.Itoa(department_id))

		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(results2)
		for results2.Next() {
			//err = results2.Scan(&contract.AccountID)
			err = results2.Scan(&account_id)
			if err != nil {
				panic(err.Error())
			}
		}
		contract.AccountID = account_id

		//account_idをもとに、customer_nameの取得
		results3, err := db.Query("SELECT corporation_name FROM customers WHERE account_id=" + strconv.Itoa(account_id))
		
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(results3)
		for results3.Next() {
			err = results3.Scan(&contract.CustomerName)
			if err != nil {
				panic(err.Error())
			}
		}

		//battery optionの種別取得
		//contract_idをもとにbattery_optionsを取得
		var contract_id = contract.Contract.ContractID
		results4, err := db.Query("SELECT battery_option_id,info_type FROM battery_options WHERE contract_id=" + strconv.Itoa(contract_id))
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(results4)
		for results4.Next() {
			var battery_option_id int
			var info_type string
			err = results4.Scan(&battery_option_id,&info_type)
			if err != nil {
				panic(err.Error())
			}

			if (info_type=="field" || info_type==""){
				contract.BatteryFieldID=append(contract.BatteryFieldID,battery_option_id)
			} else if info_type=="request"{
				contract.BatteryRequestID=append(contract.BatteryRequestID,battery_option_id)
			} else if info_type=="manufacture"{
				contract.BatteryManuID=append(contract.BatteryManuID,battery_option_id)
			} else if info_type=="product"{
				contract.BatteryProductID=append(contract.BatteryProductID,battery_option_id)
			}
		}

		//chargerの種別取得
		results5, err := db.Query("SELECT charger_id,info_type FROM chargers WHERE contract_id=" + strconv.Itoa(contract_id))
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(results5)
		for results5.Next() {
			var charger_id int
			var info_type string
			err = results5.Scan(&charger_id,&info_type)
			if err != nil {
				panic(err.Error())
			}

			if (info_type=="field" || info_type==""){
				contract.ChargerFieldID=append(contract.ChargerFieldID,charger_id)
			} else if info_type=="request"{
				contract.ChargerRequestID=append(contract.ChargerRequestID,charger_id)
			} else if info_type=="manufacture"{
				contract.ChargerManuID=append(contract.ChargerManuID,charger_id)
			} else if info_type=="product"{
				contract.ChargerProductID=append(contract.ChargerProductID,charger_id)
			}
		}

		contracts = append(contracts, contract)
	}
	fmt.Println(contracts)
	send(contracts, w)
}

func CustomerContractView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "department_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM contracts WHERE department_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var contracts []contractElm
	for results1.Next() {
		var contract contractElm
		Columns := columns(&contract)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		contracts = append(contracts, contract)
	}
	fmt.Println(contracts)
	send(contracts, w)
}

func CreateContract(w http.ResponseWriter, r *http.Request) {
	//var customer customerElm
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	id := keyVal["contract_id"]
	field := keyVal["department_id"]
	name := keyVal["contract_name"]
	ctype := keyVal["contract_type"]
	st_date := keyVal["execution_date"]
	ed_date := keyVal["expiration_date"]
	fmt.Print(id,field,name,ctype,st_date,ed_date)
	//fmt.Println("new customer name: ", name)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO contracts(contract_id) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE contracts SET department_id = ? WHERE contract_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(field,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE contracts SET contract_name = ? WHERE contract_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE contracts SET contract_type = ? WHERE contract_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(ctype,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE contracts SET execution_date = ? WHERE contract_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(st_date,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE contracts SET expiration_date = ? WHERE contract_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(ed_date,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("create!", w)
}

func DeleteContract(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "contract_id")
	id := idtmp[0]

	db := open()
	defer db.Close()
	
	stmt, err := db.Prepare("DELETE FROM contracts WHERE contract_id = ?")

	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "contract ID = %s was deleted",id)
}