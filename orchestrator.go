package main

import (
	"fmt"
	"github.com/owulveryck/gorchestrator/structure"
	"math/rand"
	"time"
)

const (
	ToRun   = 1
	Running = 2
	Success = 3
	Failure = 4
)

type Message struct {
	id    int
	state int
	wait  chan structure.Matrix
}

func main() {
	// Allocate a zeroed array of size 8×8
	m := structure.Matrix{
		0, 1, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 0, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0,
	}

	n := m.Dim()
	cs := make([]<-chan Message, n)
	for i := 0; i < n; i++ {
		cs[i] = run(i, time.Duration(rand.Intn(1e3))*time.Millisecond)
		node := <-cs[i]
		go func() {
			for {
				node.wait <- m
			}
		}()
	}
	c := fanIn(cs...)
	timeout := time.After(5 * time.Second)
	for {
		select {
		case node := <-c:
			if node.state >= Running {
				fmt.Printf("%v has finished\n", node.id)
				// 0 its row in the matrix
				for c := 0; c < n; c++ {
					m.Set(node.id, c, int64(node.state))
				}
			}
		case <-timeout:
			fmt.Println("Timeout")
			return
		default:
			if m.Sum() >= int64(Success*n*n) {
				fmt.Println("All done!")
				return
			}
		}
	}
	fmt.Println("This is the end!")
}

func run(id int, duration time.Duration) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan structure.Matrix) // Shared between all messages.
	go func() {
		state := ToRun
		for state <= ToRun {
			c <- Message{id, state, waitForIt}
			m := <-waitForIt
			s := m.Dim()
			state = Running
			for i := 0; i < s; i++ {
				if m.At(i, id) != Success && m.At(i, id) != 0 {
					state = ToRun
				}
			}
			if state == Running {
				fmt.Printf("I am %v, and I am running\n", id)
				time.Sleep(duration)
				state = Success
				// Now send the message that I'm done...
				c <- Message{id, state, waitForIt}
			}
		}
		close(c)
	}()
	return c
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
