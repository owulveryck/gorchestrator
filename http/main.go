package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"log"
	"net/http"
)

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

// This will hold all the requested tasks
var tasks map[string]orchestrator.Input

func main() {

	tasks = make(map[string]orchestrator.Input, 0)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
