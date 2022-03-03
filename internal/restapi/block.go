package db

import (
	//"fmt"
	"net/http"
)

//bms 一覧
func GetPacks(w http.ResponseWriter, r *http.Request) {

	//db := open_block_db()
	//defer db.Close()

	//query := "SELECT p1, p2 FROM [table]"//
	/*
	results1, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var packs []Pack
	for results1.Next() {//userが取得できるprj_idについてそれぞれ処理
		var pack_tmp Pack
		err := results1.Scan(&pack_tmp.BmsID,&device_name.State)//
		if err != nil {
			panic(err.Error())
		}
	}
	*/
	dummy1 := Pack{"AAAA-00000001","on",12.5,76.5,370}
	dummy2 := Pack{"AAAA-00000002","off",0,12.5,370}
	var packs_dummy []Pack
	//packs_dummy = [dummy1,dummy2]
	packs_dummy = append(packs_dummy,dummy1)
	packs_dummy = append(packs_dummy,dummy2)

	send(packs_dummy,w)
}

//bms詳細
/*
func GetPackDetail(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータの取得
	q_device_id := query(r, "device_id")
	if (len(q_device_id)==0){
		//send("please specify query parameter : device_id", w)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please specify query parameter : device_id"))
		return
	}

	db := open()
	defer db.Close()

	query := "SELECT device_id, device_name FROM device WHERE device_id = $1"
	results1, err := db.Query(query,q_device_id[0])
	if err != nil {
		panic(err.Error())
	}

	var device_name IoTDeviceElm
	for results1.Next() {//userが取得できるprj_idについてそれぞれ処理
		err := results1.Scan(&device_name.DeviceID,&device_name.DeviceName)
		if err != nil {
			panic(err.Error())
		}
	}

	send(device_name,w)
}
*/