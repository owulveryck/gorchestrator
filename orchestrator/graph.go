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
	"log"
	"time"
)

// Graph is the input of the orchestrator
type Graph struct {
	Name    string           `json:"name,omitempty"`
	State   int              `json:"state"`
	Digraph structure.Matrix `json:"digraph"`
	Nodes   []Node           `json:"nodes"`
	Timeout <-chan time.Time `json:"-"`
}

func (v *Graph) getNodeFromName(n string) (Node, error) {
	var nn Node
	return nn, nil
}

// Run executes the Graph structure
func (v *Graph) Run() {
	v.State = Running
	m := v.Digraph

	n := m.Dim()
	cs := make([]<-chan Message, n)
	for i := 0; i < n; i++ {
		cs[i] = v.Nodes[i].Run()
	}
	bcast := func() chan bool {
		bcast := make(chan bool)
		go func() {
			for {
				for i := 0; i < n; i++ {
					node := <-cs[i]
					go func() {
						log.Println("Advertising ", node.ID)
						node.Wait <- *v
					}()
				}
				<-bcast
				log.Println("Received an information on the broadcast channel")
			}
		}()
		return bcast
	}()
	c := fanIn(cs...)
	for {
		select {
		case node := <-c:
			log.Println("Received a message")
			if node.State >= Running {
				for c := 0; c < n; c++ {
					if m.At(node.ID, c) != 0 {
						m.Set(node.ID, c, int64(node.State))
					}
				}
			}
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
			log.Println("Sending true to the broadcast")
			bcast <- true
		case <-v.Timeout:
			v.State = Canceled
			return
			/*
				default:
					log.Println("Defaut received")
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
			*/
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
