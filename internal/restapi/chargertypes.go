package db

import (
	"fmt"
	"net/http"
)

// BattteriesView はバッテリー情報(全件)をDBから取得してJSONでフロントに返す
func ChargerTypesView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM charger_types")
	if err != nil {
		panic(err.Error())
	}
	var chargertypes []chargerTypesElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var chargertype chargerTypesElm
		columns := columns(&chargertype)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		chargertypes = append(chargertypes, chargertype) //各customerをcustomersに格納
	}
	fmt.Println(chargertypes)
	send(chargertypes, w)
}