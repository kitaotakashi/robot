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

func GetErrorList(w http.ResponseWriter, r *http.Request) {
	
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

	var errors []errorsListElm
	results1, err := db.Query("SELECT error_code,error_category,error_message,required_action FROM errors ORDER BY error_code")
	if err != nil {
		panic(err.Error())
	}	
	for results1.Next() {
		var error_elm errorsListElm
		err = results1.Scan(&error_elm.ErrorCode,&error_elm.ErrorCategory,&error_elm.ErrorMessage,&error_elm.RequiredAction)
		if err != nil {
			panic(err.Error())
		}
		errors = append(errors,error_elm)
	}

	send(errors, w)
}