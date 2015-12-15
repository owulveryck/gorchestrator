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
	"github.com/owulveryck/toscalib"
	"log"
	"os"
	"reflect"
)

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
	for i, im := range t.Imports {
		log.Println("Importing ", im)
		r, err := os.Open(im)
		if err != nil {
			log.Fatal(err)
		}
		err = t.Parse(r)
		if err != nil {
			log.Fatal(err)
		}
		imports[i] = t
	}
	s, _ := t.AdjacencyMatrix.Dims()
	v.Digraph = make([]int64, s*s+s)
	v.Nodes = make([]orchestrator.Node, s)
	for i := 0; i < s; i++ {
		v.Nodes[i].ID = i
		n := t.GetNodeTemplateFromId(i)
		// Fill in a map with method as key and artifact as value
		interfaces := make(map[string]string, 0)
		for _, intf := range n.Interfaces {
			for val, vv := range intf {
				switch reflect.TypeOf(vv).Kind() {
				case reflect.String:
					interfaces[val] = reflect.ValueOf(vv).String()
				case reflect.Map:
					interfaces[val] = fmt.Sprintf("%v", reflect.ValueOf(vv).MapIndex(reflect.ValueOf("implementation")))
				default:
					interfaces[val] = "Found"
				}
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
