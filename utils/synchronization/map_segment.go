package synchronization

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/utils/hash"
	"sync"
)

func NewMapSegment[Key comparable, value any](segmentCount int) *MapSegment[Key, value] {
	ms := &MapSegment[Key, value]{
		segments:    map[int]*Map[Key, value]{},
		cache:       map[Key]int{},
		consistency: hash.NewConsistency(segmentCount),
	}
	for i := 0; i < segmentCount; i++ {
		ms.consistency.AddNode(i)
		ms.segments[i] = NewMap[Key, value]()
	}

	return ms
}

// MapSegment 基于分段锁实现的并发安全的字典数据结构map
type MapSegment[Key comparable, Value any] struct {
	segments    map[int]*Map[Key, Value]
	cache       map[Key]int
	consistency *hash.Consistency
	lock        sync.RWMutex
}

func (slf *MapSegment[Key, Value]) Atom(handle func(m hash.Map[Key, Value])) {
	panic("this function is currently not supported")
}

func (slf *MapSegment[Key, Value]) Set(key Key, value Value) {
	slf.lock.RLock()
	s, exist := slf.cache[key]
	slf.lock.RUnlock()
	if !exist {
		s = slf.consistency.PickNode(key)
		slf.lock.Lock()
		slf.cache[key] = s
		slf.lock.Unlock()
	}
	slf.segments[s].Set(key, value)
}

func (slf *MapSegment[Key, Value]) Get(key Key) (value Value) {
	slf.lock.RLock()
	s, exist := slf.cache[key]
	slf.lock.RUnlock()
	if !exist {
		return value
	}
	return slf.segments[s].Get(key)
}

// AtomGetSet 原子方式获取一个值并在之后进行赋值
func (slf *MapSegment[Key, Value]) AtomGetSet(key Key, handle func(value Value, exist bool) (newValue Value, isSet bool)) {
	var value Value
	slf.lock.RLock()
	s, exist := slf.cache[key]
	slf.lock.RUnlock()
	if !exist {
		if newValue, isSet := handle(value, exist); isSet {
			slf.Set(key, newValue)
		}
		return
	}
	slf.segments[s].AtomGetSet(key, handle)
}

func (slf *MapSegment[Key, Value]) Exist(key Key) bool {
	slf.lock.RLock()
	_, exist := slf.cache[key]
	slf.lock.RUnlock()
	return exist
}

func (slf *MapSegment[Key, Value]) GetExist(key Key) (value Value, exist bool) {
	slf.lock.RLock()
	s, exist := slf.cache[key]
	slf.lock.RUnlock()
	if !exist {
		return value, false
	}
	return slf.segments[s].GetExist(key)
}

func (slf *MapSegment[Key, Value]) Delete(key Key) {
	slf.lock.Lock()
	s, exist := slf.cache[key]
	delete(slf.cache, key)
	slf.lock.Unlock()
	if exist {
		slf.segments[s].Delete(key)
	}
}

func (slf *MapSegment[Key, Value]) DeleteGet(key Key) (value Value) {
	slf.lock.Lock()
	s, exist := slf.cache[key]
	delete(slf.cache, key)
	slf.lock.Unlock()
	if exist {
		return slf.segments[s].DeleteGet(key)
	}
	return
}

func (slf *MapSegment[Key, Value]) DeleteGetExist(key Key) (value Value, exist bool) {
	slf.lock.Lock()
	s, exist := slf.cache[key]
	delete(slf.cache, key)
	slf.lock.Unlock()
	if exist {
		return slf.segments[s].DeleteGetExist(key)
	}
	return value, exist
}

func (slf *MapSegment[Key, Value]) DeleteExist(key Key) bool {
	slf.lock.Lock()
	s, exist := slf.cache[key]
	delete(slf.cache, key)
	slf.lock.Unlock()
	if exist {
		return slf.segments[s].DeleteExist(key)
	}
	return false
}

func (slf *MapSegment[Key, Value]) Clear() {
	slf.lock.Lock()
	for k := range slf.cache {
		delete(slf.cache, k)
	}
	for _, m := range slf.segments {
		m.Clear()
	}
	slf.lock.Unlock()
}

func (slf *MapSegment[Key, Value]) ClearHandle(handle func(key Key, value Value)) {
	slf.lock.Lock()
	for k := range slf.cache {
		delete(slf.cache, k)
	}
	for _, m := range slf.segments {
		m.ClearHandle(handle)
	}
	slf.lock.Unlock()
}

func (slf *MapSegment[Key, Value]) Range(handle func(key Key, value Value)) {
	for _, m := range slf.segments {
		m.Range(handle)
	}
}

func (slf *MapSegment[Key, Value]) RangeSkip(handle func(key Key, value Value) bool) {
	for _, m := range slf.segments {
		m.RangeSkip(handle)
	}
}

func (slf *MapSegment[Key, Value]) RangeBreakout(handle func(key Key, value Value) bool) {
	for _, m := range slf.segments {
		if m.rangeBreakout(handle) {
			break
		}
	}
}

func (slf *MapSegment[Key, Value]) RangeFree(handle func(key Key, value Value, skip func(), breakout func())) {
	for _, m := range slf.segments {
		if m.rangeFree(handle) {
			break
		}
	}
}

func (slf *MapSegment[Key, Value]) Keys() []Key {
	var s = make([]Key, 0, len(slf.cache))
	slf.lock.RLock()
	for k, _ := range slf.cache {
		s = append(s, k)
	}
	defer slf.lock.RUnlock()
	return s
}

func (slf *MapSegment[Key, Value]) Slice() []Value {
	slf.lock.RLock()
	var s = make([]Value, 0, len(slf.cache))
	slf.lock.RUnlock()
	for _, m := range slf.segments {
		s = append(s, m.Slice()...)
	}
	return s
}

func (slf *MapSegment[Key, Value]) Map() map[Key]Value {
	slf.lock.RLock()
	var s = map[Key]Value{}
	slf.lock.RUnlock()
	for _, m := range slf.segments {
		for k, v := range m.Map() {
			s[k] = v
		}
	}
	return s
}

func (slf *MapSegment[Key, Value]) Size() int {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	return len(slf.cache)
}

// GetOne 获取一个
func (slf *MapSegment[Key, Value]) GetOne() (value Value) {
	for k, s := range slf.cache {
		return slf.segments[s].Get(k)
	}
	return value
}

func (slf *MapSegment[Key, Value]) MarshalJSON() ([]byte, error) {
	var ms struct {
		Segments     map[int]*Map[Key, Value]
		Cache        map[Key]int
		SegmentCount int
	}
	ms.Segments = slf.segments
	ms.Cache = slf.cache
	ms.SegmentCount = len(slf.segments)
	return json.Marshal(ms)
}

func (slf *MapSegment[Key, Value]) UnmarshalJSON(bytes []byte) error {
	var ms struct {
		Segments     map[int]*Map[Key, Value]
		Cache        map[Key]int
		SegmentCount int
	}
	if err := json.Unmarshal(bytes, &ms); err != nil {
		return err
	}
	slf.lock.Lock()
	slf.consistency = hash.NewConsistency(ms.SegmentCount)
	for i := 0; i < ms.SegmentCount; i++ {
		slf.consistency.AddNode(i)
	}
	slf.segments = ms.Segments
	slf.cache = ms.Cache
	slf.lock.Unlock()
	return nil
}
