package stream

import (
	"github.com/kercylan98/minotaur/utils/hash"
)

// WithMap 将映射转换为流
func WithMap[K comparable, V any](m map[K]V) Map[K, V] {
	return m
}

type Map[K comparable, V any] map[K]V

// Map 返回映射
func (slf Map[K, V]) Map() map[K]V {
	return slf
}

// Keys 返回键列表
func (slf Map[K, V]) Keys() Slice[K] {
	var res = make([]K, 0, len(slf))
	for key := range slf {
		res = append(res, key)
	}
	return res
}

// Values 返回值列表
func (slf Map[K, V]) Values() Slice[V] {
	var res = make([]V, 0, len(slf))
	for _, value := range slf {
		res = append(res, value)
	}
	return res
}

// Filter 是 hash.Filter 的快捷方式
func (slf Map[K, V]) Filter(reserve bool, expression func(key K, value V) bool) Map[K, V] {
	return hash.Filter(reserve, slf, expression)
}

// FilterT 是 hash.FilterT 的快捷方式
func (slf Map[K, V]) FilterT(expression func(key K, value V) bool) Map[K, V] {
	return hash.FilterT(slf, expression)
}

// FilterF 是 hash.FilterF 的快捷方式
func (slf Map[K, V]) FilterF(expression func(key K, value V) bool) Map[K, V] {
	return hash.FilterF(slf, expression)
}

// FilterCopy 是 hash.FilterCopy 的快捷方式
func (slf Map[K, V]) FilterCopy(reserve bool, expression func(key K, value V) bool) Map[K, V] {
	return hash.FilterCopy(reserve, slf, expression)
}

// FilterTCopy 是 hash.FilterTCopy 的快捷方式
func (slf Map[K, V]) FilterTCopy(expression func(key K, value V) bool) Map[K, V] {
	return hash.FilterTCopy(slf, expression)
}

// FilterFCopy 是 hash.FilterFCopy 的快捷方式
func (slf Map[K, V]) FilterFCopy(expression func(key K, value V) bool) Map[K, V] {
	return hash.FilterFCopy(slf, expression)
}

// Chunk 是 hash.Chunk 的快捷方式
func (slf Map[K, V]) Chunk(size int) Maps[K, V] {
	chunks := hash.Chunk(slf, size)
	var res = make([]Map[K, V], 0, len(chunks))
	for i := range chunks {
		res = append(res, chunks[i])
	}
	return res
}

// Each 是 hash.Each 的快捷方式
func (slf Map[K, V]) Each(abort bool, iterator func(i int, key K, item V) bool) Map[K, V] {
	hash.Each(abort, slf, iterator)
	return slf
}

// EachT 是 hash.EachT 的快捷方式
func (slf Map[K, V]) EachT(iterator func(i int, key K, item V) bool) Map[K, V] {
	hash.EachT(slf, iterator)
	return slf
}

// EachF 是 hash.EachF 的快捷方式
func (slf Map[K, V]) EachF(iterator func(i int, key K, item V) bool) Map[K, V] {
	hash.EachF(slf, iterator)
	return slf
}
