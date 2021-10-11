package db

import (
	"fmt"
	"net/http"
)

// BattteriesView はバッテリー情報(全件)をDBから取得してJSONでフロントに返す
func BatteryTypesView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM battery_types")
	if err != nil {
		panic(err.Error())
	}
	var batterytypes []batteryTypesElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var batterytype batteryTypesElm
		columns := columns(&batterytype)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		batterytypes = append(batterytypes, batterytype) //各customerをcustomersに格納
	}
	fmt.Println(batterytypes)
	send(batterytypes, w)
}