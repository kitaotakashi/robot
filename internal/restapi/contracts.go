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

