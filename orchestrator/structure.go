package orchestrator

import (
	"encoding/json"
)

// Input is the input of the orchestrator
type Input struct {
	Name    string `json:"name",omitempty`
	State   int    `json:"state"`
	Digraph Matrix `json:"digraph"`
	Nodes   []Node `json:"nodes"`
}

// Node is a "runable" node description
type Node struct {
	ID       int               `json:"id"`
	Name     string            `json:"name",omitempty`
	Engine   string            `json:"engine",omitempty` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string            `json:"artifact"`
	Args     []string          `json:"args",omitempty`   // the arguments of the artifact, if needed
	Outputs  map[string]string `json:"output",omitempty` // the key is the name of the parameter, the value its value (always a string)
}

type Message struct {
	ID    int
	State int
	Wait  chan Matrix
}

// Error is a type used when any error related to the input or node structure occurs
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Check is the structure is coherent, (a squared matrix with as many nodes as needed)
func (i *Input) Check() Error {
	if len(i.Nodes)*len(i.Nodes) != len(i.Digraph) {
		return Error{1, "Structure is not coherent"}
	}
	return Error{0, ""}
}

func (e *Error) Error() string {
	o, _ := json.Marshal(e)
	return string(o)
}
