//https://code-database.com/knowledges/87 htmlをgoで立てたサーバで展開する方法
//https://qiita.com/TakahiRoyte/items/949f4e88caecb02119aa#:~:text=REST(REpresentational%20State%20Transfer)%E3%81%AF,%E3%81%AE%E9%80%81%E5%8F%97%E4%BF%A1%E3%82%92%E8%A1%8C%E3%81%84%E3%81%BE%E3%81%99%E3%80%82
//↑RESTについて
package main

import (
	"fmt"      //標準入力など(デバッグ用なので最終的にはいらない...?)
	"net/http" //サーバを立てるために必要

	//"../../internal/view"
	//"../../db"
	db "../../internal/restapi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	Server()
}

// Server はhttpリクエスト毎の処理を登録してサーバーを立てる
func Server() error {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", OpenHtml.MainHandler)
	router.Handle("/", http.FileServer(http.Dir("../../front/build")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../front/build/static"))))
	router.HandleFunc("/api/v1/units", db.UnitsView).Methods("GET")
	router.HandleFunc("/api/v1/customers", db.CustomersView).Methods("GET")
	router.HandleFunc("/api/v1/detaile", db.DetaileView).Methods("GET")
	router.HandleFunc("/api/v1/detail", db.DetailedView).Methods("GET")
	router.HandleFunc("/api/v1/customer", db.CustomerView).Methods("GET")
	//router.HandleFunc("/api/v1/unit/", db.UnitView).Methods("GET")
	fmt.Println("RoBOT Server Started Port 443")

	//return http.ListenAndServe(fmt.Sprintf(":%d", 80), router)
	//return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "../../ssl/fullchain.pem", "../../ssl/server.key", router) //kitao追加 https
	return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "../../ssl/fullchain.pem", "../../ssl/server.key", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}), handlers.AllowCredentials())(router))

	// http://18.180.144.98:80/
	// https://jugem.live/
	//nohup go run - &
}
