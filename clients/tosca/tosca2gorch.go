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
	"github.com/owulveryck/gorchestrator/orchestrator"
	"github.com/owulveryck/gorchestrator/structure"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/toscalib/toscaexec"
	"log"
	"os"
)

func togorch(t toscalib.ServiceTemplateDefinition) orchestrator.Graph {
	e := toscaexec.GeneratePlaybook(t)
	var g orchestrator.Graph
	g.Digraph = structure.Matrix(e.AdjacencyMatrix)
	g.Name = t.Description
	for i, n := range e.Index {
		var node orchestrator.Node
		node.ID = i
		node.Name = fmt.Sprintf("%v:%v", n.NodeTemplate.Name, n.OperationName)
		node.Engine = "shell"
		node.Artifact = n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Implementation
		//node.Args = n.NodeTemplate.Interfaces[n.InterfaceName].Inputs
		g.Nodes = append(g.Nodes, node)
	}
	return g
}

func main() {
	var t toscalib.ServiceTemplateDefinition
	var v orchestrator.Graph
	args := os.Args[1:]
	switch len(args) {
	case 0:
		err := t.Parse(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	case 1:
		r, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = t.Parse(r)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Too many arguments")
	}
	v = togorch(t)
	r, _ := json.MarshalIndent(v, "  ", "  ")
	fmt.Printf("%s\n", string(r))
}
