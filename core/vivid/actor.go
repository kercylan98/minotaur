package vivid

import (
	"github.com/kercylan98/minotaur/core"
)

type ActorOptionDefiner func(options *ActorOptions)
type ActorProducer func() Actor
type FunctionalActor = OnReceiveFunc

type OnReceiveFunc func(ctx ActorContext)

func (f OnReceiveFunc) OnReceive(ctx ActorContext) {
	f(ctx)
}

type Actor interface {
	OnReceive(ctx ActorContext)
}

type ActorId = core.Address
