package collection

// DeduplicateSliceInPlace 去除切片中的重复元素
func DeduplicateSliceInPlace[S ~[]V, V comparable](s *S) {
	if s == nil || len(*s) < 2 {
		return
	}

	var m = make(map[V]struct{}, len(*s))
	var writeIndex int
	for readIndex, v := range *s {
		if _, ok := m[v]; !ok {
			(*s)[writeIndex] = (*s)[readIndex]
			writeIndex++
			m[v] = struct{}{}
		}
	}
	*s = (*s)[:writeIndex]
}

// DeduplicateSlice 去除切片中的重复元素，返回新切片
func DeduplicateSlice[S ~[]V, V comparable](s S) S {
	if s == nil || len(s) < 2 {
		return any(s).(S)
	}

	var r = make([]V, 0, len(s))
	var m = make(map[V]struct{}, len(s))
	for _, v := range s {
		if _, ok := m[v]; !ok {
			r = append(r, v)
			m[v] = struct{}{}
		}
	}
	return r
}

// DeduplicateSliceInPlaceWithCompare 去除切片中的重复元素，使用自定义的比较函数
func DeduplicateSliceInPlaceWithCompare[S ~[]V, V any](s *S, compare func(a, b V) bool) {
	if s == nil || len(*s) < 2 {
		return
	}
	seen := make(map[int]struct{})
	resultIndex := 0
	for i := range *s {
		unique := true
		for j := range seen {
			if compare((*s)[i], (*s)[j]) {
				unique = false // Found a duplicate
				break
			}
		}
		if unique {
			seen[i] = struct{}{}
			(*s)[resultIndex] = (*s)[i]
			resultIndex++
		}
	}
	*s = (*s)[:resultIndex]
}

// DeduplicateSliceWithCompare 去除切片中的重复元素，使用自定义的比较函数，返回新的切片
func DeduplicateSliceWithCompare[S ~[]V, V any](s S, compare func(a, b V) bool) S {
	if s == nil || compare == nil || len(s) < 2 {
		return s
	}
	seen := make(map[int]struct{})
	var result = make([]V, 0, len(s))
	for i := range s {
		unique := true
		for j := range result {
			if compare(s[i], result[j]) {
				unique = false
				break
			}
		}
		if unique {
			result = append(result, s[i])
			seen[i] = struct{}{}
		}
	}
	return result
}
