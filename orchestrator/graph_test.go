/*
Olivier Wulveryck - author of Gorchestrator
Copyright (C) 2015 Olivier Wulveryck

This file is part of the Gorchestrator project and
is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package orchestrator

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRun(t *testing.T) {
	e := valid.Check()
	if e.Code != 0 {
		t.Errorf("Struct should be valid, error is: %v", e.Error())
	}
	e = notValid.Check()
	if e.Code == 0 {
		t.Errorf("Struct should not be valid, error is: %v", e.Error())
	}

	tasks = make(map[string](*Node), 0)
	router := NewRouter()
	// Load client cert
	cert, err := tls.LoadX509KeyPair("./test/executor.pem", "./test/executor_key.pem")
	if err != nil {
		t.Fatal(err)

	}

	caCert, err := ioutil.ReadFile("./test/orchestrator.pem")
	if err != nil {
		t.Fatal(err)

	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	ts := httptest.NewUnstartedServer(router)
	ts.TLS = &tls.Config{
		ClientCAs:    caCertPool,
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	ts.StartTLS()
	defer ts.Close()
	exe := ExecutorBackend{
		"self",
		fmt.Sprintf("%v/v1", ts.URL),
		"./test/orchestrator.pem",
		"./test/orchestrator_key.pem",
		"./test/executor.pem",
		"/ping",
		nil,
	}
	err = exe.Init()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Launching tests")
	allValid := []Graph{valid, validAndNoArtifact, validAndSleep, validAndExecSuccess}
	allValid = []Graph{
		valid,
		valid,
		valid,
		valid,
		valid,
		valid,
		valid,
	}

	var wg sync.WaitGroup
	for index, _ := range allValid {
		wg.Add(1)
		go func(v Graph) {
			v.Run([]ExecutorBackend{exe})
			if v.GetState() != Success {
				t.Fatalf("Failed: %v", v)
			}
			t.Logf("[%v] Test Finished", v.Name)
			wg.Done()
		}(allValid[index])
	}
	/*
		allInvalid := []Graph{validAndTimeout, validAndExecFailure}
		for _, v := range allInvalid {
			wg.Add(1)
			go func(v *Graph) {
				v.Run([]ExecutorBackend{exe})
				if v.State <= Success {
					t.Fatalf("Failed: %v", v)
				}
				wg.Done(v)
			}()
		}
	*/
	wg.Wait()
}
func BenchmarkRun(b *testing.B) {
	e := valid.Check()
	if e.Code != 0 {
		b.Errorf("Struct should be valid, error is: %v", e.Error())
	}
	e = notValid.Check()
	if e.Code == 0 {
		b.Errorf("Struct should not be valid, error is: %v", e.Error())
	}
	exe := ExecutorBackend{
		"self",
		"https://localhost:8585/v1",
		"./security/certs/orchestrator/orchestrator.pem",
		"./security/certs/orchestrator/orchestrator_key.pem",
		"./security/certs/executor/executor.pem",
		"/ping",
		nil,
	}
	exe.Init()

	var wg sync.WaitGroup
	count := b.N
	vs := make([]Graph, count)
	wg.Add(count)
	for i := 0; i < count; i++ {
		vs[i] = valid
		vs[i].Name = fmt.Sprintf("%v", i)
		go func(v *Graph, wg *sync.WaitGroup) {
			v.Run([]ExecutorBackend{exe})
			wg.Done()
		}(&vs[i], &wg)
	}
	wg.Wait()
}
