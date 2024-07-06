package ecs_test

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/experiment/internal/ecs"
	"testing"
)

type Position struct {
	X float64
	Y float64
}

func TestSendRegisterComponentMessage(t *testing.T) {
	system := vivid.NewActorSystem()
	world := system.ActorOf(func() vivid.Actor {
		return ecs.NewWorld()
	})

	posId := ecs.SendRegisterComponentMessage[*Position](system, world)

	eid := ecs.SendCreateEntityMessage(system, world, posId)

	pos := ecs.SendLoadComponentMessage[*Position](system, world, eid)

	pos.X = 1

	pos = ecs.SendLoadComponentMessage[*Position](system, world, eid)

	t.Log(pos.X, pos.Y)

	system.Shutdown()
}
