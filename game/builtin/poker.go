package builtin

import "github.com/kercylan98/minotaur/utils/maths"

func NewPoker[PlayerID comparable](pile *PokerCardPile, options ...PokerOption[PlayerID]) *Poker[PlayerID] {
	poker := &Poker[PlayerID]{
		pile:      pile,
		handCards: map[PlayerID][][]PokerCard{},
	}
	for _, option := range options {
		option(poker)
	}
	return poker
}

type Poker[PlayerID comparable] struct {
	pile         *PokerCardPile
	comparePoint map[PokerPoint]int
	compareColor map[PokerColor]int
	handCards    map[PlayerID][][]PokerCard
}

// HandCard 获取玩家特定索引的手牌组
func (slf *Poker[PlayerID]) HandCard(playerId PlayerID, index int) []PokerCard {
	return slf.handCards[playerId][index]
}

// HandCards 获取玩家所有手牌
//   - 获取结果为多份手牌
func (slf *Poker[PlayerID]) HandCards(playerId PlayerID) [][]PokerCard {
	return slf.handCards[playerId]
}

// HandCardGroupCount 获取玩家共有多少副手牌
func (slf *Poker[PlayerID]) HandCardGroupCount(playerId PlayerID) int {
	return len(slf.handCards[playerId])
}

// GetPile 获取牌堆
func (slf *Poker[PlayerID]) GetPile() *PokerCardPile {
	return slf.pile
}

// Compare 比较两张扑克牌大小
func (slf *Poker[PlayerID]) Compare(card1 PokerCard, expression maths.CompareExpression, card2 PokerCard) bool {
	var point1, point2 int
	if slf.comparePoint == nil {
		point1, point2 = int(card1.GetPoint()), int(card2.GetPoint())
	} else {
		point1, point2 = slf.comparePoint[card1.GetPoint()], slf.comparePoint[card2.GetPoint()]
	}
	if maths.Compare(point1, expression, point2) {
		return true
	}
	var color1, color2 int
	if slf.comparePoint == nil {
		color1, color2 = int(card1.GetColor()), int(card2.GetColor())
	} else {
		color1, color2 = slf.compareColor[card1.GetColor()], slf.compareColor[card2.GetColor()]
	}
	return maths.Compare(color1, expression, color2)
}
