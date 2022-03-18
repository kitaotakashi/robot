package db

import (
	//"fmt"
	"net/http"
)

//bms 一覧
func GetPacks(w http.ResponseWriter, r *http.Request) {

	db := open_block_db()
	defer db.Close()

	query := "SELECT bms_id, work_state, current12, soc FROM bms"//
	
	results1, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var packs []Pack
	for results1.Next() {//userが取得できるprj_idについてそれぞれ処理
		var pack_tmp Pack
		err := results1.Scan(&pack_tmp.BmsID,&pack_tmp.State,&pack_tmp.Current,&pack_tmp.Energy)//
		if err != nil {
			panic(err.Error())
		}
		pack_tmp.Capacity = 24

		packs = append(packs,pack_tmp)
	}
	
	/*
	dummy1 := Pack{"AAAA-00000001","on",12.5,76.5,370}
	dummy2 := Pack{"AAAA-00000002","off",0,12.5,370}
	var packs_dummy []Pack
	//packs_dummy = [dummy1,dummy2]
	packs_dummy = append(packs_dummy,dummy1)
	packs_dummy = append(packs_dummy,dummy2)
	send(packs_dummy,w)
	*/

	send(packs,w)
}

//bms詳細
func GetPackDetail(w http.ResponseWriter, r *http.Request) {
	q_bms_id := query(r, "bms_id")
	if (len(q_bms_id)==0){
		//send("please specify query parameter : device_id", w)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please specify query parameter : bms_id"))
		return
	}

	db := open_block_db()
	defer db.Close()

	query := "SELECT bms_id, work_state, current12, soc FROM bms WHERE bms_id = ?"
	results1, err := db.Query(query,q_bms_id[0])
	if err != nil {
		panic(err.Error())
	}

	var pack_detail PackDetail
	for results1.Next() {//userが取得できるprj_idについてそれぞれ処理
		err := results1.Scan(&pack_detail.BmsID,&pack_detail.State,&pack_detail.Current,&pack_detail.Energy)
		if err != nil {
			panic(err.Error())
		}
		pack_detail.Capacity = 24

		//bmuを取得
		var bmus []Bmu
		query := "SELECT bmu_id, soc FROM bms WHERE bms_id=?"
		results2, err := db.Query(query,pack_detail.BmsID)
		if err != nil {
			panic(err.Error())
		}
		for results2.Next() {
			var bmu_tmp Bmu
			err := results2.Scan(&bmu_tmp.BmuID,&bmu_tmp.Energy)
			if err != nil {
				panic(err.Error())
			}
			bmus= append(bmus, bmu_tmp)
		}
		pack_detail.BmuData = bmus
	}
	
	/*
	dummy1 := Bmu{"BBBB-00000001",60.4}
	dummy2 := Bmu{"BBBB-00000002",23.4}
	var bmu_dummy []Bmu
	bmu_dummy = append(bmu_dummy,dummy1)
	bmu_dummy = append(bmu_dummy,dummy2)

	pack_detail_dummy := PackDetail{"AAAA-00000001","on",12.5,76.5,370,bmu_dummy}

	send(pack_detail_dummy,w)
	*/
	send(pack_detail,w)
}

//bmu 詳細
func GetModuleDetail(w http.ResponseWriter, r *http.Request) {
	q_bmu_id := query(r, "bmu_id")
	if (len(q_bmu_id)==0){
		//send("please specify query parameter : device_id", w)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please specify query parameter : bmu_id"))
		return
	}

	db := open_block_db()
	defer db.Close()

	query := "SELECT bmu_id, work_state, current12, soc FROM bms WHERE bmu_id = ?"
	results1, err := db.Query(query,q_bmu_id[0])
	if err != nil {
		panic(err.Error())
	}

	var module_detail ModuleDetail
	for results1.Next() {//userが取得できるprj_idについてそれぞれ処理
		err := results1.Scan(&module_detail.BmuID,&module_detail.State,&module_detail.Current,&module_detail.Energy)//
		if err != nil {
			panic(err.Error())
		}
		module_detail.Capacity = 24
	}
	
	/*
	module_dummy := ModuleDetail{"BBBB-00000001","on",24.4,60.4,370}
	send(module_dummy,w)
	*/

	send(module_detail,w)
}
