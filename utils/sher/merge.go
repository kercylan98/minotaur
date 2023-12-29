package sher

// MergeSlices 合并切片
func MergeSlices[S ~[]V, V any](slices ...S) (result S) {
	if len(slices) == 0 {
		return nil
	}

	var length int
	for _, slice := range slices {
		length += len(slice)
	}

	result = make(S, 0, length)
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return
}

// MergeMaps 合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会覆盖前面的 map 中的 key
func MergeMaps[M ~map[K]V, K comparable, V any](maps ...M) (result M) {
	if len(maps) == 0 {
		return nil
	}

	var length int
	for _, m := range maps {
		length += len(m)
	}

	result = make(M, length)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return
}

// MergeMapsWithSkip 合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会被跳过
func MergeMapsWithSkip[M ~map[K]V, K comparable, V any](maps ...M) (result M) {
	if len(maps) == 0 {
		return nil
	}

	result = make(M)
	for _, m := range maps {
		for k, v := range m {
			if _, ok := result[k]; !ok {
				result[k] = v
			}
		}
	}
	return
}
