package synchronization

import "sync"

func NewMap[Key comparable, value any]() *Map[Key, value] {
	return &Map[Key, value]{
		data: make(map[Key]value),
	}
}

// Map 并发安全的字典数据结构
type Map[Key comparable, Value any] struct {
	lock sync.RWMutex
	data map[Key]Value
}

func (slf *Map[Key, Value]) Set(key Key, value Value) {
	slf.lock.Lock()
	defer slf.lock.Unlock()
	slf.data[key] = value
}

func (slf *Map[Key, Value]) Get(key Key) Value {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	return slf.data[key]
}

func (slf *Map[Key, Value]) Exist(key Key) bool {
	slf.lock.RLock()
	_, exist := slf.data[key]
	slf.lock.RUnlock()
	return exist
}

func (slf *Map[Key, Value]) GetExist(key Key) (Value, bool) {
	slf.lock.RLock()
	value, exist := slf.data[key]
	slf.lock.RUnlock()
	return value, exist
}

func (slf *Map[Key, Value]) Length() int {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	return len(slf.data)
}

func (slf *Map[Key, Value]) Delete(key Key) {
	slf.lock.Lock()
	defer slf.lock.Unlock()
	delete(slf.data, key)
}

func (slf *Map[Key, Value]) DeleteGet(key Key) Value {
	slf.lock.Lock()
	defer slf.lock.Unlock()
	v := slf.data[key]
	delete(slf.data, key)
	return v
}

func (slf *Map[Key, Value]) DeleteGetExist(key Key) (Value, bool) {
	slf.lock.Lock()
	defer slf.lock.Unlock()
	v, exist := slf.data[key]
	delete(slf.data, key)
	return v, exist
}

func (slf *Map[Key, Value]) DeleteExist(key Key) bool {
	slf.lock.Lock()
	if _, exist := slf.data[key]; !exist {
		slf.lock.Unlock()
		return exist
	}
	delete(slf.data, key)
	slf.lock.Unlock()
	return true
}

func (slf *Map[Key, Value]) Clear() {
	slf.lock.Lock()
	defer slf.lock.Unlock()
	for k := range slf.data {
		delete(slf.data, k)
	}
}

func (slf *Map[Key, Value]) Range(handle func(key Key, value Value)) {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	for k, v := range slf.data {
		key, value := k, v
		handle(key, value)
	}
}

func (slf *Map[Key, Value]) RangeSkip(handle func(key Key, value Value) bool) {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	for k, v := range slf.data {
		key, value := k, v
		if !handle(key, value) {
			continue
		}
	}
}

func (slf *Map[Key, Value]) RangeBreakout(handle func(key Key, value Value) bool) {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	for k, v := range slf.data {
		key, value := k, v
		if !handle(key, value) {
			break
		}
	}
}

func (slf *Map[Key, Value]) RangeFree(handle func(key Key, value Value, skip func(), breakout func())) {
	var skipExec, breakoutExec bool
	var skip = func() {
		skipExec = true
	}
	var breakout = func() {
		breakoutExec = true
	}
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	for k, v := range slf.data {
		key, value := k, v
		handle(key, value, skip, breakout)
		if skipExec {
			continue
		}
		if breakoutExec {
			break
		}
	}
}

func (slf *Map[Key, Value]) Keys() []Key {
	slf.lock.RLock()
	var s = make([]Key, 0, len(slf.data))
	for k, _ := range slf.data {
		s = append(s, k)
	}
	slf.lock.RUnlock()
	return s
}

func (slf *Map[Key, Value]) Slice() []Value {
	slf.lock.RLock()
	var s = make([]Value, 0, len(slf.data))
	for _, v := range slf.data {
		s = append(s, v)
	}
	slf.lock.RUnlock()
	return s
}

func (slf *Map[Key, Value]) Map() map[Key]Value {
	var m = make(map[Key]Value)
	slf.lock.RLock()
	for k, v := range slf.data {
		m[k] = v
	}
	slf.lock.RUnlock()
	return m
}
