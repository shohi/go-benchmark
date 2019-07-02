package benchmark

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// refer, https://gist.github.com/dtjm/c6ebc86abe7515c988ec
func BenchmarkStringJoin(b *testing.B) {
	b.Run("join", benchJoin)
	b.Run("sprintf", benchSprintf)
	b.Run("builder", benchStringBuilder)

	b.Run("concat", benchConcat)
	b.Run("concat-oneline", benchConcatOneLine)

	b.Run("buffer", benchBuffer)
	b.Run("buffer-reset", benchBufferWithReset)
	b.Run("buffer-fprintf", benchBufferFprintf)

}

var (
	testData = []string{"a", "b", "c", "d", "e"}
)

func benchJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testData, ":")
		_ = s
	}
}

func benchSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%s:%s:%s:%s", testData[0], testData[1], testData[2], testData[3], testData[4])
		_ = s
	}
}

func benchConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testData[0] + ":"
		s += testData[1] + ":"
		s += testData[2] + ":"
		s += testData[3] + ":"
		s += testData[4]
		_ = s
	}
}

func benchConcatOneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testData[0] + ":" +
			testData[1] + ":" +
			testData[2] + ":" +
			testData[3] + ":" +
			testData[4]
		_ = s
	}
}

func benchBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		b.WriteString(testData[0])
		b.WriteByte(':')
		b.WriteString(testData[1])
		b.WriteByte(':')
		b.WriteString(testData[2])
		b.WriteByte(':')
		b.WriteString(testData[3])
		b.WriteByte(':')
		b.WriteString(testData[4])
		s := b.String()
		_ = s
	}
}

func benchBufferWithReset(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		buf.Reset()

		buf.WriteString(testData[0])
		buf.WriteByte(':')
		buf.WriteString(testData[1])
		buf.WriteByte(':')
		buf.WriteString(testData[2])
		buf.WriteByte(':')
		buf.WriteString(testData[3])
		buf.WriteByte(':')
		buf.WriteString(testData[4])
		s := buf.String()
		_ = s
	}
}

func benchBufferFprintf(b *testing.B) {
	buf := &bytes.Buffer{}

	for i := 0; i < b.N; i++ {
		buf.Reset()

		fmt.Fprintf(buf, "%s:%s:%s:%s:%s", testData[0], testData[1], testData[2], testData[3], testData[4])
		s := buf.String()
		_ = s
	}

}

func benchStringBuilder(b *testing.B) {
	var buf strings.Builder

	for i := 0; i < b.N; i++ {
		buf.Reset()

		buf.WriteString(testData[0])
		buf.WriteByte(':')
		buf.WriteString(testData[1])
		buf.WriteByte(':')
		buf.WriteString(testData[2])
		buf.WriteByte(':')
		buf.WriteString(testData[3])
		buf.WriteByte(':')
		buf.WriteString(testData[4])
		s := buf.String()
		_ = s
	}
}
