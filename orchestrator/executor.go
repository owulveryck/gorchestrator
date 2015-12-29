package orchestrator

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ExecutorBackend represent an executor
type ExecutorBackend struct {
	Url         string // eg: https://localhost:8585/v1
	Certificate string //
	Key         string
	CACert      string
	Ping        string // eg /ping
	Client      *http.Client
}

func (e *ExecutorBackend) Init() error {
	var client *http.Client
	// Load client cert
	cert, err := tls.LoadX509KeyPair(e.Certificate, e.Key)
	if err != nil {
		return err
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(e.CACert)
	if err != nil {
		return err

	}
	caCertPool := x509.NewCertPool()
	r := caCertPool.AppendCertsFromPEM(caCert)
	if r == false {
		return fmt.Errorf("No certificate found in %v", e.CACert)
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client = &http.Client{Transport: transport}

	e.Client = client
	// Now check the ping url
	_, err = client.Get(fmt.Sprintf("%v/%v", e.Url, e.Ping))
	if err != nil {
		return err
	}

	return nil
}
