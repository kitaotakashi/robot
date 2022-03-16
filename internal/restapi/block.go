package db

import (
	"fmt"
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
func GetPackDetail(w http.ResponseWriter, r *http.Request) {
	q_bms_id := query(r, "bms_id")
	if (len(q_bms_id)==0){
		//send("please specify query parameter : device_id", w)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please specify query parameter : bms_id"))
		return
	}else{
		fmt.Println(q_bms_id)
	}
	/*
	//クエリパラメータの取得
	q_device_id := query(r, "dms_id")
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
	*/

	dummy1 := Bmu{"BBBB-00000001",60.4}
	dummy2 := Bmu{"BBBB-00000002",23.4}
	var bmu_dummy []Bmu
	bmu_dummy = append(bmu_dummy,dummy1)
	bmu_dummy = append(bmu_dummy,dummy2)

	pack_detail_dummy := PackDetail{"AAAA-00000001","on",12.5,76.5,370,bmu_dummy}

	send(pack_detail_dummy,w)
}

//bmu 詳細
func GetModuleDetail(w http.ResponseWriter, r *http.Request) {
	q_bmu_id := query(r, "bmu_id")
	if (len(q_bmu_id)==0){
		//send("please specify query parameter : device_id", w)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please specify query parameter : bmu_id"))
		return
	}else{
		fmt.Println(q_bmu_id)
	}
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
	module_dummy := ModuleDetail{"BBBB-00000001","on",24.4,60.4,370}

	send(module_dummy,w)
}
