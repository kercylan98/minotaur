package ecs_test

import (
	ecs2 "github.com/kercylan98/minotaur/engine/ecs"
	"testing"
)

type Position struct {
	X float64
	Y float64
}

func TestWorld_Spawn(t *testing.T) {
	w := ecs2.NewWorld()
	pos := w.RegComponent(new(Position))

	w.Spawn(pos)

	iter := w.Query(ecs2.Equal(pos)).Iterator()
	for iter.Next() {
		t.Log(iter.Entity())

		position := iter.Get(pos).(*Position)
		position.X = 1
		position = iter.Get(pos).(*Position)
		t.Log(position)
	}
}
