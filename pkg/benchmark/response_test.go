package benchmark

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

var probBuf [1]byte

type nopWriter struct{}

func (w nopWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func readWithIOUtil(resp *http.Response) []byte {
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}

	return data
}

func readWithBodyRead(resp *http.Response) []byte {
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.Body == nil || resp.ContentLength <= 0 {
		return nil
	}

	data := make([]byte, resp.ContentLength)

	_, err := io.ReadFull(resp.Body, data)
	if err != nil {
		return nil
	}

	if err == nil {
		_, err = io.ReadFull(resp.Body, probBuf[0:])
		if err == nil {
			return nil
		}
	}

	return data
}

func readWithCopyBuffer(resp *http.Response) []byte {
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.Body == nil || resp.ContentLength <= 0 {
		return nil
	}

	data := make([]byte, resp.ContentLength)

	_, err := io.CopyBuffer(nopWriter{}, resp.Body, data)
	if err != nil {
		return nil
	}

	if err == nil {
		_, err = io.ReadFull(resp.Body, probBuf[0:])
		if err == nil {
			return nil
		}
	}

	return data
}

func BenchmarkResponse(b *testing.B) {
	cases := []struct {
		name string
		fn   func(resp *http.Response) []byte
	}{
		{"ReadWithIOUtil", readWithIOUtil},
		{"ReadWithBody", readWithBodyRead},
		{"ReadWithCopyBuffer", readWithCopyBuffer},
	}

	var resp http.Response

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			// r := rand.New(rand.NewSource(1))
			// size := r.Int63n(1024*1024*2) + 1
			var size int64 = 1024 * 1024 * 2
			for k := 0; k < b.N; k++ {
				resp.Body = ioutil.NopCloser(bytes.NewReader(make([]byte, size)))
				resp.ContentLength = size
				c.fn(&resp)
			}
		})
	}
}
