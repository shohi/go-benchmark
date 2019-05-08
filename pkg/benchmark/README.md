# Benchmark Result

## defer

```
version: go1.12.4

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkDefer/MutexDeferUnlock-8         	30000000	        51.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefer/MutexUnlock-8              	100000000	        15.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	3.146s

```

## String

```
version: go1.12.4

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkString/BytesBufferOutput-100-8         	  300000	      4833 ns/op	   14896 B/op	     103 allocs/op
BenchmarkString/StringsBuilderOutput-100-8      	 2000000	       895 ns/op	     256 B/op	       5 allocs/op
BenchmarkString/BytesBufferWrite-100-8          	  500000	      2594 ns/op	    2928 B/op	       6 allocs/op
BenchmarkString/StringsBuilderWrite-100-8       	 1000000	      1746 ns/op	    2048 B/op	       8 allocs/op
BenchmarkString/GrownBytesBufferWrite-100-8     	 1000000	      2052 ns/op	    2624 B/op	       3 allocs/op
BenchmarkString/GrownStringsBuilderWrite-100-8  	 2000000	      1008 ns/op	    1920 B/op	       2 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	226.269

```