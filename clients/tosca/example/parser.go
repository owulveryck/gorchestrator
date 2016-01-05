package main

import (
	"fmt"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/toscalib/toscaexec"
	//"gopkg.in/yaml.v2"
	"github.com/awalterschulze/gographviz"
	"log"
	"os"
)

func main() {
	var t toscalib.ServiceTemplateDefinition
	err := t.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	// Creates a new graph
	g := gographviz.NewGraph()
	g.AddAttr("", "rankdir", "LR")
	g.SetName("G")
	g.SetDir(true)
	e := toscaexec.GeneratePlaybook(t)
	for i, p := range e.Index {
		g.AddNode("G", fmt.Sprintf("%v", i),
			map[string]string{
				"id":    fmt.Sprintf("\"%v\"", i),
				"label": fmt.Sprintf("\"%v|%v\"", p.NodeTemplate.Name, p.OperationName),
				"shape": "\"record\"",
			})
	}
	l := e.AdjacencyMatrix.Dim()
	for r := 0; r < l; r++ {
		for c := 0; c < l; c++ {
			if e.AdjacencyMatrix.At(r, c) == 1 {
				g.AddEdge(fmt.Sprintf("%v", r), fmt.Sprintf("%v", c), true, nil)

			}

		}

	}
	log.Println("here")
	s := g.String()
	fmt.Println(s)
	/*
		d, _ := yaml.Marshal(e)
		fmt.Println(string(d))
	*/
	/*
		for i, n := range e.Index {
			log.Printf("[%v] %v:%v -> %v %v",
				i,
				n.NodeTemplate.Name,
				n.OperationName,
				n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Implementation,
				n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Inputs,
			)
		}
	*/
}
