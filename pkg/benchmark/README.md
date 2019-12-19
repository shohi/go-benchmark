# Benchmark Result

`go test -bench=BenchmarkXXX -benchmem`

## defer

```
go version go1.12.4 darwin/amd64

BenchmarkDefer/MutexDeferUnlock-8         	30000000	        51.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefer/MutexUnlock-8              	100000000	        15.3 ns/op	       0 B/op	       0 allocs/op
```

## String

```
go version go1.12.4 darwin/amd64

BenchmarkString/BytesBufferOutput-100-8         	  300000	      4833 ns/op	   14896 B/op	     103 allocs/op
BenchmarkString/StringsBuilderOutput-100-8      	 2000000	       895 ns/op	     256 B/op	       5 allocs/op
BenchmarkString/BytesBufferWrite-100-8          	  500000	      2594 ns/op	    2928 B/op	       6 allocs/op
BenchmarkString/StringsBuilderWrite-100-8       	 1000000	      1746 ns/op	    2048 B/op	       8 allocs/op
BenchmarkString/GrownBytesBufferWrite-100-8     	 1000000	      2052 ns/op	    2624 B/op	       3 allocs/op
BenchmarkString/GrownStringsBuilderWrite-100-8  	 2000000	      1008 ns/op	    1920 B/op	       2 allocs/op
```

## Bytes

```
go version go1.12.4 darwin/amd64

BenchmarkBytes/StringToBytes-8         	50000000	        23.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkBytes/BytesToString-Plain-8   	100000000	        17.5 ns/op	       5 B/op	       1 allocs/op
BenchmarkBytes/BytesToString-Reflect-8 	500000000	         3.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytes/BytesToString-Cast-8    	1000000000	         2.01 ns/op	       0 B/op	       0 allocs/op
```


## Date

```
go version go1.12.4 darwin/amd64

BenchmarkDateParse/Time-8         	10000000	       144 ns/op	     112 B/op	       2 allocs/op
BenchmarkDateParse/JinzhuNow-8    	  300000	      5854 ns/op	    2087 B/op	      35 allocs/op
```

## Pool

```
go version go1.12.6 darwin/amd64

BenchmarkPool/sync-8         	    3000	    377230 ns/op	  553402 B/op	    1012 allocs/op
BenchmarkPool/commons-8      	    2000	    589136 ns/op	   16336 B/op	     500 allocs/op
BenchmarkPool/ring-8         	   10000	    201252 ns/op	      55 B/op	       0 allocs/op
BenchmarkPool/channel-8      	   10000	    203237 ns/op	      54 B/op	       0 allocs/op
```

## IntToString

```
go version go1.12.5 darwin/amd64

pkg: github.com/shohi/go-benchmark/pkg/benchmark
BenchmarkIntToString/fmt.Sprintf-8         	10000000	       123 ns/op	      40 B/op	       2 allocs/op
BenchmarkIntToString/fmt.Sprint-8          	10000000	       118 ns/op	      40 B/op	       2 allocs/op
BenchmarkIntToString/strconv.FormatInt-8   	30000000	        48.1 ns/op	      32 B/op	       1 allocs/op
```

## StringJoin

```
go version go1.12.6 darwin/amd64

BenchmarkStringJoin/join-8         	20000000	        69.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkStringJoin/sprintf-8      	 5000000	       348 ns/op	      96 B/op	       6 allocs/op
BenchmarkStringJoin/builder-8      	20000000	        77.6 ns/op	      24 B/op	       2 allocs/op
BenchmarkStringJoin/concat-8       	10000000	       173 ns/op	      32 B/op	       4 allocs/op
BenchmarkStringJoin/concat-oneline-8         	20000000	        63.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringJoin/buffer-8                 	20000000	        88.1 ns/op	      64 B/op	       1 allocs/op
BenchmarkStringJoin/buffer-reset-8           	30000000	        53.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringJoin/buffer-fprintf-8         	 5000000	       330 ns/op	      80 B/op	       5 allocs/op
```

## Read

```
go version go1.13 darwin/amd64

BenchmarkReadConcrete-8    	1000000000	         0.267 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadInterface-8   	31164573	        37.0 ns/op	     112 B/op	       1 allocs/op
BenchmarkReadNoEscape-8    	100000000	        10.4 ns/op	       0 B/op	       0 allocs/op
```

refer, https://github.com/lukechampine/noescape

## Switch

```
go version go1.13 darwin/amd64

BenchmarkSwitchString/FromSwitch-8         	18041079	        66.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchString/FromMap-8            	42159949	        26.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchString/FromRedix-8          	 9148819	       130 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchString/FromiRedix-8         	 8326143	       143 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchInt/FromSwith-8             	94503271	        11.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchInt/FromMap-8               	49495755	        23.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchInt/FromSlice-8             	100000000	        10.4 ns/op	       0 B/op	       0 allocs/op
```

## Select

```
go version go1.13 darwin/amd64

BenchmarkSelect/Contention-0-8         	178574319	         7.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelect/Contention-50-8        	134048511	         8.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelect/Contention-100-8       	128771427	         9.71 ns/op	       0 B/op	       0 allocs/op
```

## CAS

```
go version go1.13 darwin/amd64

BenchmarkCAS/int32-8         	100000000	        11.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkCAS/int64-8         	98332533	        11.5 ns/op	       0 B/op	       0 allocs/op
```

## Chan

```
go version go1.13.4 darwin/amd64

BenchmarkChanStruct-8   	 6576400	       179 ns/op	       0 B/op	       0 allocs/op
BenchmarkChanBool-8     	 6600854	       177 ns/op	       0 B/op	       0 allocs/op
BenchmarkChanInt-8      	 6620025	       180 ns/op	       0 B/op	       0 allocs/op
BenchmarkChanChan-8     	25154895	        46.1 ns/op	       0 B/op	       0 allocs/op
```

refer, https://gist.github.com/atotto/9342938

## File

```
go version go1.13.5 darwin/amd64

BenchmarkFileOpenClose_ZeroSize-8   	  106500	     10967 ns/op	     136 B/op	       3 allocs/op
```
