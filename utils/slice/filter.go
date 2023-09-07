package slice

// Filter 根据特定的表达式过滤切片成员
//   - reserve: 是否保留符合条件的成员
//   - slice: 待过滤的切片
//   - expression: 过滤表达式
//
// 这个函数的作用是遍历输入切片 slice，然后根据 expression 函数的返回值来决定是否保留每个元素。具体来说
//   - 如果 expression 返回 true 并且 reserve 也是 true，那么元素会被保留
//   - 如果 expression 返回 false 并且 reserve 是 false，那么元素也会被保留
//
// 该没有创建新的内存空间或进行元素复制，所以整个操作相当高效。同时，由于 s 和 slice 实际上共享底层的数组，因此这个函数会改变传入的 slice。如果不希望改变原切片，需要在函数调用之前手动复制一份或者使用 FilterCopy 函数。
func Filter[V any](reserve bool, slice []V, expression func(index int, item V) bool) []V {
	if len(slice) == 0 {
		return slice
	}

	var (
		i int
		j int
	)
	for i = 0; i < len(slice); i++ {
		if expression(i, slice[i]) {
			if reserve {
				slice[j] = slice[i]
				j++
			}
		} else {
			if !reserve {
				slice[j] = slice[i]
				j++
			}
		}
	}

	return slice[:j]
}

// FilterT 与 Filter 的功能相同，但是 reserve 被默认为 true
func FilterT[V any](slice []V, expression func(index int, item V) bool) []V {
	return Filter(true, slice, expression)
}

// FilterF 与 Filter 的功能相同，但是 reserve 被默认为 false
func FilterF[V any](slice []V, expression func(index int, item V) bool) []V {
	return Filter(false, slice, expression)
}

// FilterCopy 与 Filter 的功能相同，但是不会改变原切片，而是返回一个新的切片
func FilterCopy[V any](reserve bool, slice []V, expression func(index int, item V) bool) []V {
	var s = make([]V, len(slice))
	copy(s, slice)
	return Filter(reserve, s, expression)
}

// FilterCopyT 与 FilterCopy 的功能相同，但是 reserve 被默认为 true
func FilterCopyT[V any](slice []V, expression func(index int, item V) bool) []V {
	return FilterCopy(true, slice, expression)
}

// FilterCopyF 与 FilterCopy 的功能相同，但是 reserve 被默认为 false
func FilterCopyF[V any](slice []V, expression func(index int, item V) bool) []V {
	return FilterCopy(false, slice, expression)
}
