package main

import (
	"fmt"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"github.com/owulveryck/toscalib"
	"os"
)

func main() {
	var t toscalib.ToscaDefinition
	var v orchestrator.Graph
	err := t.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	// Fill the digraph
	r, c := t.AdjacencyMatrix.Dims()
	v.Digraph = make([]int64, r*r+r)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v.Digraph[r*i+j] = int64(t.AdjacencyMatrix.At(i, j))
		}
	}
	fmt.Println(v)

}
