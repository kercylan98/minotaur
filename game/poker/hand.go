package poker

import "github.com/kercylan98/minotaur/utils/generic"

const (
	HandNone = "None" // 无牌型
)

// HandHandle 扑克牌型验证函数
type HandHandle[P, C generic.Number, T Card[P, C]] func(rule *Rule[P, C, T], cards []T) bool

// HandSingle 单牌
func HandSingle[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		return len(cards) == 1
	}
}

// HandPairs 对子
func HandPairs[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		return len(cards) == 2 && rule.IsPointContinuity(2, cards...)
	}
}

// HandThreeOfKind 三张
func HandThreeOfKind[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		return len(cards) == 3 && rule.IsPointContinuity(3, cards...)
	}
}

// HandThreeOfKindWithOne 三带一
func HandThreeOfKindWithOne[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		if len(group) != 2 {
			return false
		}
		var hasThree bool
		var count int
		for _, cards := range group {
			if len(cards) == 3 {
				hasThree = true
			} else {
				count = len(cards)
			}
		}
		return hasThree && count == 1
	}
}

// HandThreeOfKindWithTwo 三带二
func HandThreeOfKindWithTwo[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		if len(group) != 2 {
			return false
		}
		var hasThree bool
		var count int
		for _, cards := range group {
			if len(cards) == 3 {
				hasThree = true
			} else {
				count = len(cards)
			}
		}
		return hasThree && count == 2
	}
}

// HandOrderSingle 顺子
func HandOrderSingle[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		return len(cards) >= count && rule.IsPointContinuity(1, cards...)
	}
}

// HandOrderPairs 对子顺子
func HandOrderPairs[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) < count*2 || len(cards)%2 != 0 {
			return false
		}
		return rule.IsPointContinuity(2, cards...)
	}
}

// HandOrderSingleThree 三张顺子
func HandOrderSingleThree[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) < count*3 || len(cards)%3 != 0 {
			return false
		}
		return rule.IsPointContinuity(3, cards...)
	}
}

// HandOrderSingleFour 四张顺子
func HandOrderSingleFour[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) < count*4 || len(cards)%4 != 0 {
			return false
		}
		return rule.IsPointContinuity(4, cards...)
	}
}

// HandOrderThreeWithOne 三带一顺子
func HandOrderThreeWithOne[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var continuous []T
		var other int
		for _, cards := range group {
			if len(cards) == 3 {
				continuous = append(continuous, cards...)
			} else {
				other += len(cards)
			}
		}
		if !rule.IsPointContinuity(3, continuous...) {
			return false
		}
		return other == len(continuous)/3
	}
}

// HandOrderThreeWithTwo 三带二顺子
func HandOrderThreeWithTwo[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var continuous []T
		var other int
		for _, cards := range group {
			if len(cards) == 3 {
				continuous = append(continuous, cards...)
			} else if len(cards)%2 == 0 {
				other += len(cards) / 2
			} else {
				return false
			}
		}
		if !rule.IsPointContinuity(3, continuous...) {
			return false
		}
		return other == len(continuous)/3
	}
}

// HandOrderFourWithOne 四带一顺子
func HandOrderFourWithOne[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var continuous []T
		var other int
		for _, cards := range group {
			if len(cards) == 4 {
				continuous = append(continuous, cards...)
			} else {
				other += len(cards)
			}
		}
		if !rule.IsPointContinuity(4, continuous...) {
			return false
		}
		return other == len(continuous)/4
	}
}

// HandOrderFourWithTwo 四带二顺子
func HandOrderFourWithTwo[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var continuous []T
		var other int
		for _, cards := range group {
			if len(cards) == 4 {
				continuous = append(continuous, cards...)
			} else if len(cards)%2 == 0 {
				other += len(cards) / 2
			} else {
				return false
			}
		}
		if !rule.IsPointContinuity(4, continuous...) {
			return false
		}
		return other == len(continuous)/4
	}
}

// HandOrderFourWithThree 四带三顺子
func HandOrderFourWithThree[P, C generic.Number, T Card[P, C]](count int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var continuous []T
		var other int
		for _, cards := range group {
			if len(cards) == 4 {
				continuous = append(continuous, cards...)
			} else if len(cards)%3 == 0 {
				other += len(cards) / 3
			} else {
				return false
			}
		}
		if !rule.IsPointContinuity(4, continuous...) {
			return false
		}
		return other == len(continuous)/4
	}
}

// HandFourWithOne 四带一
func HandFourWithOne[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var hasFour bool
		var count int
		for _, cards := range group {
			if len(cards) == 4 {
				hasFour = true
			} else {
				count = len(cards)
			}
		}
		return hasFour && count == 1
	}
}

// HandFourWithTwo 四带二
func HandFourWithTwo[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var hasFour bool
		var count int
		for _, cards := range group {
			if len(cards) == 4 {
				hasFour = true
			} else {
				count = len(cards)
			}
		}
		return hasFour && count == 2
	}
}

// HandFourWithThree 四带三
func HandFourWithThree[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var hasFour bool
		var count int
		for _, cards := range group {
			if len(cards) == 4 {
				hasFour = true
			} else {
				count = len(cards)
			}
		}
		return hasFour && count == 3
	}
}

// HandFourWithTwoPairs 四带两对
func HandFourWithTwoPairs[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var hasFour bool
		var count int
		for _, cards := range group {
			length := len(cards)
			if length == 4 && !hasFour {
				hasFour = true
			} else if length%2 == 0 {
				count += len(cards) / 2
				if count > 2 {
					return false
				}
			} else {
				return false
			}
		}
		return hasFour && count == 2
	}
}

// HandBomb 炸弹
func HandBomb[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		return len(cards) == 4 && rule.IsPointContinuity(4, cards...)
	}
}

// HandStraightPairs 连对
func HandStraightPairs[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) < 6 || len(cards)%2 != 0 {
			return false
		}
		return rule.IsPointContinuity(2, cards...)
	}
}

// HandPlane 飞机
//   - 表示三张点数相同的牌组成的连续的牌
func HandPlane[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) < 6 || len(cards)%3 != 0 {
			return false
		}
		return rule.IsPointContinuity(3, cards...)
	}
}

// HandPlaneWithOne 飞机带单
func HandPlaneWithOne[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		if len(group) < 2 {
			return false
		}
		var hasThree bool
		var count int
		for _, cards := range group {
			if len(cards) == 3 {
				hasThree = true
			} else {
				count = len(cards)
			}
		}
		return hasThree && count == 1
	}
}

// HandRocket 王炸
//   - 表示一对王牌，即大王和小王
func HandRocket[P, C generic.Number, T Card[P, C]](pile *CardPile[P, C, T]) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) != 2 {
			return false
		}
		return IsRocket[P, C, T](pile, cards[0], cards[1])
	}
}

// HandFlush 同花
//   - 表示所有牌的花色都相同
func HandFlush[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		return IsFlush[P, C, T](cards...)
	}
}

// HandFlushStraight 同花顺
//   - count: 顺子的对子数量，例如当 count = 2 时，可以是 334455、445566、556677、667788、778899
//   - lower: 顺子的最小连续数量
//   - limit: 顺子的最大连续数量
func HandFlushStraight[P, C generic.Number, T Card[P, C]](count, lower, limit int) HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) < lower*count || len(cards) > limit*count || len(cards)%count != 0 {
			return false
		}
		if !IsFlush[P, C, T](cards...) {
			return false
		}
		return rule.IsPointContinuity(count, cards...)
	}
}

// HandLeopard 豹子
//   - 表示三张点数相同的牌
//   - 例如：333、444、555、666、777、888、999、JJJ、QQQ、KKK、AAA
//   - 大小王不能用于豹子，因为他们没有点数
func HandLeopard[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		if len(cards) == 0 {
			return false
		}
		if len(cards) == 1 {
			return true
		}
		var card = cards[0]
		for i := 1; i < len(cards); i++ {
			if !Equal[P, C, T](card, cards[1]) {
				return false
			}
		}
		return true
	}
}

// HandTwoWithOne 二带一
//   - 表示两张点数相同的牌，加上一张其他点数的牌
//   - 例如：334、445、556、667、778、889、99J、TTQ、JJK、QQA、AA2
//   - 大小王不能用于二带一，因为他们没有点数
//   - 通常用于炸金花玩法中检查对子
func HandTwoWithOne[P, C generic.Number, T Card[P, C]]() HandHandle[P, C, T] {
	return func(rule *Rule[P, C, T], cards []T) bool {
		group := GroupByPoint[P, C, T](cards...)
		var hasTwo bool
		var count int
		for _, cards := range group {
			if len(cards) == 2 {
				hasTwo = true
			} else {
				count = len(cards)
			}
		}
		return hasTwo && count == 1
	}
}
