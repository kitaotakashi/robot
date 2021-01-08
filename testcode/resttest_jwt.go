//https://code-database.com/knowledges/87 htmlをgoで立てたサーバで展開する方法
//https://qiita.com/TakahiRoyte/items/949f4e88caecb02119aa#:~:text=REST(REpresentational%20State%20Transfer)%E3%81%AF,%E3%81%AE%E9%80%81%E5%8F%97%E4%BF%A1%E3%82%92%E8%A1%8C%E3%81%84%E3%81%BE%E3%81%99%E3%80%82
//↑RESTについて
package main

import (
	"./OpenHtml"
	"./db"
	"encoding/json"
	"errors"
	"fmt" //標準入力など(デバッグ用なので最終的にはいらない...?)
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http" //サーバを立てるために必要
)

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
			aud := "https://jugem.live"
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			// fmt.Printf("%v", token)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}
			// Verify 'iss' claim
			iss := "https://robot.jp.auth0.com/"
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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", OpenHtml.MainHandler)
	router.Handle("/api/v1/units/", jwtMiddleware.Handler(http.HandlerFunc(db.UnitsView))).Methods("GET")
	router.Handle("/api/v1/customers/", jwtMiddleware.Handler(http.HandlerFunc(db.CustomersView))).Methods("GET")
	router.Handle("/api/v1/detaile/", jwtMiddleware.Handler(http.HandlerFunc(db.DetailedView))).Methods("GET")
	router.Handle("/api/v1/customer/", jwtMiddleware.Handler(http.HandlerFunc(db.CustomerView))).Methods("GET")

	/*
		router.HandleFunc("/", OpenHtml.MainHandler)
		router.HandleFunc("/api/v1/units/", db.UnitsView).Methods("GET")
		router.HandleFunc("/api/v1/customers/", db.CustomersView).Methods("GET")
		router.HandleFunc("/api/v1/detaile/", db.DetaileView).Methods("GET")
		router.HandleFunc("/api/v1/customer/", db.CustomerView).Methods("GET")
		//router.HandleFunc("/api/v1/unit/", db.UnitView).Methods("GET")
	*/
	fmt.Println("Server Started Port 443")

	//return http.ListenAndServe(fmt.Sprintf(":%d", 80), router)
	//return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "ssl/fullchain.pem", "ssl/server.key", router) //kitao追加 https
	return http.ListenAndServeTLS(fmt.Sprintf(":%d", 443), "ssl/fullchain.pem", "ssl/server.key", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*", "http://localhost:3000/"}))(router))

	// http://18.180.144.98:80/
	// https://jugem.live/
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://robot.jp.auth0.com/.well-known/jwks.json")

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
