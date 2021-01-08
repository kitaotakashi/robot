package db

import (
	"fmt"
	"net/http"
)

// CustomersView は顧客情報(全件)をDBから取得してJSONでフロントに返す
func CustomersView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM customers")
	if err != nil {
		panic(err.Error())
	}
	var customers []customerElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var customer customerElm
		columns := columns(&customer)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer) //各customerをcustomersに格納
	}
	fmt.Println(customers)
	send(customers, w)
}
