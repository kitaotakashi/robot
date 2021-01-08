package db

import (
	"fmt"
	"net/http"
)

// onlyUnitsView はバッテリー情報(全件)をDBから取得してJSONで返す
func onlyUnitsView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close() //エラーが起きてもdb.Close()は実行する
	results, err := db.Query("SELECT * FROM units")
	if err != nil {
		panic(err.Error())
	}
	var units []unitElm //unitElmを複数格納するunits作成
	for results.Next() {
		var unit unitElm
		columns := columns(&unit)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		units = append(units, unit) //各unitをunitsに格納
	}
	fmt.Println(units)
	send(units, w)
}
