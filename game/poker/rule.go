package poker

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/maths"
	"sort"
)

func NewRule(options ...Option) *Rule {
	poker := &Rule{
		pokerHand:      map[string]HandHandle{},
		pointSort:      hash.Copy(defaultPointSort),
		colorSort:      hash.Copy(defaultColorSort),
		pokerHandValue: map[string]int{},
		restraint:      map[string]map[string]struct{}{},
	}
	for _, option := range options {
		option(poker)
	}
	if poker.pointValue == nil {
		poker.pointValue = poker.pointSort
	}
	return poker
}

type Rule struct {
	pokerHand              map[string]HandHandle
	pokerHandValue         map[string]int
	pointValue             map[Point]int
	colorValue             map[Color]int
	pointSort              map[Point]int
	colorSort              map[Color]int
	excludeContinuityPoint map[Point]struct{}
	restraint              map[string]map[string]struct{}
}

// PokerHandIsMatch 检查两组扑克牌牌型是否匹配
func (slf *Rule) PokerHandIsMatch(cardsA, cardsB []Card) bool {
	handA, exist := slf.PokerHand(cardsA...)
	if !exist {
		return false
	}
	handB, exist := slf.PokerHand(cardsB...)
	if !exist {
		return false
	}
	if hash.Exist(slf.restraint[handA], handB) || hash.Exist(slf.restraint[handB], handA) {
		return false
	}
	return len(cardsA) == len(cardsB)
}

// PokerHand 获取一组扑克的牌型
func (slf *Rule) PokerHand(cards ...Card) (pokerHand string, hit bool) {
	for phn := range slf.pokerHandValue {
		if slf.pokerHand[phn](slf, cards) {
			return phn, true
		}
	}
	return HandNone, false
}

// IsPointContinuity 检查一组扑克牌点数是否连续，count 表示了每个点数的数量
func (slf *Rule) IsPointContinuity(count int, cards ...Card) bool {
	if len(cards) == 0 {
		return false
	}
	group := GroupByPoint(cards...)
	var values []int
	for point, cards := range group {
		if _, exist := slf.excludeContinuityPoint[point]; exist {
			return false
		}
		if count != len(cards) {
			return false
		}
		values = append(values, slf.GetValueWithPoint(point))
	}
	return maths.IsContinuityWithSort(values)
}

// IsSameColor 检查一组扑克牌是否同花
func (slf *Rule) IsSameColor(cards ...Card) bool {
	length := len(cards)
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	var color = cards[0].GetColor()
	for i := 1; i < length; i++ {
		if cards[i].GetColor() != color {
			return false
		}
	}
	return true
}

// IsSamePoint 检查一组扑克牌是否同点
func (slf *Rule) IsSamePoint(cards ...Card) bool {
	length := len(cards)
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	var point = cards[0].GetPoint()
	for i := 1; i < length; i++ {
		if cards[i].GetPoint() != point {
			return false
		}
	}
	return true
}

// SortByPointDesc 将扑克牌按照点数从大到小排序
func (slf *Rule) SortByPointDesc(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return slf.pointSort[cards[i].GetPoint()] > slf.pointSort[cards[j].GetPoint()]
	})
	return cards
}

// SortByPointAsc 将扑克牌按照点数从小到大排序
func (slf *Rule) SortByPointAsc(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return slf.pointSort[cards[i].GetPoint()] < slf.pointSort[cards[j].GetPoint()]
	})
	return cards
}

// SortByColorDesc 将扑克牌按照花色从大到小排序
func (slf *Rule) SortByColorDesc(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return slf.colorSort[cards[i].GetColor()] > slf.colorSort[cards[j].GetColor()]
	})
	return cards
}

// SortByColorAsc 将扑克牌按照花色从小到大排序
func (slf *Rule) SortByColorAsc(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return slf.colorSort[cards[i].GetColor()] < slf.colorSort[cards[j].GetColor()]
	})
	return cards
}

// GetValueWithPokerHand 获取扑克牌的牌型牌值
func (slf *Rule) GetValueWithPokerHand(hand string) int {
	return slf.pokerHandValue[hand]
}

// GetValueWithCards 获取扑克牌的牌值
func (slf *Rule) GetValueWithCards(cards ...Card) int {
	var value int
	for _, card := range cards {
		value += slf.pointValue[card.GetPoint()]
		value += slf.colorValue[card.GetColor()]
	}
	return value
}

// GetValueWithPoint 获取扑克牌的点数牌值
func (slf *Rule) GetValueWithPoint(point Point) int {
	return slf.pointValue[point]
}

// GetValueWithColor 获取扑克牌的花色牌值
func (slf *Rule) GetValueWithColor(color Color) int {
	return slf.colorValue[color]
}

// CompareValueWithCards 根据特定的条件表达式比较两组扑克牌的牌值
func (slf *Rule) CompareValueWithCards(cards1 []Card, expression maths.CompareExpression, cards2 []Card) bool {
	return maths.Compare(slf.GetValueWithCards(cards1...), expression, slf.GetValueWithCards(cards2...))
}
