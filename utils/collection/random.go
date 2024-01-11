package collection

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/random"
)

// ChooseRandomSliceElementRepeatN 获取切片中的 inputN 个随机元素，允许重复
//   - 如果 inputN 大于切片长度或小于 0 时将会发生 panic
func ChooseRandomSliceElementRepeatN[S ~[]V, V any](slice S, n int) (result []V) {
	if slice == nil {
		return
	}
	if n > len(slice) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the input or less than 0, inputN: %d, length: %d", n, len(slice)))
	}
	result = make([]V, n)
	for i := 0; i < n; i++ {
		result[i] = slice[random.Int(0, len(slice)-1)]
	}
	return
}

// ChooseRandomIndexRepeatN 获取切片中的 inputN 个随机元素的索引，允许重复
//   - 如果 inputN 大于切片长度或小于 0 时将会发生 panic
func ChooseRandomIndexRepeatN[S ~[]V, V any](slice S, n int) (result []int) {
	if slice == nil {
		return
	}
	if n > len(slice) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the input or less than 0, inputN: %d, length: %d", n, len(slice)))
	}
	result = make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = random.Int(0, len(slice)-1)
	}
	return
}

// ChooseRandomSliceElement 获取切片中的随机元素
func ChooseRandomSliceElement[S ~[]V, V any](slice S) (v V) {
	if slice == nil {
		return
	}
	return slice[random.Int(0, len(slice)-1)]
}

// ChooseRandomIndex 获取切片中的随机元素的索引
func ChooseRandomIndex[S ~[]V, V any](slice S) (index int) {
	if slice == nil {
		return
	}
	return random.Int(0, len(slice)-1)
}

// ChooseRandomSliceElementN 获取切片中的 inputN 个随机元素
//   - 如果 inputN 大于切片长度或小于 0 时将会发生 panic
func ChooseRandomSliceElementN[S ~[]V, V any](slice S, n int) (result []V) {
	if slice == nil {
		return
	}
	if n > len(slice) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the input or less than 0, inputN: %d, length: %d", n, len(slice)))
	}
	result = make([]V, n)
	for i := 0; i < n; i++ {
		result[i] = slice[random.Int(0, len(slice)-1)]
	}
	return
}

// ChooseRandomIndexN 获取切片中的 inputN 个随机元素的索引
//   - 如果 inputN 大于切片长度或小于 0 时将会发生 panic
func ChooseRandomIndexN[S ~[]V, V any](slice S, n int) (result []int) {
	if slice == nil {
		return
	}
	if n > len(slice) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the input or less than 0, inputN: %d, length: %d", n, len(slice)))
	}
	result = make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = random.Int(0, len(slice)-1)
	}
	return
}

// ChooseRandomMapKeyRepeatN 获取 map 中的 inputN 个随机 key，允许重复
//   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic
func ChooseRandomMapKeyRepeatN[M ~map[K]V, K comparable, V any](m M, n int) (result []K) {
	if m == nil {
		return
	}
	if n > len(m) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the map or less than 0, inputN: %d, length: %d", n, len(m)))
	}
	result = make([]K, n)
	for i := 0; i < n; i++ {
		for k := range m {
			result[i] = k
			break
		}
	}
	return
}

// ChooseRandomMapValueRepeatN 获取 map 中的 inputN 个随机 inputV，允许重复
//   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic
func ChooseRandomMapValueRepeatN[M ~map[K]V, K comparable, V any](m M, n int) (result []V) {
	if m == nil {
		return
	}
	if n > len(m) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the map or less than 0, inputN: %d, length: %d", n, len(m)))
	}
	result = make([]V, n)
	for i := 0; i < n; i++ {
		for _, v := range m {
			result[i] = v
			break
		}
	}
	return
}

// ChooseRandomMapKeyAndValueRepeatN 获取 map 中的 inputN 个随机 key 和 inputV，允许重复
//   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic
func ChooseRandomMapKeyAndValueRepeatN[M ~map[K]V, K comparable, V any](m M, n int) M {
	if m == nil {
		return nil
	}
	if n > len(m) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the map or less than 0, inputN: %d, length: %d", n, len(m)))
	}
	result := make(M, n)
	for i := 0; i < n; i++ {
		for k, v := range m {
			result[k] = v
			break
		}
	}
	return result
}

// ChooseRandomMapKey 获取 map 中的随机 key
func ChooseRandomMapKey[M ~map[K]V, K comparable, V any](m M) (k K) {
	if m == nil {
		return
	}
	for k = range m {
		return
	}
	return
}

// ChooseRandomMapValue 获取 map 中的随机 inputV
func ChooseRandomMapValue[M ~map[K]V, K comparable, V any](m M) (v V) {
	if m == nil {
		return
	}
	for _, v = range m {
		return
	}
	return
}

// ChooseRandomMapKeyN 获取 map 中的 inputN 个随机 key
//   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic
func ChooseRandomMapKeyN[M ~map[K]V, K comparable, V any](m M, n int) (result []K) {
	if m == nil {
		return
	}
	if n > len(m) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the map or less than 0, inputN: %d, length: %d", n, len(m)))
	}
	result = make([]K, n)
	i := 0
	for k := range m {
		result[i] = k
		i++
		if i == n {
			break
		}
	}
	return
}

// ChooseRandomMapValueN 获取 map 中的 inputN 个随机 inputV
//   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic
func ChooseRandomMapValueN[M ~map[K]V, K comparable, V any](m M, n int) (result []V) {
	if m == nil {
		return
	}
	if n > len(m) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the map or less than 0, inputN: %d, length: %d", n, len(m)))
	}
	result = make([]V, n)
	i := 0
	for _, v := range m {
		result[i] = v
		i++
		if i == n {
			break
		}
	}
	return
}

// ChooseRandomMapKeyAndValue 获取 map 中的随机 key 和 inputV
func ChooseRandomMapKeyAndValue[M ~map[K]V, K comparable, V any](m M) (k K, v V) {
	if m == nil {
		return
	}
	for k, v = range m {
		return
	}
	return
}

// ChooseRandomMapKeyAndValueN 获取 map 中的 inputN 个随机 key 和 inputV
//   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic
func ChooseRandomMapKeyAndValueN[M ~map[K]V, K comparable, V any](m M, n int) M {
	if m == nil {
		return nil
	}
	if n > len(m) || n < 0 {
		panic(fmt.Errorf("inputN is greater than the length of the map or less than 0, inputN: %d, length: %d", n, len(m)))
	}
	result := make(M, n)
	i := 0
	for k, v := range m {
		result[k] = v
		i++
		if i == n {
			break
		}
	}
	return result
}
