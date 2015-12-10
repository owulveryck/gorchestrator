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
	s, _ := t.AdjacencyMatrix.Dims()
	v.Digraph = make([]int64, s*s+s)
	v.Nodes = make([]orchestrator.Node, s)
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			v.Digraph[s*i+j] = int64(t.AdjacencyMatrix.At(i, j))
		}
	}
	fmt.Println(v)

}
