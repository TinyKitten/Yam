package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
)

// IndexHandler ルートパスのハンドラ
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: もっとマシにする
	box := packr.NewBox("../template")
	s, err := box.MustString("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t := template.New("index")
	t, err = t.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
