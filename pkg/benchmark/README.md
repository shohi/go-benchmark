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

## Bytes

```
version: go1.12.4

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkBytes/StringToBytes-8         	50000000	        23.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkBytes/BytesToString-Plain-8   	100000000	        17.5 ns/op	       5 B/op	       1 allocs/op
BenchmarkBytes/BytesToString-Reflect-8 	500000000	         3.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytes/BytesToString-Cast-8    	1000000000	         2.01 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	7.041s

```


## Date

```
version: go1.12.4

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkDateParse/Time-8         	10000000	       144 ns/op	     112 B/op	       2 allocs/op
BenchmarkDateParse/JinzhuNow-8    	  300000	      5854 ns/op	    2087 B/op	      35 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	3.413s

```

## Pool

```
version: go1.12.4

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkPool/sync-8         	    3000	    370376 ns/op	  553408 B/op	    1012 allocs/op
BenchmarkPool/commons-8      	    3000	    560896 ns/op	   16223 B/op	     500 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	3.037s

```
