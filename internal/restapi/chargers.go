package db

import (
	"fmt"
	"net/http"
)

func ChargersView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM chargers")
	if err != nil {
		panic(err.Error())
	}
	var chargers []chargerElm
	for results.Next() {
		var charger chargerElm
		columns := columns(&charger)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		chargers = append(chargers, charger)
	}
	fmt.Println(chargers)
	send(chargers, w)
}