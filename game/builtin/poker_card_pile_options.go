package builtin

type PokerCardPileOption func(pile *PokerCardPile)

// WithPokerCardPileShuffle 通过特定的洗牌算法创建牌堆
//   - 需要保证洗牌后的牌堆剩余扑克数量与之前相同，否则将会引发 panic
func WithPokerCardPileShuffle(shuffleHandle func(pile []PokerCard) []PokerCard) PokerCardPileOption {
	return func(pile *PokerCardPile) {
		if shuffleHandle == nil {
			return
		}
		pile.shuffleHandle = shuffleHandle
	}
}

// WithPokerCardPileExcludeColor 通过排除特定花色的方式创建牌堆
func WithPokerCardPileExcludeColor(colors ...PokerColor) PokerCardPileOption {
	return func(pile *PokerCardPile) {
		if pile.excludeColor == nil {
			pile.excludeColor = map[PokerColor]struct{}{}
		}
		for _, color := range colors {
			pile.excludeColor[color] = struct{}{}
		}
	}
}

// WithPokerCardPileExcludePoint 通过排除特定点数的方式创建牌堆
func WithPokerCardPileExcludePoint(points ...PokerPoint) PokerCardPileOption {
	return func(pile *PokerCardPile) {
		if pile.excludePoint == nil {
			pile.excludePoint = map[PokerPoint]struct{}{}
		}
		for _, point := range points {
			pile.excludePoint[point] = struct{}{}
		}
	}
}

// WithPokerCardPileExcludeCard 通过排除特定扑克牌的方式创建牌堆
func WithPokerCardPileExcludeCard(cards ...PokerCard) PokerCardPileOption {
	return func(pile *PokerCardPile) {
		if pile.excludeCard == nil {
			pile.excludeCard = map[PokerPoint]map[PokerColor]struct{}{}
		}
		for _, card := range cards {
			cs, exist := pile.excludeCard[card.GetPoint()]
			if !exist {
				cs = map[PokerColor]struct{}{}
			}
			cs[card.GetColor()] = struct{}{}
		}
	}
}
