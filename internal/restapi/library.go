package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

// query はkeyがkのクエリパラメータを返す
func query(r *http.Request, k string) []string {
	v := r.URL.Query() // map[string][]string
	if v == nil {
		e := []string{"0"}
		return e
	}
	i, _ := v[k]
	return i
}

// open はデータベースと接続する
func open() *sql.DB {
	//db, err := sql.Open("mysql", "test_user:test_pass@tcp(10.0.1.229:3306)/test_db?parseTime=True")
	db, err := sql.Open("mysql", "test_user:test_pass@tcp(10.0.1.229:3306)/robot_db?parseTime=True")
	//db, err := sql.Open("mysql", "test_user:test_pass@tcp(10.0.1.229:3306)/mico_test?parseTime=True")
	if err != nil {
		panic(err.Error())
	}
	return db
}

// send はフロントにjsonデータを送る
func send(data interface{}, w http.ResponseWriter) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// columns は各Elm構造体のフィールドを格納した配列を返す
func columns(i interface{}) []interface{} {
	s := reflect.ValueOf(i).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	return columns
}

//func create()
