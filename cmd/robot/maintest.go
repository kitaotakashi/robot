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
    fmt.Fprintf(w, "0aHIYjg2h4AQ0bu6MJOBe5-Qgm1jBpiWTbUSiptyZ40.XDi_5t34FW25GEQBQJPUAU2OKcjJutOUYefqngHTYxk")
}

func forCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		//origin := "http://localhost:3000"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Origin", origin)

        w.Header().Set("Access-Control-Allow-Headers", "origin, X-Requested-With, Content-Type, Accept")
		//w.Header().Set("Access-Control-Allow-Headers", "*")
		
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
        // プリフライトリクエストの対応
        if r.Method == "OPTIONS" {
			fmt.Println("pf2")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Methods", "DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "*")
    		//w.WriteHeader(200)
            w.WriteHeader(http.StatusOK)
            return
		}
		/*
		if r.Method == http.MethodOptions {
			fmt.Println("pf")

			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}*/
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
	//router.HandleFunc("/api/v1/customer/",db.CreateCustomer)
	router.HandleFunc("/api/v1/customer/", db.CustomerView).Methods("GET")
	//router.HandleFunc("/api/v1/customer/post/",db.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customer/post/",db.CreateCustomer)
	router.HandleFunc("/api/v1/customer/delete/",db.DeleteCustomer)

	router.HandleFunc("/api/v1/contracts/", db.ContractsView).Methods("GET")
	router.HandleFunc("/api/v1/contractshome/", db.ContractsHomeView).Methods("GET")
	router.HandleFunc("/api/v1/contractshomedefault/", db.ContractsHomeDefaultView).Methods("GET")
	router.HandleFunc("/api/v1/contract/", db.ContractView).Methods("GET")
	router.HandleFunc("/api/v1/contractdefault/", db.ContractDefaultView).Methods("GET")
	router.HandleFunc("/api/v1/customercontract/", db.CustomerContractView).Methods("GET")
	router.HandleFunc("/api/v1/customercontractdefault/", db.CustomerContractDefaultView).Methods("GET")
	//router.HandleFunc("/api/v1/contract/", db.CreateContract).Methods("POST")
	router.HandleFunc("/api/v1/contract/post/", db.CreateContract)
	router.HandleFunc("/api/v1/contract/delete/",db.DeleteContract)

	router.HandleFunc("/api/v1/batteries/", db.BatteriesView).Methods("GET")
	router.HandleFunc("/api/v1/battery/", db.BatteryView).Methods("GET")
	//router.HandleFunc("/api/v1/battery/", db.CreateBattery).Methods("POST")
	router.HandleFunc("/api/v1/battery/post/", db.CreateBattery)

	router.HandleFunc("/api/v1/departments/", db.DepartmentsView).Methods("GET")
	router.HandleFunc("/api/v1/department/", db.DepartmentView).Methods("GET")
	router.HandleFunc("/api/v1/customerdepartment/", db.CustomerDepartmentView).Methods("GET")
	//router.HandleFunc("/api/v1/department/", db.CreateDepartment).Methods("POST")
	router.HandleFunc("/api/v1/department/post/", db.CreateDepartment)
	router.HandleFunc("/api/v1/department/delete/",db.DeleteDepartment)

	router.HandleFunc("/api/v1/batteryoptions/", db.BatteryOptionsView).Methods("GET")
	router.HandleFunc("/api/v1/batteryoption/", db.BatteryOptionView).Methods("GET")
	router.HandleFunc("/api/v1/customerbatteryoption/", db.CustomerBatteryOptionView).Methods("GET")
	router.HandleFunc("/api/v1/contractbatteryoption/", db.ContractBatteryOptionView).Methods("GET")
	router.HandleFunc("/api/v1/contractbatteryoption/delete/", db.DeleteContractBatteryOption)
	router.HandleFunc("/api/v1/batteryoption/post/", db.CreateBatteryOption)
	router.HandleFunc("/api/v1/batteryoption/put/", db.UpdateBatteryOption)
	router.HandleFunc("/api/v1/batteryoption/delete/",db.DeleteBatteryOption)

	router.HandleFunc("/api/v1/batteryrequests/", db.BatteryRequestView).Methods("GET")

	router.HandleFunc("/api/v1/chargers/", db.ChargersView).Methods("GET")
	router.HandleFunc("/api/v1/charger/", db.ChargerView).Methods("GET")
	router.HandleFunc("/api/v1/customercharger/", db.CustomerChargerView).Methods("GET")
	router.HandleFunc("/api/v1/contractcharger/", db.ContractChargerView).Methods("GET")
	router.HandleFunc("/api/v1/contractcharger/delete/", db.DeleteContractCharger)
	router.HandleFunc("/api/v1/charger/post/", db.CreateCharger)
	router.HandleFunc("/api/v1/charger/delete/",db.DeleteCharger)

	router.HandleFunc("/api/v1/errors/",db.ErrorsView).Methods("GET")
	router.HandleFunc("/api/v1/error/",db.ErrorView).Methods("GET")
	router.HandleFunc("/api/v1/errorstates/",db.ErrorStatesView).Methods("GET")
	router.HandleFunc("/api/v1/errorstate/",db.ErrorStateView).Methods("GET")

	//others
	//router.HandleFunc("/api/v1/unit/", db.UnitView).Methods("GET")
	//router.HandleFunc("/.well-known/acme-challenge/0aHIYjg2h4AQ0bu6MJOBe5-Qgm1jBpiWTbUSiptyZ40", challengetoken)//encryptの証明

	fmt.Println("RoBOT Server Started Port 443")

	//return http.ListenAndServe(fmt.Sprintf(":%d", 80), router)
	return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "../../ssl/fullchain.pem", "../../ssl/server.key", router) //kitao追加 https
	
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