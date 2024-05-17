package unsafevivid

import vivid "github.com/kercylan98/minotaur/vivid/vivids"

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
