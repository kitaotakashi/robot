package db

import (
	"fmt"
	"net/http"
)

// DepartmantsView は顧客の事業所情報(全件)をDBから取得してJSONでフロントに返す
func DepartmentsView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM departments")
	if err != nil {
		panic(err.Error())
	}
	var departments []departmentElm //customerElmを複数格納するcustomers作成
	for results.Next() {
		var department departmentElm
		columns := columns(&department)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		departments = append(departments, department) //各customerをcustomersに格納
	}
	fmt.Println(departments)
	send(departments, w)
}
