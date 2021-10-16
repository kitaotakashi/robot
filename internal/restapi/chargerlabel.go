package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// BatteryView はbatteryに格納しているデータをDBから取得し、JSONで返す
func ChargerLabelView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "charger_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM charger_labels WHERE charger_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var chargerlabels []chargerLabelsElm
	for results1.Next() {
		var chargerlabel chargerLabelsElm
		Columns := columns(&chargerlabel)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		chargerlabels = append(chargerlabels, chargerlabel)
	}
	fmt.Println(chargerlabels)
	send(chargerlabels, w)
}

// BatteryView はbatteryに格納しているデータをDBから取得し、JSONで返す
func SerialChargerLabelView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "contract_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM charger_labels WHERE serial_number=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var chargerlabels []chargerLabelsElm
	for results1.Next() {
		var chargerlabel chargerLabelsElm
		Columns := columns(&chargerlabel)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		chargerlabels = append(chargerlabels, chargerlabel)
	}
	fmt.Println(chargerlabels)
	send(chargerlabels, w)
}

func CreateChargerLabels(w http.ResponseWriter, r *http.Request) {
	//var battery batteryElm
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
  	id := keyVal["charger_id"]
	serial := keyVal["serial_number"]
	typeid := keyVal["charger_type_id"]
	fmt.Print(id,serial,typeid)
	//fmt.Println("new customer name: ", name)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO charger_labels(charger_id) VALUES(?)")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(id)
  	if err != nil {
    	panic(err.Error())
  	}

	stmt, err = db.Prepare("UPDATE charger_labels SET serial_number = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(serial,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE charger_labels SET charger_type_id = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(typeid,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("create!", w)
}

func UpdateChargerLabels(w http.ResponseWriter, r *http.Request) {
	idtmp := query(r, "charger_id")
	id := idtmp[0]

	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)
	serial := keyVal["serial_number"]
	typeid := keyVal["charger_type_id"]
	fmt.Print(serial,typeid)
	//fmt.Println("new customer name: ", name)

	db := open()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE charger_labels SET serial_number = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(serial,id)
  	if err != nil {
    	panic(err.Error())
  	}
	
	stmt, err = db.Prepare("UPDATE charger_labels SET charger_type_id = ? WHERE charger_id = ?")
  	if err != nil {
    	panic(err.Error())
  	} 
	_, err = stmt.Exec(typeid,id)
  	if err != nil {
    	panic(err.Error())
  	}

	send("update!", w)
}