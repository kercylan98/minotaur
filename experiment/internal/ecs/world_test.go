package ecs_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/ecs"
	"github.com/kercylan98/minotaur/experiment/internal/ecs/query"
	"reflect"
	"testing"
)

type Position struct {
	X float64
	Y float64
}

type Name struct {
	name string
}

type Velocity struct {
	X float64
	Y float64
}

func TestNewWorld(t *testing.T) {
	world := ecs.NewWorld()

	c1 := world.RegisterComponent(reflect.TypeOf((*Position)(nil)).Elem())
	c2 := world.RegisterComponent(reflect.TypeOf((*Velocity)(nil)).Elem())
	c3 := world.RegisterComponent(reflect.TypeOf((*Name)(nil)).Elem())

	world.Spawn(c1)         // 1
	world.Spawn(c2)         // 2
	world.Spawn(c3)         // 3
	world.Spawn(c1, c2)     // 4
	world.Spawn(c1, c3)     // 5
	world.Spawn(c2, c3)     // 6
	world.Spawn(c1, c2, c3) // 7

	q := query.And(
		query.In(c1, c2),
		query.Or(
			query.Equal(c1),
			query.NotIn(c3),
		),
	)
	t.Log(q.String())

	iterator := world.Query(q).Iterator()
	for iterator.Next() {
		t.Log(iterator.Entity())
	}

	world.GenerateDotFile("./dot.dot")
}
