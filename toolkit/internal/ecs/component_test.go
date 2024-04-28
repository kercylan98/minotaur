package ecs

import (
	"reflect"
	"testing"
)

func TestComponent_Append(t *testing.T) {
	var c = newComponent(reflect.TypeOf(Position{}))
	ExceptNum[int](t, c.Len(), 0)
	ExceptNum[int](t, c.Cap(), 0)

	for i := 0; i < 100; i++ {
		c.Append(1)
	}

	ExceptNum[int](t, c.Len(), 100)
	ExceptNum[int](t, c.Cap(), 128)

	c.Append(100)
	ExceptNum[int](t, c.Len(), 200)
	ExceptNum[int](t, c.Cap(), 256)
}

func TestComponent_Delete(t *testing.T) {
	var c = newComponent(reflect.TypeOf(Position{}))
	c.Append(10)
	c.Delete(2)

	ExceptNum[int](t, c.Len(), 9)
	ExceptNum[int](t, c.Cap(), 16)
}

func TestComponent_DeleteRange(t *testing.T) {
	var c = newComponent(reflect.TypeOf(Position{}))
	c.Append(10)
	c.DeleteRange(2, 5)

	ExceptNum[int](t, c.Len(), 7)
	ExceptNum[int](t, c.Cap(), 16)
}

func TestComponent_Get(t *testing.T) {
	var c = newComponent(reflect.TypeOf(Position{}))
	c.Append(10)

	for i := 0; i < 10; i++ {
		_ = c.Get(i)
	}
}

func TestComponent_Len(t *testing.T) {
	var c = newComponent(reflect.TypeOf(Position{}))
	ExceptNum[int](t, c.Len(), 0)
	c.Append(10)
	ExceptNum[int](t, c.Len(), 10)
}

func TestComponent_Cap(t *testing.T) {
	var c = newComponent(reflect.TypeOf(Position{}))
	ExceptNum[int](t, c.Cap(), 0)
	c.Append(10)
	ExceptNum[int](t, c.Cap(), 16)
}
