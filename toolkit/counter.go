package toolkit

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"sync"
)

// Counter 树计数器
type Counter[T constraints.Number] struct {
	v  T
	p  *Counter[T]
	rw sync.RWMutex
}

// Sub 生成子计数器，子计数器的增减操作会影响父计数器
func (c *Counter[T]) Sub() *Counter[T] {
	return &Counter[T]{
		p: c,
	}
}

// Add 增加计数
func (c *Counter[T]) Add(delta T) {
	c.rw.Lock()
	c.v += delta
	c.rw.Unlock()
	if c.p != nil {
		c.p.Add(delta)
	}
}

// Val 获取计数
func (c *Counter[T]) Val() T {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.v
}
