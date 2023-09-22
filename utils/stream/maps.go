package stream

import "github.com/kercylan98/minotaur/utils/slice"

// WithMaps 将多个 map 转换为流
func WithMaps[K comparable, V any](maps ...map[K]V) Maps[K, V] {
	var ms = make(Maps[K, V], len(maps))
	for i, m := range maps {
		ms[i] = m
	}
	return ms
}

// WithStreamMaps 将多个 map 转换为流
func WithStreamMaps[K comparable, V any](maps ...Map[K, V]) Maps[K, V] {
	var ms = make(Maps[K, V], len(maps))
	for i, m := range maps {
		ms[i] = m
	}
	return ms
}

type Maps[K comparable, V any] []Map[K, V]

// Merge 将多个 map 合并为一个 map，返回该 map 的流
//   - 当多个 map 中存在相同 key 的元素时，将会发生随机覆盖
func (slf Maps[K, V]) Merge() Map[K, V] {
	var m = make(map[K]V)
	for _, stream := range slf {
		for k, v := range stream {
			m[k] = v
		}
	}
	return m
}

// Drop 是 slice.Drop 的快捷方式
func (slf Maps[K, V]) Drop(start, n int) Maps[K, V] {
	return slice.Drop(start, n, slf)
}

// DropBy 是 slice.DropBy 的快捷方式
func (slf Maps[K, V]) DropBy(fn func(index int, value Map[K, V]) bool) Maps[K, V] {
	return slice.DropBy(slf, fn)
}

// Each 是 slice.Each 的快捷方式
func (slf Maps[K, V]) Each(abort bool, iterator func(index int, item Map[K, V]) bool) Maps[K, V] {
	slice.Each(abort, slf, iterator)
	return slf
}

// EachT 是 slice.EachT 的快捷方式
func (slf Maps[K, V]) EachT(iterator func(index int, item Map[K, V]) bool) Maps[K, V] {
	slice.EachT(slf, iterator)
	return slf
}

// EachF 是 slice.EachF 的快捷方式
func (slf Maps[K, V]) EachF(iterator func(index int, item Map[K, V]) bool) Maps[K, V] {
	slice.EachF(slf, iterator)
	return slf
}

// EachReverse 是 slice.EachReverse 的快捷方式
func (slf Maps[K, V]) EachReverse(abort bool, iterator func(index int, item Map[K, V]) bool) Maps[K, V] {
	slice.EachReverse(abort, slf, iterator)
	return slf
}

// EachReverseT 是 slice.EachReverseT 的快捷方式
func (slf Maps[K, V]) EachReverseT(iterator func(index int, item Map[K, V]) bool) Maps[K, V] {
	slice.EachReverseT(slf, iterator)
	return slf
}

// EachReverseF 是 slice.EachReverseF 的快捷方式
func (slf Maps[K, V]) EachReverseF(iterator func(index int, item Map[K, V]) bool) Maps[K, V] {
	slice.EachReverseF(slf, iterator)
	return slf
}

// FillBy 是 slice.FillBy 的快捷方式
func (slf Maps[K, V]) FillBy(fn func(index int, value Map[K, V]) Map[K, V]) Maps[K, V] {
	return slice.FillBy(slf, fn)
}

// FillByCopy 是 slice.FillByCopy 的快捷方式
func (slf Maps[K, V]) FillByCopy(fn func(index int, value Map[K, V]) Map[K, V]) Maps[K, V] {
	return slice.FillByCopy(slf, fn)
}

// FillUntil 是 slice.FillUntil 的快捷方式
func (slf Maps[K, V]) FillUntil(abort bool, fn func(index int, value Map[K, V]) (Map[K, V], bool)) Maps[K, V] {
	return slice.FillUntil(abort, slf, fn)
}

// FillUntilCopy 是 slice.FillUntilCopy 的快捷方式
func (slf Maps[K, V]) FillUntilCopy(abort bool, fn func(index int, value Map[K, V]) (Map[K, V], bool)) Maps[K, V] {
	return slice.FillUntilCopy(abort, slf, fn)
}

// FillUntilT 是 slice.FillUntilT 的快捷方式
func (slf Maps[K, V]) FillUntilT(fn func(index int, value Map[K, V]) (Map[K, V], bool)) Maps[K, V] {
	return slice.FillUntilT(slf, fn)
}

// FillUntilF 是 slice.FillUntilF 的快捷方式
func (slf Maps[K, V]) FillUntilF(fn func(index int, value Map[K, V]) (Map[K, V], bool)) Maps[K, V] {
	return slice.FillUntilF(slf, fn)
}

// FillUntilTCopy 是 slice.FillUntilTCopy 的快捷方式
func (slf Maps[K, V]) FillUntilTCopy(fn func(index int, value Map[K, V]) (Map[K, V], bool)) Maps[K, V] {
	return slice.FillUntilTCopy(slf, fn)
}

// FillUntilFCopy 是 slice.FillUntilFCopy 的快捷方式
func (slf Maps[K, V]) FillUntilFCopy(fn func(index int, value Map[K, V]) (Map[K, V], bool)) Maps[K, V] {
	return slice.FillUntilFCopy(slf, fn)
}

// Filter 是 slice.Filter 的快捷方式
func (slf Maps[K, V]) Filter(reserve bool, expression func(index int, item Map[K, V]) bool) Maps[K, V] {
	return slice.Filter(reserve, slf, expression)
}

// FilterT 是 slice.FilterT 的快捷方式
func (slf Maps[K, V]) FilterT(expression func(index int, item Map[K, V]) bool) Maps[K, V] {
	return slice.FilterT(slf, expression)
}

// FilterF 是 slice.FilterF 的快捷方式
func (slf Maps[K, V]) FilterF(expression func(index int, item Map[K, V]) bool) Maps[K, V] {
	return slice.FilterF(slf, expression)
}

// FilterCopy 是 slice.FilterCopy 的快捷方式
func (slf Maps[K, V]) FilterCopy(reserve bool, expression func(index int, item Map[K, V]) bool) Maps[K, V] {
	return slice.FilterCopy(reserve, slf, expression)
}

// FilterTCopy 是 slice.FilterTCopy 的快捷方式
func (slf Maps[K, V]) FilterTCopy(expression func(index int, item Map[K, V]) bool) Maps[K, V] {
	return slice.FilterTCopy(slf, expression)
}

// FilterFCopy 是 slice.FilterFCopy 的快捷方式
func (slf Maps[K, V]) FilterFCopy(expression func(index int, item Map[K, V]) bool) Maps[K, V] {
	return slice.FilterFCopy(slf, expression)
}

// Shuffle 是 slice.Shuffle 的快捷方式
func (slf Maps[K, V]) Shuffle() Maps[K, V] {
	return slice.Shuffle(slf)
}

// ShuffleCopy 是 slice.ShuffleCopy 的快捷方式
func (slf Maps[K, V]) ShuffleCopy() Maps[K, V] {
	return slice.ShuffleCopy(slf)
}

// UniqueBy 是 slice.UniqueBy 的快捷方式
func (slf Maps[K, V]) UniqueBy(fn func(Map[K, V]) any) Maps[K, V] {
	return slice.UniqueBy(slf, fn)
}
