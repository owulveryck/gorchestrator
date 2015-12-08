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

	go v.Run()
	n := v.Digraph.Dim()
	for v.Digraph.Sum() <= int64(orchestrator.Success*n*n) {
		fmt.Println(v.Digraph)
		time.Sleep(1 * time.Second)
	}
}
