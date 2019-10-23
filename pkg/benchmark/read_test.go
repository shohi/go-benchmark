package benchmark

import (
	"io"
	"testing"

	"github.com/lukechampine/noescape"
)

type yesReader struct{}

func (yesReader) Read(p []byte) (int, error) {
	return copy(p, "yes"), nil
}

/*
func BenchmarkRead(b *testing.B) {

	cases := []struct {
		name string
		fn   func(b *testing.B)
	}{

		{"Concrete", benchReadConcrete},
		{"Interface", benchReadInterface},
		{"NoEscape", benchReadNoEscape},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			c.fn(100, b)
		})
	}
}
*/

func BenchmarkReadConcrete(b *testing.B) {
	r := yesReader{}
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 100)
		r.Read(buf)
	}
}

func BenchmarkReadInterface(b *testing.B) {
	var r io.Reader = yesReader{}
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 100)
		r.Read(buf)
	}
}

func BenchmarkReadNoEscape(b *testing.B) {
	var r io.Reader = yesReader{}
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 100)
		noescape.Read(r, buf)
	}
}
