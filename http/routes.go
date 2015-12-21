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
	}, Route{
		"TaskList",
		"GET",
		"/v1/tasks/",
		TaskList,
	}, Route{
		"TaskDelete",
		"DELETE",
		"/v1/tasks/{id}",
		TaskDelete,
	},
}
