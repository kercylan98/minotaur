package dispatcher

var _ Dispatcher = (*Goroutine)(nil)

func NewGoroutine() Dispatcher {
	return new(Goroutine)
}

type Goroutine int64

func (g *Goroutine) Dispatch(f func()) {
	go f()
}
