package benchmark

import "testing"

func BenchmarkInterfaceCast(b *testing.B) {
	b.Run("plain", benchValuePlain)
	b.Run("cast", benchValueCast)
}

func benchValuePlain(b *testing.B) {
	var s string = "value"
	var result string
	for i := 0; i < b.N; i++ {
		result = s
	}
	_ = result
}

func benchValueCast(b *testing.B) {
	var s interface{}
	s = "value"
	var result string
	for i := 0; i < b.N; i++ {
		result = s.(string)
	}
	_ = result
}
