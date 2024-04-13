package ecs

import "testing"

type TestComponent struct {
	x int
	y int
}

type TestComponent2 struct {
	x int
}

func BenchmarkWorldAddEntity(b *testing.B) {
	var world = NewWorld()
	var cid = Component[TestComponent](&world)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		world.CreateEntity(cid)
	}
}

func BenchmarkWorldAddEntityWithTwoComponents(b *testing.B) {
	var world = NewWorld()
	var cid1 = Component[TestComponent](&world)
	var cid2 = Component[TestComponent2](&world)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		world.CreateEntity(cid1, cid2)
	}
}
