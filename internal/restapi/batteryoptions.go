package db

import (
	"fmt"
	"net/http"
)

// BattteriesView はバッテリー情報(全件)をDBから取得してJSONでフロントに返す
func BatteryOptionsView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM battery_options")
	if err != nil {
		panic(err.Error())
	}
	var batteryoptions []batteryOptionElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var batteryoption batteryOptionElm
		columns := columns(&batteryoption)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		batteryoptions = append(batteryoptions, batteryoption) //各customerをcustomersに格納
	}
	fmt.Println(batteryoptions)
	send(batteryoptions, w)
}

func BatteryManufactureView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM battery_options")
	if err != nil {
		panic(err.Error())
	}
	var batteryoptions []batteryOptionElm
	for results1.Next() {
		var batteryoption batteryOptionElm
		Columns := columns(&batteryoption)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(batteryoption.InfoType.String)
		if batteryoption.InfoType.String == "manufacture"{
			batteryoptions = append(batteryoptions, batteryoption)
		}
	}
	fmt.Println(batteryoptions)
	send(batteryoptions, w)
}