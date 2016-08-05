package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Methods(Routes[0].Method).Path(Routes[0].Path).Handler(Routes[0].Handler)
	r.Methods(Routes[1].Method).Path(Routes[1].Path).Handler(Routes[1].Handler)

	// r.HandleFunc("/{username}", homeHandler)
	http.ListenAndServe(":3000", r)
}
