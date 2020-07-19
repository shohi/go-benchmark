package benchmark

import (
	"sync"
	"testing"
)

func BenchmarkSync_WithLock(b *testing.B) {
	var s sync.Mutex

	b.SetParallelism(10000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Lock()
			s.Unlock()
		}
	})
}

func BenchmarkSync_WithChannel(b *testing.B) {
	ch := make(chan byte, 4096)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
		}
	}()
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(1000)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ch <- byte(0)
		}
	})

	close(ch)
	wg.Wait()
}
