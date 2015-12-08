package orchestrator

import (
	"encoding/json"
	"github.com/owulveryck/gorchestrator/structure"
)

type Message struct {
	ID    int
	State int
	Wait  chan structure.Matrix
}

// Error is a type used when any error related to the input or node structure occurs
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	o, _ := json.Marshal(e)
	return string(o)
}
