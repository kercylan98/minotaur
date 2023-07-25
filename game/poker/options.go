package poker

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
)

type Option[P, C generic.Number, T Card[P, C]] func(rule *Rule[P, C, T])

// WithHand 通过绑定特定牌型的方式创建扑克玩法
//   - 牌型顺序决定了牌型的优先级
func WithHand[P, C generic.Number, T Card[P, C]](pokerHand string, value int, handle HandHandle[P, C, T]) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		if _, exist := rule.pokerHand[pokerHand]; exist {
			panic(fmt.Errorf("same poker hand name: %s", pokerHand))
		}
		rule.pokerHand[pokerHand] = handle
		rule.pokerHandValue[pokerHand] = value

		restraint, exist := rule.restraint[pokerHand]
		if !exist {
			restraint = map[string]struct{}{}
			rule.restraint[pokerHand] = restraint
		}
		restraint[pokerHand] = struct{}{}
	}
}

// WithHandRestraint 通过绑定特定克制牌型的方式创建扑克玩法
func WithHandRestraint[P, C generic.Number, T Card[P, C]](pokerHand, restraint string) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		r, exist := rule.restraint[pokerHand]
		if !exist {
			r = map[string]struct{}{}
			rule.restraint[pokerHand] = r
		}
		r[restraint] = struct{}{}
	}
}

// WithHandRestraintFull 通过绑定所有克制牌型的方式创建扑克玩法
//   - 需要确保在牌型声明之后调用
func WithHandRestraintFull[P, C generic.Number, T Card[P, C]](pokerHand string) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		for hand := range rule.pokerHand {
			r, exist := rule.restraint[pokerHand]
			if !exist {
				r = map[string]struct{}{}
				rule.restraint[pokerHand] = r
			}
			r[hand] = struct{}{}
		}
	}
}

// WithPointValue 通过特定的扑克点数牌值创建扑克玩法
func WithPointValue[P, C generic.Number, T Card[P, C]](pointValues map[P]int) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		rule.pointValue = pointValues
	}
}

// WithColorValue 通过特定的扑克花色牌值创建扑克玩法
func WithColorValue[P, C generic.Number, T Card[P, C]](colorValues map[C]int) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		rule.colorValue = colorValues
	}
}

// WithPointSort 通过特定的扑克点数顺序创建扑克玩法，顺序必须为连续的
func WithPointSort[P, C generic.Number, T Card[P, C]](pointSort map[P]int) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		for k, v := range pointSort {
			rule.pointSort[k] = v
		}
	}
}

// WithColorSort 通过特定的扑克花色顺序创建扑克玩法，顺序必须为连续的
func WithColorSort[P, C generic.Number, T Card[P, C]](colorSort map[C]int) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		for k, v := range colorSort {
			rule.colorSort[k] = v
		}
	}
}

// WithExcludeContinuityPoint 排除连续的点数
func WithExcludeContinuityPoint[P, C generic.Number, T Card[P, C]](points ...P) Option[P, C, T] {
	return func(rule *Rule[P, C, T]) {
		if rule.excludeContinuityPoint == nil {
			rule.excludeContinuityPoint = make(map[P]struct{})
		}
		for _, point := range points {
			rule.excludeContinuityPoint[point] = struct{}{}
		}
	}
}
