package structure

import (
	"encoding/json"
	"fmt"
	"math"
)

// Matrix is a list representation of a squared matrix
type Matrix []int64

// Input is the input of the orchestrator
type Input struct {
	Name    string `json:"name",omitempty`
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

// isValid check if the matrix is squared
func (m *Matrix) isValid() error {
	l := math.Sqrt(float64(len(*m)))
	if float64(int64(l)) != l {
		return fmt.Errorf("Matrix is not a squared one")
	}
	return nil
}

// Dim returns the dimension of the matrix
func (m *Matrix) Dim() int64 {
	err := m.isValid()
	if err != nil {
		return 0
	}
	return int64(math.Sqrt(float64(len(*m))))
}

// Get sets the value v in row r and column c
func (m *Matrix) Set(v, r, c int64) {
	i := m.Dim()
	(*m)[r*i+c] = v
}

// Get returns the value in row r and column c
func (m *Matrix) Get(r, c int64) int64 {
	i := m.Dim()
	return (*m)[r*i+c]
}
