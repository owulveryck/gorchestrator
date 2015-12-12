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
package orchestrator

import (
	"github.com/owulveryck/gorchestrator/structure"
	"sync"
	"time"
)

// Graph is the input of the orchestrator
type Graph struct {
	Name    string           `json:"name",omitempty`
	State   int              `json:"state"`
	Digraph structure.Matrix `json:"digraph"`
	Nodes   []Node           `json:"nodes"`
}

// Run executes the Graph structure
func (v *Graph) Run(stop <-chan time.Time) {
	v.State = Running
	m := v.Digraph

	n := m.Dim()
	cs := make([]<-chan Message, n)
	l := new(lock)
	co := sync.NewCond(l)
	for i := 0; i < n; i++ {
		cs[i] = v.Nodes[i].Run()
		node := <-cs[i]
		go func() {
			for {
				node.Wait <- *v
				l.Lock()
				defer l.Unlock()
				co.Wait()
			}
		}()
	}
	c := fanIn(cs...)
	for {
		select {
		case node := <-c:
			if node.State >= Running {
				//fmt.Printf("%v has finished (%v)\n", node.ID, node.State)
				for c := 0; c < n; c++ {
					if m.At(node.ID, c) != 0 {
						m.Set(node.ID, c, int64(node.State))
					}
				}
				co.Broadcast()
			}
		case <-stop:
			v.State = Canceled
			return
		default:
			stop := true
			v.State = Success
			for r := 0; r < n; r++ {
				for c := 0; c < n; c++ {
					switch {
					case m.At(r, c) == ToRun:
						stop = false
						v.State = Failure
					case m.At(r, c) == Running:
						stop = false
						v.State = Failure
					case m.At(r, c) > Success:
						v.State = Failure
					}
				}
			}
			if stop {
				return
			}
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
