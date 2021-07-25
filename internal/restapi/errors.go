package db

import (
	"fmt"
	"net/http"
	"strconv"
)

type ErrorTips struct{
	ErrorState 	errorStatesElm	`json:"error_state"`
	Errors		errorsElm		`json:"errors"`
	CorporationName	string		`json:"corporation_name"`
	DepartmentName	string		`json:"department_name"`
}

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

	var error_tips []ErrorTips

	results, err := db.Query("SELECT * FROM error_states")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var error_tip ErrorTips
		var error_state errorStatesElm
		columns := columns(&error_state)
		err = results.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}
		error_tip.ErrorState = error_state
		
		var id = error_state.ErrorCode
		fmt.Println("?")
		fmt.Println(strconv.Itoa(id))

		results2, err := db.Query("SELECT error_code,error_category,error_message,required_action FROM errors WHERE error_code=" + strconv.Itoa(id))
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			var error errorsElm
			fmt.Println(results2)
			//columns := columns(&error)
			//err = results2.Scan(&error)
			err = results2.Scan(&error.ErrorCode,&error.ErrorCategory,&error.ErrorMessage,&error.RequiredAction)
			if err != nil {
				panic(err.Error())
			}
			error_tip.Errors.ErrorCode = error.ErrorCode
			error_tip.Errors.ErrorCategory = error.ErrorCategory
			error_tip.Errors.ErrorMessage = error.ErrorMessage
			error_tip.Errors.RequiredAction = error.RequiredAction
		}
		error_tip.CorporationName=""
		error_tip.DepartmentName=""

		error_tips = append(error_tips, error_tip)
	}
	fmt.Println(error_tips)
	send(error_tips, w)
}
