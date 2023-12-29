package sher

import "github.com/kercylan98/minotaur/utils/generic"

// FindMinimumInSlice 获取切片中的最小值
func FindMinimumInSlice[S ~[]V, V generic.Number](slice S, handler ComparisonHandler[V]) (result V) {
	if slice == nil {
		return
	}
	result = slice[0]
	for i := 1; i < len(slice); i++ {
		if handler(slice[i], result) {
			result = slice[i]
		}
	}
	return
}

// FindMaximumInSlice 获取切片中的最大值
func FindMaximumInSlice[S ~[]V, V generic.Number](slice S, handler ComparisonHandler[V]) (result V) {
	if slice == nil {
		return
	}
	result = slice[0]
	for i := 1; i < len(slice); i++ {
		if handler(result, slice[i]) {
			result = slice[i]
		}
	}
	return
}

// FindMin2MaxInSlice 获取切片中的最小值和最大值
func FindMin2MaxInSlice[S ~[]V, V generic.Number](slice S, handler ComparisonHandler[V]) (min, max V) {
	if slice == nil {
		return
	}
	min = slice[0]
	max = slice[0]
	for i := 1; i < len(slice); i++ {
		if handler(slice[i], min) {
			min = slice[i]
		}
		if handler(max, slice[i]) {
			max = slice[i]
		}
	}
	return
}

// FindMinFromMap 获取 map 中的最小值
func FindMinFromMap[M ~map[K]V, K comparable, V generic.Number](m M, handler ComparisonHandler[V]) (result V) {
	if m == nil {
		return
	}
	for _, v := range m {
		if handler(v, result) {
			result = v
		}
	}
	return
}

// FindMaxFromMap 获取 map 中的最大值
func FindMaxFromMap[M ~map[K]V, K comparable, V generic.Number](m M, handler ComparisonHandler[V]) (result V) {
	if m == nil {
		return
	}
	for _, v := range m {
		if handler(result, v) {
			result = v
		}
	}
	return
}

// FindMin2MaxFromMap 获取 map 中的最小值和最大值
func FindMin2MaxFromMap[M ~map[K]V, K comparable, V generic.Number](m M, handler ComparisonHandler[V]) (min, max V) {
	if m == nil {
		return
	}
	for _, v := range m {
		if handler(v, min) {
			min = v
		}
		if handler(max, v) {
			max = v
		}
	}
	return
}
