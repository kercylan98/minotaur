package poker

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/maths"
	"sort"
)

func NewRule[P, C generic.Number, T Card[P, C]](options ...Option[P, C, T]) *Rule[P, C, T] {
	poker := &Rule[P, C, T]{
		pokerHand: map[string]HandHandle[P, C, T]{},
		//pointSort:      hash.Copy(defaultPointSort),
		//colorSort:      hash.Copy(defaultColorSort),
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

type Rule[P, C generic.Number, T Card[P, C]] struct {
	pokerHand              map[string]HandHandle[P, C, T]
	pokerHandValue         map[string]int
	pointValue             map[P]int
	colorValue             map[C]int
	pointSort              map[P]int
	colorSort              map[C]int
	excludeContinuityPoint map[P]struct{}
	restraint              map[string]map[string]struct{}
}

// GetCardCountWithPointMaximumNumber 获取指定点数的牌的数量
func (slf *Rule[P, C, T]) GetCardCountWithPointMaximumNumber(cards []T, point P, maximumNumber int) int {
	count := 0
	for _, card := range cards {
		if card.GetPoint() == point {
			count++
			if count >= maximumNumber {
				return count
			}
		}
	}
	return count
}

// GetCardCountWithColorMaximumNumber 获取指定花色的牌的数量
func (slf *Rule[P, C, T]) GetCardCountWithColorMaximumNumber(cards []T, color C, maximumNumber int) int {
	count := 0
	for _, card := range cards {
		if card.GetColor() == color {
			count++
			if count >= maximumNumber {
				return count
			}
		}
	}
	return count
}

// GetCardCountWithMaximumNumber 获取指定牌的数量
func (slf *Rule[P, C, T]) GetCardCountWithMaximumNumber(cards []T, card T, maximumNumber int) int {
	count := 0
	for _, c := range cards {
		if Equal[P, C, T](c, card) {
			count++
			if count >= maximumNumber {
				return count
			}
		}
	}
	return count
}

// GetCardCountWithPoint 获取指定点数的牌的数量
func (slf *Rule[P, C, T]) GetCardCountWithPoint(cards []T, point P) int {
	count := 0
	for _, card := range cards {
		if card.GetPoint() == point {
			count++
		}
	}
	return count
}

// GetCardCountWithColor 获取指定花色的牌的数量
func (slf *Rule[P, C, T]) GetCardCountWithColor(cards []T, color C) int {
	count := 0
	for _, card := range cards {
		if card.GetColor() == color {
			count++
		}
	}
	return count
}

// GetCardCount 获取指定牌的数量
func (slf *Rule[P, C, T]) GetCardCount(cards []T, card T) int {
	count := 0
	for _, c := range cards {
		if Equal[P, C, T](c, card) {
			count++
		}
	}
	return count
}

// PokerHandIsMatch 检查两组扑克牌牌型是否匹配
func (slf *Rule[P, C, T]) PokerHandIsMatch(cardsA, cardsB []T) bool {
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
func (slf *Rule[P, C, T]) PokerHand(cards ...T) (pokerHand string, hit bool) {
	for phn := range slf.pokerHandValue {
		if slf.pokerHand[phn](slf, cards) {
			return phn, true
		}
	}
	return HandNone, false
}

// IsPointContinuity 检查一组扑克牌点数是否连续，count 表示了每个点数的数量
func (slf *Rule[P, C, T]) IsPointContinuity(count int, cards ...T) bool {
	if len(cards) == 0 {
		return false
	}
	group := GroupByPoint[P, C, T](cards...)
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
func (slf *Rule[P, C, T]) IsSameColor(cards ...T) bool {
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
func (slf *Rule[P, C, T]) IsSamePoint(cards ...T) bool {
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
func (slf *Rule[P, C, T]) SortByPointDesc(cards []T) []T {
	sort.Slice(cards, func(i, j int) bool {
		return slf.pointSort[cards[i].GetPoint()] > slf.pointSort[cards[j].GetPoint()]
	})
	return cards
}

// SortByPointAsc 将扑克牌按照点数从小到大排序
func (slf *Rule[P, C, T]) SortByPointAsc(cards []T) []T {
	sort.Slice(cards, func(i, j int) bool {
		return slf.pointSort[cards[i].GetPoint()] < slf.pointSort[cards[j].GetPoint()]
	})
	return cards
}

// SortByColorDesc 将扑克牌按照花色从大到小排序
func (slf *Rule[P, C, T]) SortByColorDesc(cards []T) []T {
	sort.Slice(cards, func(i, j int) bool {
		return slf.colorSort[cards[i].GetColor()] > slf.colorSort[cards[j].GetColor()]
	})
	return cards
}

// SortByColorAsc 将扑克牌按照花色从小到大排序
func (slf *Rule[P, C, T]) SortByColorAsc(cards []T) []T {
	sort.Slice(cards, func(i, j int) bool {
		return slf.colorSort[cards[i].GetColor()] < slf.colorSort[cards[j].GetColor()]
	})
	return cards
}

// GetValueWithPokerHand 获取扑克牌的牌型牌值
func (slf *Rule[P, C, T]) GetValueWithPokerHand(hand string, cards ...T) int {
	hv := slf.pokerHandValue[hand]
	return hv * slf.GetValueWithCards(cards...)
}

// GetValueWithCards 获取扑克牌的牌值
func (slf *Rule[P, C, T]) GetValueWithCards(cards ...T) int {
	var value int
	for _, card := range cards {
		value += slf.GetValueWithPoint(card.GetPoint())
		value += slf.GetValueWithColor(card.GetColor())
	}
	return value
}

// GetValueWithPoint 获取扑克牌的点数牌值
func (slf *Rule[P, C, T]) GetValueWithPoint(point P) int {
	return slf.pointValue[point]
}

// GetValueWithColor 获取扑克牌的花色牌值
func (slf *Rule[P, C, T]) GetValueWithColor(color C) int {
	return slf.colorValue[color]
}

// CompareValueWithCards 根据特定的条件表达式比较两组扑克牌的牌值
func (slf *Rule[P, C, T]) CompareValueWithCards(cards1 []T, expression maths.CompareExpression, cards2 []T) bool {
	return maths.Compare(slf.GetValueWithCards(cards1...), expression, slf.GetValueWithCards(cards2...))
}
