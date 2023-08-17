package mico2

import (
	//"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"log"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"encoding/json"
	"strings"
)

func ManageInfoView(w http.ResponseWriter, r *http.Request) {
	//クエリ取得
	//id := query(r, "unit_id")
	q_page := query(r, "page")
	var _q_page int = 1
	if len(q_page)>0{
		_q_page,_ = strconv.Atoi(q_page[0])
	}

	q_serial_number := query(r, "serial_number")
	var _q_serial_number string
	if len(q_serial_number)>0{
		if CheckInt(q_serial_number[0]){
			_q_serial_number = q_serial_number[0]
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing Parameters or Incorrect Format:serial_number(int)"))
			return
		}	
	}

	q_error := query(r, "error")
	var _q_error int = 0//0:エラーなし、1:エラーあり
	if len(q_error)>0 {
		tmp,_ := strconv.Atoi(q_error[0])
		if tmp==1{
			_q_error = 1
		}
	}
	q_reg := query(r, "registered")
	var _q_reg int = 0//0:登録なし、1:登録あり
	if len(q_reg)>0 {
		tmp,_ := strconv.Atoi(q_reg[0])
		if tmp==1{
			_q_reg = 1
		}
	}

	q_car_model_id := query(r, "car_model_id")
	var _q_car_model_id int
	if len(q_car_model_id)>0{
		_q_car_model_id,_ = strconv.Atoi(q_car_model_id[0]) 
	}

	//env読み込み
	err := godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
    }
	max_data_num,_ := strconv.Atoi(os.Getenv("MAX_DATA_NUM"))
	manage_info_table := os.Getenv("MANAGE_INFO_TABLE")
	car_model_table := os.Getenv("CAR_MODEL_TABLE")
	battery_table := os.Getenv("BATTERY_TABLE")

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

	var res_data []manageInfoPnt
	var manageInfoParent manageInfoPnt
	var page pageElm
	
	//ページ数を取得
	results1, err := db.Query("SELECT unit_id FROM manage_info")
	if err != nil {
		panic(err.Error())
	}
	var manage_info_num int = 0
	var units_list []int
	for results1.Next() {
		var unit_tmp sql.NullInt64
		err = results1.Scan(&unit_tmp)
		if err != nil {
			panic(err.Error())
		}
		if unit_tmp.Valid{
			units_list = append(units_list,int(unit_tmp.Int64))
		}
		manage_info_num += 1
	}

	results1, err = db.Query("SELECT unit_id FROM "+battery_table)
	if err != nil {
		panic(err.Error())
	}
	units_num := 0
	for results1.Next() {
		var unit_id_tmp int
		err = results1.Scan(&unit_id_tmp)
		if err != nil {
			panic(err.Error())
		}
		if contains(units_list,unit_id_tmp){
			continue
		}
		units_num += 1
	}

	max_page := int((manage_info_num+units_num)/max_data_num)+1
	if _q_page > max_page{
		_q_page = max_page
	}
	if _q_page <= 0{
		_q_page = 1
	}
	//sql用にデータの開始と終了を取得
	offset := (_q_page-1)*max_data_num
	//fmt.Println(_q_page,max_page)

	//test
	page.PageNow = _q_page
	page.PageMax = max_page
	if len(q_serial_number)>0{
		page.PageNow = 1
		page.PageMax = 1
	}

	var manage_infos []manageInfoData
	var unit_id_list []int

	results1, err = db.Query("SELECT serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment FROM "+manage_info_table+" ORDER BY serial_number LIMIT "+strconv.Itoa(max_data_num)+" OFFSET "+strconv.Itoa(offset))
	//TODO:カーモデルidが指定された場合
	if len(q_car_model_id)>0{
		results1, err = db.Query("SELECT serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment FROM "+manage_info_table+" WHERE car_model_id = "+strconv.Itoa(_q_car_model_id)+" ORDER BY serial_number LIMIT "+strconv.Itoa(max_data_num)+" OFFSET "+strconv.Itoa(offset))
	}else if len(q_serial_number)>0{
		results1, err = db.Query("SELECT serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment FROM "+manage_info_table+" WHERE serial_number = "+_q_serial_number+" ORDER BY serial_number LIMIT "+strconv.Itoa(max_data_num)+" OFFSET "+strconv.Itoa(offset))
	}
	if err != nil {
		panic(err.Error())
	}
	for results1.Next() {
		var manage_info manageInfoData

		var car_model_id sql.NullInt32
		var unit_id sql.NullInt64
		var battery_type,customer,charger,seller,comment sql.NullString

		err = results1.Scan(&manage_info.SerialNumber,&unit_id,&battery_type,&manage_info.CreateAt,&customer,&car_model_id,&charger,&seller,&comment)
		if err != nil {
			panic(err.Error())
		}
		
		//ust->jst表記に変換
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		manage_info.CreateAt = manage_info.CreateAt.In(jst).Add(-9*time.Hour)

		//unit id
		if unit_id.Valid {
			manage_info.UnitID = strconv.FormatInt(unit_id.Int64, 10)
		}else{
			manage_info.UnitID = ""
		}

		//バッテリータイプ
		if battery_type.Valid {
			manage_info.BatteryType = battery_type.String
		}else{
			manage_info.BatteryType = ""
		}

		//顧客
		if customer.Valid {
			manage_info.Customer = customer.String
		}else{
			manage_info.Customer = ""
		}

		//car model id
		if car_model_id.Valid {
			//TODO:car_model listから名前を取得
			query2 := "SELECT car_model_name FROM "+car_model_table+" WHERE car_model_id = "+strconv.Itoa(int(car_model_id.Int32))
			results2,err := db.Query(query2)
			//results2,err := db.Query(query2,strconv.Itoa(int(car_model_id.Int32)))
			if err != nil {
				panic(err.Error())
			}
			for results2.Next() {
				err = results2.Scan(&manage_info.CarModel)
				if err != nil {
					panic(err.Error())
				}
			}
		}else{
			manage_info.CarModel = ""
		}

		//charger
		if charger.Valid {
			manage_info.Charger = charger.String
		}else{
			manage_info.Charger = ""
		}

		//seller
		if seller.Valid {
			manage_info.Seller = seller.String
		}else{
			manage_info.Seller = ""
		}

		//comment
		if comment.Valid {
			manage_info.Comment = comment.String
		}else{
			manage_info.Comment = ""
		}
		
		//TODO:errorデータやregisterデータを取得しておく
		var is_error bool
		var is_registered bool

		var is_error_cnt int
		if unit_id.Valid{
			is_registered = true
			unit_id_list = append(unit_id_list,int(unit_id.Int64))

			//unit_idからvoltage,current,socを取得
			query2 := "SELECT soc,battery_voltage,battery_current,output_voltage,output_current FROM "+battery_table+" WHERE unit_id = "+strconv.Itoa(int(unit_id.Int64))
			results2,err := db.Query(query2)
			if err != nil {
				panic(err.Error())
			}
			for results2.Next() {
				err = results2.Scan(&manage_info.SoC,&manage_info.Voltage,&manage_info.Current,&manage_info.OutputVoltage,&manage_info.OutputCurrent)
				if err != nil {
					panic(err.Error())
				}
			}

			query2 = "SELECT COUNT(error_code) FROM error_states WHERE object_id = "+strconv.Itoa(int(unit_id.Int64))
			results2,err = db.Query(query2)
			if err != nil {
				panic(err.Error())
			}
			for results2.Next() {
				err = results2.Scan(&is_error_cnt)
				if err != nil {
					panic(err.Error())
				}
			}
			//fmt.Println(is_error_cnt)
			if is_error_cnt>0{
				is_error = true
			}
			manage_info.IsError = is_error
			manage_info.State = "登録済み"
		}else{
			manage_info.State = "実機未登録"
		}
		var is_error_var int = 0
		if is_error{
			is_error_var = 1
		}
		is_error = false
		//エラークエリが指定された場合
		if len(q_error)>0{
			if _q_error != is_error_var{
				continue
			}
		}

		var is_reg_var int = 0
		if is_registered{
			is_reg_var = 1
		}
		//登録ずみクエリが指定された場合
		if len(q_reg)>0{
			if _q_reg!= is_reg_var{
				continue
			}
		}

		manage_infos = append(manage_infos,manage_info)
	}

	//情報登録されていないunitを取得
	offset = (_q_page-1)*max_data_num - manage_info_num
	rest_unit_max_num := max_data_num - offset
	is_get_unit := false

	//ページ数でoffsetを調整
	if offset < 0{
		if offset > -1*max_data_num{
			offset = 0
			is_get_unit = true
		}
	}else{
		is_get_unit = true
	}
	if len(q_car_model_id)>0{
		is_get_unit = false
	}
	if len(q_serial_number)>0{
		is_get_unit = false
	}

	if is_get_unit{
		var is_error bool
		var is_error_cnt int
		
		results1, err = db.Query("SELECT unit_id,soc,output_voltage,output_current FROM "+battery_table+" ORDER BY unit_id LIMIT "+strconv.Itoa(rest_unit_max_num)+" OFFSET "+strconv.Itoa(offset))
		if err != nil {
			panic(err.Error())
		}
		for results1.Next() {
			var manage_info manageInfoData
			var unit_id_tmp int64
			err = results1.Scan(&unit_id_tmp,&manage_info.SoC,&manage_info.Voltage,&manage_info.Current)
			if err != nil {
				panic(err.Error())
			}
			//fmt.Println(contains(unit_id_list,unit_id_tmp))
			if contains(unit_id_list,int(unit_id_tmp)){
				continue
			}else{
				manage_info.UnitID=strconv.FormatInt(unit_id_tmp,10)
				manage_info.State = "情報未登録"
				
				query2 := "SELECT COUNT(error_code) FROM error_states WHERE object_id = "+strconv.FormatInt(unit_id_tmp,10)
				results2,err := db.Query(query2)
				if err != nil {
					panic(err.Error())
				}
				for results2.Next() {
					err = results2.Scan(&is_error_cnt)
					if err != nil {
						panic(err.Error())
					}
				}
				if is_error_cnt>0{
					is_error = true
				}
				var is_error_var int = 0
				if is_error{
					is_error_var = 1
				}
				//エラークエリが指定された場合
				if len(q_error)>0{
					if _q_error != is_error_var{
						continue
					}
				}
				manage_info.IsError = is_error

				//登録ずみクエリが指定された場合
				if len(q_reg)>0{
					if _q_reg==1{
						continue
					}
				}

				manage_infos = append(manage_infos,manage_info)
			}
		}
	}

	page.DataNum = len(manage_infos)

	manageInfoParent.Page = page
	manageInfoParent.Data=manage_infos
	res_data = append(res_data,manageInfoParent)

	send(res_data, w)
}

func AddManageInfo(w http.ResponseWriter, r *http.Request) {
	//ヘッダからAuthorizationを取得する
	h := r.Header["Authorization"]

	//tokenをdecode
	tokenString := h[0][7:]//Baerer以下を取り出し

	token, err := jwt.Parse(tokenString, nil)
    if token == nil {
        panic(err.Error())
    }
    claims, _ := token.Claims.(jwt.MapClaims)

	//user mali取得
	//user_mail := claims["https://classmethod.jp/email"]
	//roll取得
	user_role := claims["https://classmethod.jp/roles"].([]interface{})[0]

	if user_role != "admin"{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You Dont have access right"))
		return
	}

	db := open()
	defer db.Close()

	//env読み込み
	err = godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
    }
	manage_info_table := os.Getenv("MANAGE_INFO_TABLE")

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	var serial_number,unit_id,battery_type,customer,car_model_id,charger,seller,comment string

	//serial_numberがあるかどうか
	if len(keyVal["serial_number"]) > 0{
		if CheckInt(keyVal["serial_number"]){
			serial_number = keyVal["serial_number"]
			//シリアルナンバーが既に使われていないか
			results1, err := db.Query("SELECT count(serial_number) FROM "+manage_info_table+" WHERE serial_number = "+serial_number)
			if err != nil {
				panic(err.Error())
			}
			var cnt int
			for results1.Next() {
				err = results1.Scan(&cnt)
				if err != nil {
					panic(err.Error())
				}
			}
			if cnt>0{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("This serial number has already exists"))
				return
			}
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please Specify Parameters:serial_number with CORRECT type(int)"))
			return
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing Parameters:serial_number"))
		return
	}

	if len(keyVal["unit_id"]) > 0 {
		if CheckInt(keyVal["unit_id"]){
			unit_id = keyVal["unit_id"]

			//unit_idが別のmanage_infoに登録されていないかchk
			results1, err := db.Query("SELECT serial_number FROM "+manage_info_table+" WHERE unit_id = "+unit_id)
			if err != nil {
				panic(err.Error())
			}
			var serial_number_already []int
			for results1.Next() {
				var tmp int
				err = results1.Scan(&tmp)
				if err != nil {
					panic(err.Error())
				}
				serial_number_already = append(serial_number_already,tmp)
			}
			if len(serial_number_already)>0{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("This unit_id has already registered (serial number : "+strconv.Itoa(serial_number_already[0])+")"))
				return
			}
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please Specify Parameters:unit_id with CORRECT type(int)"))
			return
		}
	}
	
	if len(keyVal["battery_type"])>0{
		battery_type = keyVal["battery_type"]
	}

	if len(keyVal["customer"])>0{
		customer = keyVal["customer"]

		//cutomer_listに名前を登録
		//同じcustomer_nameが存在するかchk
		//fmt.Println("SELECT count(customer_id) FROM customer_list WHERE customer_name = '"+customer+"'")
		if strings.Contains(customer,";")==false{
			results1, err := db.Query("SELECT count(customer_id) FROM customer_list WHERE customer_name = '"+customer+"'")
			if err != nil {
				panic(err.Error())
			}
			var tmp int = 0
			for results1.Next() {
				err = results1.Scan(&tmp)
				if err != nil {
					panic(err.Error())
				}
			}
			//存在しない場合、登録
			if tmp==0{
				stmtIns, err := db.Prepare(fmt.Sprintf("INSERT INTO %s (customer_name) VALUES (?)", "customer_list"))
				if err != nil {
					panic(err.Error())
				}
				defer stmtIns.Close()
				_, err = stmtIns.Exec(customer)
			}
		}
	}

	if len(keyVal["car_model_id"])>0{
		if CheckInt(keyVal["car_model_id"]){
			car_model_id = keyVal["car_model_id"]
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please Specify Parameters:car_model_id with CORRECT type(int)"))
			return
		}
	}

	if len(keyVal["charger"])>0{
		charger = keyVal["charger"]
	}

	if len(keyVal["seller"])>0{
		seller = keyVal["seller"]
	}

	if len(keyVal["comment"])>0{
		comment = keyVal["comment"]
	}

	//mage_infoの作成時はcreate_atをサーバー側で追加
	var create_at string = TransTimestampToString(time.Now().In(jst))

	//query = "INSERT INTO "+manage_info_table+" (serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment) VALUES ()
	if len(keyVal["unit_id"]) == 0{
		stmtIns, err := db.Prepare(fmt.Sprintf("INSERT INTO %s (serial_number,battery_type,create_at,customer,car_model_id,charger,seller,comment) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", manage_info_table))
		if err != nil {
			panic(err.Error())
		}
		defer stmtIns.Close()
		_, err = stmtIns.Exec(serial_number,battery_type,create_at,customer,car_model_id,charger,seller,comment)
	}else{
		stmtIns, err := db.Prepare(fmt.Sprintf("INSERT INTO %s (serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", manage_info_table))
		if err != nil {
		panic(err.Error())
		}
		defer stmtIns.Close()
		_, err = stmtIns.Exec(serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment)
	}

	send("add manage information",w)
}

func EditManageInfo(w http.ResponseWriter, r *http.Request) {
	q_serial_number := query(r, "serial_number")
	var _q_serial_number string
	if len(q_serial_number)>0 && CheckInt(q_serial_number[0]){
		_q_serial_number = q_serial_number[0]
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing Parameters or Incorrect Format:serial_number(int)"))
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

	//user mali取得
	//user_mail := claims["https://classmethod.jp/email"]
	//roll取得
	user_role := claims["https://classmethod.jp/roles"].([]interface{})[0]

	if user_role != "admin"{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You Dont have access right"))
		return
	}

	db := open()
	defer db.Close()

	//env読み込み
	err = godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
    }
	manage_info_table := os.Getenv("MANAGE_INFO_TABLE")

	//serial numberが存在するかchkする
	results1, err := db.Query("SELECT count(serial_number) FROM "+manage_info_table+" WHERE serial_number = "+_q_serial_number)
	if err != nil {
		panic(err.Error())
	}
	var cnt int
	for results1.Next() {
		err = results1.Scan(&cnt)
		if err != nil {
			panic(err.Error())
		}
	}
	if cnt==0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There is no manage_info with such serial_number"))
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	var serial_number,unit_id,battery_type,customer,car_model_id,charger,seller,comment string

	if len(keyVal["serial_number"])>0{
		if CheckInt(keyVal["serial_number"]){
			serial_number = keyVal["serial_number"]
			//シリアルナンバーが既に使われていないか
			results1, err := db.Query("SELECT count(serial_number) FROM "+manage_info_table+" WHERE serial_number = "+serial_number)
			if err != nil {
				panic(err.Error())
			}
			var cnt int
			for results1.Next() {
				err = results1.Scan(&cnt)
				if err != nil {
					panic(err.Error())
				}
			}
			if cnt>0{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("This serial number has already exists"))
				return
			}
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please Specify Parameters:serial_number with CORRECT type(int)"))
			return
		}
	}else{
		serial_number = _q_serial_number
	}

	if len(keyVal["unit_id"]) > 0 {
		if CheckInt(keyVal["unit_id"]){
			unit_id = keyVal["unit_id"]

			//unit_idが別のmanage_infoに登録されていないかchk
			results1, err := db.Query("SELECT serial_number FROM "+manage_info_table+" WHERE unit_id = "+unit_id)
			if err != nil {
				panic(err.Error())
			}
			var serial_number_already []int
			for results1.Next() {
				var tmp int
				err = results1.Scan(&tmp)
				if err != nil {
					panic(err.Error())
				}
				if _q_serial_number != strconv.Itoa(tmp){
					serial_number_already = append(serial_number_already,tmp)
				}
			}
			if len(serial_number_already)>0{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("This unit_id has already registered (serial number : "+strconv.Itoa(serial_number_already[0])+")"))
				return
			}
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please Specify Parameters:unit_id with CORRECT type(int)"))
			return
		}
	}
	
	if len(keyVal["battery_type"])>0{
		battery_type = keyVal["battery_type"]
	}

	if len(keyVal["customer"])>0{
		customer = keyVal["customer"]
	}

	if len(keyVal["car_model_id"])>0{
		if CheckInt(keyVal["car_model_id"]){
			car_model_id = keyVal["car_model_id"]
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please Specify Parameters:car_model_id with CORRECT type(int)"))
			return
		}
	}

	if len(keyVal["charger"])>0{
		charger = keyVal["charger"]
	}

	if len(keyVal["seller"])>0{
		seller = keyVal["seller"]
	}

	if len(keyVal["comment"])>0{
		comment = keyVal["comment"]
	}

	//mage_infoの作成時はcreate_atをサーバー側で追加
	var create_at string = TransTimestampToString(time.Now().In(jst))

	//query = "INSERT INTO "+manage_info_table+" (serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment) VALUES ()
	if len(keyVal["unit_id"]) == 0{
		stmtIns, err := db.Prepare(fmt.Sprintf("UPDATE %s SET unit_id = NULL, serial_number = ?, battery_type = ?,create_at = ?,customer = ?,car_model_id = ?,charger = ?,seller = ?,comment = ? WHERE (serial_number = ?)", manage_info_table))
		if err != nil {
			panic(err.Error())
		}
		defer stmtIns.Close()
		_, err = stmtIns.Exec(serial_number, battery_type,create_at,customer,car_model_id,charger,seller,comment,_q_serial_number)
	}else{
		stmtIns, err := db.Prepare(fmt.Sprintf("UPDATE %s SET unit_id = ?, serial_number = ?, battery_type = ?,create_at = ?,customer = ?,car_model_id = ?,charger = ?,seller = ?,comment = ? WHERE (serial_number = ?)", manage_info_table))
		if err != nil {
		panic(err.Error())
		}
		defer stmtIns.Close()
		_, err = stmtIns.Exec(unit_id,serial_number, battery_type,create_at,customer,car_model_id,charger,seller,comment,_q_serial_number)
	}

	send("edit manage information:serial number="+_q_serial_number,w)
}

func DeleteManageInfo(w http.ResponseWriter, r *http.Request) {
	q_serial_number := query(r, "serial_number")
	var _q_serial_number string
	if len(q_serial_number)>0 && CheckInt(q_serial_number[0]){
		_q_serial_number = q_serial_number[0]
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing Parameters or Incorrect Format:serial_number(int)"))
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

	//user mali取得
	//user_mail := claims["https://classmethod.jp/email"]
	//roll取得
	user_role := claims["https://classmethod.jp/roles"].([]interface{})[0]

	if user_role != "admin"{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You Dont have access right"))
		return
	}

	db := open()
	defer db.Close()

	//env読み込み
	err = godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
    }
	manage_info_table := os.Getenv("MANAGE_INFO_TABLE")

	//serial numberが存在するかchkする
	results1, err := db.Query("SELECT count(serial_number) FROM "+manage_info_table+" WHERE serial_number = "+_q_serial_number)
	if err != nil {
		panic(err.Error())
	}
	var cnt int
	for results1.Next() {
		err = results1.Scan(&cnt)
		if err != nil {
			panic(err.Error())
		}
	}
	if cnt==0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There is no manage_info with such serial_number"))
		return
	}

	_, err = db.Query("DELETE FROM "+manage_info_table+" WHERE serial_number = "+_q_serial_number)
	if err != nil {
		panic(err.Error())
	}

	send("delete manage information:serial number="+_q_serial_number,w)
}