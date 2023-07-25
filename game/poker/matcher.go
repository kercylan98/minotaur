package poker

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
)

// NewMatcher 创建一个新的匹配器
//   - evaluate: 用于评估一组扑克牌的分值，分值最高的组合将被选中
func NewMatcher[P, C generic.Number, T Card[P, C]]() *Matcher[P, C, T] {
	matcher := &Matcher[P, C, T]{
		filter: map[string]*MatcherFilter[P, C, T]{},
	}
	return matcher
}

// Matcher 匹配器
//   - 用于匹配扑克牌型，筛选分组等
type Matcher[P, C generic.Number, T Card[P, C]] struct {
	filter map[string]*MatcherFilter[P, C, T]
	sort   []string
}

// RegType 注册一个新的牌型
//   - name: 牌型名称
//   - evaluate: 用于评估一组扑克牌的分值，分值最高的组合将被选中
//   - options: 牌型选项
func (slf *Matcher[P, C, T]) RegType(name string, evaluate func([]T) int64, options ...MatcherOption[P, C, T]) *Matcher[P, C, T] {
	if hash.Exist(slf.filter, name) {
		panic("exist of the same type")
	}
	filter := &MatcherFilter[P, C, T]{
		evaluate: evaluate,
	}
	for _, option := range options {
		option(filter)
	}
	slf.filter[name] = filter
	slf.sort = append(slf.sort, name)
	return slf
}

// Group 将一组扑克牌按照匹配器的规则分组，并返回最佳组合及其牌型名称
func (slf *Matcher[P, C, T]) Group(cards []T) (name string, result []T) {
	for _, n := range slf.sort {
		result = slf.filter[n].group(cards)
		if len(result) > 0 {
			return n, result
		}
	}
	return
}
