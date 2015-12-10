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
		n := t.GetNodeTemplateFromId(i)
		// Fill in a map with method as key and artifact as value
		interfaces := make(map[string]string, 0)
		for _, intf := range n.Interfaces {
			for val, _ := range intf {
				interfaces[val] = "Found"

			}
		} // FIXME
		var op string
		switch i {
		case n.GetConfigureIndex():
			op = "configure"
		case n.GetStartIndex():
			op = "start"
		case n.GetStopIndex():
			op = "stop"
		case n.GetCreateIndex():
			op = "create"
		case n.GetDeleteIndex():
			op = "delete"
		case n.GetInitialIndex():
			op = "initial"
		case n.GetPostConfigureSourceIndex():
			op = "postConfigureSource"
		case n.GetPostConfigureTargetIndex():
			op = "postConfigureTarget"
		case n.GetPreConfigureSourceIndex():
			op = "preConfigureSource"
		case n.GetPreConfigureTargetIndex():
			op = "preConfiguretarget"
		}
		v.Nodes[i].Artifact = interfaces[op]

		v.Nodes[i].Name = fmt.Sprintf("%v:%v", n.Name, op)
		for j := 0; j < s; j++ {
			v.Digraph[s*i+j] = int64(t.AdjacencyMatrix.At(i, j))
		}
	}

	r, _ := json.MarshalIndent(v, "  ", "  ")
	fmt.Print("%s\n", string(r))
}
