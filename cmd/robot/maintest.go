//https://code-database.com/knowledges/87 htmlをgoで立てたサーバで展開する方法
//https://qiita.com/TakahiRoyte/items/949f4e88caecb02119aa#:~:text=REST(REpresentational%20State%20Transfer)%E3%81%AF,%E3%81%AE%E9%80%81%E5%8F%97%E4%BF%A1%E3%82%92%E8%A1%8C%E3%81%84%E3%81%BE%E3%81%99%E3%80%82
//↑RESTについて
package main

import (
	"fmt"      //標準入力など(デバッグ用なので最終的にはいらない...?)
	"net/http" //サーバを立てるために必要
	//"log"

	"../../internal/view"
	//"../../db"
	//db "app/internal/restapi"
	db "../../internal/restapi"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/rs/cors"
)

func challengetoken(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "rt0cYZ6L4gPOOYaPQgipkdG2ELQ0uQ21Ao46YjxsS98.XDi_5t34FW25GEQBQJPUAU2OKcjJutOUYefqngHTYxk")
}

func forCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		//origin := "http://localhost:3000"
        w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
        w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Methods", "GET,POST, DELETE, OPTIONS")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
        // プリフライトリクエストの対応
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
        return
    })
}

func main() {
	Server()
}

// Server はhttpリクエスト毎の処理を登録してサーバーを立てる
func Server() error {//logの場合はreturnがいらないのでerrorを消す
	//router := mux.NewRouter().StrictSlash(true)->corsが動かない原因かも
	router := mux.NewRouter()
	router.Use(forCORS)
	//front
	router.HandleFunc("/test/", OpenHtml.MainHandler)
	router.Handle("/", http.FileServer(http.Dir("../../front/build")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../front/build/static"))))
	//api
	router.HandleFunc("/api/v1/units/", db.UnitsView).Methods("GET")
	//router.HandleFunc("/api/v1/detaile/", db.DetaileView).Methods("GET")
	router.HandleFunc("/api/v1/unit/", db.DetailView).Methods("GET")

	router.HandleFunc("/api/v1/customers/", db.CustomersView).Methods("GET")
	router.HandleFunc("/api/v1/customer/", db.CustomerView).Methods("GET")
	router.HandleFunc("/api/v1/customer/",db.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customer/",db.DeleteCustomer).Methods("DELETE")

	router.HandleFunc("/api/v1/contracts/", db.ContractsView).Methods("GET")
	router.HandleFunc("/api/v1/contract/", db.ContractView).Methods("GET")
	router.HandleFunc("/api/v1/contract/", db.CreateContract).Methods("POST")

	router.HandleFunc("/api/v1/batteries/", db.BatteriesView).Methods("GET")
	router.HandleFunc("/api/v1/battery/", db.BatteryView).Methods("GET")
	router.HandleFunc("/api/v1/battery/", db.CreateBattery).Methods("POST")

	//others
	//router.HandleFunc("/api/v1/unit/", db.UnitView).Methods("GET")
	//router.HandleFunc("/.well-known/acme-challenge/rt0cYZ6L4gPOOYaPQgipkdG2ELQ0uQ21Ao46YjxsS98", challengetoken)//encryptの証明

	fmt.Println("RoBOT Server Started Port 443")

	//log.Fatal(http.ListenAndServeTLS(":443", "../../ssl/fullchain_new.pem", "../../ssl/server_new.key",router))
	//return http.ListenAndServe(fmt.Sprintf(":%d", 80), router)
	return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "../../ssl/fullchain.pem", "../../ssl/server.key", router) //kitao追加 https
	//return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "../../ssl/fullchain.pem", "../../ssl/server.key", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Origin", "application/json"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*", "http://localhost:3000"}))(router))
	
	/*
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000", "*"},
        AllowCredentials: true,
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization","Origin"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
    })

	handler := c.Handler(router)
    log.Fatal(http.ListenAndServeTLS(":443", "../../ssl/fullchain_new.pem", "../../ssl/server_new.key",handler))
	*/

	// http://18.180.144.98:80/
	// https://jugem.live/
	//nohup go run - &
}