package main

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/examples/internal/room/actor"
	"github.com/kercylan98/minotaur/examples/internal/room/messages"
)

func main() {
	system := vivid.NewActorSystem()
	room := system.ActorOf(func() vivid.Actor {
		return actor.NewRoom()
	})

	system.Ask(room, &messages.JoinRoomAsk{EntityId: "user_001"})
	system.Ask(room, &messages.JoinRoomAsk{EntityId: "user_002"})

	system.ShutdownGracefully()
}
