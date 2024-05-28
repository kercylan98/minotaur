package vivid

import "reflect"

// Behavior 表示 Actor 的行为
type Behavior interface {
	getMessageType() reflect.Type
	onHandler(ctx MessageContext, message Message)
}

type behaviorAdapter[T Message] struct {
	typ     reflect.Type
	handler func(ctx MessageContext, message T)
}

func (b *behaviorAdapter[T]) onHandler(ctx MessageContext, message Message) {
	b.handler(ctx, message.(T))
}

func (b *behaviorAdapter[T]) getMessageType() reflect.Type {
	return b.typ
}
