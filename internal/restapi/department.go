package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

// department は顧客の事業所情報を格納する
type department struct {
	// 契約情報
	//Contracts []contractElm `json:"contracts"`
	// バッテリー情報
	//Units []unitElm `json:"units"`
	//　顧客情報
	//Customer []customerElm `json:"customer"`

	//事業所情報
	Department []departmentElm `json:"department"`
}

// CustomerView はCustomerページに必要なデータをDBから取得し、JSONで返す
func DepartmentView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "department_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM departments WHERE department_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var departments []departmentElm
	for results1.Next() {
		var department departmentElm
		//Columns := columns(&department_input)
		Columns := columns(&department)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		//department.Department = append(department.Department, department_input)
		//department.Department = 
		/*
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
		
		results2, err = db.Query("SELECT * FROM customers WHERE account_id= " +department.Department.parent_id + ")")
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			var customer customerElm
			Columns = columns(&customer)
			err = results2.Scan(Columns...)
			if err != nil {
				panic(err.Error())
			}
			department.Customer = append(department.Customer, customer)
		}
		*/
		departments=append(departments,department)
	}
	fmt.Println(departments)
	send(departments, w)
}

func CustomerDepartmentView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "parent_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM departments WHERE parent_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var departments []departmentElm
	for results1.Next() {
		var department departmentElm
		//var department_input departmentElm
		//Columns := columns(&department_input)
		Columns := columns(&department)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		departments = append(departments, department)
	}
	fmt.Println(departments)
	send(departments, w)
}

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	//var customer customerElm
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	id := keyVal["department_id"]
	dname := keyVal["department_name"]
	pid := keyVal["parent_id"]//customer->account_id
	name := keyVal["name"]
	position := keyVal["position"]
	postal := keyVal["postal_code"]
	address := keyVal["address"]
	mail := keyVal["mail"]
	phone := keyVal["phone"]
	dwh := keyVal["daily_working_hour"]
	wh := keyVal["weekly_holiday"]

	fmt.Println(id,dname,pid,name,position,postal,address,mail,phone,dwh,wh)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO departments(department_id) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(id)
  	if err != nil {
    	panic(err.Error())
  	}

	  stmt, err = db.Prepare("UPDATE departments SET department_name = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dname,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET parent_id = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pid,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET name = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	}
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET position = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(position,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET postal_code = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(postal,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET address = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(address,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET mail = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(mail,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET phone = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(phone,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET daily_working_hour = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dwh,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE departments SET weekly_holiday = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(wh,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("create!", w)
}

func UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "department_id")
	id := idtmp[0]

	//var customer customerElm
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
	dname := keyVal["department_name"]
	pid := keyVal["parent_id"]//customer->account_id
	name := keyVal["name"]
	position := keyVal["position"]
	postal := keyVal["postal_code"]
	address := keyVal["address"]
	mail := keyVal["mail"]
	phone := keyVal["phone"]
	dwh := keyVal["daily_working_hour"]
	wh := keyVal["weekly_holiday"]

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE departments SET department_name = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dname,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET parent_id = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pid,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET name = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	}
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET position = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(position,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET postal_code = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(postal,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET address = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(address,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET mail = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(mail,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET phone = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(phone,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE departments SET daily_working_hour = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(dwh,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE departments SET weekly_holiday = ? WHERE department_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(wh,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("update!", w)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "department_id")
	id := idtmp[0]

	db := open()
	defer db.Close()
	
	//stmt, err := db.Prepare("DELETE FROM customers WHERE name = ?")
	stmt, err := db.Prepare("DELETE FROM departments WHERE department_id = ?")
	//stmt, err := db.Prepare("DELETE FROM customers WHERE date_of_birth = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Department with department ID = %s was deleted",id)
}

func DeleteFromDepartment(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "department_id")
	id := idtmp[0]

	db := open()
	defer db.Close()
	
	stmt, err := db.Prepare("DELETE FROM departments WHERE department_id = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Department with department ID = %s was deleted",id)

	//department_idに紐づいたcontractの削除
	results, err := db.Query("SELECT contract_id FROM contracts WHERE department_id="+ id)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var contract_id int
		err = results.Scan(&contract_id)
		
		//かつ各contract_idに紐づいたbatteries,charger_labelsの削除
		stmt, err := db.Prepare("DELETE FROM batteries WHERE contract_id = ?")
		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(contract_id)
		if err != nil {
			panic(err.Error())
		}

		//charger_idでmanuのものを取得
		results2, err := db.Query("SELECT charger_id FROM chargers WHERE department_id="+ id)
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			var charger_id int
			err = results2.Scan(&charger_id)

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
		fmt.Fprintf(w, "contract ID = %s was deleted",strconv.Itoa(contract_id))
	}

	//department_idに紐づいたbattery_optionとchargerの削除
	stmt, err = db.Prepare("DELETE FROM battery_options WHERE department_id = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}

	stmt, err = db.Prepare("DELETE FROM chargers WHERE department_id = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "battery & charger options with department ID = %s was deleted",id)
}