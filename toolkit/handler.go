package toolkit

type Handler interface {
	Handle()
}

type FunctionalHandler func()

func (f FunctionalHandler) Handle() {
	f()
}
