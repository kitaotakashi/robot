package db

import (
	"fmt"
	"net/http"
)

// detaile はバッテリーの詳細情報を格納する
type detaile struct {
	// 契約情報
	Contract contractElm `json:"contract"`
	// バッテリー情報
	Unit unitElm `json:"unit"`
	//　顧客情報
	Customer customerElm `json:"customer"`
}

// DetaileView はdetaileページに必要なデータをDBから取得し、JSONで返す
func DetaileView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "unit_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM units WHERE unit_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var detailes []detaile //detaileを複数格納するdata作成
	for results1.Next() {
		var detaile detaile
		Columns := columns(&detaile.Unit)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		results2, err := db.Query("SELECT * FROM customers WHERE account_id=(SELECT account_id FROM contracts WHERE unit_id=" + id[0] + ")")
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			Columns = columns(&detaile.Customer)
			err = results2.Scan(Columns...)
			if err != nil {
				panic(err.Error())
			}
		}
		results2, err = db.Query("SELECT * FROM contracts WHERE unit_id=" + id[0])
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			Columns = columns(&detaile.Contract)
			err = results2.Scan(Columns...)
		}
		detailes = append(detailes, detaile)
	}
	fmt.Println(detailes)
	send(detailes, w)
}
