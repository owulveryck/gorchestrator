/*
Olivier Wulveryck - author of Gorchestrator
Copyright (C) 2015 Olivier Wulveryck

This file is part of the Gorchestrator project and
is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"os"
	"strconv"
	"strings"
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
		n.Name = strings.Replace(n.Name, "-", "_", -1)
		n.Name = strings.Replace(n.Name, ":", "_Method", -1)
		g.AddNode("G", n.Name,
			map[string]string{
				"id":    fmt.Sprintf("\"%v\"", strconv.Itoa(n.ID)),
				"label": fmt.Sprintf("\"%v|%v|%v|%v\"", n.Name, n.Engine, n.Artifact, n.Args),
				"shape": "\"record\"",
			})
		m[n.ID] = n.Name
	}
	l := v.Digraph.Dim()
	for r := 0; r < l; r++ {
		for c := 0; c < l; c++ {
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
