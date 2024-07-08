package ecs_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/ecs"
	"reflect"
	"testing"
)

func BenchmarkWorld_Spawn(b *testing.B) {
	world := ecs.NewWorld()
	cmp := world.RegisterComponent(reflect.TypeOf((*Position)(nil)).Elem())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		world.Spawn(cmp)
	}

}
