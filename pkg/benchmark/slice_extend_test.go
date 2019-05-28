package benchmark

import "testing"

// refer, https://gist.github.com/xogeny/b819af6a0cf8ba1caaef

func BenchmarkSliceExtend(b *testing.B) {
	existing := make([]byte, 2048, 2048)

	cases := []struct {
		name string
		fn   func(b *testing.B, buf []byte)
	}{
		{"append", benchAppend},
		{"append-alloc", benchAppendAlloc},
		{"copy", benchCopy},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			c.fn(b, existing)
		})
	}
}

func doCopy(useCopy bool, preAlloc bool, existing []byte) []byte {
	length := len(existing)
	var y []byte
	if useCopy {
		y = make([]byte, 0, 2*length)
		copy(y[:length], existing)
	} else {
		var init []byte
		if preAlloc {
			init = make([]byte, 0, 2*length)
		} else {
			init = []byte{}
		}
		y = append(init, existing...)
	}

	// append more one byte to simulator real cases
	_ = append(y, 'b')
	return y
}

func benchAppend(b *testing.B, buf []byte) {
	for i := 0; i < b.N; i++ {
		doCopy(false, false, buf)
	}
}

func benchAppendAlloc(b *testing.B, buf []byte) {
	for i := 0; i < b.N; i++ {
		doCopy(false, true, buf)
	}
}

func benchCopy(b *testing.B, buf []byte) {
	for i := 0; i < b.N; i++ {
		doCopy(true, true, buf)
	}
}
