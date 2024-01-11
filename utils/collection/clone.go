package collection

import (
	"slices"
)

// CloneSlice 克隆切片，该函数是 slices.Clone 的快捷方式
func CloneSlice[S ~[]V, V any](slice S) S {
	return slices.Clone(slice)
}

// CloneMap 克隆 map
func CloneMap[M ~map[K]V, K comparable, V any](m M) M {
	if m == nil {
		return nil
	}

	var result = make(M, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// CloneSliceN 克隆 slice 为 n 个切片进行返回
func CloneSliceN[S ~[]V, V any](slice S, n int) []S {
	if slice == nil {
		return nil
	}
	if n <= 0 {
		return []S{}
	}

	var result = make([]S, n)
	for i := 0; i < n; i++ {
		result[i] = CloneSlice(slice)
	}
	return result
}

// CloneMapN 克隆 map 为 n 个 map 进行返回
func CloneMapN[M ~map[K]V, K comparable, V any](m M, n int) []M {
	if m == nil {
		return nil
	}

	if n <= 0 {
		return []M{}
	}

	var result = make([]M, n)
	for i := 0; i < n; i++ {
		result[i] = CloneMap(m)
	}
	return result
}

// CloneSlices 克隆多个切片
func CloneSlices[S ~[]V, V any](slices ...S) []S {
	if slices == nil {
		return nil
	}

	var result = make([]S, len(slices))
	for i, slice := range slices {
		result[i] = CloneSlice(slice)
	}
	return result
}

// CloneMaps 克隆多个 map
func CloneMaps[M ~map[K]V, K comparable, V any](maps ...M) []M {
	if maps == nil {
		return nil
	}

	var result = make([]M, len(maps))
	for i, m := range maps {
		result[i] = CloneMap(m)
	}
	return result
}
