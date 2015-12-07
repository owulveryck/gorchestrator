package main

import (
	"fmt"
	"github.com/owulveryck/gorchestrator/structure"
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

func main() {
	// Allocate a zeroed array of size 8×8
	m := valid.Digraph

	n := m.Dim()
	cs := make([]<-chan structure.Message, n)
	l := new(lock)
	co := sync.NewCond(l)
	for i := 0; i < n; i++ {
		cs[i] = valid.Nodes[i].Run()
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
	timeout := time.After(5 * time.Second)
	for {
		select {
		case node := <-c:
			if node.State >= structure.Running {
				fmt.Printf("%v has finished (%v)\n", node.ID, node.State)
				for c := 0; c < n; c++ {
					m.Set(node.ID, c, int64(node.State))
				}
				co.Broadcast()
			}
		case <-timeout:
			fmt.Println("Timeout")
			return
		default:
			if m.Sum() >= int64(structure.Success*n*n) {
				fmt.Println("All done!")
				return
			}
		}
	}
}

func fanIn(inputs ...<-chan structure.Message) <-chan structure.Message {
	c := make(chan structure.Message)
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
