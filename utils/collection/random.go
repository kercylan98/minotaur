package collection

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/random"
)

// ChooseRandomSliceElementRepeatN 返回 slice 中的 n 个可重复随机元素
//   - 当 slice 长度为 0 或 n 小于等于 0 时将会返回 nil
func ChooseRandomSliceElementRepeatN[S ~[]V, V any](slice S, n int) (result []V) {
	if len(slice) == 0 || n <= 0 {
		return
	}
	result = make([]V, n)
	m := len(slice) - 1
	for i := 0; i < n; i++ {
		result[i] = slice[random.Int(0, m)]
	}
	return
}

// ChooseRandomIndexRepeatN 返回 slice 中的 n 个可重复随机元素的索引
//   - 当 slice 长度为 0 或 n 小于等于 0 时将会返回 nil
func ChooseRandomIndexRepeatN[S ~[]V, V any](slice S, n int) (result []int) {
	if len(slice) == 0 || n <= 0 {
		return
	}
	result = make([]int, n)
	m := len(slice) - 1
	for i := 0; i < n; i++ {
		result[i] = random.Int(0, m)
	}
	return
}

// ChooseRandomSliceElement 返回 slice 中随机一个元素，当 slice 长度为 0 时将会得到 V 的零值
func ChooseRandomSliceElement[S ~[]V, V any](slice S) (v V) {
	if len(slice) == 0 {
		return
	}
	return slice[random.Int(0, len(slice)-1)]
}

// ChooseRandomIndex 返回 slice 中随机一个元素的索引，当 slice 长度为 0 时将会得到 -1
func ChooseRandomIndex[S ~[]V, V any](slice S) (index int) {
	if len(slice) == 0 {
		return -1
	}
	return random.Int(0, len(slice)-1)
}

// ChooseRandomSliceElementN 返回 slice 中的 n 个不可重复的随机元素
//   - 当 slice 长度为 0 或 n 大于 slice 长度或小于 0 时将会发生 panic
func ChooseRandomSliceElementN[S ~[]V, V any](slice S, n int) (result []V) {
	if len(slice) == 0 || n <= 0 || n > len(slice) {
		panic(fmt.Errorf("n is greater than the length of the slice or less than 0, n: %d, length: %d", n, len(slice)))
	}
	result = make([]V, 0, n)
	valid := ConvertSliceToIndexOnlyMap(slice)
	for i := range valid {
		result = append(result, slice[i])
		if len(result) == n {
			break
		}
	}
	return
}

// ChooseRandomIndexN 获取切片中的 n 个随机元素的索引
//   - 如果 n 大于切片长度或小于 0 时将会发生 panic
func ChooseRandomIndexN[S ~[]V, V any](slice S, n int) (result []int) {
	if len(slice) == 0 {
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

// ChooseRandomMapKeyRepeatN 获取 map 中的 n 个随机 key，允许重复
//   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
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

// ChooseRandomMapValueRepeatN 获取 map 中的 n 个随机 inputV，允许重复
//   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
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

// ChooseRandomMapKeyAndValueRepeatN 获取 map 中的 n 个随机 key 和 v，允许重复
//   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
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

// ChooseRandomMapKeyAndValue 获取 map 中的随机 key 和 v
func ChooseRandomMapKeyAndValue[M ~map[K]V, K comparable, V any](m M) (k K, v V) {
	if m == nil {
		return
	}
	for k, v = range m {
		return
	}
	return
}

// ChooseRandomMapKeyAndValueN 获取 map 中的 inputN 个随机 key 和 v
//   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
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
