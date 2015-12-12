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
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Node is a "runable" node description
type Node struct {
	ID       int               `json:"id"`
	State    int               `json:"state,omitempty"`
	Name     string            `json:"name",omitempty`
	Engine   string            `json:"engine",omitempty` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string            `json:"artifact"`
	Args     []string          `json:"args",omitempty`   // the arguments of the artifact, if needed
	Outputs  map[string]string `json:"output",omitempty` // the key is the name of the parameter, the value its value (always a string)
}

// Run executes the artifact of a given node
func (n *Node) Run() <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan Graph) // Shared between all messages.
	var ga = regexp.MustCompile(`^get_attribute (.+):(.+)$`)

	go func() {
		n.State = ToRun
		for n.State <= ToRun {
			c <- Message{n.ID, n.State, waitForIt}
			g := <-waitForIt
			m := g.Digraph
			s := m.Dim()
			n.State = Running
			for i := 0; i < s; i++ {
				if m.At(i, n.ID) < Success && m.At(i, n.ID) > 0 {
					n.State = ToRun
				} else if m.At(i, n.ID) >= Failure {
					n.State = NotRunnable
					continue
				}
			}
			if n.State == NotRunnable {
				//fmt.Printf("I am %v, and I cannot run\n", n.ID)
				c <- Message{n.ID, n.State, waitForIt}
			}
			if n.State == Running {
				// Check and find the arguments
				for i, arg := range n.Args {
					// If argument is a get_attribute node:attribute
					// Then substitute it to its actual value
					subargs := ga.FindStringSubmatch(arg)
					if len(subargs) == 4 {
						nn, _ := g.getNodeFromName(subargs[2])
						n.Args[i] = nn.Outputs[subargs[3]]
					}
				}
				c <- Message{n.ID, n.State, waitForIt}
				switch n.Engine {
				case "nil":
					n.State = Success
				case "sleep": // For test purpose
					//fmt.Printf("I am %v, and I am running: the module %v, with %v %v\n", n.ID, n.Engine, n.Artifact, n.Args)
					time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
					rand.Seed(time.Now().Unix())
					n.State = Success
					n.Outputs["result"] = fmt.Sprintf("%v_%v", n.Name, time.Now().Unix())
				default:
					// Send the message to the appropriate backend
					n.State = Success

				}
				c <- Message{n.ID, n.State, waitForIt}
			}
		}
		close(c)
	}()
	return c
}
