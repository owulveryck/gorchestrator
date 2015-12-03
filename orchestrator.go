package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math/rand"
	"time"
)

type Message struct {
	id   int
	run  bool
	wait chan mat64.Dense
}

func main() {
	// Allocate a zeroed array of size 8×8
	m := mat64.NewDense(8, 8, nil)
	m.Set(0, 1, 1)
	m.Set(0, 4, 1) // First row
	m.Set(1, 6, 1)
	m.Set(1, 6, 1) // second row
	m.Set(3, 2, 1)
	m.Set(3, 6, 1) // fourth row
	m.Set(5, 0, 1)
	m.Set(5, 1, 1)
	m.Set(5, 2, 1) // fifth row
	m.Set(7, 6, 1) // seventh row
	fa := mat64.Formatted(m, mat64.Prefix("    "))
	// Display the matrix
	fmt.Printf("\nm = %v\n\n", fa)

	n, _ := m.Dims()
	cs := make([]<-chan Message, n)
	for i := 0; i < n; i++ {
		cs[i] = run(i, time.Duration(rand.Intn(1e3))*time.Millisecond)
		node := <-cs[i]
		go func() {
			for {
				node.wait <- *m
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
					m.Set(node.id, c, 0)
				}
			}
		case <-timeout:
			fmt.Println("Timeout")
			return
		default:
			if mat64.Sum(m) == 0 {
				fmt.Println("All done!")
				return
			}
		}
	}
	fmt.Println("This is the end!")
}

func run(id int, duration time.Duration) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan mat64.Dense) // Shared between all messages.
	go func() {
		run := false
		for run == false {
			c <- Message{id, run, waitForIt}
			m := <-waitForIt
			s, _ := m.Dims()
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
