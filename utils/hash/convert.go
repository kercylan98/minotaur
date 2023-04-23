package hash

// ToSlice 将 map 的 value 转换为切片
func ToSlice[K comparable, V any](m map[K]V) []V {
	var s = make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// KeyToSlice 将 map 的 key 转换为切片
func KeyToSlice[K comparable, V any](m map[K]V) []K {
	var s = make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Reversal 将 map 的 key 和 value 互换
func Reversal[K comparable, V comparable](m map[K]V) map[V]K {
	var nm = make(map[V]K)
	for k, v := range m {
		nm[v] = k
	}
	return nm
}
