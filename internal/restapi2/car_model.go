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

func GetCarModelList(w http.ResponseWriter, r *http.Request) {
	
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

	var car_model_list []carModelListElm
	results1, err := db.Query("SELECT car_model_id,car_model_name,comment FROM car_models ORDER BY car_model_id")
	if err != nil {
		panic(err.Error())
	}	
	for results1.Next() {
		var car_model_elm carModelListElm
		err = results1.Scan(&car_model_elm.CarModelID,&car_model_elm.CarModelName,&car_model_elm.Comment)
		if err != nil {
			panic(err.Error())
		}
		car_model_list = append(car_model_list,car_model_elm)
	}

	if len(car_model_list)==0{
		car_model_list=[]carModelListElm{}
	}
	send(car_model_list, w)
}