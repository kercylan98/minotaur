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

// ToMap 将切片转换为 map
func ToMap[V any](slice []V) map[int]V {
	var m = make(map[int]V)
	for i, v := range slice {
		m[i] = v
	}
	return m
}

// ToIterator 将切片转换为 Iterator
func ToIterator[V comparable](slice []V) map[V]struct{} {
	var m = make(map[V]struct{})
	for _, v := range slice {
		m[v] = struct{}{}
	}
	return m
}

// ToMapBool 将切片转换为 map，value作为Key
func ToMapBool[V comparable](slice []V) map[V]bool {
	var m = make(map[V]bool)
	for _, v := range slice {
		m[v] = true
	}
	return m
}

// ToSortMap 将切片转换为 SortMap
func ToSortMap[V any](slice []V) SortMap[int, V] {
	var m SortMap[int, V]
	for i, v := range slice {
		m.Set(i, v)
	}
	return m
}

// Copy 复制一个map
func Copy[K comparable, V any](m map[K]V) map[K]V {
	var backup = make(map[K]V)
	for k, v := range m {
		backup[k] = v
	}
	return backup
}
