package benchmark

import (
	"os"
	"testing"
)

func operateFile() {
	fp := "testdata/tmp_bench_file.txt"
	file, err := os.Open(fp)
	if err != nil {
		file, err = os.Create(fp)
	}

	if err != nil {
		panic(err)
	}

	_ = file
	_ = file.Close()
}

func BenchmarkFileOpenClose_ZeroSize(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		operateFile()
	}
}
