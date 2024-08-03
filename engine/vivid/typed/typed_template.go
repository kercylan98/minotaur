package main

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/internal/messages"
)

type RoomActorTyped interface {
	JoinRoom(ctx *vivid.ActorContext, message messages.Terminated)
}

type RoomActor struct {
}

func (r *RoomActor) OnReceive(ctx vivid.ActorContext) {

}
