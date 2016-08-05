package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	t, _ := template.ParseFiles("public/index.html")
	vars := mux.Vars(r)
	t.ExecuteTemplate(w, "index.html", vars["username"])
}
