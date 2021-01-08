package db

import (
	"fmt"
	"net/http"
)

// customer は顧客の詳細情報を格納する
type customer struct {
	// 契約情報
	Contracts []contractElm `json:"contracts"`
	// バッテリー情報
	Units []unitElm `json:"units"`
	//　顧客情報
	Customer customerElm `json:"customer"`
}

// CustomerView はCustomerページに必要なデータをDBから取得し、JSONで返す
func CustomerView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "account_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM customers WHERE account_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var customer customer
	for results1.Next() {
		Columns := columns(&customer.Customer)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		results2, err := db.Query("SELECT * FROM contracts WHERE account_id=" + id[0])
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			var contract contractElm
			Columns = columns(&contract)
			err = results2.Scan(Columns...)
			if err != nil {
				panic(err.Error())
			}
			customer.Contracts = append(customer.Contracts, contract)
		}
		results2, err = db.Query("SELECT * FROM units WHERE unit_id=(SELECT unit_id FROM contracts WHERE account_id= " + id[0] + ")")
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			var unit unitElm
			Columns = columns(&unit)
			err = results2.Scan(Columns...)
			if err != nil {
				panic(err.Error())
			}
			customer.Units = append(customer.Units, unit)
		}
	}
	fmt.Println(customer)
	send(customer, w)
}
