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
along with this prograv.Digraph.  If not, see <http://www.gnu.org/licenses/>.
*/
package orchestrator

import (
	"github.com/owulveryck/gorchestrator/structure"
	"log"
	"regexp"
	"sync"
	"time"
)

// Graph is the input of the orchestrator
type Graph struct {
	Name    string           `json:"name,omitempty"`
	State   int              `json:"state"`
	Digraph structure.Matrix `json:"digraph"`
	Nodes   []Node           `json:"nodes"`
	Timeout <-chan time.Time `json:"-"`
	mu      sync.RWMutex     `json:"-"`
	ID      string           `json:"id,omitempty"`
}

func (v *Graph) getNodesFromRegexp(n string) ([]Node, error) {
	re := regexp.MustCompile(n)
	var nn []Node
	for _, node := range v.Nodes {
		if re.MatchString(node.Name) {
			nn = append(nn, node)
		}
	}
	return nn, nil
}

func (v *Graph) getNodeFromName(n string) (Node, error) {
	var nn Node
	return nn, nil
}

var mu sync.RWMutex

// Run executes the Graph structure
func (v *Graph) Run(exe []ExecutorBackend) {
	v.State = Running

	n := v.Digraph.Dim()
	cs := make([]<-chan Message, n)
	co := make(chan Graph)
	cos := broadcast(co, n, n)

	for i := 0; i < n; i++ {
		cs[i] = v.Nodes[i].Run(exe)
		v.Nodes[i].mu.Lock()
		v.Nodes[i].waitForEvent = cos[i]
		v.Nodes[i].mu.Unlock()
	}
	c := fanIn(cs...)
	co <- *v
	for {
		select {
		case node := <-c:
			log.Printf("Received notification from node %v, its state is %v", node.ID, node.State)
			if node.State >= Running {
				mu.Lock()
				for c := 0; c < n; c++ {
					if v.Digraph.At(node.ID, c) != 0 {
						v.Digraph.Set(node.ID, c, int64(node.State))
					}
				}
				mu.Unlock()
			}
			state := Success
			for r := 0; r < n; r++ {
				for c := 0; c < n; c++ {
					switch {
					case v.Digraph.At(r, c) == ToRun:
						state = Running
					case v.Digraph.At(r, c) == Running:
						state = Running
					case v.Digraph.At(r, c) > Success:
						state = Failure
					}
				}
			}
			v.State = state
			if v.State >= Success {
				return
			}
			co <- *v
		case <-v.Timeout:
			co <- *v
			v.State = Timeout
			return
		}
	}
}

// Check is the structure is coherent, (a squared matrix with as many nodes as needed)
func (i *Graph) Check() Error {
	if len(i.Nodes)*len(i.Nodes) != len(i.Digraph) {
		return Error{1, "Structure is not coherent"}
	}
	return Error{0, ""}
}
