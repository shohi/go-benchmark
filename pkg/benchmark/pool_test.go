package benchmark

import (
	"bytes"
	"context"
	"math/rand"
	"runtime"
	"sync"
	"testing"

	pool "github.com/jolestar/go-commons-pool"
	ipool "github.com/shohi/gocode/pkg/pool"
)

// refer, https://github.com/golang/go/issues/22950
func BenchmarkPool(b *testing.B) {
	var a, z [500]*bytes.Buffer

	b.Run("sync", func(b *testing.B) {
		benchSyncPool(a[:], z[:], b)
	})

	b.Run("commons", func(b *testing.B) {
		benchCommonsPool(a[:], z[:], b)
	})

	b.Run("ring", func(b *testing.B) {
		benchRingPool(a[:], z[:], b)
	})
}

func benchRingPool(a, z []*bytes.Buffer, b *testing.B) {
	p := ipool.NewFixedPool(1024, func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = p.Get().(*bytes.Buffer)
		}
		for j := 0; j < len(a); j++ {
			p.Put(a[j])
		}
		a = z
		runtime.GC()
	}
}

func benchSyncPool(a, z []*bytes.Buffer, b *testing.B) {
	p := sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 1024))
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = p.Get().(*bytes.Buffer)
		}
		for j := 0; j < len(a); j++ {
			p.Put(a[j])
		}
		a = z
		runtime.GC()
	}
}

func benchCommonsPool(a, z []*bytes.Buffer, b *testing.B) {
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			return bytes.NewBuffer(make([]byte, 0, 1024)), nil
		})

	ctx := context.Background()
	conf := &pool.ObjectPoolConfig{
		MaxTotal: -1,
		MaxIdle:  -1,
	}
	p := pool.NewObjectPool(ctx, factory, conf)
	defer p.Close(ctx)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			obj, err := p.BorrowObject(ctx)
			if err != nil {
				b.Fail()
			}
			a[j] = obj.(*bytes.Buffer)
		}
		for j := 0; j < len(a); j++ {
			err := p.ReturnObject(ctx, a[j])
			if err != nil {
				b.Fail()
			}
		}
		a = z
		runtime.GC()
	}
}

func BenchmarkCommons(b *testing.B) {
	type BenchObject struct {
		Num int32
	}

	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, pool.NewPooledObjectFactorySimple(func(context.Context) (interface{}, error) {
		return &BenchObject{Num: rand.Int31()}, nil
	}))
	defer p.Close(ctx)

	for i := 0; i < b.N; i++ {
		o, err := p.BorrowObject(ctx)
		if err != nil {
			b.Fail()
		}
		err = p.ReturnObject(ctx, o)
		if err != nil {
			b.Fail()
		}
		runtime.GC()
	}
}
