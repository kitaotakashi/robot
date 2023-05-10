package mico2

import (
	"fmt"
	"net/url"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	//リクエストボディからusernameとpasswordを取得
	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
    	panic(err.Error())
  	}
	keyVal := make(map[string]string)
  	json.Unmarshal(body, &keyVal)

	username := keyVal["username"]
	password := keyVal["password"]

	//fmt.Println(username,password)

	//token取得用のauth0 domainおよびclient_idをenvファイルから取得
	err = godotenv.Load("../../.env")
	//もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
		return
	}
	auth_domain := os.Getenv("AUTH_DOMAIN")
	client_id := os.Getenv("AUTH_CLIENT_ID")
	//client_secret := os.Getenv("AUTH_CLIENT_SECRET")

	values := url.Values{}

	values.Set("grant_type", "password")
    values.Add("client_id", client_id)
	//values.Add("client_secret", client_secret)
	values.Add("username", username)
	values.Add("password", password)
	//values.Add("connection", "Robot-User-DB")
	//values.Add("scope","offline_access")
	//values.Add("audience","https://rev-fl988hpj.us.auth0.com/api/v2/")

	//auth0へget tokenリクエスト
	req, err := http.NewRequest(
        "POST",
        auth_domain+"oauth/token",//url
        strings.NewReader(values.Encode()),
    )
    if err != nil {
		fmt.Println("Request Failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
    }

	// Content-Type 設定
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		fmt.Println("Request Failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
        return
    }
    defer resp.Body.Close()

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	// Log the request body 
	
	fmt.Println(body)
	bodyString := string(body)
	fmt.Println(bodyString)
	

	//get result
	token := Token{}
  	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("Reading body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Reading body failed"))
		return
	}
	if token.Token=="" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong username or password"))
		return
	}else{
		send(token, w)
	}
}