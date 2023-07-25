package poker

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math"
)

// IsContainJoker 检查扑克牌是否包含大小王
func IsContainJoker[P, C generic.Number, T Card[P, C]](pile *CardPile[P, C, T], cards ...T) bool {
	for _, card := range cards {
		if IsJoker[P, C, T](pile, card) {
			return true
		}
	}
	return false
}

// GroupByPoint 将扑克牌按照点数分组
func GroupByPoint[P, C generic.Number, T Card[P, C]](cards ...T) map[P][]T {
	group := map[P][]T{}
	for _, card := range cards {
		group[card.GetPoint()] = append(group[card.GetPoint()], card)
	}
	return group
}

// GroupByColor 将扑克牌按照花色分组
func GroupByColor[P, C generic.Number, T Card[P, C]](cards ...T) map[C][]T {
	group := map[C][]T{}
	for _, card := range cards {
		group[card.GetColor()] = append(group[card.GetColor()], card)
	}
	return group
}

// IsRocket 两张牌能否组成红黑 Joker
func IsRocket[P, C generic.Number, T Card[P, C]](pile *CardPile[P, C, T], cardA, cardB T) bool {
	var num int
	for _, joker := range pile.jokers {
		if cardA.GetPoint() == joker || cardB.GetPoint() == joker {
			num++
		}
	}
	return num == 2
}

// IsFlush 判断是否是同花
func IsFlush[P, C generic.Number, T Card[P, C]](cards ...T) bool {
	if len(cards) == 0 {
		return false
	}
	if len(cards) == 1 {
		return true
	}

	color := cards[0].GetColor()
	for i := 1; i < len(cards); i++ {
		if cards[i].GetColor() != color {
			return false
		}
	}
	return true
}

// GetCardsPoint 获取一组扑克牌的点数
func GetCardsPoint[P, C generic.Number, T Card[P, C]](cards ...T) []P {
	var points = make([]P, len(cards))
	for i, card := range cards {
		points[i] = card.GetPoint()
	}
	return points
}

// GetCardsColor 获取一组扑克牌的花色
func GetCardsColor[P, C generic.Number, T Card[P, C]](cards ...T) []C {
	var colors = make([]C, len(cards))
	for i, card := range cards {
		colors[i] = card.GetColor()
	}
	return colors
}

// IsContain 一组扑克牌是否包含某张牌
func IsContain[P, C generic.Number, T Card[P, C]](cards []T, card T) bool {
	for _, c := range cards {
		if Equal[P, C, T](c, card) {
			return true
		}
	}
	return false
}

// IsContainAll 一组扑克牌是否包含另一组扑克牌
func IsContainAll[P, C generic.Number, T Card[P, C]](cards []T, cards2 []T) bool {
	for _, card := range cards2 {
		if !IsContain[P, C, T](cards, card) {
			return false
		}
	}
	return true
}

// GetPointAndColor 返回扑克牌的点数和花色
func GetPointAndColor[P, C generic.Number, T Card[P, C]](card T) (P, C) {
	return card.GetPoint(), card.GetColor()
}

// EqualPoint 比较两张扑克牌的点数是否相同
func EqualPoint[P, C generic.Number, T Card[P, C]](card1 T, card2 T) bool {
	return card1.GetPoint() == card2.GetPoint()
}

// EqualColor 比较两张扑克牌的花色是否相同
func EqualColor[P, C generic.Number, T Card[P, C]](card1 T, card2 T) bool {
	return card1.GetColor() == card2.GetColor()
}

// Equal 比较两张扑克牌的点数和花色是否相同
func Equal[P, C generic.Number, T Card[P, C]](card1 T, card2 T) bool {
	return EqualPoint[P, C, T](card1, card2) && EqualColor[P, C, T](card1, card2)
}

// MaxPoint 返回两张扑克牌中点数较大的一张
func MaxPoint[P, C generic.Number, T Card[P, C]](card1 T, card2 T) T {
	if card1.GetPoint() > card2.GetPoint() {
		return card1
	}
	return card2
}

// MinPoint 返回两张扑克牌中点数较小的一张
func MinPoint[P, C generic.Number, T Card[P, C]](card1 T, card2 T) T {
	if card1.GetPoint() < card2.GetPoint() {
		return card1
	}
	return card2
}

// MaxColor 返回两张扑克牌中花色较大的一张
func MaxColor[P, C generic.Number, T Card[P, C]](card1 T, card2 T) T {
	if card1.GetColor() > card2.GetColor() {
		return card1
	}
	return card2
}

// MinColor 返回两张扑克牌中花色较小的一张
func MinColor[P, C generic.Number, T Card[P, C]](card1 T, card2 T) T {
	if card1.GetColor() < card2.GetColor() {
		return card1
	}
	return card2
}

// Max 返回两张扑克牌中点数和花色较大的一张
func Max[P, C generic.Number, T Card[P, C]](card1 T, card2 T) T {
	if card1.GetPoint() > card2.GetPoint() {
		return card1
	} else if card1.GetPoint() < card2.GetPoint() {
		return card2
	} else if card1.GetColor() > card2.GetColor() {
		return card1
	}
	return card2
}

// Min 返回两张扑克牌中点数和花色较小的一张
func Min[P, C generic.Number, T Card[P, C]](card1 T, card2 T) T {
	if card1.GetPoint() < card2.GetPoint() {
		return card1
	} else if card1.GetPoint() > card2.GetPoint() {
		return card2
	} else if card1.GetColor() < card2.GetColor() {
		return card1
	}
	return card2
}

// PointDifference 计算两张扑克牌的点数差
func PointDifference[P, C generic.Number, T Card[P, C]](card1 T, card2 T) int {
	return int(math.Abs(float64(card1.GetPoint()) - float64(card2.GetPoint())))
}

// ColorDifference 计算两张扑克牌的花色差
func ColorDifference[P, C generic.Number, T Card[P, C]](card1 T, card2 T) int {
	return int(math.Abs(float64(card1.GetColor()) - float64(card2.GetColor())))
}

// IsNeighborColor 判断两张扑克牌是否为相邻的花色
func IsNeighborColor[P, C generic.Number, T Card[P, C]](card1 T, card2 T) bool {
	return ColorDifference[P, C, T](card1, card2) == 1
}

// IsNeighborPoint 判断两张扑克牌是否为相邻的点数
func IsNeighborPoint[P, C generic.Number, T Card[P, C]](card1 T, card2 T) bool {
	return PointDifference[P, C, T](card1, card2) == 1
}

// IsJoker 判断扑克牌是否为大小王
func IsJoker[P, C generic.Number, T Card[P, C]](pile *CardPile[P, C, T], card T) bool {
	for _, joker := range pile.jokers {
		if card.GetPoint() == joker {
			return true
		}
	}
	return false
}
