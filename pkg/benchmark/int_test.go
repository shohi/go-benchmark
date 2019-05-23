package benchmark

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkIntToString(b *testing.B) {
	var val int64 = 123456789098764321
	var result string

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// result = fmt.Sprintf("%d", val)
			_ = fmt.Sprintf("%d", val)
		}

		_ = result
	})

	b.Run("fmt.Sprint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// result = fmt.Sprint(val)
			_ = fmt.Sprint(val)
		}

		_ = result

	})

	b.Run("strconv.FormatInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// var tmp [1024]byte
			// _ = strconv.AppendInt(tmp[:0], int64(val), 10)
			// result = strconv.FormatUint(uint64(val), 10)
			_ = strconv.FormatInt(val, 10)
		}

		_ = result

	})

}
