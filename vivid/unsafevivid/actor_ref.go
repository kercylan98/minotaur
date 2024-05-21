package unsafevivid

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	vivid "github.com/kercylan98/minotaur/vivid/vivids"
)

func NewActorRef(system *ActorSystem, actorId vivid.ActorId) *ActorRef {
	return &ActorRef{
		system:  system,
		actorId: actorId,
	}
}

type ActorRef struct {
	system  *ActorSystem
	actorId vivid.ActorId
}

func (a *ActorRef) GetId() vivid.ActorId {
	return a.actorId
}

func (a *ActorRef) Tell(msg vivid.Message, opts ...vivid.MessageOption) error {
	return a.system.Tell(a.GetId(), msg, opts...)
}

func (a *ActorRef) Ask(msg vivid.Message, opts ...vivid.MessageOption) (vivid.Message, error) {
	return a.system.Ask(a.GetId(), msg, opts...)
}

func (a *ActorRef) Subscribe(ctx *ActorContext, event vivid.Event) {
	// 订阅应该只是一个接口，实际上是通过 Ask 或者 Tell 进行订阅，这样可以实现分布式事件
	err := a.Tell(&SubscribeEventMessage{
		Subscriber: ctx.GetActorId(),
		Event:      event,
	})

	if err != nil {
		log.Error("subscribe event failed", err)
	}
}
