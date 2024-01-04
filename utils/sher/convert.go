package sher

// ConvertSliceToAny 将切片转换为任意类型的切片
func ConvertSliceToAny[S ~[]V, V any](s S) []any {
	var r = make([]any, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}

// ConvertSliceToIndexMap 将切片转换为索引为键的映射
func ConvertSliceToIndexMap[S ~[]V, V any](s S) map[int]V {
	var r = make(map[int]V, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}

// ConvertSliceToMap 将切片转换为值为键的映射
func ConvertSliceToMap[S ~[]V, V comparable](s S) map[V]struct{} {
	var r = make(map[V]struct{}, len(s))
	for _, v := range s {
		r[v] = struct{}{}
	}
	return r
}

// ConvertSliceToBoolMap 将切片转换为值为键的映射
func ConvertSliceToBoolMap[S ~[]V, V comparable](s S) map[V]bool {
	var r = make(map[V]bool, len(s))
	for _, v := range s {
		r[v] = true
	}
	return r
}

// ConvertMapKeysToSlice 将映射的键转换为切片
func ConvertMapKeysToSlice[M ~map[K]V, K comparable, V any](m M) []K {
	var r = make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// ConvertMapValuesToSlice 将映射的值转换为切片
func ConvertMapValuesToSlice[M ~map[K]V, K comparable, V any](m M) []V {
	var r = make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
