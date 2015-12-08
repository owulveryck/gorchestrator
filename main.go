package main

import (
	"encoding/json"
	"fmt"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"os"
	"time"
)

func main() {

	var v orchestrator.Input
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&v); err != nil {
		panic(err)

	}

	timeout := time.After(5 * time.Second)

	go v.Run(timeout)
	for v.State < orchestrator.Success {
		fmt.Println(v.Digraph)
		time.Sleep(1 * time.Second)
	}
}
