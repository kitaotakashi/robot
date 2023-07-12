package OpenHtml

import (
	"net/http"      //サーバを立てるために必要
	"text/template" //htmlファイルを展開するために必要
	//"fmt"
)

var MICO_APP_BASE = "/root/mico_front_2/.next"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../internal/view/resttest.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func BlockHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../front/dist/index.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../front/build/index.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func Mico404Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../front/out/404.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoTestHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../front/out/battery/detail.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoDetailHandler(w http.ResponseWriter, r *http.Request) {
	//tmpl, err := template.ParseFiles("../../front/out/battery/detail.html") //htmlからテンプレートを作成
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/detail.html")
	//tmpl, err := template.ParseFiles("/root/mico_front_2/.next/server/app/battery/detail.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoEditHandler(w http.ResponseWriter, r *http.Request) {
	//tmpl, err := template.ParseFiles("../../front/out/battery/edit.html") //htmlからテンプレートを作成
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/edit.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoAddHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/add.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoChangeUnitHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/changeunit.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

//text handler
func MicoDetailTextHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/detail.txt") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoEditTextHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/edit.txt") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoAddTextHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/add.txt") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}

func MicoChangeUnitTextHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(MICO_APP_BASE+"/server/app/battery/changeunit.txt") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}