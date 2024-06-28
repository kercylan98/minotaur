package vivid

var defaultDispatcher Dispatcher = new(goroutineDispatcher)

var _ Dispatcher = (*goroutineDispatcher)(nil)

type Dispatcher interface {
	Dispatch(f func())
}

type goroutineDispatcher int64

func (d *goroutineDispatcher) Dispatch(f func()) {
	go f()
}
