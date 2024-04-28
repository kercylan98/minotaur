package sole

import (
	"github.com/kercylan98/minotaur/utils/collection"
)

// NewOnce 创建一个用于数据取值去重的结构实例
func NewOnce[V any]() *Once[V] {
	once := &Once[V]{
		r: map[any]struct{}{},
	}
	return once
}

// Once 用于数据取值去重的结构体
type Once[V any] struct {
	r map[any]struct{}
}

// Get 获取一个值，当该值已经被获取过的时候，返回 defaultValue，否则返回 value
func (slf *Once[V]) Get(key any, value V, defaultValue V) V {
	_, exist := slf.r[key]
	if exist {
		return defaultValue
	}
	slf.r[key] = struct{}{}
	return value
}

// Reset 当 key 数量大于 0 时，将会重置对应 key 的记录，否则重置所有记录
func (slf *Once[V]) Reset(key ...any) {
	if len(key) > 0 {
		for _, k := range key {
			delete(slf.r, k)
		}
		return
	}
	collection.ClearMap(slf.r)
}
