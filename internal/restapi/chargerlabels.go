package db

import (
	"fmt"
	"net/http"
)

// BattteriesView はバッテリー情報(全件)をDBから取得してJSONでフロントに返す
func ChargerLabelsView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM charger_labels")
	if err != nil {
		panic(err.Error())
	}
	var chargerlabels []chargerLabelsElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var chargerlabel chargerLabelsElm
		columns := columns(&chargerlabel)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		chargerlabels = append(chargerlabels, chargerlabel) //各customerをcustomersに格納
	}
	fmt.Println(chargerlabels)
	send(chargerlabels, w)
}