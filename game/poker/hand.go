package poker

// HandHandle 扑克牌型验证函数
type HandHandle func(poker *Poker, cards []Card) bool

// HandPairs 对子
func HandPairs() HandHandle {
	return func(poker *Poker, cards []Card) bool {
		return len(cards) == 2 && cards[0].EqualPoint(cards[1])
	}
}

// HandFlushPairs 同花对子
func HandFlushPairs() HandHandle {
	return func(poker *Poker, cards []Card) bool {
		if len(cards) != 2 {
			return false
		}
		card1, card2 := cards[0], cards[1]
		return card1.Equal(card2)
	}
}

// HandSingle 单牌
func HandSingle() HandHandle {
	return func(poker *Poker, cards []Card) bool {
		return len(cards) == 1
	}
}

// HandThreeOfKind 三张
func HandThreeOfKind() HandHandle {
	return func(poker *Poker, cards []Card) bool {
		return len(cards) == 3 && cards[0].EqualPoint(cards[1]) && cards[1].EqualPoint(cards[2])
	}
}
