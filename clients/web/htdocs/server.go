package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Go to http://localhost:9090/starter.html")
	panic(http.ListenAndServe(":9090", http.FileServer(http.Dir("./"))))

}
