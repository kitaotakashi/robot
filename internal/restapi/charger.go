package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// customer は顧客の詳細情報を格納する
type Charger struct {
	Charger chargerElm `json:"charger"`
}

// CustomerView はCustomerページに必要なデータをDBから取得し、JSONで返す
func ChargerView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "charger_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM chargers WHERE charger_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var chargers []chargerElm
	for results1.Next() {
		var charger chargerElm
		Columns := columns(&charger)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		chargers = append(chargers, charger)
	}
	fmt.Println(chargers)
	send(chargers, w)
}

func CustomerChargerView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "department_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM chargers WHERE department_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var chargers []chargerElm
	for results1.Next() {
		var charger chargerElm
		Columns := columns(&charger)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		chargers = append(chargers, charger)
	}
	fmt.Println(chargers)
	send(chargers, w)
}

func ContractChargerView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "contract_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM chargers WHERE contract_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var chargers []chargerElm
	for results1.Next() {
		var charger chargerElm
		Columns := columns(&charger)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		chargers = append(chargers, charger)
	}
	fmt.Println(chargers)
	send(chargers, w)
}

func CreateCharger(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	id := keyVal["charger_id"]
	did := keyVal["department_id"]
	contract := keyVal["contract_id"]
	udate := keyVal["update_date"]
	info := keyVal["info_type"]
	name := keyVal["charger_name"]
	environment := keyVal["charge_environment"]
	els := keyVal["charge_environment_else"]
	how2 := keyVal["how2supply"]
	splug := keyVal["supply_plug"]
	sels := keyVal["supply_else"]
	psa := keyVal["power_supply_ampere"]
	stand := keyVal["stand"]
	//ps2c := keyVal["power_supply2charger_cable_langth"]
	//c2f := keyVal["charger2forklift_cable_langth"]
	help := keyVal["charger_setting_help"]
	comment := keyVal["comment"]
	request := keyVal["request"]
	pic1 :=	keyVal["pic_charger_stand"]
	pic2 := keyVal["pic_power_supply"]
	pic3 := keyVal["pic_supply_plug"]

	fmt.Println(id,did)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO chargers(charger_id) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET department_id = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(did,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET contract_id = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(contract,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE chargers SET update_date = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(udate,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET info_type = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(info,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET charger_name = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET charge_environment = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(environment,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET charge_environment_else = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(els,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET how2supply = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(how2,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET supply_plug = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(splug,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET supply_else = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(sels,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET power_supply_ampere = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(psa,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET stand = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(stand,id)
  	if err != nil {
    	panic(err.Error())
  	}
	/*
	stmt, err = db.Prepare("UPDATE chargers SET power_supply2charger_cable_langth = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(ps2c,id)
  	if err != nil {
    	panic(err.Error())
  	}*/
	
	/*
	stmt, err = db.Prepare("UPDATE chargers SET charger2forklift_cable_langth = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(c2f,id)
  	if err != nil {
    	panic(err.Error())
  	}*/

	stmt, err = db.Prepare("UPDATE chargers SET charger_setting_help = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(help,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET comment = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(comment,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET request = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(request,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET pic_charger_stand = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic1,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET pic_power_supply = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic2,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE chargers SET pic_supply_plug = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(pic3,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	send("create!", w)
	
}

func DeleteCharger(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "charger_id")
	id := idtmp[0]

	db := open()
	defer db.Close()
	
	stmt, err := db.Prepare("DELETE FROM chargers WHERE charger_id = ?")

	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(id)
   	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "charger ID = %s was deleted",id)
}


func UpdateCharger(w http.ResponseWriter, r *http.Request) {

	idtmp := query(r, "charger_id")
	id := idtmp[0]

	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
	did := keyVal["department_id"]
	contract := keyVal["contract_id"]
	udate := keyVal["update_date"]
	info := keyVal["info_type"]
	name := keyVal["charger_name"]
	environment := keyVal["charge_environment"]
	els := keyVal["charge_environment_else"]
	how2 := keyVal["how2supply"]
	splug := keyVal["supply_plug"]
	sels := keyVal["supply_else"]
	psa := keyVal["power_supply_ampere"]
	stand := keyVal["stand"]
	//ps2c := keyVal["power_supply2charger_cable_langth"]
	//c2f := keyVal["charger2forklift_cable_langth"]
	help := keyVal["charger_setting_help"]
	comment := keyVal["comment"]

	db := open()
	defer db.Close()
	
	stmt, err := db.Prepare("UPDATE chargers SET department_id = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(did,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET contract_id = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(contract,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE chargers SET update_date = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(udate,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET info_type = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(info,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET charger_name = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(name,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET charge_environment = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(environment,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET charge_environment_else = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(els,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET how2supply = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(how2,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET supply_plug = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(splug,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET supply_else = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(sels,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET power_supply_ampere = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(psa,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET stand = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(stand,id)
  	if err != nil {
    	panic(err.Error())
  	}
	/*
	stmt, err = db.Prepare("UPDATE chargers SET power_supply2charger_cable_langth = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(ps2c,id)
  	if err != nil {
    	panic(err.Error())
  	}*/
	
	/*
	stmt, err = db.Prepare("UPDATE chargers SET charger2forklift_cable_langth = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(c2f,id)
  	if err != nil {
    	panic(err.Error())
  	}*/
	
	stmt, err = db.Prepare("UPDATE chargers SET charger_setting_help = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(help,id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE chargers SET comment = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(comment,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("update!", w)
	
}