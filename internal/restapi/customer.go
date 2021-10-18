package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

// customer は顧客の詳細情報を格納する
type customer struct {
	// 契約情報
	//Contracts []contractElm `json:"contracts"`
	// バッテリー情報
	//Units []unitElm `json:"units"`
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
		//契約情報
		/*
		results2, err := db.Query("SELECT * FROM contracts WHERE department_id=" + id[0])
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
		//バッテリー情報
		/*
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
		}*/
	}
	fmt.Println(customer)
	send(customer, w)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	//var customer customerElm
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	id := keyVal["account_id"]
	cname := keyVal["corporation_name"]
	sector := keyVal["sector"]
	name := keyVal["name"]
	position := keyVal["position"]
	//dob := keyVal["date_of_birth"]
	postal := keyVal["postal_code"]
	address := keyVal["address"]
	mail := keyVal["mail"]
	phone := keyVal["phone"]
	fmt.Println(id,cname,sector,name,position,postal,address,mail,phone)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO customers(account_id) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(id)
  	if err != nil {
    	panic(err.Error())
  	}

	  stmt, err = db.Prepare("UPDATE customers SET corporation_name = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(cname,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET sector = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(sector,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET name = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET position = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(position,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET postal_code = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(postal,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET address = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(address,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET mail = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(mail,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET phone = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(phone,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("create!", w)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "account_id")
	id := idtmp[0]

	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
	cname := keyVal["corporation_name"]
	sector := keyVal["sector"]
	name := keyVal["name"]
	position := keyVal["position"]
	postal := keyVal["postal_code"]
	address := keyVal["address"]
	mail := keyVal["mail"]
	phone := keyVal["phone"]

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE customers SET corporation_name = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(cname,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET sector = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(sector,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET name = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET position = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(position,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET postal_code = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(postal,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET address = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(address,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET mail = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(mail,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE customers SET phone = ? WHERE account_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(phone,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("update!", w)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "account_id")
	id := idtmp[0]

	db := open()
	defer db.Close()
	
	//stmt, err := db.Prepare("DELETE FROM customers WHERE name = ?")
	stmt, err := db.Prepare("DELETE FROM customers WHERE account_id = ?")
	//stmt, err := db.Prepare("DELETE FROM customers WHERE date_of_birth = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Post with account ID = %s was deleted",id)
}

func DeleteFromCustomer(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "account_id")
	id := idtmp[0]

	db := open()
	defer db.Close()
	
	stmt, err := db.Prepare("DELETE FROM customers WHERE account_id = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Post with account ID = %s was deleted",id)

	//account_idに紐づいたdepartment_idの取得
	results, err := db.Query("SELECT department_id FROM departments WHERE parent_id="+ id)
	if err != nil {
		panic(err.Error())
	}
	//department_idに紐づいた各種削除
	for results.Next() {
		var department_id int
		err = results.Scan(&department_id)

		//fmt.Fprintf(w,"department id is"+strconv.Itoa(department_id))

		//departmentの削除
		stmt, err := db.Prepare("DELETE FROM departments WHERE department_id = ?")
		if err != nil {
		panic(err.Error())
		}
		_, err = stmt.Exec(department_id)
		if err != nil {
		panic(err.Error())
		}

		//department_idに紐づいたcontractの削除
		results2, err := db.Query("SELECT contract_id FROM contracts WHERE department_id="+ strconv.Itoa(department_id))
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			var contract_id int
			err = results2.Scan(&contract_id)
			//fmt.Fprintf(w,"contract id is"+strconv.Itoa(contract_id))

			//各contract_idに紐づいたbatteriesの削除
			stmt, err := db.Prepare("DELETE FROM batteries WHERE contract_id = ?")
			if err != nil {
				panic(err.Error())
			}
			_, err = stmt.Exec(contract_id)
			if err != nil {
				panic(err.Error())
			}

			//charger_idでmanuのものを取得
			results3, err := db.Query("SELECT charger_id FROM chargers WHERE department_id="+ strconv.Itoa(department_id))
			if err != nil {
				panic(err.Error())
			}
			//charger_labelsの削除
			for results3.Next() {
				var charger_id int
				err = results3.Scan(&charger_id)
	
				stmt, err := db.Prepare("DELETE FROM charger_labels WHERE charger_id = ?")
				if err != nil {
					panic(err.Error())
				}
				_, err = stmt.Exec(charger_id)
				if err != nil {
					panic(err.Error())
				}
			}
	
			//契約削除
			stmt, err = db.Prepare("DELETE FROM contracts WHERE contract_id = ?")
			if err != nil {
				panic(err.Error())
			}
			_, err = stmt.Exec(contract_id)
			if err != nil {
				panic(err.Error())
			}
		}

		//department_idに紐づいたbattery_optionとchargerの削除
		stmt, err = db.Prepare("DELETE FROM battery_options WHERE department_id = ?")
		if err != nil {
		panic(err.Error())
		}
		_, err = stmt.Exec(department_id)
		if err != nil {
		panic(err.Error())
		}

		stmt, err = db.Prepare("DELETE FROM chargers WHERE department_id = ?")
		if err != nil {
		panic(err.Error())
		}
		_, err = stmt.Exec(department_id)
		if err != nil {
		panic(err.Error())
		}
	}
}