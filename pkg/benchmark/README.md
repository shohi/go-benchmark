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
version: go1.12.6

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkPool/sync-8         	    3000	    377230 ns/op	  553402 B/op	    1012 allocs/op
BenchmarkPool/commons-8      	    2000	    589136 ns/op	   16336 B/op	     500 allocs/op
BenchmarkPool/ring-8         	   10000	    201252 ns/op	      55 B/op	       0 allocs/op
BenchmarkPool/channel-8      	   10000	    203237 ns/op	      54 B/op	       0 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	6.527s
```

## IntToString

```
version: go1.12.5

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkIntToString/fmt.Sprintf-8         	10000000	       123 ns/op	      40 B/op	       2 allocs/op
BenchmarkIntToString/fmt.Sprint-8          	10000000	       118 ns/op	      40 B/op	       2 allocs/op
BenchmarkIntToString/strconv.FormatInt-8   	30000000	        48.1 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	4.191s
```

## StringJoin

```
version: go1.12.6

goos: darwin
goarch: amd64
pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkStringJoin/join-8         	20000000	        69.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkStringJoin/sprintf-8      	 5000000	       348 ns/op	      96 B/op	       6 allocs/op
BenchmarkStringJoin/builder-8      	20000000	        77.6 ns/op	      24 B/op	       2 allocs/op
BenchmarkStringJoin/concat-8       	10000000	       173 ns/op	      32 B/op	       4 allocs/op
BenchmarkStringJoin/concat-oneline-8         	20000000	        63.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringJoin/buffer-8                 	20000000	        88.1 ns/op	      64 B/op	       1 allocs/op
BenchmarkStringJoin/buffer-reset-8           	30000000	        53.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringJoin/buffer-fprintf-8         	 5000000	       330 ns/op	      80 B/op	       5 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	13.943s
```

## Read

```
go version go1.13 darwin/amd64

pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkReadConcrete-8    	1000000000	         0.267 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadInterface-8   	31164573	        37.0 ns/op	     112 B/op	       1 allocs/op
BenchmarkReadNoEscape-8    	100000000	        10.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	2.951s
```

refer, https://github.com/lukechampine/noescape

## Switch

```
go version go1.13 darwin/amd64

pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkSwitchString/FromSwitch-8         	18041079	        66.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchString/FromMap-8            	42159949	        26.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchString/FromRedix-8          	 9148819	       130 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchString/FromiRedix-8         	 8326143	       143 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchInt/FromSwith-8             	94503271	        11.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchInt/FromMap-8               	49495755	        23.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchInt/FromSlice-8             	100000000	        10.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	8.813s
```

## Select

```
go version go1.13 darwin/amd64

pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkSelect/Contention-0-8         	174359425	         7.28 ns/op
BenchmarkSelect/Contention-50-8        	131247832	         7.75 ns/op
BenchmarkSelect/Contention-100-8       	155191272	         8.86 ns/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	6.979s
```

## CAS

```
go version go1.13 darwin/amd64

pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkCAS/int32-8         	98189749	        11.5 ns/op
BenchmarkCAS/int64-8         	96007542	        11.7 ns/op
PASS
ok  	github.com/shohi/go-benchmark/pkg/benchmark	3.156s
```
