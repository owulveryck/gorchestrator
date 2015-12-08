package orchestrator

import (
	"github.com/owulveryck/gorchestrator/structure"
	"sync"
	"time"
)

// Input is the input of the orchestrator
type Input struct {
	Name    string           `json:"name",omitempty`
	State   int              `json:"state"`
	Digraph structure.Matrix `json:"digraph"`
	Nodes   []Node           `json:"nodes"`
}

// Run executes the Input structure
func (v *Input) Run(stop <-chan time.Time) {
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
				node.Wait <- m
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
					m.Set(node.ID, c, int64(node.State))
				}
				co.Broadcast()
			}
		case <-stop:
			v.State = Canceled
			return
		default:
			switch {
			case m.Sum() == int64(Success*n*n):
				v.State = Success
				return
			case m.Sum() > int64(Success*n*n):
				//fmt.Println("All done!")
				v.State = Failure
				return
			}
		}
	}
}

// Check is the structure is coherent, (a squared matrix with as many nodes as needed)
func (i *Input) Check() Error {
	if len(i.Nodes)*len(i.Nodes) != len(i.Digraph) {
		return Error{1, "Structure is not coherent"}
	}
	return Error{0, ""}
}
