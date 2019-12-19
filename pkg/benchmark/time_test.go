package benchmark

import (
	"testing"
	"time"
)

var testTime time.Time

func BenchmarkTimeNow(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testTime = time.Now()
	}

	_ = testTime
}
