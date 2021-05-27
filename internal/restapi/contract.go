package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
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

// CustomerView はCustomerページに必要なデータをDBから取得し、JSONで返す
func ContractView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "contract_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM contracts WHERE contract_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var contract contractElm
	for results1.Next() {
		Columns := columns(&contract)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(contract)
	send(contract, w)
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
	field := keyVal["account_id"]
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

	stmt, err = db.Prepare("UPDATE contracts SET account_id = ? WHERE contract_id = ?")
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