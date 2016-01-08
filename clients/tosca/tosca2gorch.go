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
	"regexp"
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
		if n.OperationName == "noop" {
			node.Engine = "nil"
		} else {
			node.Engine = "toscassh"
		}
		node.Artifact = n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Implementation
		for argName, argValue := range n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Inputs {
			for get, val := range argValue {
				switch get {
				case "value":
					node.Args = append(node.Args, fmt.Sprintf("%v=%v", argName, val[0]))
				case "get_input":
					value := e.Inputs[val[0]]
					node.Args = append(node.Args, fmt.Sprintf("%v=%v", argName, value))
				case "get_property":
					log.Println("DEBUG:", argValue)
					prop, err := t.GetProperty(val[0], val[1])
					node.Args = append(node.Args, fmt.Sprintf("%v=%v", argName, prop))
					if err != nil {
						log.Println("Cannot find property %v on %v", val[1], val[0])
					}
				case "get_attribute":
					log.Println("DEBUG (get_attribute):", val)
					node.Args = append(node.Args, fmt.Sprintf("%v=get_attribute %v*:%v", argName, val[0], val[1]))
				default:
					node.Args = append(node.Args, fmt.Sprintf("DEBUG: %v=%v", argName, val))

				}
			}
		}
		// Sets the output
		// For every node, the attributes of the node or its type is an output
		node.Outputs = make(map[string]string, 0)
		for k, _ := range n.NodeTemplate.Refs.Type.Attributes {
			node.Outputs[k] = ""
		}
		//for k, v := range n.NodeTemplate.Attributes {
		//node.Outputs[k] = v
		//}
		// Sets the target
		// Find the "host" requirement
		compute := regexp.MustCompile(`[cC]ompute$`)
		var target string
		target = "self"
		targetType := "none"
		curr := n.NodeTemplate.Requirements
		for i := 0; i < len(t.TopologyTemplate.NodeTemplates); i++ {
			for _, req := range curr {
				if name, ok := req["host"]; ok {
					// Get the NodeTemplate "name"
					nt := t.TopologyTemplate.NodeTemplates[name.Node]
					// Get the required node's type
					targetType = nt.Type
					// If targetType is a node compute
					if compute.MatchString(targetType) {
						target = name.Node
						break
					}
					curr = nt.Requirements
				}
				if target != "self" {
					break
				}
			}
		}
		node.Target = target
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
	for _, n := range v.Nodes {
		log.Printf("[%v]\t%v\t%v\t%v", n.Name, n.Artifact, n.Args, n.Outputs)

	}
	r, _ := json.MarshalIndent(v, "  ", "  ")
	fmt.Printf("%s\n", string(r))
}
