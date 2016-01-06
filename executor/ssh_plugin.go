package executor

import (
	"fmt"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type sshConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	PublicKeyFile  string `yaml:"public_key"`
	PrivateKeyFile string `yaml:"private_key"`
	User           string `yaml:"remote_user"`
}

type SSHCommand struct {
	Path   string
	Env    []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

func (client *SSHClient) RunCommand(cmd *SSHCommand) error {
	var (
		session *ssh.Session
		err     error
	)

	if session, err = client.newSession(); err != nil {
		return err

	}
	defer session.Close()

	if err = client.prepareCommand(session, cmd); err != nil {
		return err

	}

	err = session.Run(cmd.Path)
	return err

}

func (client *SSHClient) prepareCommand(session *ssh.Session, cmd *SSHCommand) error {
	for _, env := range cmd.Env {
		variable := strings.Split(env, "=")
		if len(variable) != 2 {
			continue

		}

		if err := session.Setenv(variable[0], variable[1]); err != nil {
			return err

		}

	}

	if cmd.Stdin != nil {
		stdin, err := session.StdinPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stdin for session: %v", err)

		}
		go io.Copy(stdin, cmd.Stdin)

	}

	if cmd.Stdout != nil {
		stdout, err := session.StdoutPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stdout for session: %v", err)

		}
		go io.Copy(cmd.Stdout, stdout)

	}

	if cmd.Stderr != nil {
		stderr, err := session.StderrPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stderr for session: %v", err)

		}
		go io.Copy(cmd.Stderr, stderr)

	}

	return nil

}

func (client *SSHClient) newSession() (*ssh.Session, error) {
	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)

	}

	session, err := connection.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)

	}

	modes := ssh.TerminalModes{
		// ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud

	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		return nil, fmt.Errorf("request for pseudo terminal failed: %s", err)

	}

	return session, nil

}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil

	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil

	}
	return ssh.PublicKeys(key)

}

func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)

	}
	return nil

}

// ssh runs the node via the ssh plugin on host "target"
func (n *node) ssh() error {
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
		return fmt.Errorf("Cannot find entry for host %v in the ssh config file", n.Target)
	}
	// ssh.Password("your_password")
	sshConfig := &ssh.ClientConfig{
		User: conf[n.Target].User,
		Auth: []ssh.AuthMethod{
			PublicKeyFile(conf[n.Target].PrivateKeyFile),
			SSHAgent(),
		},
	}

	client := &SSHClient{
		Config: sshConfig,
		Host:   conf[n.Target].Host,
		Port:   conf[n.Target].Port,
	}

	cmd := &SSHCommand{
		Path:   fmt.Sprintf("%v %v", n.Artifact, n.Args[0:]),
		Env:    []string{},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	fmt.Printf("Running command: %s\n", cmd.Path)
	n.State = orchestrator.Running
	if err := client.RunCommand(cmd); err != nil {
		n.State = orchestrator.Failure
		fmt.Fprintf(os.Stderr, "command run error: %s\n", err)
		return err
	}
	n.State = orchestrator.Success
	return nil
}
