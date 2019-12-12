package benchmark

import (
	"sync/atomic"
	"testing"
)

func BenchmarkCAS(b *testing.B) {
	cases := []struct {
		name string
		fn   func(b *testing.B)
	}{
		{"int32", benchCASWithInt32},
		{"int64", benchCASWithInt64},
	}

	for _, c := range cases {
		b.Run(c.name, c.fn)
	}
}

func benchCASWithInt64(b *testing.B) {
	var val int64 = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.CompareAndSwapInt64(&val, 0, 1)
		atomic.CompareAndSwapInt64(&val, 1, 0)
	}
}

func benchCASWithInt32(b *testing.B) {

	var val int32 = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.CompareAndSwapInt32(&val, 0, 1)
		atomic.CompareAndSwapInt32(&val, 1, 0)
	}
}
