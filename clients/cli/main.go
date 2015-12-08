package main

import (
	"encoding/json"
	"fmt"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"os"
	"sync"
	"time"
)

func main() {

	var v orchestrator.Graph
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&v); err != nil {
		panic(err)

	}

	timeout := time.After(3 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		v.Run(timeout)
		wg.Done()
	}(&wg)
	fmt.Println("Waiting")
	wg.Wait()
}
