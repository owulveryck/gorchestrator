package executor

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"os"
	"os/exec"
)

type node orchestrator.Node

// shell runs the node via the shell plugin on host localhost
func (n *node) shell() error {
	d := exec.Command(n.Artifact, n.Args...)
	// Set the stdin stdout and stderr of the dot subprocess
	stdinOfDotProcess, err := d.StdinPipe()
	if err != nil {
		n.State = orchestrator.Failure
		return err

	}
	defer stdinOfDotProcess.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
	readCloser, err := d.StdoutPipe()
	if err != nil {
		n.State = orchestrator.Failure
		return err

	}
	d.Stderr = os.Stderr

	// Actually run the dot subprocess
	if err = d.Run(); err != nil { //Use start, not run
		n.State = orchestrator.Failure
		log.Error("An error occured: ", err) //replace with logger, or anything you want
		return err
	}
	//fmt.Fprintf(stdinOfDotProcess, s)
	stdinOfDotProcess.Close()

	// Read from stdout and store it in the correct structure
	var buf bytes.Buffer
	buf.ReadFrom(readCloser)
	n.State = orchestrator.Success

	return nil
}

// Run actually launch the engine to run the node
func (n *node) Run() {
	contextLogger := log.WithFields(log.Fields{
		"Service": "executor",
		"Method":  "node.Run",
		"Node":    n.Name,
	})
	contextLogger.Data["Method"] = "node.Run"
	contextLogger.Data["Node"] = n.Name
	contextLogger.Data["State"] = n.State
	if n.Artifact != "" {
		switch n.Engine {
		case "shell":
			_ = n.shell()
		case "ssh":
			_ = n.ssh()
		case "toscassh":
			contextLogger.Info("Launching toscassh engine")
			err := n.toscassh()
			if err != nil {
				contextLogger.Warning("Execution returned", err)
			} else {
				contextLogger.Info("Execution successfuled", err)
			}
		case "default":
			n.State = orchestrator.Failure
		}
	}
}
