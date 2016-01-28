package orchestrator

func broadcast(ch <-chan Graph, size, lag int) []chan Graph {
	cs := make([]chan Graph, size)
	for i, _ := range cs {
		// The size of the channels buffer controls how far behind the recievers
		// of the fanOut channels can lag the other channels.
		//cs[i] = make(chan Graph)
		cs[i] = make(chan Graph, lag)

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

func fanOut(outputs ...chan<- Graph) chan<- Graph {
	c := make(chan Graph)
	for i := range outputs {
		output := outputs[i] // New instance of 'input' for each loop.
		go func(i int) {
			for {
				output <- <-c
			}
		}(i)
	}
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
