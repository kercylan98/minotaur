package sher

import "sort"

type ComparisonHandler[V any] func(source, target V) bool

// InSlice 判断切片中是否包含某个元素
func InSlice[S ~[]V, V any](slice S, v V, handler ComparisonHandler[V]) bool {
	if slice == nil {
		return false
	}
	for _, value := range slice {
		if handler(v, value) {
			return true
		}
	}
	return false
}

// InSliceByBinarySearch 判断切片中是否包含某个元素，使用二分搜索
func InSliceByBinarySearch[S ~[]V, V any](slice S, v V, handler ComparisonHandler[V]) bool {
	return sort.Search(len(slice), func(i int) bool {
		return handler(v, slice[i])
	}) != len(slice)
}

// AllInSlice 判断切片中是否包含所有元素
func AllInSlice[S ~[]V, V any](slice S, values []V, handler ComparisonHandler[V]) bool {
	if slice == nil {
		return false
	}
	for _, value := range values {
		if !InSlice(slice, value, handler) {
			return false
		}
	}
	return true
}

// AllInSliceByBinarySearch 判断切片中是否包含所有元素，使用二分搜索
func AllInSliceByBinarySearch[S ~[]V, V any](slice S, values []V, handler ComparisonHandler[V]) bool {
	if slice == nil {
		return false
	}
	for _, value := range values {
		if !InSliceByBinarySearch(slice, value, handler) {
			return false
		}
	}
	return true
}

// AnyInSlice 判断切片中是否包含任意一个元素
func AnyInSlice[S ~[]V, V any](slice S, values []V, handler ComparisonHandler[V]) bool {
	if slice == nil {
		return false
	}
	for _, value := range values {
		if InSlice(slice, value, handler) {
			return true
		}
	}
	return false
}

// AnyInSliceByBinarySearch 判断切片中是否包含任意一个元素，使用二分搜索
func AnyInSliceByBinarySearch[S ~[]V, V any](slice S, values []V, handler ComparisonHandler[V]) bool {
	if slice == nil {
		return false
	}
	for _, value := range values {
		if InSliceByBinarySearch(slice, value, handler) {
			return true
		}
	}
	return false
}

// InSlices 判断多个切片中是否包含某个元素
func InSlices[S ~[]V, V any](slices []S, v V, handler ComparisonHandler[V]) bool {
	return InSlice(MergeSlices(slices...), v, handler)
}

// InSlicesByBinarySearch 判断多个切片中是否包含某个元素，使用二分搜索
func InSlicesByBinarySearch[S ~[]V, V any](slices []S, v V, handler ComparisonHandler[V]) bool {
	return InSliceByBinarySearch(MergeSlices(slices...), v, handler)
}

// AllInSlices 判断多个切片中是否包含所有元素
func AllInSlices[S ~[]V, V any](slices []S, values []V, handler ComparisonHandler[V]) bool {
	return AllInSlice(MergeSlices(slices...), values, handler)
}

// AllInSlicesByBinarySearch 判断多个切片中是否包含所有元素，使用二分搜索
func AllInSlicesByBinarySearch[S ~[]V, V any](slices []S, values []V, handler ComparisonHandler[V]) bool {
	return AllInSliceByBinarySearch(MergeSlices(slices...), values, handler)
}

// AnyInSlices 判断多个切片中是否包含任意一个元素
func AnyInSlices[S ~[]V, V any](slices []S, values []V, handler ComparisonHandler[V]) bool {
	return AnyInSlice(MergeSlices(slices...), values, handler)
}

// AnyInSlicesByBinarySearch 判断多个切片中是否包含任意一个元素，使用二分搜索
func AnyInSlicesByBinarySearch[S ~[]V, V any](slices []S, values []V, handler ComparisonHandler[V]) bool {
	return AnyInSliceByBinarySearch(MergeSlices(slices...), values, handler)
}

// InAllSlices 判断元素是否在所有切片中都存在
func InAllSlices[S ~[]V, V any](slices []S, v V, handler ComparisonHandler[V]) bool {
	if slices == nil {
		return false
	}
	for _, slice := range slices {
		if !InSlice(slice, v, handler) {
			return false
		}
	}
	return true
}

// InAllSlicesByBinarySearch 判断元素是否在所有切片中都存在，使用二分搜索
func InAllSlicesByBinarySearch[S ~[]V, V any](slices []S, v V, handler ComparisonHandler[V]) bool {
	if slices == nil {
		return false
	}
	for _, slice := range slices {
		if !InSliceByBinarySearch(slice, v, handler) {
			return false
		}
	}
	return true
}

// AnyInAllSlices 判断元素是否在所有切片中都存在任意至少一个
func AnyInAllSlices[S ~[]V, V any](slices []S, v []V, handler ComparisonHandler[V]) bool {
	if slices == nil {
		return false
	}
	for _, slice := range slices {
		if AnyInSlice(slice, v, handler) {
			return true
		}
	}
	return false
}

// AnyInAllSlicesByBinarySearch 判断元素是否在所有切片中都存在任意至少一个，使用二分搜索
func AnyInAllSlicesByBinarySearch[S ~[]V, V any](slices []S, v []V, handler ComparisonHandler[V]) bool {
	if slices == nil {
		return false
	}
	for _, slice := range slices {
		if AnyInSliceByBinarySearch(slice, v, handler) {
			return true
		}
	}
	return false
}

// KeyInMap 判断 map 中是否包含某个 key
func KeyInMap[M ~map[K]V, K comparable, V any](m M, key K) bool {
	_, ok := m[key]
	return ok
}

// ValueInMap 判断 map 中是否包含某个 value
func ValueInMap[M ~map[K]V, K comparable, V any](m M, value V, handler ComparisonHandler[V]) bool {
	if m == nil {
		return false
	}
	for _, v := range m {
		if handler(value, v) {
			return true
		}
	}
	return false
}

// AllKeyInMap 判断 map 中是否包含所有 key
func AllKeyInMap[M ~map[K]V, K comparable, V any](m M, keys ...K) bool {
	if m == nil {
		return false
	}
	for _, key := range keys {
		if !KeyInMap(m, key) {
			return false
		}
	}
	return true
}

// AllValueInMap 判断 map 中是否包含所有 value
func AllValueInMap[M ~map[K]V, K comparable, V any](m M, values []V, handler ComparisonHandler[V]) bool {
	if m == nil {
		return false
	}
	for _, value := range values {
		if !ValueInMap(m, value, handler) {
			return false
		}
	}
	return true
}

// AnyKeyInMap 判断 map 中是否包含任意一个 key
func AnyKeyInMap[M ~map[K]V, K comparable, V any](m M, keys ...K) bool {
	if m == nil {
		return false
	}
	for _, key := range keys {
		if KeyInMap(m, key) {
			return true
		}
	}
	return false
}

// AnyValueInMap 判断 map 中是否包含任意一个 value
func AnyValueInMap[M ~map[K]V, K comparable, V any](m M, values []V, handler ComparisonHandler[V]) bool {
	if m == nil {
		return false
	}
	for _, value := range values {
		if ValueInMap(m, value, handler) {
			return true
		}
	}
	return false
}

// AllKeyInMaps 判断多个 map 中是否包含所有 key
func AllKeyInMaps[M ~map[K]V, K comparable, V any](maps []M, keys ...K) bool {
	if maps == nil {
		return false
	}
	for _, m := range maps {
		if !AllKeyInMap(m, keys...) {
			return false
		}
	}
	return true
}

// AllValueInMaps 判断多个 map 中是否包含所有 value
func AllValueInMaps[M ~map[K]V, K comparable, V any](maps []M, values []V, handler ComparisonHandler[V]) bool {
	if maps == nil {
		return false
	}
	for _, m := range maps {
		if !AllValueInMap(m, values, handler) {
			return false
		}
	}
	return true
}

// AnyKeyInMaps 判断多个 map 中是否包含任意一个 key
func AnyKeyInMaps[M ~map[K]V, K comparable, V any](maps []M, keys ...K) bool {
	if maps == nil {
		return false
	}
	for _, m := range maps {
		if AnyKeyInMap(m, keys...) {
			return true
		}
	}
	return false
}

// AnyValueInMaps 判断多个 map 中是否包含任意一个 value
func AnyValueInMaps[M ~map[K]V, K comparable, V any](maps []M, values []V, handler ComparisonHandler[V]) bool {
	if maps == nil {
		return false
	}
	for _, m := range maps {
		if AnyValueInMap(m, values, handler) {
			return true
		}
	}
	return false
}

// InAllMaps 判断元素是否在所有 map 中都存在
func InAllMaps[M ~map[K]V, K comparable, V any](maps []M, key K) bool {
	if maps == nil {
		return false
	}
	for _, m := range maps {
		if !KeyInMap(m, key) {
			return false
		}
	}
	return true
}

// AnyInAllMaps 判断元素是否在所有 map 中都存在任意至少一个
func AnyInAllMaps[M ~map[K]V, K comparable, V any](maps []M, keys []K) bool {
	if maps == nil {
		return false
	}
	for _, m := range maps {
		if AnyKeyInMap(m, keys...) {
			return true
		}
	}
	return false
}
