package ecs

import (
	"reflect"
	"testing"
)

func BenchmarkComponentAppend(b *testing.B) {
	c := newComponent(reflect.TypeOf(Position{}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Append(1)
	}
}
