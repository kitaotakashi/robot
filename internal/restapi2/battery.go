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

	//env読み込み
	err := godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
    }
	max_data_num,_ := strconv.Atoi(os.Getenv("MAX_DATA_NUM"))

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
	results1, err = db.Query("SELECT * FROM units ORDER BY unit_id LIMIT "+strconv.Itoa(max_data_num)+" OFFSET "+strconv.Itoa(offset))
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
		batteries = append(batteries,battery)
	}
	page.DataNum = len(batteries)

	batteryParent.Page = page
	batteryParent.Data=batteries
	res_data = append(res_data,batteryParent)

	send(res_data, w)
}