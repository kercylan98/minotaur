package poker

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/slice"
	"reflect"
	"sort"
)

// MatcherOption 匹配器选项
type MatcherOption[P, C generic.Number, T Card[P, C]] func(matcher *MatcherFilter[P, C, T])

// WithMatcherScoreAsc 通过升序评估分数创建匹配器
//   - 用于评估一组扑克牌的分值，分值最低的组合将被选中
//   - 默认为分数最高的组合将被选中
func WithMatcherScoreAsc[P, C generic.Number, T Card[P, C]]() MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.asc = true
	}
}

// WithMatcherLeastLength 通过匹配最小长度的扑克牌创建匹配器
//   - length: 牌型的长度，表示需要匹配的扑克牌最小数量
func WithMatcherLeastLength[P, C generic.Number, T Card[P, C]](length int) MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			combinations = slice.LimitedCombinations(cards, length, len(cards))
			return combinations
		})
	}
}

// WithMatcherLength 通过匹配长度的扑克牌创建匹配器
//   - length: 牌型的长度，表示需要匹配的扑克牌数量
func WithMatcherLength[P, C generic.Number, T Card[P, C]](length int) MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			combinations = slice.LimitedCombinations(cards, length, length)
			return combinations
		})
	}
}

// WithMatcherContinuity 通过匹配连续的扑克牌创建匹配器
func WithMatcherContinuity[P, C generic.Number, T Card[P, C]]() MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			n := len(cards)

			if n <= 0 {
				return combinations
			}

			// 对扑克牌按点数进行排序
			sort.Slice(cards, func(i, j int) bool {
				return cards[i].GetPoint() < cards[j].GetPoint()
			})

			// 查找连续的牌型组合
			for i := 0; i < n; i++ {
				combination := []T{cards[i]}
				for j := i + 1; j < n; j++ {
					if cards[j].GetPoint()-combination[len(combination)-1].GetPoint() == 1 {
						combination = append(combination, cards[j])
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

// WithMatcherContinuityPointOrder 通过匹配连续的扑克牌创建匹配器，与 WithMatcherContinuity 不同的是，该选项将按照自定义的点数顺序进行匹配
func WithMatcherContinuityPointOrder[P, C generic.Number, T Card[P, C]](order map[P]int) MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		var getOrder = func(card T) P {
			if v, ok := order[card.GetPoint()]; ok {
				return P(v)
			}
			return card.GetPoint()
		}
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			n := len(cards)

			if n <= 0 {
				return combinations
			}

			// 对扑克牌按点数进行排序
			sort.Slice(cards, func(i, j int) bool {
				return getOrder(cards[i]) < getOrder(cards[j])
			})

			// 查找连续的牌型组合
			for i := 0; i < n; i++ {
				combination := []T{cards[i]}
				for j := i + 1; j < n; j++ {
					if getOrder(cards[j])-getOrder(combination[len(combination)-1]) == 1 {
						combination = append(combination, cards[j])
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

// WithMatcherFlush 通过匹配同花的扑克牌创建匹配器
func WithMatcherFlush[P, C generic.Number, T Card[P, C]]() MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T

			groups := GroupByColor[P, C, T](cards...)
			for _, group := range groups {
				combinations = append(combinations, slice.Combinations(group)...)
			}

			return combinations
		})
	}
}

// WithMatcherTie 通过匹配相同点数的扑克牌创建匹配器
func WithMatcherTie[P, C generic.Number, T Card[P, C]]() MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			groups := GroupByPoint[P, C, T](cards...)
			for _, group := range groups {
				for _, ts := range slice.Combinations(group) {
					combinations = append(combinations, ts)
				}
			}
			return combinations
		})
	}
}

// WithMatcherTieCount 通过匹配相同点数的特定数量的扑克牌创建匹配器
//   - count: 牌型中相同点数的牌的数量
func WithMatcherTieCount[P, C generic.Number, T Card[P, C]](count int) MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			groups := GroupByPoint[P, C, T](cards...)
			for _, group := range groups {
				if len(group) < count {
					continue
				}
				for _, ts := range slice.Combinations(group) {
					if len(ts) == count {
						combinations = append(combinations, ts)
					}
				}
			}
			return combinations
		})
	}
}

// WithMatcherTieCountNum 通过匹配相同点数的特定数量的扑克牌创建匹配器
//   - count: 牌型中相同点数的牌的数量
//   - num: 牌型中相同点数的牌的数量
func WithMatcherTieCountNum[P, C generic.Number, T Card[P, C]](count, num int) MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			cs := slice.LimitedCombinations(cards, count*num, count*num)
			var pointCount = make(map[P]int)
			for _, group := range cs {
				var ok = false
				for _, t := range group {
					pointCount[t.GetPoint()]++
					if len(pointCount) == 2 {
						var matchCount = true
						for _, n := range pointCount {
							if n != num {
								matchCount = false
								break
							}
						}
						if matchCount {
							ok = true
							break
						}
					}
				}
				if ok {
					combinations = append(combinations, group)
				}
				for point := range pointCount {
					delete(pointCount, point)
				}
			}
			return combinations
		})
	}
}

// WithMatcherNCarryM 通过匹配N带相同点数M的扑克牌创建匹配器
//   - n: 需要匹配的主牌数量
//   - m: 需要匹配的附加牌数量
func WithMatcherNCarryM[P, C generic.Number, T Card[P, C]](n, m int) MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			groups := GroupByPoint[P, C, T](cards...)
			for _, group := range groups {
				if len(group) != n {
					continue
				}
				ms := slice.Combinations(slice.SubWithCheck(cards, group, func(a, b T) bool { return reflect.DeepEqual(a, b) }))
				for i := 0; i < len(ms); i++ {
					ts := GroupByPoint[P, C, T](ms[i]...)
					for _, cs := range ts {
						if len(cs) == m {
							combinations = append(combinations, slice.Merge(group, cs))
						}
					}
				}
			}
			return combinations
		})
	}
}

// WithMatcherNCarryMSingle 通过匹配N带M的扑克牌创建匹配器
//   - n: 需要匹配的主牌数量
//   - m: 需要匹配的附加牌数量
func WithMatcherNCarryMSingle[P, C generic.Number, T Card[P, C]](n, m int) MatcherOption[P, C, T] {
	return func(matcher *MatcherFilter[P, C, T]) {
		matcher.AddHandle(func(cards []T) [][]T {
			var combinations [][]T
			groups := GroupByPoint[P, C, T](cards...)
			for _, group := range groups {
				if len(group) != n {
					continue
				}
				ms := slice.Combinations(slice.SubWithCheck(cards, group, func(a, b T) bool { return reflect.DeepEqual(a, b) }))
				for i := 0; i < len(ms); i++ {
					ts := ms[i]
					if len(ts) == m {
						combinations = append(combinations, slice.Merge(group, ts))
					}
				}
			}
			if len(combinations) > 0 {
				fmt.Println(len(combinations))
			}
			return combinations
		})
	}
}
