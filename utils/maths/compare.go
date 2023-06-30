package maths

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/slice"
	"sort"
)

const (
	CompareGreaterThan        CompareExpression = 1 // 大于
	CompareGreaterThanOrEqual CompareExpression = 2 // 大于等于
	CompareLessThan           CompareExpression = 3 // 小于
	CompareLessThanOrEqual    CompareExpression = 4 // 小于等于
	CompareEqual              CompareExpression = 5 // 等于
)

// CompareExpression 比较表达式
type CompareExpression int

// Compare 根据特定表达式比较两个值
func Compare[V generic.Ordered](a V, expression CompareExpression, b V) bool {
	switch expression {
	case CompareGreaterThan:
		return a > b
	case CompareGreaterThanOrEqual:
		return a >= b
	case CompareLessThan:
		return a < b
	case CompareLessThanOrEqual:
		return a <= b
	case CompareEqual:
		return a == b
	}
	panic("unknown expression")
}

// IsContinuity 检查一组值是否连续
func IsContinuity[V generic.Integer](values []V) bool {
	length := len(values)
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	for i := 1; i < length; i++ {
		if values[i] != values[i-1]+1 {
			return false
		}
	}
	return true
}

// IsContinuityWithSort 检查一组值排序后是否连续
func IsContinuityWithSort[V generic.Integer](values []V) bool {
	sli := slice.Copy(values)
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	return IsContinuity(sli)
}
