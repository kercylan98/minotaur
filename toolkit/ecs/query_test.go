package ecs_test

import (
	"github.com/kercylan98/minotaur/toolkit/ecs"
	"testing"
)

type Position struct {
	X, Y float64
}

func TestQueryComponent(t *testing.T) {
	var world = ecs.NewWorld()
	var id = ecs.Component[Position](&world)

	var eid = world.CreateEntity(id)
	Change(&world, eid)
	Load(&world, eid, t)
}

func Change(world *ecs.World, eid ecs.EntityId) {
	var pos = ecs.QueryComponent[Position](world, eid)
	pos.X = 1
	pos.Y = 2
}

func Load(world *ecs.World, eid ecs.EntityId, t *testing.T) {
	var pos = ecs.QueryComponent[Position](world, eid)
	t.Log(pos.X, pos.Y)
}
