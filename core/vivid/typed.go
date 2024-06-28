package vivid

type ActorTyped interface {
	Close()
}

func NewTyped[T Message]() Typed[T] {
	return make(chan T)
}

type Typed[T Message] chan T

func (t Typed[T]) Close() {
	close(t)
}
