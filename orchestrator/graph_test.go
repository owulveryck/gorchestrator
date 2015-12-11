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

func BenchmarkRun(b *testing.B) {
	e := valid.Check()
	if e.Code != 0 {
		b.Errorf("Struct should be valid, error is: %v", e.Error())
	}
	e = notValid.Check()
	if e.Code == 0 {
		b.Errorf("Struct should not be valid, error is: %v", e.Error())
	}

	var wg sync.WaitGroup
	count := b.N
	vs := make([]Graph, count)
	wg.Add(count)
	for i := 0; i < count; i++ {
		vs[i] = valid
		vs[i].Name = fmt.Sprintf("%v", i)
		go func(v Graph, wg *sync.WaitGroup) {
			v.Run(nil)
			wg.Done()
		}(vs[i], &wg)
	}
	wg.Wait()
}
