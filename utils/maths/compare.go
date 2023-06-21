package maths

import "github.com/kercylan98/minotaur/utils/generic"

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
