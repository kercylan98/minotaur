package combination

import (
	"github.com/kercylan98/minotaur/utils/random"
)

// NewMatcher 创建一个新的匹配器
func NewMatcher[T Item](options ...MatcherOption[T]) *Matcher[T] {
	matcher := &Matcher[T]{
		filter: make([]func(items []T) [][]T, 0),
	}
	for _, option := range options {
		option(matcher)
	}
	if matcher.evaluate == nil {
		matcher.evaluate = func(items []T) float64 {
			return random.Float64()
		}
	}
	return matcher
}

// Matcher 用于从一组数据内提取组合的数据结构
type Matcher[T Item] struct {
	evaluate func(items []T) float64 // 用于对组合进行评估，返回一个分值的评价函数
	filter   []func(items []T) [][]T // 用于对组合进行筛选的函数
}

// AddFilter 添加一个筛选器
//   - 筛选器用于对组合进行筛选，返回一个二维数组，每个数组内的元素都是一个组合
func (slf *Matcher[T]) AddFilter(filter func(items []T) [][]T) {
	slf.filter = append(slf.filter, filter)
}

// Combinations 从一组数据中提取所有符合筛选器规则的组合
func (slf *Matcher[T]) Combinations(items []T) [][]T {
	var combinations = [][]T{items}
	for _, handle := range slf.filter {
		combinations = append(combinations, handle(items)...)
	}
	return combinations
}

// Best 从一组数据中提取符筛选器规则的最佳组合
func (slf *Matcher[T]) Best(items []T) []T {
	var bestCombination = items

	for _, handle := range slf.filter {
		var bestScore float64 = -1
		filteredCombinations := handle(bestCombination)
		if len(filteredCombinations) == 0 {
			return nil
		}
		for _, combination := range filteredCombinations {
			score := slf.evaluate(combination)
			if score > bestScore || bestScore == -1 {
				bestCombination = combination
				bestScore = score
			}
		}
	}

	return bestCombination
}

// Worst 从一组数据中提取符筛选器规则的最差组合
func (slf *Matcher[T]) Worst(items []T) []T {
	var worstCombination = items

	for _, handle := range slf.filter {
		var worstScore float64 = -1
		filteredCombinations := handle(worstCombination)
		if len(filteredCombinations) == 0 {
			return nil
		}
		for _, combination := range filteredCombinations {
			score := slf.evaluate(combination)
			if score < worstScore || worstScore == -1 {
				worstCombination = combination
				worstScore = score
			}
		}
	}

	return worstCombination
}
