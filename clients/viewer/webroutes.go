package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

// NewRouter is a new router instance.
func NewRouter() *mux.Router {

	// Definition des routes
	var routes = routes{
		route{
			"SVG diagram",
			"GET",
			"/graph/{id}.svg",
			displaySvg,
		},
		route{
			"Main page ",
			"GET",
			"/view/{id}",
			displayMain,
		},
		route{
			"Main page ",
			"GET",
			"/view/",
			displayMain,
		},
	}
	router := mux.NewRouter().StrictSlash(true)
	//router.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	// Define the access to the root of the web
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("htdocs/")))
	router.Headers("Cache-Control", "no-cache, no-store, must-revalidate")

	return router
}
