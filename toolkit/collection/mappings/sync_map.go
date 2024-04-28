package mappings

import (
	"encoding/json"
"github.com/kercylan98/minotaur/toolkit/collection"
"sync"
)

// NewSyncMap 创建一个 SyncMap
func NewSyncMap[K comparable, V any](source ...map[K]V) *SyncMap[K, V] {
	m := &SyncMap[K, V]{
		data: make(map[K]V),
	}
	if len(source) > 0 {
		m.data = collection.MergeMaps(source...)
	}
	return m
}

// SyncMap 是基于 sync.RWMutex 实现的线程安全的 map
//   - 适用于要考虑并发读写但是并发读写的频率不高的情况
type SyncMap[K comparable, V any] struct {
	lock sync.RWMutex
	data map[K]V
	atom bool
}

// Set 设置一个值
func (sm *SyncMap[K, V]) Set(key K, value V) {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	sm.data[key] = value
}

// Get 获取一个值
func (sm *SyncMap[K, V]) Get(key K) V {
	if !sm.atom {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
	}
	return sm.data[key]
}

// Atom 原子操作
func (sm *SyncMap[K, V]) Atom(handle func(m map[K]V)) {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	handle(sm.data)
}

// Exist 判断是否存在
func (sm *SyncMap[K, V]) Exist(key K) bool {
	if !sm.atom {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
	}
	_, exist := sm.data[key]
	return exist
}

// GetExist 获取一个值并判断是否存在
func (sm *SyncMap[K, V]) GetExist(key K) (V, bool) {
	if !sm.atom {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
	}
	value, exist := sm.data[key]
	return value, exist
}

// Delete 删除一个值
func (sm *SyncMap[K, V]) Delete(key K) {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	delete(sm.data, key)
}

// DeleteGet 删除一个值并返回
func (sm *SyncMap[K, V]) DeleteGet(key K) V {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	v := sm.data[key]
	delete(sm.data, key)
	return v
}

// DeleteGetExist 删除一个值并返回是否存在
func (sm *SyncMap[K, V]) DeleteGetExist(key K) (V, bool) {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	v, exist := sm.data[key]
	delete(sm.data, key)
	return v, exist
}

// DeleteExist 删除一个值并返回是否存在
func (sm *SyncMap[K, V]) DeleteExist(key K) bool {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	if _, exist := sm.data[key]; !exist {
		sm.lock.Unlock()
		return exist
	}
	delete(sm.data, key)
	return true
}

// Clear 清空
func (sm *SyncMap[K, V]) Clear() {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	for k := range sm.data {
		delete(sm.data, k)
	}
}

// ClearHandle 清空并处理
func (sm *SyncMap[K, V]) ClearHandle(handle func(key K, value V)) {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	for k, v := range sm.data {
		handle(k, v)
		delete(sm.data, k)
	}
}

// Range 遍历所有值，如果 handle 返回 true 则停止遍历
func (sm *SyncMap[K, V]) Range(handle func(key K, value V) bool) {
	if !sm.atom {
		sm.lock.Lock()
		defer sm.lock.Unlock()
	}
	for k, v := range sm.data {
		key, value := k, v
		if handle(key, value) {
			break
		}
	}
}

// Keys 获取所有的键
func (sm *SyncMap[K, V]) Keys() []K {
	if !sm.atom {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
	}
	var s = make([]K, 0, len(sm.data))
	for k := range sm.data {
		s = append(s, k)
	}
	return s
}

// Slice 获取所有的值
func (sm *SyncMap[K, V]) Slice() []V {
	if !sm.atom {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
	}
	var s = make([]V, 0, len(sm.data))
	for _, v := range sm.data {
		s = append(s, v)
	}
	return s
}

// Map 转换为普通 map
func (sm *SyncMap[K, V]) Map() map[K]V {
	if !sm.atom {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
	}
	var m = make(map[K]V)
	for k, v := range sm.data {
		m[k] = v
	}
	return m
}

// Size 获取数量
func (sm *SyncMap[K, V]) Size() int {
	if !sm.atom {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
	}
	return len(sm.data)
}

func (sm *SyncMap[K, V]) MarshalJSON() ([]byte, error) {
	m := sm.Map()
	return json.Marshal(m)
}

func (sm *SyncMap[K, V]) UnmarshalJSON(bytes []byte) error {
	var m = make(map[K]V)
	if !sm.atom {
		sm.lock.Lock()
		sm.lock.Unlock()
	}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	sm.data = m
	return nil
}
