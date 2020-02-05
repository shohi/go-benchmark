package benchmark

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/acomagu/bufpipe"
)

func BenchmarkPipe(b *testing.B) {
	osReader, osWriter, _ := os.Pipe()
	ioReader, ioWriter := io.Pipe()
	bReader, bWriter := bufpipe.New(nil)

	cases := []struct {
		name   string
		reader io.Reader
		writer io.WriteCloser
	}{
		{"os", osReader, osWriter},
		{"io", ioReader, ioWriter},
		{"buf", bReader, bWriter},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchPipe(c.reader, c.writer, 512)
			}
		})
	}
}

// FIXME: more reasonable test
func benchPipe(r io.Reader, w io.WriteCloser, n int) {
	done := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			_ = line
		}

		close(done)
	}()

	for i := 0; i < n; i++ {
		w.Write([]byte("hello"))
	}

	w.Close()

	<-done
}
