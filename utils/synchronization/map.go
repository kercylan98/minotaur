package synchronization

import (
	"encoding/json"
	"sync"
)

func NewMap[Key comparable, value any]() *Map[Key, value] {
	return &Map[Key, value]{
		data: make(map[Key]value),
	}
}

// Map 并发安全的字典数据结构
type Map[Key comparable, Value any] struct {
	lock sync.RWMutex
	data map[Key]Value
	atom bool
}

func (slf *Map[Key, Value]) Set(key Key, value Value) {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	slf.data[key] = value
}

func (slf *Map[Key, Value]) Get(key Key) Value {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	return slf.data[key]
}

// AtomGetSet 原子方式获取一个值并在之后进行赋值
func (slf *Map[Key, Value]) AtomGetSet(key Key, handle func(value Value, exist bool) (newValue Value, isSet bool)) {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	value, exist := slf.data[key]
	if newValue, isSet := handle(value, exist); isSet {
		slf.data[key] = newValue
	}
}

// Atom 原子操作
func (slf *Map[Key, Value]) Atom(handle func(m *Map[Key, Value])) {
	slf.lock.Lock()
	slf.atom = true
	handle(slf)
	slf.atom = false
	slf.lock.Unlock()
}

func (slf *Map[Key, Value]) Exist(key Key) bool {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	_, exist := slf.data[key]
	return exist
}

func (slf *Map[Key, Value]) GetExist(key Key) (Value, bool) {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	value, exist := slf.data[key]
	return value, exist
}

func (slf *Map[Key, Value]) Length() int {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	return len(slf.data)
}

func (slf *Map[Key, Value]) Delete(key Key) {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	delete(slf.data, key)
}

func (slf *Map[Key, Value]) DeleteGet(key Key) Value {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	v := slf.data[key]
	delete(slf.data, key)
	return v
}

func (slf *Map[Key, Value]) DeleteGetExist(key Key) (Value, bool) {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	v, exist := slf.data[key]
	delete(slf.data, key)
	return v, exist
}

func (slf *Map[Key, Value]) DeleteExist(key Key) bool {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	if _, exist := slf.data[key]; !exist {
		return exist
	}
	delete(slf.data, key)
	return true
}

func (slf *Map[Key, Value]) Clear() {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	for k := range slf.data {
		delete(slf.data, k)
	}
}

func (slf *Map[Key, Value]) ClearHandle(handle func(key Key, value Value)) {
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	for k, v := range slf.data {
		handle(k, v)
		delete(slf.data, k)
	}
}

func (slf *Map[Key, Value]) Range(handle func(key Key, value Value)) {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	for k, v := range slf.data {
		key, value := k, v
		handle(key, value)
	}
}

func (slf *Map[Key, Value]) RangeSkip(handle func(key Key, value Value) bool) {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	for k, v := range slf.data {
		key, value := k, v
		if !handle(key, value) {
			continue
		}
	}
}

func (slf *Map[Key, Value]) RangeBreakout(handle func(key Key, value Value) bool) {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
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
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
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
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	var s = make([]Key, 0, len(slf.data))
	for k, _ := range slf.data {
		s = append(s, k)
	}
	return s
}

func (slf *Map[Key, Value]) Slice() []Value {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	var s = make([]Value, 0, len(slf.data))
	for _, v := range slf.data {
		s = append(s, v)
	}
	return s
}

func (slf *Map[Key, Value]) Map() map[Key]Value {
	var m = make(map[Key]Value)
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	for k, v := range slf.data {
		m[k] = v
	}
	return m
}

func (slf *Map[Key, Value]) Size() int {
	if !slf.atom {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	return len(slf.data)
}

func (slf *Map[Key, Value]) MarshalJSON() ([]byte, error) {
	m := slf.Map()
	return json.Marshal(m)
}

func (slf *Map[Key, Value]) UnmarshalJSON(bytes []byte) error {
	var m = make(map[Key]Value)
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	if !slf.atom {
		slf.lock.Lock()
		defer slf.lock.Unlock()
	}
	slf.data = m
	return nil
}
