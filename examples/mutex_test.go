package mutex

import "testing"
import "sync"

type Map struct {
	sync.Mutex
	m map[string]int
}

func (m *Map) Put(key string, val int) {
	m.Lock()
	m.m[key] = val
	m.Unlock()
}

func BenchmarkMutex(b *testing.B) {
	const N = 16
	for n := 0; n < b.N; n++ {
		var done sync.WaitGroup
		done.Add(N)
		m := Map{m: make(map[string]int)}
		for i := 0; i < N; i++ {
			go func(i int) {
				defer done.Done()
				m.Put("foo", i)
			}(i)
		}
		done.Wait()
	}
}
