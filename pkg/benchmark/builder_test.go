package benchmark

import (
	"testing"
	"time"
)

type TaskOptions struct {
	name     string
	timeout  time.Duration
	priority int
	cost     float64

	workFn func()
}

type TaskOption func(*TaskOptions) error

func Timeout(d time.Duration) TaskOption {
	return func(opts *TaskOptions) error {
		opts.timeout = d
		return nil
	}
}

func Name(name string) TaskOption {
	return func(opts *TaskOptions) error {
		opts.name = name
		return nil
	}
}

func Priority(p int) TaskOption {
	return func(opts *TaskOptions) error {
		opts.priority = p
		return nil
	}
}

func WorkFn(fn func()) TaskOption {
	return func(opts *TaskOptions) error {
		opts.workFn = fn
		return nil
	}
}

func Cost(c float64) TaskOption {
	return func(opts *TaskOptions) error {
		opts.cost = c
		return nil
	}
}

type Task struct {
	opts TaskOptions
}

func newTaskWithOptions(options ...TaskOption) (Task, error) {
	opts := TaskOptions{}
	for _, opt := range options {
		if err := opt(&opts); err != nil {
			return Task{}, err
		}
	}

	return Task{opts: opts}, nil
}

func newTaskWithConfig(opts TaskOptions) (Task, error) {
	return Task{opts: opts}, nil
}

func (s *Task) run() error {
	s.opts.workFn()
	// disable optimization
	_ = s.opts.cost
	return nil
}

func BenchmarkBuilder(b *testing.B) {
	opts := TaskOptions{
		name:     "task",
		priority: 1,
		workFn: func() {
			time.Sleep(2 * time.Millisecond)
		},
		cost:    10.24,
		timeout: 2 * time.Second,
	}

	cases := []struct {
		skip bool
		name string
		fn   func(*testing.B, TaskOptions)
	}{
		{false, "WithOption", benchBuilderWithOption},
		{true, "WithConfig", benchBuilderWithConfig},
	}

	for _, c := range cases {
		if c.skip {
			continue
		}
		b.Run(c.name, func(b *testing.B) {
			c.fn(b, opts)
		})
	}
}

func benchBuilderWithOption(b *testing.B, opts TaskOptions) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		task, _ := newTaskWithOptions(
			Name(opts.name),
			Timeout(opts.timeout),
			Priority(opts.priority),
			Cost(opts.cost),
			WorkFn(opts.workFn))

		_ = task.run()
	}
}
func benchBuilderWithConfig(b *testing.B, opts TaskOptions) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o := opts
		task, _ := newTaskWithConfig(o)
		_ = task.run()
	}
}
