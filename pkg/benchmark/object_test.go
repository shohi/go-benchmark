package benchmark

import (
	"context"
	"runtime"
	"sync"
	"testing"

	pool "github.com/jolestar/go-commons-pool"
)

type session struct {
	name string

	url    string
	header map[string]string

	ptr interface{}
	typ string
}

func newSession() *session {
	var flag = 0
	return &session{
		name:   "session",
		url:    "http://localhost/path/to/test",
		header: make(map[string]string, 3),
		ptr:    &flag,
		typ:    "sync",
	}
}

func BenchmarkObjectPooled(b *testing.B) {
	var a, z [500]*session

	b.Run("sync", func(b *testing.B) {
		benchObjectSyncPool(a[:], z[:], b)
	})

	b.Run("commons", func(b *testing.B) {
		benchObjectCommonsPool(a[:], z[:], b)
	})
}

func benchObjectSyncPool(a, z []*session, b *testing.B) {
	// var s *session
	p := sync.Pool{
		New: func() interface{} {
			return newSession()
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = p.Get().(*session)
		}
		for j := 0; j < len(a); j++ {
			p.Put(a[j])
		}

		a = z
		runtime.GC()
	}
}

func benchObjectCommonsPool(a, z []*session, b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			o, _ := p.BorrowObject(ctx)
			a[j] = o.(*session)
		}
		for j := 0; j < len(a); j++ {
			p.ReturnObject(ctx, a[j])
		}

		a = z
		runtime.GC()
	}
}
