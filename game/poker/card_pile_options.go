package poker

import "github.com/kercylan98/minotaur/utils/generic"

type CardPileOption[P, C generic.Number, T Card[P, C]] func(pile *CardPile[P, C, T])

// WithCardPileShuffle 通过特定的洗牌算法创建牌堆
//   - 需要保证洗牌后的牌堆剩余扑克数量与之前相同，否则将会引发 panic
func WithCardPileShuffle[P, C generic.Number, T Card[P, C]](shuffleHandle func(pile []T) []T) CardPileOption[P, C, T] {
	return func(pile *CardPile[P, C, T]) {
		if shuffleHandle == nil {
			return
		}
		pile.shuffleHandle = shuffleHandle
	}
}

// WithCardPileExcludeColor 通过排除特定花色的方式创建牌堆
func WithCardPileExcludeColor[P, C generic.Number, T Card[P, C]](colors ...C) CardPileOption[P, C, T] {
	return func(pile *CardPile[P, C, T]) {
		if pile.excludeColor == nil {
			pile.excludeColor = map[C]struct{}{}
		}
		for _, color := range colors {
			pile.excludeColor[color] = struct{}{}
		}
	}
}

// WithCardPileExcludePoint 通过排除特定点数的方式创建牌堆
func WithCardPileExcludePoint[P, C generic.Number, T Card[P, C]](points ...P) CardPileOption[P, C, T] {
	return func(pile *CardPile[P, C, T]) {
		if pile.excludePoint == nil {
			pile.excludePoint = map[P]struct{}{}
		}
		for _, point := range points {
			pile.excludePoint[point] = struct{}{}
		}
	}
}

// WithCardPileExcludeCard 通过排除特定扑克牌的方式创建牌堆
func WithCardPileExcludeCard[P, C generic.Number, T Card[P, C]](cards ...Card[P, C]) CardPileOption[P, C, T] {
	return func(pile *CardPile[P, C, T]) {
		if pile.excludeCard == nil {
			pile.excludeCard = map[P]map[C]struct{}{}
		}
		for _, card := range cards {
			point := card.GetPoint()
			cs, exist := pile.excludeCard[point]
			if !exist {
				cs = map[C]struct{}{}
				pile.excludeCard[point] = cs
			}
			for _, joker := range pile.jokers {
				if point != joker {
					cs[card.GetColor()] = struct{}{}
				}
			}
		}
	}
}
