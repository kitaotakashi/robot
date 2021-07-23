package db

import (
	"fmt"
	"net/http"
)

func ErrorsView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM errors")
	if err != nil {
		panic(err.Error())
	}
	var errors []errorsElm
	for results.Next() {
		var error errorsElm
		columns := columns(&error)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		errors = append(errors, error)
	}
	fmt.Println(errors)
	send(errors, w)
}

func ErrorStatesView(w http.ResponseWriter, r *http.Request) {
	db := open()
	defer db.Close()
	results, err := db.Query("SELECT * FROM error_states")
	if err != nil {
		panic(err.Error())
	}
	var error_states []errorStatesElm
	for results.Next() {
		var error_state errorStatesElm
		columns := columns(&error_state)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		error_states = append(error_states, error_state)
	}
	fmt.Println(error_states)
	send(error_states, w)
}
