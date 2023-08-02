package combination

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/maths"
)

type ValidatorOption[T Item] func(validator *Validator[T])

// WithValidatorHandle 通过特定的验证函数对组合进行验证
func WithValidatorHandle[T Item](handle func(items []T) bool) ValidatorOption[T] {
	return func(validator *Validator[T]) {
		validator.vh = append(validator.vh, handle)
	}
}

// WithValidatorHandleLength 校验组合的长度是否符合要求
func WithValidatorHandleLength[T Item](length int) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		return len(items) == length
	})
}

// WithValidatorHandleLengthRange 校验组合的长度是否在指定的范围内
func WithValidatorHandleLengthRange[T Item](min, max int) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		return len(items) >= min && len(items) <= max
	})
}

// WithValidatorHandleLengthMin 校验组合的长度是否大于等于指定的最小值
func WithValidatorHandleLengthMin[T Item](min int) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		return len(items) >= min
	})
}

// WithValidatorHandleLengthMax 校验组合的长度是否小于等于指定的最大值
func WithValidatorHandleLengthMax[T Item](max int) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		return len(items) <= max
	})
}

// WithValidatorHandleLengthNot 校验组合的长度是否不等于指定的值
func WithValidatorHandleLengthNot[T Item](length int) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		return len(items) != length
	})
}

// WithValidatorHandleTypeLength 校验组合成员类型数量是否为指定的值
func WithValidatorHandleTypeLength[T Item, E generic.Ordered](length int, getType func(item T) E) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		var types = make(map[E]bool)
		for _, item := range items {
			types[getType(item)] = true
			if len(types) > length {
				return false
			}
		}
		return len(items) == length
	})
}

// WithValidatorHandleTypeLengthRange 校验组合成员类型数量是否在指定的范围内
func WithValidatorHandleTypeLengthRange[T Item, E generic.Ordered](min, max int, getType func(item T) E) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		var types = make(map[E]bool)
		for _, item := range items {
			types[getType(item)] = true
			if len(types) > max {
				return false
			}
		}
		return len(types) >= min && len(types) <= max
	})
}

// WithValidatorHandleTypeLengthMin 校验组合成员类型数量是否大于等于指定的最小值
func WithValidatorHandleTypeLengthMin[T Item, E generic.Ordered](min int, getType func(item T) E) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		var types = make(map[E]bool)
		for _, item := range items {
			types[getType(item)] = true
			if len(types) > min {
				return false
			}
		}
		return len(types) >= min
	})
}

// WithValidatorHandleTypeLengthMax 校验组合成员类型数量是否小于等于指定的最大值
func WithValidatorHandleTypeLengthMax[T Item, E generic.Ordered](max int, getType func(item T) E) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		var types = make(map[E]bool)
		for _, item := range items {
			types[getType(item)] = true
			if len(types) > max {
				return false
			}
		}
		return len(types) <= max
	})
}

// WithValidatorHandleTypeLengthNot 校验组合成员类型数量是否不等于指定的值
func WithValidatorHandleTypeLengthNot[T Item, E generic.Ordered](length int, getType func(item T) E) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		var types = make(map[E]bool)
		for _, item := range items {
			types[getType(item)] = true
			if len(types) > length {
				return false
			}
		}
		return len(types) != length
	})
}

// WithValidatorHandleContinuous 校验组合成员是否连续
func WithValidatorHandleContinuous[T Item, Index generic.Integer](getIndex func(item T) Index) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		index := make([]Index, len(items))
		for i, item := range items {
			index[i] = getIndex(item)
		}
		return maths.IsContinuityWithSort(index)
	})
}

// WithValidatorHandleContinuousNot 校验组合成员是否不连续
func WithValidatorHandleContinuousNot[T Item, Index generic.Integer](getIndex func(item T) Index) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		index := make([]Index, len(items))
		for i, item := range items {
			index[i] = getIndex(item)
		}
		return !maths.IsContinuityWithSort(index)
	})
}

// WithValidatorHandleGroupContinuous 校验组合成员是否能够按类型分组并且连续
func WithValidatorHandleGroupContinuous[T Item, E generic.Ordered, Index generic.Integer](getType func(item T) E, getIndex func(item T) Index) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		var types = make(map[E]int)
		var index = make([]Index, len(types))
		for _, item := range items {
			e := getType(item)
			types[e]++
			if types[e] == 1 {
				index = append(index, getIndex(item))
			}
		}
		var n = -1
		for _, i := range types {
			if n == -1 {
				n = i
			} else if n != i {
				return false
			}
		}
		return maths.IsContinuityWithSort(index)
	})
}

// WithValidatorHandleGroupContinuousN 校验组合成员是否能够按分组为 n 组类型并且连续
func WithValidatorHandleGroupContinuousN[T Item, E generic.Ordered, Index generic.Integer](n int, getType func(item T) E, getIndex func(item T) Index) ValidatorOption[T] {
	return WithValidatorHandle[T](func(items []T) bool {
		var types = make(map[E]int)
		var index = make([]Index, len(types))
		for _, item := range items {
			e := getType(item)
			types[e]++
			if types[e] == 1 {
				index = append(index, getIndex(item))
			}
		}
		if n != len(types) {
			return false
		}
		var n = -1
		for _, i := range types {
			if n == -1 {
				n = i
			} else if n != i {
				return false
			}
		}
		return maths.IsContinuityWithSort(index)
	})
}
