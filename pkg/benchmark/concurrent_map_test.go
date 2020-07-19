package benchmark

import (
	"strconv"
	"sync"
	"testing"

	cmap "github.com/orcaman/concurrent-map"
)

func initCMap(m *cmap.ConcurrentMap, n int) {
	for i := 0; i < n; i++ {
		m.Set(strconv.Itoa(i), i)
	}
}

func initMap(m map[string]int, n int) {
	for i := 0; i < n; i++ {
		m[strconv.Itoa(i)] = i
	}
}
func BenchmarkMap_Concurrent(b *testing.B) {
	m := cmap.New()
	initCMap(&m, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Pop(strconv.Itoa(i))
	}
}

func BenchmarkMap_Lock(b *testing.B) {
	var s sync.Mutex
	m := make(map[string]int, b.N)
	initMap(m, b.N)

	for i := 0; i < b.N; i++ {
		s.Lock()
		delete(m, strconv.Itoa(i))
		s.Unlock()
	}
}
