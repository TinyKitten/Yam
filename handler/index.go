package handler

import (
	"html/template"
	"log"
	"net/http"
)

// IndexHandler ルートパスのハンドラ
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: もっとマシにする
	t := template.Must(template.ParseFiles("template/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatal(err)
	}
}
