package db

import (
	"fmt"
	"net/http"
)

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