package db

import (
	"fmt"
	"net/http"
)

// BattteriesView はバッテリー情報(全件)をDBから取得してJSONでフロントに返す
func BatteriesView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM batteries")
	if err != nil {
		panic(err.Error())
	}
	var batteries []batteryElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var battery batteryElm
		columns := columns(&battery)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		batteries = append(batteries, battery) //各customerをcustomersに格納
	}
	fmt.Println(batteries)
	send(batteries, w)
}