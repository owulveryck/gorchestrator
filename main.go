package main

import (
	"encoding/json"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"os"
)

func main() {

	var v orchestrator.Input
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&v); err != nil {
		panic(err)

	}

	v.Run()
	//valid.Run()
}
