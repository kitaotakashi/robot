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
)

func ManageInfoView(w http.ResponseWriter, r *http.Request) {
	//クエリ取得
	//id := query(r, "unit_id")
	q_page := query(r, "page")
	var _q_page int = 1
	if len(q_page)>0{
		_q_page,_ = strconv.Atoi(q_page[0])
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

	db := open()
	defer db.Close()

	var res_data []manageInfoPnt
	var manageInfoParent manageInfoPnt
	var page pageElm
	//ページ数を取得
	results1, err := db.Query("SELECT count(serial_number) FROM manage_info")
	if err != nil {
		panic(err.Error())
	}
	var manage_info_num int
	for results1.Next() {
		err = results1.Scan(&manage_info_num)
		if err != nil {
			panic(err.Error())
		}
	}
	max_page := int(manage_info_num/max_data_num)+1
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

	var manage_infos []manageInfoData
	results1, err = db.Query("SELECT serial_number,unit_id,battery_type,create_at,customer,car_model_id,charger,seller,comment FROM manage_info ORDER BY serial_number LIMIT "+strconv.Itoa(max_data_num)+" OFFSET "+strconv.Itoa(offset))
	if err != nil {
		panic(err.Error())
	}
	for results1.Next() {
		var manage_info manageInfoData

		var car_model_id sql.NullInt32
		var battery_type,customer,charger,seller,comment sql.NullString

		err = results1.Scan(&manage_info.SerialNumber,&manage_info.UnitID,&battery_type,&manage_info.CreateAt,&customer,&car_model_id,&charger,&seller,&comment)
		if err != nil {
			panic(err.Error())
		}
		
		//ust->jst表記に変換
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		manage_info.CreateAt = manage_info.CreateAt.In(jst).Add(-9*time.Hour)

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
			fmt.Println(car_model_id.Int32)
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
			manage_info.Seller = charger.String
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
		//var is_error bool
		//var is_registered bool

		//エラークエリが指定された場合
		if len(q_error)>0{
			fmt.Println(_q_error)
			continue
			/*
			if _q_error != is_error_var{
				continue
			}*/
		}
		//登録ずみクエリが指定された場合
		if len(q_reg)>0{
			fmt.Println(_q_reg)
			continue
			/*
			if _q_reg!= is_reg_var{
				continue
			}*/
		}

		//TODO:カーモデルidが指定された場合
		if len(q_car_model_id)>0{
			fmt.Println(_q_car_model_id)
			continue
		}
		manage_infos = append(manage_infos,manage_info)
	}
	page.DataNum = len(manage_infos)

	manageInfoParent.Page = page
	manageInfoParent.Data=manage_infos
	res_data = append(res_data,manageInfoParent)

	send(res_data, w)
}
