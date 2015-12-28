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
	"fmt"
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
	exe := ExecutorBackend{
		"https://localhost:8585/v1",
		"./security/certs/orchestrator/orchestrator.pem",
		"./security/certs/orchestrator/orchestrator_key.pem",
		"./security/certs/executor/executor.pem",
		"/ping",
		nil,
	}
	exe.Init()

	vs := []Graph{valid, validAndTimeout, validAndNoArtifact, validAndSleep}

	for _, v := range vs {
		v.Run(exe)
	}
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
		go func(v Graph, wg *sync.WaitGroup) {
			v.Run(exe)
			wg.Done()
		}(vs[i], &wg)
	}
	wg.Wait()
}
