package main

import (
	"encoding/json"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"os"
	"strconv"
)

func main() {
	// Creates a new graph
	g := gographviz.NewGraph()
	g.AddAttr("", "rankdir", "LR")
	// Now read the json input
	var v orchestrator.Graph
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&v); err != nil {
		panic(err)

	}

	// Now for each node, create a node
	g.SetName(v.Name)
	g.SetDir(true)
	m := make(map[int]string)
	// Now add every node
	for _, n := range v.Nodes {
		g.AddNode("G", n.Name,
			map[string]string{
				"id":    fmt.Sprintf("\"%v\"", strconv.Itoa(n.ID)),
				"label": fmt.Sprintf("\"%v|%v|%v\"", n.Name, n.Engine, n.Artifact),
				"shape": "\"record\"",
			})
		m[n.ID] = n.Name
	}
	for r := 0; r < v.Digraph.Dim(); r++ {
		for c := 0; c < v.Digraph.Dim(); c++ {
			if v.Digraph.At(r, c) == 1 {
				g.AddEdge(m[r], m[c], true, nil)
			}
		}
	}
	// Now add the edges
	//g.AddEdge("Hello", "World", true, nil)
	s := g.String()
	fmt.Println(s)
}
