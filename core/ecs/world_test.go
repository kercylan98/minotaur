package ecs_test

import (
	"github.com/kercylan98/minotaur/core/ecs"
	"testing"
)

type Position struct {
	X float64
	Y float64
}

func TestWorld_Spawn(t *testing.T) {
	w := ecs.NewWorld()
	pos := w.RegComponent(new(Position))

	w.Spawn(pos)

	iter := w.Query(ecs.Equal(pos)).Iterator()
	for iter.Next() {
		t.Log(iter.Entity())

		position := iter.Get(pos).(*Position)
		position.X = 1
		position = iter.Get(pos).(*Position)
		t.Log(position)
	}
}
