package asynchronization

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/utils/hash"
)

func NewMap[Key comparable, value any]() *Map[Key, value] {
	return &Map[Key, value]{
		data: make(map[Key]value),
	}
}

// Map 非并发安全的字典数据结构
//   - 可用于对 synchronization.Map 的替代
type Map[Key comparable, Value any] struct {
	data map[Key]Value
}

func (slf *Map[Key, Value]) Set(key Key, value Value) {
	slf.data[key] = value
}

func (slf *Map[Key, Value]) Get(key Key) Value {
	return slf.data[key]
}

// AtomGetSet 原子方式获取一个值并在之后进行赋值
func (slf *Map[Key, Value]) AtomGetSet(key Key, handle func(value Value, exist bool) (newValue Value, isSet bool)) {
	value, exist := slf.data[key]
	if newValue, isSet := handle(value, exist); isSet {
		slf.data[key] = newValue
	}
}

// Atom 原子操作
func (slf *Map[Key, Value]) Atom(handle func(m hash.Map[Key, Value])) {
	handle(slf)
}

func (slf *Map[Key, Value]) Exist(key Key) bool {
	_, exist := slf.data[key]
	return exist
}

func (slf *Map[Key, Value]) GetExist(key Key) (Value, bool) {
	value, exist := slf.data[key]
	return value, exist
}

func (slf *Map[Key, Value]) Delete(key Key) {
	delete(slf.data, key)
}

func (slf *Map[Key, Value]) DeleteGet(key Key) Value {
	v := slf.data[key]
	delete(slf.data, key)
	return v
}

func (slf *Map[Key, Value]) DeleteGetExist(key Key) (Value, bool) {
	v, exist := slf.data[key]
	delete(slf.data, key)
	return v, exist
}

func (slf *Map[Key, Value]) DeleteExist(key Key) bool {
	if _, exist := slf.data[key]; !exist {
		return exist
	}
	delete(slf.data, key)
	return true
}

func (slf *Map[Key, Value]) Clear() {
	for k := range slf.data {
		delete(slf.data, k)
	}
}

func (slf *Map[Key, Value]) ClearHandle(handle func(key Key, value Value)) {
	for k, v := range slf.data {
		handle(k, v)
		delete(slf.data, k)
	}
}

func (slf *Map[Key, Value]) Range(handle func(key Key, value Value)) {
	for k, v := range slf.data {
		key, value := k, v
		handle(key, value)
	}
}

func (slf *Map[Key, Value]) RangeSkip(handle func(key Key, value Value) bool) {
	for k, v := range slf.data {
		key, value := k, v
		if !handle(key, value) {
			continue
		}
	}
}

func (slf *Map[Key, Value]) RangeBreakout(handle func(key Key, value Value) bool) {
	slf.rangeBreakout(handle)
}

func (slf *Map[Key, Value]) rangeBreakout(handle func(key Key, value Value) bool) bool {
	for k, v := range slf.data {
		key, value := k, v
		if !handle(key, value) {
			return true
		}
	}
	return false
}

func (slf *Map[Key, Value]) RangeFree(handle func(key Key, value Value, skip func(), breakout func())) {
	slf.rangeFree(handle)
}

func (slf *Map[Key, Value]) rangeFree(handle func(key Key, value Value, skip func(), breakout func())) bool {
	var skipExec, breakoutExec bool
	var skip = func() {
		skipExec = true
	}
	var breakout = func() {
		breakoutExec = true
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
	return breakoutExec
}

func (slf *Map[Key, Value]) Keys() []Key {
	var s = make([]Key, 0, len(slf.data))
	for k, _ := range slf.data {
		s = append(s, k)
	}
	return s
}

func (slf *Map[Key, Value]) Slice() []Value {
	var s = make([]Value, 0, len(slf.data))
	for _, v := range slf.data {
		s = append(s, v)
	}
	return s
}

func (slf *Map[Key, Value]) Map() map[Key]Value {
	var m = make(map[Key]Value)
	for k, v := range slf.data {
		m[k] = v
	}
	return m
}

func (slf *Map[Key, Value]) Size() int {
	return len(slf.data)
}

// GetOne 获取一个
func (slf *Map[Key, Value]) GetOne() (value Value) {
	for _, v := range slf.data {
		return v
	}
	return value
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
	slf.data = m
	return nil
}
