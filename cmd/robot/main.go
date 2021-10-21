//https://code-database.com/knowledges/87 htmlをgoで立てたサーバで展開する方法
//https://qiita.com/TakahiRoyte/items/949f4e88caecb02119aa#:~:text=REST(REpresentational%20State%20Transfer)%E3%81%AF,%E3%81%AE%E9%80%81%E5%8F%97%E4%BF%A1%E3%82%92%E8%A1%8C%E3%81%84%E3%81%BE%E3%81%99%E3%80%82
//↑RESTについて
package main

import (
	//"../../OpenHtml"
	//"../../db"
	//"../../internal/restapi"
	db "../../internal/restapi"
	//"../../internal/view"
	"encoding/json"
	"errors"
	"fmt" //標準入力など(デバッグ用なので最終的にはいらない...?)
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http" //サーバを立てるために必要

	"../../internal/view"
)

func challengetoken(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "token")
}

func forCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		//origin := "http://localhost:3000"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Origin", origin)

        w.Header().Set("Access-Control-Allow-Headers", "origin, X-Requested-With, Content-Type, Accept")
		//w.Header().Set("Access-Control-Allow-Headers", "*")
		
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, PUT")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
        // プリフライトリクエストの対応
        if r.Method == "OPTIONS" {
			fmt.Println("get pf")
			//w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			//w.Header().Set("Access-Control-Allow-Headers", "origin, X-Requested-With, Content-Type, Accept")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, PUT")
			/*
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Methods", "DELETE")
			w.Header().Set("Access-Control-Allow-Methods", "PUT")
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
			*/
    		//w.WriteHeader(200)
            w.WriteHeader(http.StatusOK)
			fmt.Println("pf ok")
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

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func main() {
	Server()
}

// Server はhttpリクエスト毎の処理を登録してサーバーを立てる
func Server() error {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			//aud := "https://mico.center/"
			aud := "https://dev-transkit.jp.auth0.com/api/v2/"
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			// fmt.Printf("%v", token)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}
			// Verify 'iss' claim
			iss := "https://dev-transkit.jp.auth0.com/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := getPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	router := mux.NewRouter()
	router.Use(forCORS)
	//front
	router.HandleFunc("/test/", OpenHtml.MainHandler)
	router.Handle("/", http.FileServer(http.Dir("../../front/build")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../front/build/static"))))
	//api
	//router.HandleFunc("/api/v1/units/", db.UnitsView).Methods("GET")
	router.Handle("/api/v1/units/", jwtMiddleware.Handler(http.HandlerFunc(db.UnitsView))).Methods("GET")
	//router.HandleFunc("/api/v1/detaile/", db.DetaileView).Methods("GET")
	router.HandleFunc("/api/v1/unit/", db.DetailView).Methods("GET")
	router.HandleFunc("/api/v1/contractunit/", db.ContractDetailView).Methods("GET")

	router.HandleFunc("/api/v1/customers/", db.CustomersView).Methods("GET")
	//router.HandleFunc("/api/v1/customer/",db.CreateCustomer)
	router.HandleFunc("/api/v1/customer/", db.CustomerView).Methods("GET")
	//router.HandleFunc("/api/v1/customer/post/",db.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customer/post/",db.CreateCustomer)
	router.HandleFunc("/api/v1/customer/put/",db.UpdateCustomer)
	router.HandleFunc("/api/v1/customer/delete/",db.DeleteCustomer)
	router.HandleFunc("/api/v1/fromcustomer/delete/",db.DeleteFromCustomer)

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
	router.HandleFunc("/api/v1/contractbattery/", db.ContractBatteryView).Methods("GET")
	router.HandleFunc("/api/v1/batteryoptionbattery/", db.BatteryOptionBatteryView).Methods("GET")
	//router.HandleFunc("/api/v1/battery/", db.CreateBattery).Methods("POST")
	router.HandleFunc("/api/v1/battery/post/", db.CreateBattery)
	router.HandleFunc("/api/v1/battery/put/", db.UpdateBattery)
	router.HandleFunc("/api/v1/battery/delete/", db.DeleteBattery)

	router.HandleFunc("/api/v1/chargerlabels/", db.ChargerLabelsView).Methods("GET")
	router.HandleFunc("/api/v1/chargerlabel/", db.ChargerLabelView).Methods("GET")
	router.HandleFunc("/api/v1/serialchargerlabel/", db.SerialChargerLabelView).Methods("GET")
	router.HandleFunc("/api/v1/chargerlabel/post/", db.CreateChargerLabels)
	router.HandleFunc("/api/v1/chargerlabel/put/", db.UpdateChargerLabels)
	router.HandleFunc("/api/v1/chargerlabel/delete/", db.DeleteChargerLabels)

	router.HandleFunc("/api/v1/departments/", db.DepartmentsView).Methods("GET")
	router.HandleFunc("/api/v1/department/", db.DepartmentView).Methods("GET")
	router.HandleFunc("/api/v1/customerdepartment/", db.CustomerDepartmentView).Methods("GET")
	//router.HandleFunc("/api/v1/department/", db.CreateDepartment).Methods("POST")
	router.HandleFunc("/api/v1/department/post/", db.CreateDepartment)
	router.HandleFunc("/api/v1/department/put/", db.UpdateDepartment)
	router.HandleFunc("/api/v1/department/delete/",db.DeleteDepartment)
	router.HandleFunc("/api/v1/fromdepartment/delete/",db.DeleteFromDepartment)

	router.HandleFunc("/api/v1/batteryoptions/", db.BatteryOptionsView).Methods("GET")
	router.HandleFunc("/api/v1/batteryoption/", db.BatteryOptionView).Methods("GET")
	router.HandleFunc("/api/v1/customerbatteryoption/", db.CustomerBatteryOptionView).Methods("GET")
	router.HandleFunc("/api/v1/contractbatteryoption/", db.ContractBatteryOptionView).Methods("GET")
	router.HandleFunc("/api/v1/batterymanufacture/", db.BatteryManufactureView).Methods("GET")
	router.HandleFunc("/api/v1/contractbatteryoption/delete/", db.DeleteContractBatteryOption)
	router.HandleFunc("/api/v1/batteryoption/post/", db.CreateBatteryOption)
	router.HandleFunc("/api/v1/batteryoption/put/", db.UpdateBatteryOption)
	router.HandleFunc("/api/v1/batteryoption/delete/",db.DeleteBatteryOption)

	router.HandleFunc("/api/v1/batteryrequests/", db.BatteryRequestView).Methods("GET")

	router.HandleFunc("/api/v1/chargers/", db.ChargersView).Methods("GET")
	router.HandleFunc("/api/v1/charger/", db.ChargerView).Methods("GET")
	router.HandleFunc("/api/v1/customercharger/", db.CustomerChargerView).Methods("GET")
	router.HandleFunc("/api/v1/contractcharger/", db.ContractChargerView).Methods("GET")
	router.HandleFunc("/api/v1/chargermanufacture/", db.ChargerManufactureView).Methods("GET")
	router.HandleFunc("/api/v1/contractcharger/delete/", db.DeleteContractCharger)
	router.HandleFunc("/api/v1/charger/post/", db.CreateCharger)
	router.HandleFunc("/api/v1/charger/put/", db.UpdateCharger)
	router.HandleFunc("/api/v1/charger/delete/",db.DeleteCharger)

	router.HandleFunc("/api/v1/errors/",db.ErrorsView).Methods("GET")
	router.HandleFunc("/api/v1/error/",db.ErrorView).Methods("GET")
	router.HandleFunc("/api/v1/errorstates/",db.ErrorStatesView).Methods("GET")
	router.HandleFunc("/api/v1/errorstate/",db.ErrorStateView).Methods("GET")

	router.HandleFunc("/api/v1/batterytypes/", db.BatteryTypesView).Methods("GET")
	router.HandleFunc("/api/v1/chargertypes/", db.ChargerTypesView).Methods("GET")
	//others
	//router.HandleFunc("/api/v1/unit/", db.UnitView).Methods("GET")
	//router.HandleFunc("/.well-known/acme-challenge/[token]", challengetoken)//encryptの証明
	
	/*
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", OpenHtml.MainHandler)
	router.Handle("/", http.FileServer(http.Dir("../../front/build")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../../front/build/static"))))
	router.Handle("/api/v1/units/", jwtMiddleware.Handler(http.HandlerFunc(db.UnitsView))).Methods("GET")
	router.Handle("/api/v1/customers/", jwtMiddleware.Handler(http.HandlerFunc(db.CustomersView))).Methods("GET")
	router.Handle("/api/v1/detaile/", jwtMiddleware.Handler(http.HandlerFunc(db.DetaileView))).Methods("GET")
	router.Handle("/api/v1/customer/", jwtMiddleware.Handler(http.HandlerFunc(db.CustomerView))).Methods("GET")
	/*
		router.HandleFunc("/", OpenHtml.MainHandler)
		router.HandleFunc("/api/v1/units/", db.UnitsView).Methods("GET")
		router.HandleFunc("/api/v1/customers/", db.CustomersView).Methods("GET")
		router.HandleFunc("/api/v1/detaile/", db.DetaileView).Methods("GET")
		router.HandleFunc("/api/v1/customer/", db.CustomerView).Methods("GET")
		//router.HandleFunc("/api/v1/unit/", db.UnitView).Methods("GET")
	*/
	fmt.Println("RoBOT JWT Server Started Port 443")

	//return http.ListenAndServe(fmt.Sprintf(":%d", 80), router)
	//return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "ssl/fullchain.pem", "ssl/server.key", router) //kitao追加 https
	return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "../../ssl/fullchain.pem", "../../ssl/server.key", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*", "http://localhost:3000/"}))(router))

	// http://18.180.144.98:80/
	// https://jugem.live/
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	//resp, err := http.Get("https://robot.jp.auth0.com/.well-known/jwks.json")
	resp, err := http.Get("https://dev-transkit.jp.auth0.com/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
