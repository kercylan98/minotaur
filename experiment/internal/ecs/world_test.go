package ecs_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/ecs"
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

	world.Spawn(c1)
	world.Spawn(c1, c3)
	world.Spawn(c3)
	world.Spawn(c1, c2)
	world.Spawn(c2)
	e := world.Spawn(c1, c2, c3)
	world.Spawn(c2, c3)

	world.DelComponent(e, c2)

	world.GenerateDotFile("./dot.dot")
}
