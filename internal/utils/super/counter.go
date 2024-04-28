package super

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"sync"
)

type Counter[T generic.Number] struct {
	v  T
	p  *Counter[T]
	rw sync.RWMutex
}

func (c *Counter[T]) Sub() *Counter[T] {
	return &Counter[T]{
		p: c,
	}
}

func (c *Counter[T]) Add(delta T) {
	c.rw.Lock()
	c.v += delta
	c.rw.Unlock()
	if c.p != nil {
		c.p.Add(delta)
	}
}

func (c *Counter[T]) Val() T {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.v
}
