package orchestrator

import (
	"log"
	"sync"
)

func broadcast(ch <-chan *Graph, size, lag int) []chan *Graph {
	cs := make([]chan *Graph, size)
	for i, _ := range cs {
		// The size of the channels buffer controls how far behind the recievers
		// of the fanOut channels can lag the other channels.
		//cs[i] = make(chan Graph)
		cs[i] = make(chan *Graph, lag)

	}
	go func() {
		for i := range ch {
			for _, c := range cs {
				c <- i
			}
		}
		for _, c := range cs {
			// close all our fanOut channels when the input channel is exhausted.
			close(c)
		}
	}()
	return cs

}

func merge(done <-chan struct{}, cs ...<-chan Message) <-chan Message {
	var wg sync.WaitGroup
	out := make(chan Message, 1)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan Message) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)

	}()
	return out

}

func fanIn(inputs ...<-chan Message) <-chan Message {

	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() {
			for {
				log.Println("fanIn", input)
				c <- <-input
				log.Println("end fanIn", input)
			}
		}()
	}
	return c
}
