package dispatcher

var _ Dispatcher = (*Goroutine)(nil)

type Goroutine int64

func NewGoroutine() Dispatcher {
	return new(Goroutine)
}

func (g *Goroutine) Dispatch(f func()) {
	go f()
}
