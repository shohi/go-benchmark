package benchmark

import (
	"testing"
	"time"

	"github.com/jinzhu/now"
)

func BenchmarkDateParse(b *testing.B) {
	const dateFormat = "2019-05-08 15:04:05.006Z"

	b.Run("Time", func(b *testing.B) {
		var result time.Time
		for i := 0; i < b.N; i++ {
			result, _ = time.Parse(dateFormat, dateFormat)
		}

		_ = result
	})

	b.Run("JinzhuNow", func(b *testing.B) {
		now.TimeFormats = append(now.TimeFormats, dateFormat)

		var result time.Time
		for i := 0; i < b.N; i++ {
			result, _ = now.Parse(dateFormat)
		}

		_ = result
	})
}
