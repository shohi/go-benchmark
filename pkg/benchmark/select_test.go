package benchmark

import (
	"math/rand"
	"testing"
)

func BenchmarkSelect(b *testing.B) {
	cases := []struct {
		name string
		p    int
	}{
		{"Contention-0", 0},
		{"Contention-50", 50},
		{"Contention-100", 100},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			benchSelect(b, c.p)
		})
	}
}

func benchSelect(b *testing.B, p int) {
	ch := make(chan bool, 1)
	go func() {
		for {
			p := rand.Intn(100)
			if p < 0 {
				select {
				case ch <- true:
				default:
				}
			}
		}
	}()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		select {
		case <-ch:
			// do nothing
		default:
			// do nothing
		}
	}
}
