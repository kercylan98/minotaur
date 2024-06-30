package vivid

import "github.com/panjf2000/ants/v2"

var defaultDispatcher Dispatcher

func init() {
	pool, err := ants.NewPool(ants.DefaultAntsPoolSize, ants.WithNonblocking(true))
	if err != nil {
		defaultDispatcher = new(goroutineDispatcher)
	}
	defaultDispatcher = &antsDispatcher{pool}
}

var _ Dispatcher = (*goroutineDispatcher)(nil)

type Dispatcher interface {
	Dispatch(f func())
}

type goroutineDispatcher int64

func (d *goroutineDispatcher) Dispatch(f func()) {
	go f()
}

type antsDispatcher struct {
	*ants.Pool
}

func (a *antsDispatcher) Dispatch(f func()) {
	if err := a.Submit(f); err != nil {
		go f()
	}
}
