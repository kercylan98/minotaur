package collection

// MappingFromSlice 将切片中的元素进行转换
func MappingFromSlice[S ~[]V, NS []N, V, N any](slice S, handler func(value V) N) NS {
	if slice == nil {
		return nil
	}
	result := make(NS, len(slice))
	for i, v := range slice {
		result[i] = handler(v)
	}
	return result
}

// MappingFromMap 将 map 中的元素进行转换
func MappingFromMap[M ~map[K]V, NM map[K]N, K comparable, V, N any](m M, handler func(value V) N) NM {
	if m == nil {
		return nil
	}
	result := make(NM, len(m))
	for k, v := range m {
		result[k] = handler(v)
	}
	return result
}
