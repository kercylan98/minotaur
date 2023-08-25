package counter

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"sync"
)

// NewSimpleDeduplication 创建一个简单去重计数器
//   - 该计数器不会记录每个 key 的计数，只会记录 key 的存在与否
//   - 当 key 不存在时，计数器会将 key 记录为存在，并将计数器增加特定的值
func NewSimpleDeduplication[K comparable, V generic.Number]() *SimpleDeduplication[K, V] {
	return &SimpleDeduplication[K, V]{
		r: make(map[K]struct{}),
	}
}

// SimpleDeduplication 简单去重计数器
type SimpleDeduplication[K comparable, V generic.Number] struct {
	c V
	r map[K]struct{}
	l sync.RWMutex
}

// Add 添加计数，根据 key 判断是否重复计数
func (slf *SimpleDeduplication[K, V]) Add(key K, value V) {
	slf.l.Lock()
	defer slf.l.Unlock()
	if _, exist := slf.r[key]; !exist {
		slf.r[key] = struct{}{}
		slf.c += value
	}
}

// Get 获取计数
func (slf *SimpleDeduplication[K, V]) Get() V {
	slf.l.RLock()
	defer slf.l.RUnlock()
	return slf.c
}
