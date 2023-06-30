package poker

import "fmt"

type Option func(rule *Rule)

// WithHand 通过绑定特定牌型的方式创建扑克玩法
//   - 牌型顺序决定了牌型的优先级
func WithHand(pokerHand string, value int, handle HandHandle) Option {
	return func(rule *Rule) {
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
func WithHandRestraint(pokerHand, restraint string) Option {
	return func(rule *Rule) {
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
func WithHandRestraintFull(pokerHand string) Option {
	return func(rule *Rule) {
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
func WithPointValue(pointValues map[Point]int) Option {
	return func(rule *Rule) {
		rule.pointValue = pointValues
	}
}

// WithColorValue 通过特定的扑克花色牌值创建扑克玩法
func WithColorValue(colorValues map[Color]int) Option {
	return func(rule *Rule) {
		rule.colorValue = colorValues
	}
}

// WithPointSort 通过特定的扑克点数顺序创建扑克玩法，顺序必须为连续的
func WithPointSort(pointSort map[Point]int) Option {
	return func(rule *Rule) {
		for k, v := range pointSort {
			rule.pointSort[k] = v
		}
	}
}

// WithColorSort 通过特定的扑克花色顺序创建扑克玩法，顺序必须为连续的
func WithColorSort(colorSort map[Color]int) Option {
	return func(rule *Rule) {
		for k, v := range colorSort {
			rule.colorSort[k] = v
		}
	}
}

// WithExcludeContinuityPoint 排除连续的点数
func WithExcludeContinuityPoint(points ...Point) Option {
	return func(rule *Rule) {
		if rule.excludeContinuityPoint == nil {
			rule.excludeContinuityPoint = make(map[Point]struct{})
		}
		for _, point := range points {
			rule.excludeContinuityPoint[point] = struct{}{}
		}
	}
}
