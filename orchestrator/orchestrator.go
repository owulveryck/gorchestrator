package orchestrator

import (
	//"fmt"
	"sync"
	"time"
)

type lock int

func (l *lock) Lock() {
	*l = 1
}
func (l *lock) Unlock() {
	*l = 0
}

// Run executes the Input structure
func (v *Input) Run() {
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
	timeout := time.After(50 * time.Second)
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
		case <-timeout:
			v.State = Timeout
			//fmt.Println("Timeout")
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

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}
