package orchestrator

import (
	"github.com/owulveryck/gorchestrator/structure"
	"math/rand"
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
	waitForIt := make(chan structure.Matrix) // Shared between all messages.
	go func() {
		n.State = ToRun
		for n.State <= ToRun {
			c <- Message{n.ID, n.State, waitForIt}
			m := <-waitForIt
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
				c <- Message{n.ID, n.State, waitForIt}
				switch n.Engine {
				case "nil":
					n.State = Success
				case "sleep": // For test purpose
					//fmt.Printf("I am %v, and I am running: the module %v, with %v %v\n", n.ID, n.Engine, n.Artifact, n.Args)
					time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
					rand.Seed(time.Now().Unix())
					n.State = Success
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
