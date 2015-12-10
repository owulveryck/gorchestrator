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
		var op string
		switch i {
		case t.GetNodeTemplateFromId(i).GetConfigureIndex():
			op = "Configure"
		case t.GetNodeTemplateFromId(i).GetStartIndex():
			op = "Start"
		case t.GetNodeTemplateFromId(i).GetStopIndex():
			op = "Stop"
		case t.GetNodeTemplateFromId(i).GetCreateIndex():
			op = "Create"
		case t.GetNodeTemplateFromId(i).GetDeleteIndex():
			op = "Delete"
		case t.GetNodeTemplateFromId(i).GetInitialIndex():
			op = "Initial"
		case t.GetNodeTemplateFromId(i).GetPostConfigureSourceIndex():
			op = "PostConfigureSource"
		case t.GetNodeTemplateFromId(i).GetPostConfigureTargetIndex():
			op = "PostConfigureTarget"
		case t.GetNodeTemplateFromId(i).GetPreConfigureSourceIndex():
			op = "PreConfigureSource"
		case t.GetNodeTemplateFromId(i).GetPreConfigureTargetIndex():
			op = "PreConfiguretarget"
		}

		v.Nodes[i].Name = fmt.Sprintf("%v:%v", t.GetNodeTemplateFromId(i).Name, op)
		for j := 0; j < s; j++ {
			v.Digraph[s*i+j] = int64(t.AdjacencyMatrix.At(i, j))
		}
	}

	r, _ := json.MarshalIndent(v, "  ", "  ")
	fmt.Print("%s\n", string(r))
}
