package main

import (
	"fmt"
	"github.com/owulveryck/gorchestrator/structure"
	"math/rand"
	"time"
)

type Message struct {
	id   int
	run  bool
	wait chan structure.Matrix
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
			if node.run == true {
				fmt.Printf("%v has finished\n", node.id)
				// 0 its row in the matrix
				for c := 0; c < n; c++ {
					m.Set(node.id, c, int64(0))
				}
			}
		case <-timeout:
			fmt.Println("Timeout")
			return
		default:
			if m.Sum() == 0 {
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
		run := false
		for run == false {
			c <- Message{id, run, waitForIt}
			m := <-waitForIt
			s := m.Dim()
			run = true
			for i := 0; i < s; i++ {
				if m.At(i, id) == 1 {
					run = false
				}
			}
			if run == true {
				fmt.Printf("I am %v, and I am running\n", id)
				time.Sleep(duration)
				// Now send the message that I'm done...
				c <- Message{id, run, waitForIt}
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
