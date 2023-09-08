package stream

import "github.com/kercylan98/minotaur/utils/slice"

type Slice[V any] []V

// Slice 返回切片
func (slf Slice[V]) Slice() []V {
	return slf
}

// Copy 复制一份切片
func (slf Slice[V]) Copy() Slice[V] {
	return slice.Copy(slf)
}

// Zoom 是 slice.Zoom 的快捷方式
func (slf Slice[V]) Zoom(newSize int) Slice[V] {
	return slice.Zoom(newSize, slf)
}

// Chunk 是 slice.Chunk 的快捷方式
func (slf Slice[V]) Chunk(size int) Slices[V] {
	chunks := slice.Chunk(slf, size)
	result := make(Slices[V], len(chunks))
	for i, chunk := range chunks {
		result[i] = chunk
	}
	return result
}

// Drop 是 slice.Drop 的快捷方式
func (slf Slice[V]) Drop(start, n int) Slice[V] {
	return slice.Drop(start, n, slf)
}

// DropBy 是 slice.DropBy 的快捷方式
func (slf Slice[V]) DropBy(fn func(index int, value V) bool) Slice[V] {
	return slice.DropBy(slf, fn)
}

// Each 是 slice.Each 的快捷方式
func (slf Slice[V]) Each(abort bool, iterator func(index int, item V) bool) Slice[V] {
	slice.Each(abort, slf, iterator)
	return slf
}

// EachT 是 slice.EachT 的快捷方式
func (slf Slice[V]) EachT(iterator func(index int, item V) bool) Slice[V] {
	slice.EachT(slf, iterator)
	return slf
}

// EachF 是 slice.EachF 的快捷方式
func (slf Slice[V]) EachF(iterator func(index int, item V) bool) Slice[V] {
	slice.EachF(slf, iterator)
	return slf
}

// EachReverse 是 slice.EachReverse 的快捷方式
func (slf Slice[V]) EachReverse(abort bool, iterator func(index int, item V) bool) Slice[V] {
	slice.EachReverse(abort, slf, iterator)
	return slf
}

// EachReverseT 是 slice.EachReverseT 的快捷方式
func (slf Slice[V]) EachReverseT(iterator func(index int, item V) bool) Slice[V] {
	slice.EachReverseT(slf, iterator)
	return slf
}

// EachReverseF 是 slice.EachReverseF 的快捷方式
func (slf Slice[V]) EachReverseF(iterator func(index int, item V) bool) Slice[V] {
	slice.EachReverseF(slf, iterator)
	return slf
}

// FillBy 是 slice.FillBy 的快捷方式
func (slf Slice[V]) FillBy(fn func(index int, value V) V) Slice[V] {
	return slice.FillBy(slf, fn)
}

// FillByCopy 是 slice.FillByCopy 的快捷方式
func (slf Slice[V]) FillByCopy(fn func(index int, value V) V) Slice[V] {
	return slice.FillByCopy(slf, fn)
}

// FillUntil 是 slice.FillUntil 的快捷方式
func (slf Slice[V]) FillUntil(abort bool, fn func(index int, value V) (V, bool)) Slice[V] {
	return slice.FillUntil(abort, slf, fn)
}

// FillUntilCopy 是 slice.FillUntilCopy 的快捷方式
func (slf Slice[V]) FillUntilCopy(abort bool, fn func(index int, value V) (V, bool)) Slice[V] {
	return slice.FillUntilCopy(abort, slf, fn)
}

// FillUntilT 是 slice.FillUntilT 的快捷方式
func (slf Slice[V]) FillUntilT(fn func(index int, value V) (V, bool)) Slice[V] {
	return slice.FillUntilT(slf, fn)
}

// FillUntilF 是 slice.FillUntilF 的快捷方式
func (slf Slice[V]) FillUntilF(fn func(index int, value V) (V, bool)) Slice[V] {
	return slice.FillUntilF(slf, fn)
}

// FillUntilTCopy 是 slice.FillUntilTCopy 的快捷方式
func (slf Slice[V]) FillUntilTCopy(fn func(index int, value V) (V, bool)) Slice[V] {
	return slice.FillUntilTCopy(slf, fn)
}

// FillUntilFCopy 是 slice.FillUntilFCopy 的快捷方式
func (slf Slice[V]) FillUntilFCopy(fn func(index int, value V) (V, bool)) Slice[V] {
	return slice.FillUntilFCopy(slf, fn)
}

// Filter 是 slice.Filter 的快捷方式
func (slf Slice[V]) Filter(reserve bool, expression func(index int, item V) bool) Slice[V] {
	return slice.Filter(reserve, slf, expression)
}

// FilterT 是 slice.FilterT 的快捷方式
func (slf Slice[V]) FilterT(expression func(index int, item V) bool) Slice[V] {
	return slice.FilterT(slf, expression)
}

// FilterF 是 slice.FilterF 的快捷方式
func (slf Slice[V]) FilterF(expression func(index int, item V) bool) Slice[V] {
	return slice.FilterF(slf, expression)
}

// FilterCopy 是 slice.FilterCopy 的快捷方式
func (slf Slice[V]) FilterCopy(reserve bool, expression func(index int, item V) bool) Slice[V] {
	return slice.FilterCopy(reserve, slf, expression)
}

// FilterTCopy 是 slice.FilterTCopy 的快捷方式
func (slf Slice[V]) FilterTCopy(expression func(index int, item V) bool) Slice[V] {
	return slice.FilterTCopy(slf, expression)
}

// FilterFCopy 是 slice.FilterFCopy 的快捷方式
func (slf Slice[V]) FilterFCopy(expression func(index int, item V) bool) Slice[V] {
	return slice.FilterFCopy(slf, expression)
}

// Shuffle 是 slice.Shuffle 的快捷方式
func (slf Slice[V]) Shuffle() Slice[V] {
	return slice.Shuffle(slf)
}

// ShuffleCopy 是 slice.ShuffleCopy 的快捷方式
func (slf Slice[V]) ShuffleCopy() Slice[V] {
	return slice.ShuffleCopy(slf)
}

// UniqueBy 是 slice.UniqueBy 的快捷方式
func (slf Slice[V]) UniqueBy(fn func(V) any) Slice[V] {
	return slice.UniqueBy(slf, fn)
}
