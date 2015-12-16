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
	//"github.com/owulveryck/gorchestrator/structure"
	"github.com/owulveryck/toscalib"
	"log"
	"os"
	//"reflect"
)

func togorch(t toscalib.ToscaDefinition) orchestrator.Graph {
	var v orchestrator.Graph
	var exclusion []int
	exclusion = make([]int, 0)
	s, _ := t.AdjacencyMatrix.Dims()
	v.Digraph = make([]int64, s*s)
	v.Nodes = make([]orchestrator.Node, s)
	for i := 0; i < s; i++ {
		v.Nodes[i].ID = i
		n := t.GetNodeTemplateFromId(i)
		// Fill in a map with method as key and artifact as value
		interfaces := make(map[string]string, 0)
		args := make(map[string][]string, 0)
		for _, intf := range n.Interfaces {
			for method, interfaceDefinition := range intf {
				interfaces[method] = interfaceDefinition.Implementation
				// Fill the args
				for key, value := range interfaceDefinition.Inputs {
					for v, vv := range value {
						switch v {
						case "value":
							args[method] = append(args[method], fmt.Sprintf("%v=%v", key, vv))
						case "get_input":
							args[method] = append(args[method], fmt.Sprintf("%v=%v", key, t.TopologyTemplate.Inputs[vv].Default))
						default:
							args[method] = append(args[method], fmt.Sprintf("%v=%v(%v)", key, v, vv))
						}
					}
				}
			}
		}
		// FIXME
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
		if _, ok := interfaces[op]; !ok {
			// Look for the implementation in the type
			for _, intf := range t.NodeTypes[n.Type].Interfaces {
				for method, interfaceDefinition := range intf {
					interfaces[method] = interfaceDefinition.Implementation
					// key holds the input name
					for key, _ := range interfaceDefinition.Inputs {
						if prop, ok := n.Properties[key]; !ok {
							log.Println("No Value found for input", key)
						} else {
							for v, vv := range prop {
								switch v {
								case "value":
									args[method] = append(args[method], fmt.Sprintf("%v=%v", key, vv))
								case "get_input":
									args[method] = append(args[method], fmt.Sprintf("%v=%v", key, t.TopologyTemplate.Inputs[vv].Default))
								default:
									args[method] = append(args[method], fmt.Sprintf("%v=%v(%v)", key, v, vv))
								}
							}
						}
					}
				}
			}
		}
		v.Nodes[i].Artifact = interfaces[op]
		v.Nodes[i].Args = args[op]
		//v.Nodes[i].Args = args[op]

		if op != "" {
			v.Nodes[i].Name = fmt.Sprintf("%v:%v", n.Name, op)
			if interfaces[op] == "nil" || interfaces[op] == "" {
				exclusion = append(exclusion, i)
			}
		} else {
			v.Nodes[i].Name = fmt.Sprintf("%v", n.Name)
		}
		log.Printf("[%v] %v", v.Nodes[i].Name, interfaces[op])
		for j := 0; j < s; j++ {
			v.Digraph[s*i+j] = int64(t.AdjacencyMatrix.At(i, j))
		}
	}
	// now shrink the matrix... (remove the exclusion)
	/*
		log.Println("Exclusions", exclusion)
		for _, i := range exclusion {

			l := v.Digraph.Dim()
			//Remove the i-th line
			//= copy all the values if pos-i%l!=0
			target := make(structure.Matrix, 0)
			target2 := make(structure.Matrix, 0)
			var value int64
			for pos := 0; pos < l; pos++ {
				value = v.Digraph[pos]
				if (pos-i)%l != 0 {
					log.Printf("Pos: %v, value: %v, i:%v, l:%v", pos, value, i, l)
					target = append(target, value)
				}
			}
			// Remove the i-th column
			log.Println("TARGET", target)
			for pos, value := range target {
				if pos < i*(l-1) && pos > (i+1)*(l-1) {
					target2 = append(target2, value)
				}
			}

			// Now copy the new structure
			v.Digraph = target2
			// Now remove the node i

			// And finally renum all the node > i to i-1
		}
	*/
	return v
}

func shrink() {

}

func main() {
	var t toscalib.ToscaDefinition
	var imports []toscalib.ToscaDefinition
	imports = make([]toscalib.ToscaDefinition, 0)
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
	// Fill the digraph
	// Deal with the imports
	for _, im := range t.Imports {
		var tt toscalib.ToscaDefinition
		log.Println("Importing ", im)
		r, err := os.Open(im)
		if err != nil {
			log.Fatal(err)
		}
		err = tt.Parse(r)
		if err != nil {
			log.Fatal(err)
		}
		imports = append(imports, tt)
	}

	// Now reconstruct the global definition (only the types by now)
	for _, i := range imports {
		for key, m := range i.NodeTypes {
			if _, ok := t.NodeTypes[key]; !ok {
				t.NodeTypes[key] = m
			}
		}
	}

	v = togorch(t)
	// Convert it to gorch
	//r, _ := json.Marshal(v)
	r, _ := json.MarshalIndent(v, "  ", "  ")
	fmt.Printf("%s\n", string(r))
}
