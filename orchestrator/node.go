package orchestrator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	//"log"
	"math/rand"
	"net/http"
	//	"regexp"
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
	execID   string
	sync.RWMutex
}

func (n *Node) GetState() int {
	n.RLock()
	defer n.RUnlock()
	var state int
	state = n.State
	return state
}

func (n *Node) SetState(s int) {
	n.Lock()
	defer n.Unlock()
	n.State = s
}

// Actually executes the node (via the executor)
func (n *Node) Execute(exe ExecutorBackend) error {
	//n.//log.ebug("Entering the Execute function")
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
	//n.//log.nfo("Running")

	// Now loop for the result
	var res Node
	for err == nil && res.GetState() < Success {
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
	n.Outputs = res.Outputs
	return nil
}

// Run the node
func (n *Node) Run(exe []ExecutorBackend, done <-chan struct{}, in <-chan Information) <-chan Message {
	out := make(chan Message)
	//var ga = regexp.MustCompile(`^(.*)=get_attribute (.+):(.+)$`)

	var g Information
	go func() {
		defer close(out)
		n.SetState(ToRun)
		if len(n.Outputs) == 0 {
			n.Outputs = make(map[string]string, 0)
		}
		if n.Artifact == "" && n.Engine == "" {
			n.Engine = "nil"
		}

		message := Message{n.ID, n.GetState()}

		// This "angle" functions, waits for an information and change the state of the node regarding the information received
		stateChan := make(chan int)
		currentState := n.GetState()
		go func() {
			for g = range in {
				//log.Printf("[%v] Received a message (sequence number %v)", n.ID, g.Sequence)
				s := g.Matrix.Dim()
				newState := Running
				for i := 0; i < s; i++ {
					val := g.Matrix.At(i, n.ID)
					if val < Success && val > 0 {
						newState = ToRun
					} else if val >= Failure {
						newState = NotRunnable
					}
					if newState == NotRunnable {
						continue
					}
				}
				if currentState != newState {
					// State has changed, tell it to the main
					currentState = newState
					stateChan <- currentState
					message.State = currentState
					out <- message
				}
			}
		}()

		for state := range stateChan {
			//log.Printf("[%v] state has changed => %v", n.ID, state)
			n.SetState(state)
			switch {
			case n.GetState() == Running:
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
					}
				}
			}
			//log.Printf("[%v] Done", n.ID)
			if message.State != n.GetState() {
				message.State = n.GetState()
				out <- message
			}
		}
	}()
	return out
}
