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

package orchestrator

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInit(t *testing.T) {
	// OK test
	router := NewRouter()
	// Load client cert
	cert, err := tls.LoadX509KeyPair("./test/executor.pem", "./test/executor_key.pem")
	if err != nil {
		t.Fatal(err)

	}

	caCert, err := ioutil.ReadFile("./test/orchestrator.pem")
	if err != nil {
		t.Fatal(err)

	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	ts := httptest.NewUnstartedServer(router)
	ts.TLS = &tls.Config{
		ClientCAs:    caCertPool,
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	ts.StartTLS()
	defer ts.Close()
	// This test should work
	exe := ExecutorBackend{
		fmt.Sprintf("%v/v1", ts.URL),
		"./test/orchestrator.pem",
		"./test/orchestrator_key.pem",
		"./test/executor.pem",
		"/ping",
		nil,
	}
	err = exe.Init()
	if err != nil {
		t.Fatal(err)
	}
	// This test shoud not work because client has not the correct certificate
	exe = ExecutorBackend{
		fmt.Sprintf("%v/v1", ts.URL),
		"./test/executor.pem",
		"./test/orchestrator_key.pem",
		"./test/executor.pem",
		"/ping",
		nil,
	}
	err = exe.Init()
	if err == nil {
		t.Fatal(err)
	}

	// This test shoud not work because client has not the correct certificate
	exe = ExecutorBackend{
		fmt.Sprintf("%v/v1", ts.URL),
		"./test/executor.pem",
		"./test/orchestrator_key.pem",
		"./test/executor.pem",
		"/pong",
		nil,
	}
	err = exe.Init()
	if err == nil {
		t.Fatal(err)
	}

	exe = ExecutorBackend{
		"https://localhost:8585/v1",
		"./test/cert_nil.pem",
		"./test/orchestrator_key.pem",
		"./test/executor.pem",
		"/ping",
		nil,
	}
	err = exe.Init()
	if err == nil {
		t.Error(err)
	}
	exe = ExecutorBackend{
		"https://localhost:8585/v1",
		"./test/orchestrator.pem",
		"./test/orchestrator_key.pem",
		"./test/cert_nil.pem",
		"/ping",
		nil,
	}
	err = exe.Init()
	if err == nil {
		t.Error(err)
	}
}

type jsonErr struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	type status struct {
		Status string `json:"status"`
	}
	success := status{"success"}
	if err := json.NewEncoder(w).Encode(success); err != nil {
		panic(err)
	}

}

func TaskShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	if v, ok := tasks[id]; ok {
		switch v.Artifact {
		case "success":
			v.State = Success
		case "notfound":
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: "Not Found"}); err != nil {
				panic(err)

			}
		case "failure":
			v.State = Failure
		default:
			v.State = Success
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(v); err != nil {
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

/*
Test with this curl command:

*/
func TaskCreate(w http.ResponseWriter, r *http.Request) {
	var v Node
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
	tasks[uuid.ID] = &v

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(uuid); err != nil {
		panic(err)
	}
	return
}

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
		"Ping",
		"GET",
		"/v1/ping",
		Ping,
	}, Route{
		"TaskShow",
		"GET",
		"/v1/tasks/{id}",
		TaskShow,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var tasks map[string](*Node)

type Task struct {
	ID string `json:"id"`
}

func uuid() Task {
	var t Task
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return t
	}
	u[8] = (u[8] | 0x80) & 0xBF // what does this do?
	u[6] = (u[6] | 0x40) & 0x4F // what does this do?
	t.ID = hex.EncodeToString(u)
	return t
}
