package benchmark

import (
	"context"
	"testing"
	"time"
)

func BenchmarkContext(b *testing.B) {
	b.Run("WithCancel", benchContextWithCancel)
	b.Run("WithoutCancel", benchContextWithoutCancel)
}

func benchContextWithCancel(b *testing.B) {
	bgCtx := context.Background()

	var ch <-chan struct{}

	for i := 0; i < b.N; i++ {
		ctx, cancelFn := context.WithTimeout(bgCtx, 10*time.Millisecond)
		ch = ctx.Done()
		cancelFn()
	}

	_ = ch
}

func benchContextWithoutCancel(b *testing.B) {
	bgCtx := context.Background()

	var ch <-chan struct{}

	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(bgCtx, 10*time.Millisecond)
		ch = ctx.Done()
		_ = cancel
	}

	_ = ch
}
