package executor

import (
	"bytes"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"log"
	"os"
	"os/exec"
)

type node orchestrator.Node

// Run actually launch the engine to run the node
func (n *node) Run() {
	log.Printf("Running %v with engine %v, artifact %v and args %v", n.Name, n.Engine, n.Artifact, n.Args)
	if n.Artifact != "" {
		d := exec.Command(n.Artifact, n.Args...)
		// Set the stdin stdout and stderr of the dot subprocess
		stdinOfDotProcess, err := d.StdinPipe()
		if err != nil {
			n.State = orchestrator.Failure
			return

		}
		defer stdinOfDotProcess.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
		readCloser, err := d.StdoutPipe()
		if err != nil {
			n.State = orchestrator.Failure
			return

		}
		d.Stderr = os.Stderr

		// Actually run the dot subprocess
		if err = d.Start(); err != nil { //Use start, not run
			n.State = orchestrator.Failure
			log.Println("An error occured: ", err) //replace with logger, or anything you want
		}
		//fmt.Fprintf(stdinOfDotProcess, s)
		stdinOfDotProcess.Close()

		// Read from stdout and store it in the correct structure
		var buf bytes.Buffer
		buf.ReadFrom(readCloser)
		d.Wait()
		n.State = orchestrator.Success
	}
}
