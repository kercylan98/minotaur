package vivid

type Dispatcher interface {
	Dispatch(f func())
}

type asyncDispatcher int64

func (d *asyncDispatcher) Dispatch(f func()) {
	go f()
}
