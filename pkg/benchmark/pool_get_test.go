package benchmark

import (
	"context"
	"runtime"
	"sync"
	"testing"

	pool "github.com/jolestar/go-commons-pool"
)

func BenchmarkObjectGet(b *testing.B) {

	b.Run("sync", func(b *testing.B) {
		benchGetSyncPool(b)
	})

	b.Run("commons", func(b *testing.B) {
		benchGetCommonsPool(b)
	})
}

func benchGetSyncPool(b *testing.B) {
	// var s *session
	p := sync.Pool{
		New: func() interface{} {
			return newSession()
		},
	}

	b.ResetTimer()

	var s *session

	for i := 0; i < b.N; i++ {
		ss := p.Get().(*session)
		s = ss
		p.Put(ss)
		runtime.GC()
	}

	_ = s
}

func benchGetCommonsPool(b *testing.B) {
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			return newSession(), nil
		})

	ctx := context.Background()
	conf := &pool.ObjectPoolConfig{
		MaxTotal: -1,
		MaxIdle:  -1,
	}
	p := pool.NewObjectPool(ctx, factory, conf)
	defer p.Close(ctx)

	b.ResetTimer()

	var s *session

	for i := 0; i < b.N; i++ {
		o, _ := p.BorrowObject(ctx)
		s = o.(*session)

		p.ReturnObject(ctx, o)

		runtime.GC()
	}

	_ = s
}
