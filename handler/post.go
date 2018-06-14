package handler

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
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
			dat := ErrorPageData{
				Error: "入力が不足しています。",
			}
			t := template.Must(template.ParseFiles("template/error.html"))
			if err := t.ExecuteTemplate(w, "error.html", dat); err != nil {
				log.Fatal(err)
			}
			return
		}

		unescapedText, err := url.QueryUnescape(textQuery)
		if err != nil {
			dat := ErrorPageData{
				Error: "アンエスケープできませんでした。",
			}
			t := template.Must(template.ParseFiles("template/error.html"))
			if err := t.ExecuteTemplate(w, "error.html", dat); err != nil {
				log.Fatal(err)
			}
			return
		}
		dat := TootPageData{
			Text: unescapedText,
		}
		t := template.Must(template.ParseFiles("template/toot.html"))
		if err := t.ExecuteTemplate(w, "toot.html", dat); err != nil {
			log.Fatal(err)
		}
	}

	if r.Method == "POST" {
		instance := r.Form.Get("instance")
		body := r.Form.Get("body")

		if instance == "" || body == "" {
			dat := ErrorPageData{
				Error: "入力が不足しています。",
			}
			t := template.Must(template.ParseFiles("template/error.html"))
			if err := t.ExecuteTemplate(w, "error.html", dat); err != nil {
				log.Fatal(err)
			}
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
			dat := ErrorPageData{
				Error: "インスタンスがダウンしているか、Mastodonインスタンスではないようです。",
			}
			t := template.Must(template.ParseFiles("template/error.html"))
			if err := t.ExecuteTemplate(w, "error.html", dat); err != nil {
				log.Fatal(err)
			}
			return
		}
		defer resp.Body.Close()

		http.Redirect(w, r, "https://"+instance+"/share?text="+url.QueryEscape(body), http.StatusFound)
	}
}
