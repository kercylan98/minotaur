package dispatcher

var _ Dispatcher = (*Goroutine)(nil)

func NewGoroutineDispatcher() Dispatcher {
	return new(Goroutine)
}

type Goroutine int

func (g *Goroutine) Dispatch(f func()) {
	go f()
}
