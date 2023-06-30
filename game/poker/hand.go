package poker

const (
	HandNone = "None" // 无牌型
)

// HandHandle 扑克牌型验证函数
type HandHandle func(rule *Rule, cards []Card) bool

// HandSingle 单牌
func HandSingle() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		return len(cards) == 1
	}
}

// HandPairs 对子
func HandPairs() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		return len(cards) == 2 && rule.IsPointContinuity(2, cards...)
	}
}

// HandThreeOfKind 三张
func HandThreeOfKind() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		return len(cards) == 3 && rule.IsPointContinuity(3, cards...)
	}
}

// HandThreeOfKindWithOne 三带一
func HandThreeOfKindWithOne() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
func HandThreeOfKindWithTwo() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
func HandOrderSingle(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		return len(cards) >= count && rule.IsPointContinuity(1, cards...)
	}
}

// HandOrderPairs 对子顺子
func HandOrderPairs(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) < count*2 || len(cards)%2 != 0 {
			return false
		}
		return rule.IsPointContinuity(2, cards...)
	}
}

// HandOrderSingleThree 三张顺子
func HandOrderSingleThree(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) < count*3 || len(cards)%3 != 0 {
			return false
		}
		return rule.IsPointContinuity(3, cards...)
	}
}

// HandOrderSingleFour 四张顺子
func HandOrderSingleFour(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) < count*4 || len(cards)%4 != 0 {
			return false
		}
		return rule.IsPointContinuity(4, cards...)
	}
}

// HandOrderThreeWithOne 三带一顺子
func HandOrderThreeWithOne(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
		var continuous []Card
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
func HandOrderThreeWithTwo(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
		var continuous []Card
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
func HandOrderFourWithOne(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
		var continuous []Card
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
func HandOrderFourWithTwo(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
		var continuous []Card
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
func HandOrderFourWithThree(count int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
		var continuous []Card
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
func HandFourWithOne() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
func HandFourWithTwo() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
func HandFourWithThree() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
func HandFourWithTwoPairs() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
func HandBomb() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		return len(cards) == 4 && rule.IsPointContinuity(4, cards...)
	}
}

// HandStraightPairs 连对
func HandStraightPairs() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) < 6 || len(cards)%2 != 0 {
			return false
		}
		return rule.IsPointContinuity(2, cards...)
	}
}

// HandPlane 飞机
//   - 表示三张点数相同的牌组成的连续的牌
func HandPlane() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) < 6 || len(cards)%3 != 0 {
			return false
		}
		return rule.IsPointContinuity(3, cards...)
	}
}

// HandPlaneWithOne 飞机带单
func HandPlaneWithOne() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
func HandRocket() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) != 2 {
			return false
		}
		return IsRocket(cards[0], cards[1])
	}
}

// HandFlush 同花
//   - 表示所有牌的花色都相同
func HandFlush() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		return IsFlush(cards...)
	}
}

// HandFlushStraight 同花顺
//   - count: 顺子的对子数量，例如当 count = 2 时，可以是 334455、445566、556677、667788、778899
//   - lower: 顺子的最小连续数量
//   - limit: 顺子的最大连续数量
func HandFlushStraight(count, lower, limit int) HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) < lower*count || len(cards) > limit*count || len(cards)%count != 0 {
			return false
		}
		if !IsFlush(cards...) {
			return false
		}
		return rule.IsPointContinuity(count, cards...)
	}
}

// HandLeopard 豹子
//   - 表示三张点数相同的牌
//   - 例如：333、444、555、666、777、888、999、JJJ、QQQ、KKK、AAA
//   - 大小王不能用于豹子，因为他们没有点数
func HandLeopard() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		if len(cards) == 0 {
			return false
		}
		if len(cards) == 1 {
			return true
		}
		var card = cards[0]
		for i := 1; i < len(cards); i++ {
			if !card.Equal(cards[1]) {
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
func HandTwoWithOne() HandHandle {
	return func(rule *Rule, cards []Card) bool {
		group := GroupByPoint(cards...)
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
