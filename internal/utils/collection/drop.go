package collection

// ClearSlice 清空切片
func ClearSlice[S ~[]V, V any](slice *S) {
	if slice == nil {
		return
	}
	*slice = (*slice)[0:0]
}

// ClearMap 清空 map
func ClearMap[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

// DropSliceByIndices 删除切片中特定索引的元素
func DropSliceByIndices[S ~[]V, V any](slice *S, indices ...int) {
	if slice == nil {
		return
	}
	if len(indices) == 0 {
		return
	}

	excludeMap := make(map[int]bool)
	for _, ex := range indices {
		excludeMap[ex] = true
	}

	validElements := (*slice)[:0]
	for i, v := range *slice {
		if !excludeMap[i] {
			validElements = append(validElements, v)
		}
	}

	*slice = validElements
}

// DropSliceByCondition 删除切片中符合条件的元素
//   - condition 的返回值为 true 时，将会删除该元素
func DropSliceByCondition[S ~[]V, V any](slice *S, condition func(v V) bool) {
	if slice == nil {
		return
	}
	if condition == nil {
		return
	}

	validElements := (*slice)[:0]
	for _, v := range *slice {
		if !condition(v) {
			validElements = append(validElements, v)
		}
	}

	*slice = validElements
}

// DropSliceOverlappingElements 删除切片中与另一个切片重叠的元素
func DropSliceOverlappingElements[S ~[]V, V any](slice *S, anotherSlice []V, comparisonHandler ComparisonHandler[V]) {
	if slice == nil {
		return
	}
	if anotherSlice == nil {
		return
	}
	if comparisonHandler == nil {
		return
	}

	validElements := (*slice)[:0]
	for _, v := range *slice {
		if !InSlice(anotherSlice, v, comparisonHandler) {
			validElements = append(validElements, v)
		}
	}

	*slice = validElements
}
