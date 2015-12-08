package http

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TaskCreate",
		"POST",
		"/v1/tasks",
		TaskCreate,
	}, Route{
		"TaskShow",
		"GET",
		"/v1/tasks/{id}",
		TaskShow,
	},
}