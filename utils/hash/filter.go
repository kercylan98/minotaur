package hash

// Filter 根据特定的表达式过滤哈希表成员
//   - reserve: 是否保留符合条件的成员
//   - m: 待过滤的哈希表
//   - expression: 过滤表达式
//
// 这个函数的作用是遍历输入哈希表 m，然后根据 expression 函数的返回值来决定是否保留每个元素。具体来说
//   - 如果 expression 返回 true 并且 reserve 也是 true，那么元素会被保留
//   - 如果 expression 返回 false 并且 reserve 是 false，那么元素也会被保留
//
// 该没有创建新的内存空间或进行元素复制，所以整个操作相当高效。同时，由于 m 和 map 实际上共享底层的数组，因此这个函数会改变传入的 map。如果不希望改变原哈希表，需要在函数调用之前手动复制一份或者使用 FilterCopy 函数。
func Filter[K comparable, V any](reserve bool, m map[K]V, expression func(key K, value V) bool) map[K]V {
	if len(m) == 0 {
		return m
	}

	for key, value := range m {
		if !expression(key, value) {
			delete(m, key)
		}
	}

	return m
}

// FilterT 与 Filter 的功能相同，但是 reserve 被默认为 true
func FilterT[K comparable, V any](m map[K]V, expression func(key K, value V) bool) map[K]V {
	return Filter(true, m, expression)
}

// FilterF 与 Filter 的功能相同，但是 reserve 被默认为 false
func FilterF[K comparable, V any](m map[K]V, expression func(key K, value V) bool) map[K]V {
	return Filter(false, m, expression)
}

// FilterCopy 与 Filter 的功能相同，但是不会改变原哈希表，而是返回一个新的哈希表
func FilterCopy[K comparable, V any](reserve bool, m map[K]V, expression func(key K, value V) bool) map[K]V {
	var res = map[K]V{}
	for key, value := range m {
		if expression(key, value) {
			res[key] = value
		}
	}
	return res
}

// FilterTCopy 与 FilterCopy 的功能相同，但是 reserve 被默认为 true
func FilterTCopy[K comparable, V any](m map[K]V, expression func(key K, value V) bool) map[K]V {
	return FilterCopy(true, m, expression)
}

// FilterFCopy 与 FilterCopy 的功能相同，但是 reserve 被默认为 false
func FilterFCopy[K comparable, V any](m map[K]V, expression func(key K, value V) bool) map[K]V {
	return FilterCopy(false, m, expression)
}
