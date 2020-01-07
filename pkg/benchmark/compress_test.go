package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/shohi/gocode/pkg/compress"
)

type Tag struct {
	Key   string
	Value string
}

type CompressData struct {
	Name    string
	Addr    string
	Host    string
	Visited int
	Tags    []Tag
}

func genCompressData() []byte {
	cdata := &CompressData{
		Name:    "Test",
		Addr:    "localhost:9001",
		Host:    "abcdedfhijk",
		Visited: 100,
		Tags: []Tag{
			{"k1", "v1"},
			{"k2", "v2"},
		},
	}

	d, _ := json.Marshal(cdata)
	return d
}

func BenchmarkCompression(b *testing.B) {
	cases := []struct {
		name   string
		method string
		algo   compress.Algo
	}{
		{"LZ4", "compress", compress.Lz4Algo{}},
		{"LZ4", "uncompress", compress.Lz4Algo{}},
		{"GZIP", "compress", compress.GzipAlgo{}},
		{"GZIP", "uncompress", compress.GzipAlgo{}},
	}

	for _, c := range cases {

		b.Run(c.name+"-"+c.method, func(b *testing.B) {
			if c.method == "compress" {
				benchCompress(b, c.algo)
			} else {
				benchUncompress(b, c.algo)
			}
		})
	}
}

func benchCompress(b *testing.B, algo compress.Algo) {
	data := genCompressData()
	var out []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out, _ = algo.Compress(data)
	}
	_ = out
}

func benchUncompress(b *testing.B, algo compress.Algo) {
	data := genCompressData()
	var out []byte
	cdata, _ := algo.Compress(data)
	// fmt.Printf("original: %v, compressed: %v\n", len(data), len(cdata))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out, _ = algo.Uncompress(cdata)
	}
	_ = out
}
