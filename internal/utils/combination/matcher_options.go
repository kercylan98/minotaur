package combination

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"github.com/kercylan98/minotaur/utils/generic"
"reflect"
"sort"
)

// MatcherOption 匹配器选项
type MatcherOption[T Item] func(matcher *Matcher[T])

// WithMatcherEvaluation 设置匹配器评估函数
//   - 用于对组合进行评估，返回一个分值的评价函数
//   - 通过该选项将覆盖匹配器的默认(WithEvaluation)评估函数
func WithMatcherEvaluation[T Item](evaluate func(items []T) float64) MatcherOption[T] {
	return func(m *Matcher[T]) {
		m.evaluate = evaluate
	}
}

// WithMatcherLeastLength 通过匹配最小长度的组合创建匹配器
//   - length: 组合的长度，表示需要匹配的组合最小数量
func WithMatcherLeastLength[T Item](length int) MatcherOption[T] {
	return func(m *Matcher[T]) {
		m.AddFilter(func(items []T) [][]T {
			return collection.FindCombinationsInSliceByRange(items, length, len(items))
		})
	}
}

// WithMatcherLength 通过匹配长度的组合创建匹配器
//   - length: 组合的长度，表示需要匹配的组合数量
func WithMatcherLength[T Item](length int) MatcherOption[T] {
	return func(m *Matcher[T]) {
		m.AddFilter(func(items []T) [][]T {
			return collection.FindCombinationsInSliceByRange(items, length, length)
		})
	}
}

// WithMatcherMostLength 通过匹配最大长度的组合创建匹配器
//   - length: 组合的长度，表示需要匹配的组合最大数量
func WithMatcherMostLength[T Item](length int) MatcherOption[T] {
	return func(m *Matcher[T]) {
		m.AddFilter(func(items []T) [][]T {
			return collection.FindCombinationsInSliceByRange(items, 1, length)
		})
	}
}

// WithMatcherIntervalLength 通过匹配长度区间的组合创建匹配器
//   - min: 组合的最小长度，表示需要匹配的组合最小数量
//   - max: 组合的最大长度，表示需要匹配的组合最大数量
func WithMatcherIntervalLength[T Item](min, max int) MatcherOption[T] {
	return func(m *Matcher[T]) {
		m.AddFilter(func(items []T) [][]T {
			return collection.FindCombinationsInSliceByRange(items, min, max)
		})
	}
}

// WithMatcherContinuity 通过匹配连续的组合创建匹配器
//   - index: 用于获取组合中元素的索引值，用于判断是否连续
func WithMatcherContinuity[T Item, Index generic.Number](getIndex func(item T) Index) MatcherOption[T] {
	return func(m *Matcher[T]) {
		m.AddFilter(func(items []T) [][]T {
			var combinations [][]T
			n := len(items)

			if n <= 0 {
				return combinations
			}

			sort.Slice(items, func(i, j int) bool {
				return getIndex(items[i]) < getIndex(items[j])
			})

			for i := 0; i < n; i++ {
				combination := []T{items[i]}
				for j := i + 1; j < n; j++ {
					if getIndex(items[j])-getIndex(combination[len(combination)-1]) == Index(1) {
						combination = append(combination, items[j])
					} else {
						break
					}
				}
				if len(combination) >= 2 {
					combinations = append(combinations, combination)
				}
			}

			return combinations
		})
	}
}

// WithMatcherSame 通过匹配相同的组合创建匹配器
//   - count: 组合中相同元素的数量，当 count <= 0 时，表示相同元素的数量不限
//   - getType: 用于获取组合中元素的类型，用于判断是否相同
func WithMatcherSame[T Item, E generic.Ordered](count int, getType func(item T) E) MatcherOption[T] {
	return func(m *Matcher[T]) {
		m.AddFilter(func(items []T) [][]T {
			var combinations [][]T
			groups := collection.FindCombinationsInSliceByRange(items, count, count)
			for _, items := range groups {
				if count > 0 && len(items) != count {
					continue
				}
				var e E
				var init = true
				var same = true
				for _, item := range items {
					if init {
						init = false
						e = getType(item)
					} else if getType(item) != e {
						same = false
						break
					}
				}
				if same {
					combinations = append(combinations, items)
				}
			}
			return combinations
		})
	}
}

// WithMatcherNCarryM 通过匹配 N 携带 M 的组合创建匹配器
//   - n: 组合中元素的数量，表示需要匹配的组合数量，n 的类型需要全部相同
//   - m: 组合中元素的数量，表示需要匹配的组合数量，m 的类型需要全部相同
//   - getType: 用于获取组合中元素的类型，用于判断是否相同
func WithMatcherNCarryM[T Item, E generic.Ordered](n, m int, getType func(item T) E) MatcherOption[T] {
	return func(matcher *Matcher[T]) {
		matcher.AddFilter(func(items []T) [][]T {
			var combinations [][]T

			groups := make(map[E][]T)
			for _, item := range items {
				itemType := getType(item)
				groups[itemType] = append(groups[itemType], item)
			}

			for _, group := range groups {
				if len(group) != n {
					continue
				}

				other := collection.CloneSlice(items)
				collection.DropSliceOverlappingElements(&other, group, func(source, target T) bool {
					return reflect.DeepEqual(source, target)
				})
				ms := collection.FindCombinationsInSliceByRange(other, m, m)
				for _, otherGroup := range ms {
					var t E
					var init = true
					var same = true
					for _, item := range otherGroup {
						if init {
							init = false
							t = getType(item)
						} else if getType(item) != t {
							same = false
							break
						}
					}
					if same {
						combinations = append(combinations, collection.MergeSlices(group, otherGroup))
					}
				}
			}
			return combinations
		})
	}
}

// WithMatcherNCarryIndependentM 通过匹配 N 携带独立 M 的组合创建匹配器
//   - n: 组合中元素的数量，表示需要匹配的组合数量，n 的类型需要全部相同
//   - m: 组合中元素的数量，表示需要匹配的组合数量，m 的类型无需全部相同
//   - getType: 用于获取组合中元素的类型，用于判断是否相同
func WithMatcherNCarryIndependentM[T Item, E generic.Ordered](n, m int, getType func(item T) E) MatcherOption[T] {
	return func(matcher *Matcher[T]) {
		matcher.AddFilter(func(items []T) [][]T {
			var combinations [][]T

			groups := make(map[E][]T)
			for _, item := range items {
				itemType := getType(item)
				groups[itemType] = append(groups[itemType], item)
			}

			for _, group := range groups {
				if len(group) != n {
					continue
				}

				other := collection.CloneSlice(items)
				collection.DropSliceOverlappingElements(&other, group, func(source, target T) bool {
					return reflect.DeepEqual(source, target)
				})
				ms := collection.FindCombinationsInSliceByRange(other, m, m)
				for _, otherGroup := range ms {
					combinations = append(combinations, collection.MergeSlices(group, otherGroup))
				}
			}
			return combinations
		})
	}
}
