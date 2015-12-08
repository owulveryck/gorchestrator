package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TaskShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	if v, ok := tasks[id]; ok {
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
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
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
	}

	uuid := uuid()
	go v.Run(nil)
	tasks[uuid.ID] = v

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(uuid); err != nil {
		panic(err)
	}
	return
}
