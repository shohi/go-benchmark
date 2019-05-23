package benchmark

import (
	"context"
	"runtime"
	"sync"
	"testing"

	pool "github.com/jolestar/go-commons-pool"
)

func BenchmarkStringPool(b *testing.B) {
	var result []string

	b.Run("without-pool", func(b *testing.B) {

		for i := 0; i < b.N; i++ {

			var ss []string
			for j := 0; j < 20; j++ {
				ss = append(ss, "xxxxxxxxxxxxxxxxxx")
			}

			result = ss
		}
	})

	_ = result

	b.Run("with-pool", func(b *testing.B) {
		ctx := context.Background()
		pool := newCommonsPool(func() interface{} {
			return make([]string, 0, 128)
		})

		o, _ := pool.BorrowObject(ctx)
		pool.ReturnObject(ctx, o)

		for i := 0; i < b.N; i++ {
			o, _ := pool.BorrowObject(ctx)
			ss := o.([]string)[:0]
			for j := 0; j < 20; j++ {
				ss = append(ss, "xxxxxxxxxxxxxxxxxx")
			}

			result = ss
			pool.ReturnObject(ctx, ss)
		}
	})

	b.Run("with-sync-gc", func(b *testing.B) {
		pool := newSyncPool(func() interface{} {
			return make([]string, 0, 128)
		})
		o := pool.Get()
		pool.Put(o)

		for i := 0; i < b.N; i++ {
			ss := pool.Get().([]string)[:0]
			for j := 0; j < 20; j++ {
				ss = append(ss, "xxxxxxxxxxxxxxxxxx")
			}

			result = ss
			pool.Put(ss)
			runtime.GC()
		}
	})

	_ = result
}

func newCommonsPool(fn func() interface{}) *pool.ObjectPool {
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			return fn(), nil
		})

	ctx := context.Background()
	conf := &pool.ObjectPoolConfig{
		MaxTotal: -1,
		MaxIdle:  -1,
	}
	return pool.NewObjectPool(ctx, factory, conf)
}

func newSyncPool(fn func() interface{}) *sync.Pool {
	return &sync.Pool{
		New: fn,
	}
}
