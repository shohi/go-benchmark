package benchmark

import (
	"sync"
	"testing"
)

// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func FibWait(n int, wg *sync.WaitGroup) int {
	wg.Wait()
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func BenchmarkFib10(b *testing.B) {
	var fibResult int

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibResult = Fib(10)
	}

	_ = fibResult
}

func BenchmarkFib10_WaitGroup(b *testing.B) {
	var fibResult int
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		wg.Wait()
		fibResult = Fib(10)
	}

	_ = fibResult
}
