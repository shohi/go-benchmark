package benchmark

import (
	"sync"
	"testing"
)

// refer, https://github.com/jeromefroe/golang_benchmarks/blob/master/defer_test.go
func BenchmarkMutexDeferUnlock(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			defer mu.Unlock()
		}()
	}
}

func BenchmarkMutexUnlock(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			mu.Unlock()
		}()
	}
}
