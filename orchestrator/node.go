package orchestrator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/owulveryck/gorchestrator/structure"
	"math/rand"
	"net/http"
	"regexp"
	"sync"
	"time"
)

// Node is a "runable" node description
type Node struct {
	ID       int               `json:"id"`
	State    int               `json:"state,omitempty"`
	Name     string            `json:"name,omitempty"`   // The targeted host
	Target   string            `json:"target,omitempty"` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Engine   string            `json:"engine,omitempty"` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string            `json:"artifact"`
	Args     []string          `json:"args,omitempty"`   // the arguments of the artifact, if needed
	Outputs  map[string]string `json:"output,omitempty"` // the key is the name of the parameter, the value its value (always a string)
	GraphID  string            `json:"graph_id,omitempty"`
	execID   string            `json:"-"`
	mu       sync.RWMutex      `json:"-"`
}

func (n Node) GetState() int {
	var state int
	n.mu.RLock()
	defer n.mu.RUnlock()
	state = n.State
	return state
}

func (n *Node) SetState(s int) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.State = s
}

// Actually executes the node (via the executor)
func (n *Node) Execute(exe ExecutorBackend) error {
	n.LogDebug("Entering the Execute function")
	if exe.Client == nil {
		err := exe.Init()
		if err != nil {
			return err
		}
	}
	var id string
	var err error
	var t struct {
		ID string `json:"id"`
	}
	url := fmt.Sprintf("%v/tasks", exe.Url)
	b, _ := json.Marshal(n)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	//client := &http.Client{}
	client := exe.Client
	// Do a ping before for testing purpose
	resp, err := client.Do(req)
	if err != nil {
		//n.SetState(NotRunnable)
		return err

	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		//n.SetState(NotRunnable)
		return errors.New("Error in the executor")

	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&t); err != nil {
		//n.SetState(Failure)
		return err
	}
	n.execID = t.ID
	id = t.ID
	n.LogInfo("Running")

	// Now loop for the result
	var res Node
	for err == nil && res.State < Success {
		r, err := client.Get(fmt.Sprintf("%v/%v", url, id))
		if err != nil {
			//n.SetState(NotRunnable)
			return err
		}
		defer r.Body.Close()
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&res); err != nil {
			//n.SetState(Failure)
			return err
		}
		time.Sleep(2 * time.Second)
	}
	*n = res
	return nil
}

// Run the node
func (n *Node) Run(exe []ExecutorBackend) <-chan Message {
	n.LogDebug("Entering the Run function")
	c := make(chan Message)
	waitForIt := make(chan Graph) // Shared between all messages.
	var ga = regexp.MustCompile(`^(.*)=get_attribute (.+):(.+)$`)

	var g Graph
	go func() {
		n.SetState(ToRun)
		if len(n.Outputs) == 0 {
			n.Outputs = make(map[string]string, 0)
		}
		if n.Artifact == "" && n.Engine == "" {
			n.Engine = "nil"
		}

		for n.GetState() <= ToRun {
			n.LogDebugf("Advertize %v", 1)
			c <- Message{n.ID, n.GetState(), waitForIt}
			g = <-waitForIt
			var m structure.Matrix
			m = g.Digraph
			s := m.Dim()
			state := Running
			for i := 0; i < s; i++ {
				mu.RLock()
				if m.At(i, n.ID) < Success && m.At(i, n.ID) > 0 {
					state = ToRun
				} else if m.At(i, n.ID) >= Failure {
					state = NotRunnable
				}
				mu.RUnlock()
				if n.GetState() == NotRunnable {
					continue
				}
			}
			n.SetState(state)
			if n.GetState() == NotRunnable {
				n.LogDebug("Not runnable, advertising graph")
				c <- Message{n.ID, n.GetState(), waitForIt}
			}
			if n.GetState() == Running {
				// Check and find the arguments
				for i, arg := range n.Args {
					// If argument is a get_attribute node:attribute
					// Then substitute it to its actual value
					subargs := ga.FindStringSubmatch(arg)
					if len(subargs) == 4 {
						nn, _ := g.getNodesFromRegexp(subargs[2])
						for _, nn := range nn {
							n.Args[i] = fmt.Sprintf("%v=%v", subargs[1], nn.Outputs[subargs[3]])
						}
					}
				}

				n.LogDebugf("Advertize %v", 2)
				c <- Message{n.ID, n.GetState(), waitForIt}
				switch n.Engine {
				case "nil":
					n.SetState(Success)
				case "sleep": // For test purpose
					time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
					rand.Seed(time.Now().Unix())
					n.SetState(Success)
					n.Outputs["result"] = fmt.Sprintf("%v_%v", n.Name, time.Now().Unix())
				default:
					var executor ExecutorBackend
					executor = exe[0]
					for _, eb := range exe {
						if eb.Name == n.Target {
							executor = eb
						}
					}
					err := n.Execute(executor)
					if err != nil && n.GetState() <= Success {
						n.SetState(Failure)
					} else {
						n.SetState(Success)
					}

				}
				n.LogDebugf("Advertize %v", 3)
				c <- Message{n.ID, n.GetState(), waitForIt}
			}
		}
		close(c)
	}()
	return c
}
