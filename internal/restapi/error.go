package db

import (
	"fmt"
	"net/http"
)

func ErrorView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "error_code")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM errors WHERE error_code=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var errors []errorsElm
	for results1.Next() {
		var error errorsElm
		Columns := columns(&error)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		errors = append(errors, error)
	}
	fmt.Println(errors)
	send(errors, w)
}

func ErrorStateView(w http.ResponseWriter, r *http.Request) {
	id := query(r, "object_id")
	db := open()
	defer db.Close()
	results1, err := db.Query("SELECT * FROM error_states WHERE object_id=" + id[0])
	if err != nil {
		panic(err.Error())
	}
	var error_states []errorStatesElm
	for results1.Next() {
		var error_state errorStatesElm
		Columns := columns(&error_state)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		error_states = append(error_states, error_state)
	}
	fmt.Println(error_states)
	send(error_states, w)
}

