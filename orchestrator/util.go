package orchestrator

type lock int

func (l *lock) Lock() {
	*l = 1
}
func (l *lock) Unlock() {
	*l = 0
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
