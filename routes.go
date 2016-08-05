package main

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

var Routes = []Route{
	Route{
		Method:  "get",
		Path:    "/",
		Handler: homeHandler},
	Route{
		Method:  "get",
		Path:    "/{username}",
		Handler: homeHandler},
}
