package structure

import (
	"encoding/json"
)

type Input struct {
	Name    string `json:"name",omitempty`
	Digraph []int  `json:"digraph"`
	Nodes   []Node `json:"nodes"`
}

type Node struct {
	ID       int               `json:"id"`
	Name     string            `json:"name",omitempty`
	Engine   string            `json:"engine",omitempty` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string            `json:"artifact"`
	Args     []string          `json:"args",omitempty`   // the arguments of the artifact, if needed
	Outputs  map[string]string `json:"output",omitempty` // the key is the name of the parameter, the value its value (always a string)
}

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
