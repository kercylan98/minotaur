package poker

type CardPileOption func(pile *CardPile)

// WithCardPileShuffle 通过特定的洗牌算法创建牌堆
//   - 需要保证洗牌后的牌堆剩余扑克数量与之前相同，否则将会引发 panic
func WithCardPileShuffle(shuffleHandle func(pile []Card) []Card) CardPileOption {
	return func(pile *CardPile) {
		if shuffleHandle == nil {
			return
		}
		pile.shuffleHandle = shuffleHandle
	}
}

// WithCardPileExcludeColor 通过排除特定花色的方式创建牌堆
func WithCardPileExcludeColor(colors ...Color) CardPileOption {
	return func(pile *CardPile) {
		if pile.excludeColor == nil {
			pile.excludeColor = map[Color]struct{}{}
		}
		for _, color := range colors {
			pile.excludeColor[color] = struct{}{}
		}
	}
}

// WithCardPileExcludePoint 通过排除特定点数的方式创建牌堆
func WithCardPileExcludePoint(points ...Point) CardPileOption {
	return func(pile *CardPile) {
		if pile.excludePoint == nil {
			pile.excludePoint = map[Point]struct{}{}
		}
		for _, point := range points {
			pile.excludePoint[point] = struct{}{}
		}
	}
}

// WithCardPileExcludeCard 通过排除特定扑克牌的方式创建牌堆
func WithCardPileExcludeCard(cards ...Card) CardPileOption {
	return func(pile *CardPile) {
		if pile.excludeCard == nil {
			pile.excludeCard = map[Point]map[Color]struct{}{}
		}
		for _, card := range cards {
			point := card.GetPoint()
			cs, exist := pile.excludeCard[point]
			if !exist {
				cs = map[Color]struct{}{}
				pile.excludeCard[point] = cs
			}
			if point == PointRedJoker || point == PointBlackJoker {
				cs[ColorNone] = struct{}{}
			} else {
				cs[card.GetColor()] = struct{}{}
			}
		}
	}
}
