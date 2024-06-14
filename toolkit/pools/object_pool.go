package pools

import (
	"fmt"
	"sync"
)

// NewObjectPool 创建一个 ObjectPool
func NewObjectPool[T any](generator func() *T, releaser func(data *T)) *ObjectPool[*T] {
	if generator == nil || releaser == nil {
		panic(fmt.Errorf("generator and releaser can not be nil, generator check: %v, releaser check: %v", generator != nil, releaser != nil))
	}
	return &ObjectPool[*T]{
		releaser: releaser,
		p: sync.Pool{
			New: func() interface{} {
				return generator()
			},
		},
	}
}

// ObjectPool 基于 sync.Pool 实现的线程安全的对象池
//   - 一些高频临时生成使用的对象可以通过 ObjectPool 进行管理，例如属性计算等
type ObjectPool[T any] struct {
	p        sync.Pool
	releaser func(data T)
}

// Get 获取一个对象
func (op *ObjectPool[T]) Get() T {
	return op.p.Get().(T)
}

// Put 将使用完成的对象放回缓冲区
func (op *ObjectPool[T]) Put(data T) {
	op.releaser(data)
	op.p.Put(data)
}
