/*
Olivier Wulveryck - author of Gorchestrator
Copyright (C) 2015 Olivier Wulveryck

This file is part of the Gorchestrator project and
is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
			"Json tasks",
			"GET",
			"/tasks/tasks.json",
			getTasks,
		},
		route{
			"Display tasks",
			"GET",
			"/tasks/",
			displayTasks,
		},
		route{
			"JSon graph",
			"GET",
			"/graph/{id}.json",
			displayGraph,
		},
		route{
			"SVG diagram",
			"GET",
			"/graph/{id}.svg",
			displaySvg,
		},
		route{
			"Main page ",
			"GET",
			"/{id}",
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
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	// Define the access to the root of the web
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("htdocs/")))
	router.Headers("Cache-Control", "no-cache, no-store, must-revalidate")

	return router
}
