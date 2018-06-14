package main

import (
	"net/http"

	"github.com/TinyKitten/Yam/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/post", handler.PostHander)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8080", r)
}
