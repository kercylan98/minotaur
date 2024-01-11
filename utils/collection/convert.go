package collection

// ConvertSliceToAny 将切片转换为任意类型的切片
func ConvertSliceToAny[S ~[]V, V any](s S) []any {
	if s == nil {
		return nil
	}
	var r = make([]any, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}

// ConvertSliceToIndexMap 将切片转换为索引为键的映射
func ConvertSliceToIndexMap[S ~[]V, V any](s S) map[int]V {
	if s == nil {
		return nil
	}
	var r = make(map[int]V, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}

// ConvertSliceToMap 将切片转换为值为键的映射
func ConvertSliceToMap[S ~[]V, V comparable](s S) map[V]struct{} {
	if s == nil {
		return nil
	}
	var r = make(map[V]struct{}, len(s))
	for _, v := range s {
		r[v] = struct{}{}
	}
	return r
}

// ConvertSliceToBoolMap 将切片转换为值为键的映射
func ConvertSliceToBoolMap[S ~[]V, V comparable](s S) map[V]bool {
	if s == nil {
		return nil
	}
	var r = make(map[V]bool, len(s))
	for _, v := range s {
		r[v] = true
	}
	return r
}

// ConvertMapKeysToSlice 将映射的键转换为切片
func ConvertMapKeysToSlice[M ~map[K]V, K comparable, V any](m M) []K {
	if m == nil {
		return nil
	}
	var r = make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// ConvertMapValuesToSlice 将映射的值转换为切片
func ConvertMapValuesToSlice[M ~map[K]V, K comparable, V any](m M) []V {
	if m == nil {
		return nil
	}
	var r = make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// InvertMap 将映射的键和值互换
func InvertMap[M ~map[K]V, N ~map[V]K, K, V comparable](m M) N {
	if m == nil {
		return nil
	}
	var r = make(N, len(m))
	for k, v := range m {
		r[v] = k
	}
	return r
}

// ConvertMapValuesToBool 将映射的值转换为布尔值
func ConvertMapValuesToBool[M ~map[K]V, N ~map[K]bool, K comparable, V any](m M) N {
	if m == nil {
		return nil
	}
	var r = make(N, len(m))
	for k := range m {
		r[k] = true
	}
	return r
}

// ReverseSlice 将切片反转
func ReverseSlice[S ~[]V, V any](s *S) {
	if s == nil {
		return
	}

	var length = len(*s)
	for i := 0; i < length/2; i++ {
		(*s)[i], (*s)[length-i-1] = (*s)[length-i-1], (*s)[i]
	}
}
