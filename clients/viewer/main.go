package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Println("connect here: http://localhost:8181/svg")
	log.Fatal(http.ListenAndServe(":8181", router))
}
