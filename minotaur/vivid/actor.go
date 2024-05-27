package vivid

type Actor interface {
	OnReceive(ctx MessageContext)
}

type FreeActor[T any] struct {
	actor T
}

func (i *FreeActor[T]) OnReceive(ctx MessageContext) {

}

func (i *FreeActor[T]) GetActor() T {
	return i.actor
}
