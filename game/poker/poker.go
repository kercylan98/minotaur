package poker

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/maths"
	"sort"
)

func New(pile *CardPile, options ...Option) *Poker {
	poker := &Poker{
		pile:      pile,
		pokerHand: map[string]HandHandle{},
		pointSort: hash.Copy(defaultPointSort),
	}
	for _, option := range options {
		option(poker)
	}
	if poker.pointValue == nil {
		poker.pointValue = poker.pointSort
	}
	return poker
}

type Poker struct {
	pile              *CardPile
	pokerHand         map[string]HandHandle
	pokerHandPriority []string
	pointValue        map[Point]int
	colorValue        map[Color]int
	pointSort         map[Point]int
}

// IsContinuity 检查一组扑克牌是否连续
func (slf *Poker) IsContinuity(cards ...Card) bool {
	length := len(cards)
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	var points = make([]int, length)
	for i, card := range cards {
		points[i] = slf.pointSort[card.GetPoint()]
	}
	sort.Slice(points, func(i, j int) bool { return points[i] < points[j] })
	for i := 0; i < length-1; i++ {
		if points[i+1]-points[i] != 1 {
			return false
		}
	}
	return true
}

// CardValue 获取扑克牌的牌值
func (slf *Poker) CardValue(cards ...Card) int {
	var value int
	for _, card := range cards {
		value += slf.pointValue[card.GetPoint()]
		value += slf.colorValue[card.GetColor()]
	}
	return value
}

// Compare 根据特定的条件表达式比较两组扑克牌的牌值
func (slf *Poker) Compare(cards1 []Card, expression maths.CompareExpression, cards2 []Card) bool {
	return maths.Compare(slf.CardValue(cards1...), expression, slf.CardValue(cards2...))
}

// PokerHand 获取一组扑克的牌型
//
// 参数：
//   - cards: 扑克牌切片，类型为 []builtin.Card，表示一组扑克牌。
//
// 返回值：
//   - string: 命中的牌型名称。
//   - bool: 是否命中牌型。
func (slf *Poker) PokerHand(cards ...Card) (cardType string, hit bool) {
	for _, phn := range slf.pokerHandPriority {
		if slf.pokerHand[phn](slf, cards) {
			return phn, true
		}
	}
	return "", false
}

// GetPile 获取牌堆
func (slf *Poker) GetPile() *CardPile {
	return slf.pile
}
