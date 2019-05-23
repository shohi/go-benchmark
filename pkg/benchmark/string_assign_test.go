package benchmark

import (
	"strings"
	"testing"
)

func BenchmarkStringAssign(b *testing.B) {
	var result string
	var longStr = strings.Repeat("hi", 1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newStr := longStr
		result = newStr
	}

	_ = result
}
