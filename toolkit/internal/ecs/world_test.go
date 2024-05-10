package ecs

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit"
	"testing"
)

func TestWorld_AddComponent(t *testing.T) {
	var w = NewWorld()
	posId := GetComponentId[Position](&w)
	velId := GetComponentId[Velocity](&w)

	e := w.CreateEntity(posId)
	pos := QueryComponent[Position](&w, e)
	pos.X = 1

	t.Log(QueryComponent[Position](&w, e))

	w.AddComponent(e, velId)
	vel := QueryComponent[Velocity](&w, e)
	vel.X = 2

	t.Log(QueryComponent[Velocity](&w, e))

	w.RemoveComponent(e, posId)
	t.Log(QueryComponent[Position](&w, e))
	t.Log(QueryComponent[Velocity](&w, e))

	fmt.Println(string(toolkit.MarshalIndentJSON(NewStatus(&w), "", "  ")))
}
