package concurrent

import (
	"encoding/json"
	"sync"
)

// NewBalanceMap 创建一个并发安全且性能在普通读写和并发读写的情况下较为平衡的字典数据结构
func NewBalanceMap[Key comparable, value any](options ...BalanceMapOption[Key, value]) *BalanceMap[Key, value] {
	m := &BalanceMap[Key, value]{
		data: make(map[Key]value),
	}
	for _, option := range options {
		option(m)
	}
	return m
}

// BalanceMap 并发安全且性能在普通读写和并发读写的情况下较为平衡的字典数据结构
//   - 适用于要考虑并发读写但是并发读写的频率不高的情况
type BalanceMap[Key comparable, Value any] struct {
	lock sync.RWMutex
	data map[Key]Value
}

// Set 设置一个值
func (slf *BalanceMap[Key, Value]) Set(key Key, value Value) {
	slf.lock.Lock()
	defer slf.lock.Unlock()
	slf.data[key] = value
}

// Get 获取一个值
func (slf *BalanceMap[Key, Value]) Get(key Key) Value {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	return slf.data[key]
}

// Atom 原子操作
func (slf *BalanceMap[Key, Value]) Atom(handle func(m map[Key]Value)) {
	slf.lock.Lock()
	handle(slf.data)
	slf.lock.Unlock()
}

// Exist 判断是否存在
func (slf *BalanceMap[Key, Value]) Exist(key Key) bool {
	slf.lock.RLock()
	_, exist := slf.data[key]
	slf.lock.RUnlock()
	return exist
}

// GetExist 获取一个值并判断是否存在
func (slf *BalanceMap[Key, Value]) GetExist(key Key) (Value, bool) {
	slf.lock.RLock()
	value, exist := slf.data[key]
	slf.lock.RUnlock()
	return value, exist
}

// Delete 删除一个值
func (slf *BalanceMap[Key, Value]) Delete(key Key) {
	slf.lock.Lock()
	delete(slf.data, key)
	defer slf.lock.Unlock()
}

// DeleteGet 删除一个值并返回
func (slf *BalanceMap[Key, Value]) DeleteGet(key Key) Value {
	slf.lock.Lock()
	v := slf.data[key]
	delete(slf.data, key)
	slf.lock.Unlock()
	return v
}

// DeleteGetExist 删除一个值并返回是否存在
func (slf *BalanceMap[Key, Value]) DeleteGetExist(key Key) (Value, bool) {
	slf.lock.Lock()
	v, exist := slf.data[key]
	delete(slf.data, key)
	defer slf.lock.Unlock()
	return v, exist
}

// DeleteExist 删除一个值并返回是否存在
func (slf *BalanceMap[Key, Value]) DeleteExist(key Key) bool {
	slf.lock.Lock()
	if _, exist := slf.data[key]; !exist {
		slf.lock.Unlock()
		return exist
	}
	delete(slf.data, key)
	slf.lock.Unlock()
	return true
}

// Clear 清空
func (slf *BalanceMap[Key, Value]) Clear() {
	slf.lock.Lock()
	for k := range slf.data {
		delete(slf.data, k)
	}
	slf.lock.Unlock()
}

// ClearHandle 清空并处理
func (slf *BalanceMap[Key, Value]) ClearHandle(handle func(key Key, value Value)) {
	slf.lock.Lock()
	for k, v := range slf.data {
		handle(k, v)
		delete(slf.data, k)
	}
	slf.lock.Unlock()
}

// Range 遍历所有值，如果 handle 返回 true 则停止遍历
func (slf *BalanceMap[Key, Value]) Range(handle func(key Key, value Value) bool) {
	slf.lock.Lock()
	for k, v := range slf.data {
		key, value := k, v
		if handle(key, value) {
			break
		}
	}
	slf.lock.Unlock()
}

// Keys 获取所有的键
func (slf *BalanceMap[Key, Value]) Keys() []Key {
	slf.lock.RLock()
	var s = make([]Key, 0, len(slf.data))
	for k := range slf.data {
		s = append(s, k)
	}
	slf.lock.RUnlock()
	return s
}

// Slice 获取所有的值
func (slf *BalanceMap[Key, Value]) Slice() []Value {
	slf.lock.RLock()
	var s = make([]Value, 0, len(slf.data))
	for _, v := range slf.data {
		s = append(s, v)
	}
	slf.lock.RUnlock()
	return s
}

// Map 转换为普通 map
func (slf *BalanceMap[Key, Value]) Map() map[Key]Value {
	slf.lock.RLock()
	var m = make(map[Key]Value)
	for k, v := range slf.data {
		m[k] = v
	}
	slf.lock.RUnlock()
	return m
}

// Size 获取数量
func (slf *BalanceMap[Key, Value]) Size() int {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	return len(slf.data)
}

func (slf *BalanceMap[Key, Value]) MarshalJSON() ([]byte, error) {
	m := slf.Map()
	return json.Marshal(m)
}

func (slf *BalanceMap[Key, Value]) UnmarshalJSON(bytes []byte) error {
	var m = make(map[Key]Value)
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	slf.lock.Lock()
	slf.data = m
	slf.lock.Unlock()
	return nil
}
