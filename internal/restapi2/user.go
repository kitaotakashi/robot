package mico2

import (
	"fmt"
	//"net/url"
	"net/http"
	//"encoding/json"
	//"io/ioutil"
	"github.com/joho/godotenv"
	//"os"
	//"strings"
	"github.com/dgrijalva/jwt-go"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	
	//token取得用のauth0 domainおよびclient_idをenvファイルから取得
	err := godotenv.Load("../../.env")
	//もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
		return
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
	fmt.Println(claims)

	//user mali取得
	user_mail := claims["https://classmethod.jp/email"]
	//roll取得
	user_role := claims["https://classmethod.jp/roles"].([]interface{})[0]

	var user []userElm
	var user_elm userElm

	if user_mail != nil{
		if user_mail,ok := user_mail.(string); ok{
			user_elm.UserName = user_mail
		}
	}

	if user_role != nil{
		if user_role,ok := user_role.(string); ok{
			user_elm.UserRole = user_role
		}
	}
	
	user = append(user,user_elm)

	send(user, w)
}