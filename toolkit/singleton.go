package toolkit

import "sync"

func NewInertiaSingleton[T any](constructor func() T) *InertiaSingleton[T] {
	return &InertiaSingleton[T]{
		constructor: constructor,
	}
}

type InertiaSingleton[T any] struct {
	once        sync.Once
	v           T
	constructor func() T
}

func (s *InertiaSingleton[T]) Get() T {
	s.once.Do(func() {
		s.v = s.constructor()
	})
	return s.v
}
