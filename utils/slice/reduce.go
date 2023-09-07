package slice

import (
	"github.com/kercylan98/minotaur/utils/generic"
)

// Reduce 将切片中的多个元素组合成一个单一值
//   - start: 开始索引，如果为负数则从后往前计算，例如：-1 表示从最后一个元素开始向左遍历，1 表示从第二个元素开始
//   - slice: 待组合的切片
//   - reducer: 组合函数
func Reduce[V any, R generic.Number](start int, slice []V, reducer func(index int, item V, current R) R) (result R) {
	length := len(slice)

	if start >= length || -start > length {
		return
	}

	if start < 0 {
		for i := length + start; i >= 0; i-- {
			result = reducer(i, slice[i], result)
		}
	} else {
		for i := start; i < length; i++ {
			result = reducer(i, slice[i], result)
		}
	}
	return result
}
