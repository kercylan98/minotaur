package ecs_test

import (
	"github.com/kercylan98/minotaur/toolkit/ecs"
	"github.com/kercylan98/minotaur/utils/super"
	"reflect"
	"testing"
)

type NameComponent struct {
	Name string
}

type AgeComponent struct {
	Age int
}

func TestWorld_ComponentId(t *testing.T) {
	w := ecs.NewWorld()

	nameComponent := w.ComponentId(reflect.TypeOf(NameComponent{}))
	ageComponent := w.ComponentId(reflect.TypeOf(AgeComponent{}))

	ea := w.CreateEntity(nameComponent, ageComponent)
	eb := w.CreateEntity(nameComponent, ageComponent)

	ecs.QueryEntity[NameComponent](&w, ea).Name = "Alice"
	ecs.QueryEntity[NameComponent](&w, eb).Name = "Bob"

	ecs.QueryEntity[AgeComponent](&w, ea).Age = 20
	ecs.QueryEntity[AgeComponent](&w, eb).Age = 30

	t.Log(string(super.MarshalJSON(ecs.QueryEntity[NameComponent](&w, ea))))
	t.Log(string(super.MarshalJSON(ecs.QueryEntity[NameComponent](&w, eb))))
	t.Log(string(super.MarshalJSON(ecs.QueryEntity[AgeComponent](&w, ea))))
	t.Log(string(super.MarshalJSON(ecs.QueryEntity[AgeComponent](&w, eb))))

	merge := ecs.QueryEntity[struct {
		*NameComponent
		*AgeComponent
	}](&w, ea)

	t.Log(string(super.MarshalJSON(merge)))
}
