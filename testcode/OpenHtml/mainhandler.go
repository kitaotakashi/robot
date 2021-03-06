package OpenHtml

import (
	"net/http"      //サーバを立てるために必要
	"text/template" //htmlファイルを展開するために必要
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resttest.html") //htmlからテンプレートを作成
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil) //テンプレートを実行(ブラウザに表示)
	if err != nil {
		panic(err)
	}
}
