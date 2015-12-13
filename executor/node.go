package executor

import (
	"github.com/owulveryck/gorchestrator/orchestrator"
	"log"
)

type node orchestrator.Node

// Run actually launch the engine to run the node
func (n *node) Run() {
	log.Printf("Running %v with engine %v, artifact %v and args %v", n.Name, n.Engine, n.Artifact, n.Args)
	n.State = orchestrator.Success
	log.Println(n.State)
}
