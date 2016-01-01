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

import (
	"encoding/json"
	"fmt"
	"github.com/owulveryck/gorchestrator/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TaskList(w http.ResponseWriter, r *http.Request) {
	type Content struct {
		Id    string    `json:"id"`
		Name  string    `json:"name"`
		State string    `json:"state"`
		Start time.Time `json:"start_date,omitempty"`
		Stop  time.Time `json:"stop_date,omitempty"`
	}
	var l struct {
		C []Content `json:"tasks"`
	}
	co := make([]Content, len(tasks))

	i := 0
	for id, task := range tasks {
		c := Content{
			id,
			(task).Name,
			orchestrator.States[(task).State],
			time.Time{},
			time.Time{},
		}
		co[i] = c
		i++
	}
	l.C = co
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(l); err != nil {
		panic(err)
	}
}

func TaskShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	if v, ok := tasks[id]; ok {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(*v); err != nil {
			panic(err)

		}
		return
	} else {

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: "Not Found"}); err != nil {
			panic(err)

		}
	}
}

func TaskDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	if v, ok := tasks[id]; ok {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		v.Timeout = time.After(0)
		delete(tasks, id)
		return
	} else {

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: "Not Found"}); err != nil {
			panic(err)
		}
	}
}

/*
Test with this curl command:

*/
func TaskCreate(w http.ResponseWriter, r *http.Request) {
	var v orchestrator.Graph
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &v); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	uuid := uuid()
	exe := orchestrator.ExecutorBackend{
		"https://127.0.0.1:8585/v1",
		"orchestrator.pem",
		"orchestrator_key.pem",
		"executor.pem",
		"/ping",
		nil,
	}
	exe.Init()

	go v.Run(exe)
	v.Timeout = time.After(5 * time.Minute)
	tasks[uuid.ID] = &v

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(uuid); err != nil {
		panic(err)
	}
	return
}
