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
	vs := make([]Input, count)
	wg.Add(count)
	for i := 0; i < count; i++ {
		vs[i] = valid
		vs[i].Name = fmt.Sprintf("%v", i)
		go func(v Input, wg *sync.WaitGroup) {
			v.Run(nil)
			wg.Done()
		}(vs[i], &wg)
	}
	wg.Wait()
}
