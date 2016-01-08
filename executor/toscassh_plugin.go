package executor

import (
	"fmt"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// toscassh is a special plugin where all the inputs are passed as environment variables
func (n *node) toscassh() error {
	var conf map[string]sshConfig
	r, err := os.Open("sshConfig.yaml")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	r.Close()

	if _, ok := conf[n.Target]; !ok {
		n.State = orchestrator.Failure
		return fmt.Errorf("Cannot find entry for host %v in the ssh config file", n.Target)
	}
	// ssh.Password("your_password")
	sshConfig := &ssh.ClientConfig{
		User: conf[n.Target].User,
		Auth: []ssh.AuthMethod{
			SSHAgent(),
			PublicKeyFile(conf[n.Target].PrivateKeyFile),
		},
	}

	client := &SSHClient{
		Config: sshConfig,
		Host:   conf[n.Target].Host,
		Port:   conf[n.Target].Port,
	}

	cmd := &SSHCommand{
		Path:   n.Artifact,
		Env:    []string{},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	log.Printf("[%v] Running command: %s\n", n.Name, cmd.Path)
	n.State = orchestrator.Running
	if err := client.RunCommand(cmd); err != nil {
		n.State = orchestrator.Failure
		log.Printf("command run error: %s\n", err)
		return err
	}
	n.State = orchestrator.Success
	return nil
}
