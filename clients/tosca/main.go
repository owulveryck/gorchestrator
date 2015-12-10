package main

import (
	"encoding/json"
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
		v.Nodes[i].ID = i
		// FIXME
		n := t.GetNodeTemplateFromId(i)
		var op string
		switch i {
		case n.GetConfigureIndex():
			op = "Configure"
		case n.GetStartIndex():
			op = "Start"
		case n.GetStopIndex():
			op = "Stop"
		case n.GetCreateIndex():
			op = "Create"
		case n.GetDeleteIndex():
			op = "Delete"
		case n.GetInitialIndex():
			op = "Initial"
		case n.GetPostConfigureSourceIndex():
			op = "PostConfigureSource"
		case n.GetPostConfigureTargetIndex():
			op = "PostConfigureTarget"
		case n.GetPreConfigureSourceIndex():
			op = "PreConfigureSource"
		case n.GetPreConfigureTargetIndex():
			op = "PreConfiguretarget"
		}

		v.Nodes[i].Name = fmt.Sprintf("%v:%v", n.Name, op)
		for j := 0; j < s; j++ {
			v.Digraph[s*i+j] = int64(t.AdjacencyMatrix.At(i, j))
		}
	}

	r, _ := json.MarshalIndent(v, "  ", "  ")
	fmt.Print("%s\n", string(r))
}
