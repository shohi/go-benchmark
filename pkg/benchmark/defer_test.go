package benchmark

import (
	"sync"
	"testing"
)

// refer, https://github.com/jeromefroe/golang_benchmarks/blob/master/defer_test.go

func BenchmarkDefer(b *testing.B) {
	b.Run("MutexDeferUnlock", benchMutexDeferUnlock)
	b.Run("MutexUnlock", benchMutexUnlock)
}

func benchMutexDeferUnlock(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			defer mu.Unlock()
		}()
	}
}

func benchMutexUnlock(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			mu.Unlock()
		}()
	}
}
