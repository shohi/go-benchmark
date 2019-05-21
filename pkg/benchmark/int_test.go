package benchmark

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkIntToString(b *testing.B) {
	var val int64 = 123456789
	var result string

	b.Run("fmt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = fmt.Sprintf("%d", val)
		}

		_ = result
	})

	b.Run("strconv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = strconv.FormatUint(uint64(val), 10)
		}

		_ = result

	})

}
