package collection

import "github.com/kercylan98/minotaur/utils/generic"

// ComparisonHandler 用于比较 `source` 和 `target` 两个值是否相同的比较函数
//   - 该函数接受两个参数，分别是源值和目标值，返回 true 的情况下即表示两者相同
type ComparisonHandler[V any] func(source, target V) bool

// OrderedValueGetter 用于获取 v 的可排序字段值的函数
type OrderedValueGetter[V any, N generic.Ordered] func(v V) N
