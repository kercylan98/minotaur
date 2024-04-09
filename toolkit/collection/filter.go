package collection

// FilterOutByIndices 过滤切片中特定索引的元素，返回过滤后的切片
func FilterOutByIndices[S []V, V any](slice S, indices ...int) S {
	if slice == nil || len(slice) == 0 || len(indices) == 0 {
		return slice
	}

	excludeMap := make(map[int]bool)
	for _, ex := range indices {
		if ex >= 0 && ex < len(slice) {
			excludeMap[ex] = true
		}
	}

	if len(excludeMap) == 0 {
		return slice
	}

	validElements := make([]V, 0, len(slice)-len(excludeMap))
	for i, v := range slice {
		if !excludeMap[i] {
			validElements = append(validElements, v)
		}
	}

	return validElements
}

// FilterOutByCondition 过滤切片中符合条件的元素，返回过滤后的切片
//   - condition 的返回值为 true 时，将会过滤掉该元素
func FilterOutByCondition[S ~[]V, V any](slice S, condition func(v V) bool) S {
	if slice == nil {
		return nil
	}
	if condition == nil {
		return slice
	}

	validElements := make([]V, 0, len(slice))
	for _, v := range slice {
		if !condition(v) {
			validElements = append(validElements, v)
		}
	}

	return validElements
}

// FilterOutByKey 过滤 map 中特定的 key，返回过滤后的 map
func FilterOutByKey[M ~map[K]V, K comparable, V any](m M, key K) M {
	if m == nil {
		return nil
	}

	validMap := make(M, len(m)-1)
	for k, v := range m {
		if k != key {
			validMap[k] = v
		}
	}

	return validMap
}

// FilterOutByValue 过滤 map 中特定的 value，返回过滤后的 map
func FilterOutByValue[M ~map[K]V, K comparable, V any](m M, value V, handler ComparisonHandler[V]) M {
	if m == nil {
		return nil
	}

	validMap := make(M, len(m))
	for k, v := range m {
		if !handler(value, v) {
			validMap[k] = v
		}
	}

	return validMap
}

// FilterOutByKeys 过滤 map 中多个 key，返回过滤后的 map
func FilterOutByKeys[M ~map[K]V, K comparable, V any](m M, keys ...K) M {
	if m == nil {
		return nil
	}
	if len(keys) == 0 {
		return m
	}

	validMap := make(M, len(m)-len(keys))
	for k, v := range m {
		if !InSlice(keys, k, func(source, target K) bool {
			return source == target
		}) {
			validMap[k] = v
		}
	}

	return validMap
}

// FilterOutByValues 过滤 map 中多个 values，返回过滤后的 map
func FilterOutByValues[M ~map[K]V, K comparable, V any](m M, values []V, handler ComparisonHandler[V]) M {
	if m == nil {
		return nil
	}
	if len(values) == 0 {
		return m
	}

	validMap := make(M, len(m))
	for k, v := range m {
		if !InSlice(values, v, handler) {
			validMap[k] = v
		}
	}

	return validMap
}

// FilterOutByMap 过滤 map 中符合条件的元素，返回过滤后的 map
//   - condition 的返回值为 true 时，将会过滤掉该元素
func FilterOutByMap[M ~map[K]V, K comparable, V any](m M, condition func(k K, v V) bool) M {
	if m == nil {
		return nil
	}
	if condition == nil {
		return m
	}

	validMap := make(M, len(m))
	for k, v := range m {
		if !condition(k, v) {
			validMap[k] = v
		}
	}

	return validMap
}
