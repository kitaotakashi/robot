package mico2

import (
	//"fmt"
	"net/http"
	"strconv"
	//"time"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"log"
)

// UnitsView はunitページに必要なデータをDBから取得し、JSONで返す
func BatteriesView(w http.ResponseWriter, r *http.Request) {
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

	//env読み込み
	err := godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
    }
	max_data_num,_ := strconv.Atoi(os.Getenv("MAX_DATA_NUM"))
	manage_info_table := os.Getenv("MANAGE_INFO_TABLE")
	battery_table := os.Getenv("BATTERY_TABLE")

	db := open()
	defer db.Close()

	var res_data []batteryPnt
	var batteryParent batteryPnt
	var page pageElm
	//ページ数を取得
	results1, err := db.Query("SELECT count(unit_id) FROM units")
	if err != nil {
		panic(err.Error())
	}
	var battery_num int
	for results1.Next() {
		err = results1.Scan(&battery_num)
		if err != nil {
			panic(err.Error())
		}
	}
	max_page := int(battery_num/max_data_num)+1
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

	var batteries []batteryData
	results1, err = db.Query("SELECT * FROM "+battery_table+" ORDER BY unit_id LIMIT "+strconv.Itoa(max_data_num)+" OFFSET "+strconv.Itoa(offset))
	if err != nil {
		panic(err.Error())
	}
	for results1.Next() {
		var battery batteryData
		var unit	unitData
		Columns := columns(&unit)
		err = results1.Scan(Columns...)
		if err != nil {
			panic(err.Error())
		}
		battery.Data = unit
		
		//TODO:errorデータやregisterデータを取得しておく
		var is_error bool
		var is_registered bool

		//test
		is_error = false
		is_registered = false
		results2, err := db.Query("SELECT count(serial_number) FROM "+manage_info_table+" WHERE unit_id = "+unit.UnitID)
		if err != nil {
			panic(err.Error())
		}	
		for results2.Next() {
			var cnt int
			err = results2.Scan(&cnt)
			if err != nil {
				panic(err.Error())
			}
			if cnt>0{
				is_registered = true
			}

		}

		var is_error_var int = 0
		if is_error{
			is_error_var = 1
		}
		var is_reg_var int = 0
		if is_registered{
			is_reg_var = 1
		}

		battery.Management.IsError = is_error
		battery.Management.IsRegistered = is_registered

		//エラークエリが指定された場合
		if len(q_error)>0{
			if _q_error != is_error_var{
				continue
			}
		}
		//登録ずみクエリが指定された場合
		if len(q_reg)>0{
			if _q_reg!= is_reg_var{
				continue
			}
		}
		batteries = append(batteries,battery)
	}
	page.DataNum = len(batteries)

	batteryParent.Page = page
	batteryParent.Data=batteries
	res_data = append(res_data,batteryParent)

	send(res_data, w)
}