package db

import (
	"fmt"
	"net/http"
	"strconv"
)

type ContractHomeElm struct {
	ContractHome	contractElm		`json:"contract"`
	//DepartmentName	string		`json:"department_name"`
	AccountID		int				`json:"account_id"`
}

type ContractHomeDefaultElm struct {
	Contract		contractElm		`json:"contract"`
	//DepartmentName	string		`json:"department_name"`
	AccountID		int				`json:"account_id"`
	CustomerName	string			`json:"customer_name"`
	BatteryFieldID	[]int			`json:"battery_field"`
	ChargerFieldID	[]int			`json:"charger_field"`
	BatteryRequestID	[]int		`json:"battery_request"`
	ChargerRequestID	[]int		`json:"charger_request"`
	BatteryManuID	[]int			`json:"battery_manufacture"`
	ChargerManuID	[]int			`json:"charger_manufacture"`
	BatteryProductID	[]int		`json:"battery_product"`
	ChargerProductID	[]int		`json:"charger_product"`
}

// CustomersView は顧客情報(全件)をDBから取得してJSONでフロントに返す
func ContractsView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM contracts")
	if err != nil {
		panic(err.Error())
	}
	var contracts []contractElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var contract contractElm
		columns := columns(&contract)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		contracts = append(contracts, contract) //各customerをcustomersに格納
	}
	fmt.Println(contracts)
	send(contracts, w)
}

func ContractsHomeView(w http.ResponseWriter, r *http.Request) {

	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM contracts")
	if err != nil {
		panic(err.Error())
	}
	var contracts []ContractHomeElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var contract ContractHomeElm
		columns := columns(&contract.ContractHome)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}

		var id = contract.ContractHome.DepartmentID
		fmt.Println(strconv.Itoa(id))

		results2, err := db.Query("SELECT parent_id FROM departments WHERE department_id=" + strconv.Itoa(id))
		
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(results2)
		for results2.Next() {
			err = results2.Scan(&contract.AccountID)
			if err != nil {
				panic(err.Error())
			}
		}
		
		contracts = append(contracts, contract) //各customerをcustomersに格納
	}
	fmt.Println(contracts)
	send(contracts, w)
}

func ContractsHomeDefaultView(w http.ResponseWriter, r *http.Request) {

	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM contracts")
	if err != nil {
		panic(err.Error())
	}
	var contracts []ContractHomeDefaultElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var contract ContractHomeDefaultElm
		columns := columns(&contract.Contract)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}

		var department_id = contract.Contract.DepartmentID
		//fmt.Println(strconv.Itoa(department_id))

		//account_idの取得
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

			if info_type=="field"{
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

			if info_type=="field"{
				contract.ChargerFieldID=append(contract.ChargerFieldID,charger_id)
			} else if info_type=="request"{
				contract.ChargerRequestID=append(contract.ChargerRequestID,charger_id)
			} else if info_type=="manufacture"{
				contract.ChargerManuID=append(contract.ChargerManuID,charger_id)
			} else if info_type=="product"{
				contract.ChargerProductID=append(contract.ChargerProductID,charger_id)
			}
		}
		
		contracts = append(contracts, contract) //各customerをcustomersに格納
	}
	fmt.Println(contracts)
	send(contracts, w)
}

