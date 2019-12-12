package benchmark

import "testing"

// refer, https://gist.github.com/atotto/9342938
func BenchmarkChanStruct(b *testing.B) {
	ch := make(chan struct{})
	go func() {
		for {
			<-ch
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
}

func BenchmarkChanBool(b *testing.B) {
	ch := make(chan bool)
	go func() {
		for {
			<-ch
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- true
	}
}

func BenchmarkChanInt(b *testing.B) {
	ch := make(chan int)
	go func() {
		for {
			<-ch
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- 1
	}
}

func BenchmarkChanChan(b *testing.B) {
	ch := make(chan int, 10)
	for i := 0; i < b.N; i++ {
		ch <- 1
		_ = <-ch
	}
}
