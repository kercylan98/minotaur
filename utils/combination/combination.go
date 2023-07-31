package combination

import (
	"github.com/kercylan98/minotaur/utils/random"
)

// NewCombination 创建一个新的组合器
func NewCombination[T Item](options ...Option[T]) *Combination[T] {
	combination := &Combination[T]{
		matchers: make(map[string]*Matcher[T]),
	}
	for _, option := range options {
		option(combination)
	}
	if combination.evaluate == nil {
		combination.evaluate = func(items []T) float64 {
			return random.Float64()
		}
	}
	return combination
}

// Combination 用于从多个匹配器内提取组合的数据结构
type Combination[T Item] struct {
	evaluate func([]T) float64
	matchers map[string]*Matcher[T]
	priority []string
}

// NewMatcher 添加一个新的匹配器
func (slf *Combination[T]) NewMatcher(name string, options ...MatcherOption[T]) *Combination[T] {
	if _, exist := slf.matchers[name]; exist {
		panic("exist of the same matcher")
	}
	var matcher = &Matcher[T]{
		evaluate: slf.evaluate,
	}
	for _, option := range options {
		option(matcher)
	}
	slf.matchers[name] = matcher
	slf.priority = append(slf.priority, name)
	return slf
}

// AddMatcher 添加一个匹配器
func (slf *Combination[T]) AddMatcher(name string, matcher *Matcher[T]) *Combination[T] {
	if _, exist := slf.matchers[name]; exist {
		panic("exist of the same matcher")
	}
	slf.matchers[name] = matcher
	slf.priority = append(slf.priority, name)
	return slf
}

// RemoveMatcher 移除一个匹配器
func (slf *Combination[T]) RemoveMatcher(name string) *Combination[T] {
	delete(slf.matchers, name)
	for i := 0; i < len(slf.priority); i++ {
		if slf.priority[i] == name {
			slf.priority = append(slf.priority[:i], slf.priority[i+1:]...)
			break
		}
	}
	return slf
}

// Combinations 从一组数据中提取所有符合匹配器规则的组合
func (slf *Combination[T]) Combinations(items []T) (result [][]T) {
	for _, n := range slf.priority {
		result = append(result, slf.matchers[n].Combinations(items)...)
	}
	return result
}

// CombinationsToName 从一组数据中提取所有符合匹配器规则的组合，并返回匹配器名称
func (slf *Combination[T]) CombinationsToName(items []T) (result map[string][][]T) {
	result = make(map[string][][]T)
	for _, n := range slf.priority {
		result[n] = append(result[n], slf.matchers[n].Combinations(items)...)
	}
	return result
}

// Best 从一组数据中提取符合匹配器规则的最佳组合
func (slf *Combination[T]) Best(items []T) (name string, result []T) {
	var best float64 = -1
	for _, n := range slf.priority {
		matcher := slf.matchers[n]
		matcherBest := matcher.Best(items)
		if len(matcherBest) == 0 {
			continue
		}
		if score := matcher.evaluate(matcherBest); score > best || best == -1 {
			best = score
			name = n
			result = matcherBest
		}
	}
	return
}

// Worst 从一组数据中提取符合匹配器规则的最差组合
func (slf *Combination[T]) Worst(items []T) (name string, result []T) {
	var worst float64 = -1
	for _, n := range slf.priority {
		matcher := slf.matchers[n]
		matcherWorst := matcher.Worst(items)
		if len(matcherWorst) == 0 {
			continue
		}
		if score := matcher.evaluate(matcherWorst); score < worst || worst == -1 {
			worst = score
			name = n
			result = matcherWorst
		}
	}
	return
}
