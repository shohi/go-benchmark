package benchmark

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func bytesToStringPlain(b []byte) string {
	return string(b)
}

// refer, https://github.com/golang/go/issues/25484
func bytesToStringReflect(b []byte) string {
	var s string
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = slice.Data
	str.Len = slice.Len
	runtime.KeepAlive(&b) // this line is essential.

	return s
}

// refer, strings.Builder#String()
func bytesToStringCast(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func stringToBytes(s string) []byte {
	return []byte(s)
}

func BenchmarkBytes(b *testing.B) {
	b.Run("StringToBytes", func(b *testing.B) {
		testStr := "hello"
		var result []byte
		for i := 0; i < b.N; i++ {
			result = stringToBytes(testStr)
		}
		_ = result
	})

	cases := []struct {
		name string
		fn   func([]byte) string
	}{
		{"Plain", bytesToStringPlain},
		{"Reflect", bytesToStringReflect},
		{"Cast", bytesToStringCast},
	}

	for _, c := range cases {
		b.Run(fmt.Sprintf("BytesToString-%v", c.name), func(b *testing.B) {
			testBytes := []byte("hello")
			var result string
			for i := 0; i < b.N; i++ {
				result = c.fn(testBytes)
			}
			_ = result
		})
	}
}
