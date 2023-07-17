package hash

import (
	"sort"
)

func NewSortMap[K comparable, V any]() *SortMap[K, V] {
	return &SortMap[K, V]{
		m: map[K]V{},
		s: map[int]K{},
		r: map[K]int{},
	}
}

// SortMap 有序的 map 实现
type SortMap[K comparable, V any] struct {
	i int
	m map[K]V
	s map[int]K
	r map[K]int
}

func (slf *SortMap[K, V]) Set(key K, value V) {
	if i, exist := slf.r[key]; exist {
		slf.s[i] = key
		slf.m[key] = value
	} else {
		slf.m[key] = value
		slf.s[slf.i] = key
		slf.r[key] = slf.i
		slf.i++
	}
}

func (slf *SortMap[K, V]) Del(key K) {
	if _, exist := slf.m[key]; exist {
		delete(slf.s, slf.r[key])
		delete(slf.r, key)
		delete(slf.m, key)
	}
}

func (slf *SortMap[K, V]) Get(key K) V {
	v := slf.m[key]
	return v
}

func (slf *SortMap[K, V]) For(handle func(key K, value V) bool) {
	for k, v := range slf.m {
		if !handle(k, v) {
			break
		}
	}
}

func (slf *SortMap[K, V]) ForSort(handle func(key K, value V) bool) {
	var indexes []int
	for i := range slf.s {
		indexes = append(indexes, i)
	}
	sort.Ints(indexes)
	for _, i := range indexes {
		k := slf.s[i]
		if !handle(k, slf.m[k]) {
			break
		}
	}
}

func (slf *SortMap[K, V]) ToMap() map[K]V {
	var m = make(map[K]V)
	for k, v := range slf.m {
		m[k] = v
	}
	return m
}

func (slf *SortMap[K, V]) ToSlice() []V {
	var s = make([]V, 0, len(slf.m))
	for _, v := range slf.m {
		s = append(s, v)
	}
	return s
}

func (slf *SortMap[K, V]) ToSliceSort() []V {
	var indexes []int
	for i := range slf.s {
		indexes = append(indexes, i)
	}
	sort.Ints(indexes)
	var result []V
	for _, i := range indexes {
		k := slf.s[i]
		result = append(result, slf.m[k])
	}
	return result
}

func (slf *SortMap[K, V]) KeyToSlice() []K {
	var s = make([]K, 0, len(slf.m))
	for k := range slf.m {
		s = append(s, k)
	}
	return s
}
