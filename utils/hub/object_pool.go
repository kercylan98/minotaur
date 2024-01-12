package hub

import (
	"errors"
	"sync"
)

// NewObjectPool 创建一个线程安全的对象缓冲池
//   - 通过 Get 获取一个对象，如果缓冲区内存在可用对象则直接返回，否则新建一个进行返回
//   - 通过 Release 将使用完成的对象放回缓冲区，超出缓冲区大小的对象将被放弃
func NewObjectPool[T any](generator func() *T, releaser func(data *T)) *ObjectPool[*T] {
	if generator == nil || releaser == nil {
		panic(errors.New("generator or releaser is nil"))
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

// ObjectPool 线程安全的对象缓冲池
//   - 一些高频临时生成使用的对象可以通过 ObjectPool 进行管理，例如属性计算等
//   - 缓冲区内存在可用对象时直接返回，否则新建一个进行返回
//   - 通过 Release 将使用完成的对象放回缓冲区，超出缓冲区大小的对象将被放弃
type ObjectPool[T any] struct {
	p        sync.Pool
	releaser func(data T)
}

// Get 获取一个对象
func (op *ObjectPool[T]) Get() T {
	return op.p.Get().(T)
}

// Release 将使用完成的对象放回缓冲区
func (op *ObjectPool[T]) Release(data T) {
	op.releaser(data)
	op.p.Put(data)
}
