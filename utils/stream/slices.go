package stream

import "github.com/kercylan98/minotaur/utils/slice"

type Slices[V any] []Slice[V]

// Merge 合并为一个切片
func (slf Slices[V]) Merge() Slice[V] {
	var s = make([]V, 0)
	for _, stream := range slf {
		s = append(s, stream...)
	}
	return s
}

// Drop 的快捷方式
func (slf Slices[V]) Drop(start, n int) Slices[V] {
	return slice.Drop(start, n, slf)
}

// DropBy 的快捷方式
func (slf Slices[V]) DropBy(fn func(index int, value Slice[V]) bool) Slices[V] {
	return slice.DropBy(slf, fn)
}

// Each 的快捷方式
func (slf Slices[V]) Each(abort bool, iterator func(index int, item Slice[V]) bool) Slices[V] {
	slice.Each(abort, slf, iterator)
	return slf
}

// EachT 的快捷方式
func (slf Slices[V]) EachT(iterator func(index int, item Slice[V]) bool) Slices[V] {
	slice.EachT(slf, iterator)
	return slf
}

// EachF 的快捷方式
func (slf Slices[V]) EachF(iterator func(index int, item Slice[V]) bool) Slices[V] {
	slice.EachF(slf, iterator)
	return slf
}

// EachReverse 的快捷方式
func (slf Slices[V]) EachReverse(abort bool, iterator func(index int, item Slice[V]) bool) Slices[V] {
	slice.EachReverse(abort, slf, iterator)
	return slf
}

// EachReverseT 的快捷方式
func (slf Slices[V]) EachReverseT(iterator func(index int, item Slice[V]) bool) Slices[V] {
	slice.EachReverseT(slf, iterator)
	return slf
}

// EachReverseF 的快捷方式
func (slf Slices[V]) EachReverseF(iterator func(index int, item Slice[V]) bool) Slices[V] {
	slice.EachReverseF(slf, iterator)
	return slf
}

// FillBy 的快捷方式
func (slf Slices[V]) FillBy(fn func(index int, value Slice[V]) Slice[V]) Slices[V] {
	return slice.FillBy(slf, fn)
}

// FillByCopy 的快捷方式
func (slf Slices[V]) FillByCopy(fn func(index int, value Slice[V]) Slice[V]) Slices[V] {
	return slice.FillByCopy(slf, fn)
}

// FillUntil 的快捷方式
func (slf Slices[V]) FillUntil(abort bool, fn func(index int, value Slice[V]) (Slice[V], bool)) Slices[V] {
	return slice.FillUntil(abort, slf, fn)
}

// FillUntilCopy 的快捷方式
func (slf Slices[V]) FillUntilCopy(abort bool, fn func(index int, value Slice[V]) (Slice[V], bool)) Slices[V] {
	return slice.FillUntilCopy(abort, slf, fn)
}

// FillUntilT 的快捷方式
func (slf Slices[V]) FillUntilT(fn func(index int, value Slice[V]) (Slice[V], bool)) Slices[V] {
	return slice.FillUntilT(slf, fn)
}

// FillUntilF 的快捷方式
func (slf Slices[V]) FillUntilF(fn func(index int, value Slice[V]) (Slice[V], bool)) Slices[V] {
	return slice.FillUntilF(slf, fn)
}

// FillUntilTCopy 的快捷方式
func (slf Slices[V]) FillUntilTCopy(fn func(index int, value Slice[V]) (Slice[V], bool)) Slices[V] {
	return slice.FillUntilTCopy(slf, fn)
}

// FillUntilFCopy 的快捷方式
func (slf Slices[V]) FillUntilFCopy(fn func(index int, value Slice[V]) (Slice[V], bool)) Slices[V] {
	return slice.FillUntilFCopy(slf, fn)
}

// Filter 的快捷方式
func (slf Slices[V]) Filter(reserve bool, expression func(index int, item Slice[V]) bool) Slices[V] {
	return slice.Filter(reserve, slf, expression)
}

// FilterT 的快捷方式
func (slf Slices[V]) FilterT(expression func(index int, item Slice[V]) bool) Slices[V] {
	return slice.FilterT(slf, expression)
}

// FilterF 的快捷方式
func (slf Slices[V]) FilterF(expression func(index int, item Slice[V]) bool) Slices[V] {
	return slice.FilterF(slf, expression)
}

// FilterCopy 的快捷方式
func (slf Slices[V]) FilterCopy(reserve bool, expression func(index int, item Slice[V]) bool) Slices[V] {
	return slice.FilterCopy(reserve, slf, expression)
}

// FilterTCopy 的快捷方式
func (slf Slices[V]) FilterTCopy(expression func(index int, item Slice[V]) bool) Slices[V] {
	return slice.FilterTCopy(slf, expression)
}

// FilterFCopy 的快捷方式
func (slf Slices[V]) FilterFCopy(expression func(index int, item Slice[V]) bool) Slices[V] {
	return slice.FilterFCopy(slf, expression)
}

// Shuffle 的快捷方式
func (slf Slices[V]) Shuffle() Slices[V] {
	return slice.Shuffle(slf)
}

// ShuffleCopy 的快捷方式
func (slf Slices[V]) ShuffleCopy() Slices[V] {
	return slice.ShuffleCopy(slf)
}

// UniqueBy 的快捷方式
func (slf Slices[V]) UniqueBy(fn func(Slice[V]) any) Slices[V] {
	return slice.UniqueBy(slf, fn)
}
