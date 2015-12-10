package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type jsonErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func displayMain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]

	w.WriteHeader(http.StatusOK)
	t := template.New("Main page")
	t, err := t.ParseFiles("tmpl/index.tmpl")
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: "Not Found"}); err != nil {
			panic(err)
		}
	}
	t.Execute(w, id)
}
func displaySvg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	b, err := getSvg(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: "Not Found"}); err != nil {
			panic(err)

		}
	}
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", b)
}
