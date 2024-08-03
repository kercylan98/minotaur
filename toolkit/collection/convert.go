package collection

// ConvertSliceToBatches 将切片 s 转换为分批次的切片，当 batchSize 小于等于 0 或者 s 长度为 0 时，将会返回 nil
func ConvertSliceToBatches[S ~[]V, V any](s S, batchSize int) []S {
	if len(s) == 0 || batchSize <= 0 {
		return nil
	}
	var batches = make([]S, 0, len(s)/batchSize+1)
	for i := 0; i < len(s); i += batchSize {
		var end = i + batchSize
		if end > len(s) {
			end = len(s)
		}
		batches = append(batches, s[i:end])
	}
	return batches
}

// ConvertMapKeysToBatches 将映射的键转换为分批次的切片，当 batchSize 小于等于 0 或者 m 长度为 0 时，将会返回 nil
func ConvertMapKeysToBatches[M ~map[K]V, K comparable, V any](m M, batchSize int) [][]K {
	if len(m) == 0 || batchSize <= 0 {
		return nil
	}
	var batches = make([][]K, 0, len(m)/batchSize+1)
	var keys = ConvertMapKeysToSlice(m)
	for i := 0; i < len(keys); i += batchSize {
		var end = i + batchSize
		if end > len(keys) {
			end = len(keys)
		}
		batches = append(batches, keys[i:end])
	}
	return batches
}

// ConvertMapValuesToBatches 将映射的值转换为分批次的切片，当 batchSize 小于等于 0 或者 m 长度为 0 时，将会返回 nil
func ConvertMapValuesToBatches[M ~map[K]V, K comparable, V any](m M, batchSize int) [][]V {
	if len(m) == 0 || batchSize <= 0 {
		return nil
	}
	var batches = make([][]V, 0, len(m)/batchSize+1)
	var values = ConvertMapValuesToSlice(m)
	for i := 0; i < len(values); i += batchSize {
		var end = i + batchSize
		if end > len(values) {
			end = len(values)
		}
		batches = append(batches, values[i:end])
	}
	return batches
}

// ConvertSliceToAny 将切片转换为任意类型的切片
func ConvertSliceToAny[S ~[]V, V any](s S) []any {
	if len(s) == 0 {
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
	if len(s) == 0 {
		return make(map[int]V)
	}
	var r = make(map[int]V, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}

// ConvertSliceToIndexOnlyMap 将切片转换为索引为键的映射
func ConvertSliceToIndexOnlyMap[S ~[]V, V any](s S) map[int]struct{} {
	if len(s) == 0 {
		return nil
	}
	var r = make(map[int]struct{}, len(s))
	for i := range s {
		r[i] = struct{}{}
	}
	return r
}

// ConvertSliceToMap 将切片转换为值为键的映射
func ConvertSliceToMap[S ~[]V, V comparable](s S) map[V]struct{} {
	if len(s) == 0 {
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
	if len(s) == 0 {
		return make(map[V]bool)
	}
	var r = make(map[V]bool, len(s))
	for _, v := range s {
		r[v] = true
	}
	return r
}

// ConvertMapKeysToSlice 将映射的键转换为切片
func ConvertMapKeysToSlice[M ~map[K]V, K comparable, V any](m M) []K {
	if len(m) == 0 {
		return nil
	}
	var r = make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// ConvertMapValuesToBoolMap 将映射的值转换为 map[K]bool
func ConvertMapValuesToBoolMap[M ~map[K]V, K comparable, V any](m M) map[K]bool {
	if len(m) == 0 {
		return nil
	}
	var r = make(map[K]bool, len(m))
	for k := range m {
		r[k] = true
	}
	return r
}

// ConvertMapValuesToSlice 将映射的值转换为切片
func ConvertMapValuesToSlice[M ~map[K]V, K comparable, V any](m M) []V {
	if len(m) == 0 {
		return nil
	}
	var r = make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// InvertMap 将映射的键和值互换
func InvertMap[M ~map[K]V, N map[V]K, K, V comparable](m M) N {
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
func ConvertMapValuesToBool[M ~map[K]V, N map[K]bool, K comparable, V any](m M) N {
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
