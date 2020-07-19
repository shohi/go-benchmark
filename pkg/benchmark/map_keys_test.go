package benchmark

import (
	"testing"

	// "math"
	"reflect"
	"strconv"
)

// refer, https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map

func newMapForBenchmark(sz int) map[string]int {
	m := make(map[string]int, sz)
	for i := 0; i < sz; i++ {
		key := strconv.FormatInt(int64(i), 10)
		m[key] = i
	}

	return m
}

func BenchmarkMapKeys(b *testing.B) {
	m := newMapForBenchmark(1001000)
	cases := []struct {
		name   string
		keysFn func(map[string]int) []string
	}{
		{"traverse", keysTraverse},
		{"append", keysAppend},
		{"reflect", keysReflect},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			var keys []string
			for i := 0; i < b.N; i++ {
				keys = c.keysFn(m)
			}
			_ = keys
		})
	}

}

func keysAppend(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func keysTraverse(m map[string]int) []string {
	keys := make([]string, len(m))
	var i uint64
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}

func keysReflect(m map[string]int) []string {
	mkeys := reflect.ValueOf(m).MapKeys()
	result := make([]string, len(mkeys))
	for i, k := range mkeys {
		result[i] = k.String()
	}

	return result
}
