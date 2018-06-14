package handler

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gobuffalo/packr"
)

// ErrorPageData エラーページに渡すデータ
type ErrorPageData struct {
	Error string
}

// TootPageData トゥート画面に渡すデータ
type TootPageData struct {
	Text string
}

// PostHander /postのハンドラ
func PostHander(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		textQuery := r.URL.Query().Get("text")

		if textQuery == "" {
			handleError(w, "入力が不足しています。")
			return
		}

		unescapedText, err := url.QueryUnescape(textQuery)
		if err != nil {
			handleError(w, "アンエスケープできませんでした。")
			return
		}
		dat := TootPageData{
			Text: unescapedText,
		}
		box := packr.NewBox("../template")
		s, err := box.MustString("toot.html")
		if err != nil {
			log.Fatal(err)
		}
		t := template.New("error")
		t, err = t.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, dat); err != nil {
			log.Fatal(err)
		}
	}

	if r.Method == "POST" {
		instance := r.Form.Get("instance")
		body := r.Form.Get("body")

		if instance == "" || body == "" {
			handleError(w, "入力が不足しています。")
			return
		}

		if strings.HasPrefix(instance, "http://") {
			instance = strings.Replace(instance, "http://", "", 1)
		}
		if strings.HasPrefix(instance, "https://") {
			instance = strings.Replace(instance, "https://", "", 1)
		}

		baseURL := "https://" + instance
		resp, err := http.Get(baseURL + "/api/v1/instance")
		if err != nil || resp.StatusCode == 404 {
			handleError(w, "インスタンスがダウンしているか、Mastodonインスタンスではないようです。")
			return
		}
		defer resp.Body.Close()

		http.Redirect(w, r, "https://"+instance+"/share?text="+url.QueryEscape(body), http.StatusFound)
	}
}

func handleError(w http.ResponseWriter, msg string) {
	dat := ErrorPageData{
		Error: msg,
	}

	box := packr.NewBox("../template")
	s, err := box.MustString("error.html")
	if err != nil {
		log.Fatal(err)
	}
	t := template.New("error")
	t, err = t.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(w, dat); err != nil {
		log.Fatal(err)
	}
}
