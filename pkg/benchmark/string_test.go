package benchmark

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// refer, https://gist.github.com/knsh14/6f02536c1ef1863f9ad67f33995e52a7
func BenchmarkString(b *testing.B) {
	cases := []int{
		// 10,
		// 50,
		100,
		// 200,
		// 500,
	}

	for _, n := range cases {
		b.Run(fmt.Sprintf("BytesBufferOutput-%v", n), func(b *testing.B) {
			benchBytesBufferOutput(n, b)
		})
		b.Run(fmt.Sprintf("StringsBuilderOutput-%v", n), func(b *testing.B) {
			benchStringsBuilderOutput(n, b)
		})

		b.Run(fmt.Sprintf("BytesBufferWrite-%v", n), func(b *testing.B) {
			benchBytesBufferWrite(n, b)
		})
		b.Run(fmt.Sprintf("StringsBuilderWrite-%v", n), func(b *testing.B) {
			benchStringsBuilderWrite(n, b)
		})

		b.Run(fmt.Sprintf("GrownBytesBufferWrite-%v", n), func(b *testing.B) {
			benchGrownBytesBufferWrite(n, b)
		})
		b.Run(fmt.Sprintf("GrownStringsBuilderWrite-%v", n), func(b *testing.B) {
			benchGrownStringsBuilderWrite(n, b)
		})
	}
}

func benchBytesBufferOutput(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		var buf bytes.Buffer
		for k := 0; k < 10; k++ {
			buf.WriteString("sample\n")
		}
		for k := 0; k < 10; k++ {
			buf.WriteString("sample\n")
		}
		for k := 0; k < n; k++ {
			_ = buf.String()
		}
	}

}

func benchStringsBuilderOutput(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		var buf strings.Builder
		for k := 0; k < 10; k++ {
			buf.WriteString("sample\n")
		}
		for k := 0; k < n; k++ {
			_ = buf.String()
		}
	}
}

func benchBytesBufferWrite(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteString("sample\n")
		}
		_ = buf.String()
	}
}
func benchStringsBuilderWrite(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		var buf strings.Builder
		for k := 0; k < n; k++ {
			buf.WriteString("sample\n")
		}
		_ = buf.String()
	}
}

func benchGrownBytesBufferWrite(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		var buf bytes.Buffer
		buf.Grow(n * 6)
		for k := 0; k < n; k++ {
			buf.WriteString("sample\n")
		}
		_ = buf.String()
	}
}

func benchGrownStringsBuilderWrite(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		var buf strings.Builder
		buf.Grow(n * 6)
		for k := 0; k < n; k++ {
			buf.WriteString("sample\n")
		}
		_ = buf.String()
	}
}
