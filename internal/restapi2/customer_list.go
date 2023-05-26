package mico2

import (
	"fmt"
	//"net/url"
	"net/http"
	//"encoding/json"
	//"io/ioutil"
	"github.com/joho/godotenv"
	"os"
	//"strings"
	"github.com/dgrijalva/jwt-go"
	"log"
)

func GetCustomerList(w http.ResponseWriter, r *http.Request) {
	
	//env読み込み
	err := godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
    }

	//ヘッダからAuthorizationを取得する
	h := r.Header["Authorization"]

	//tokenをdecode
	tokenString := h[0][7:]//Baerer以下を取り出し

	token, err := jwt.Parse(tokenString, nil)
    if token == nil {
        panic(err.Error())
    }
    claims, _ := token.Claims.(jwt.MapClaims)

	//roll取得
	user_role := claims["https://classmethod.jp/roles"].([]interface{})[0]
	fmt.Println(user_role)

	db := open()
	defer db.Close()

	var customer_list []string
	results1, err := db.Query("SELECT customer_name FROM customer_list ORDER BY customer_id")
	if err != nil {
		panic(err.Error())
	}	
	for results1.Next() {
		var customer_list_elm string
		err = results1.Scan(&customer_list_elm)
		if err != nil {
			panic(err.Error())
		}
		customer_list = append(customer_list,customer_list_elm)
	}

	if len(customer_list)==0{
		customer_list=[]string{}
	}
	send(customer_list, w)
}