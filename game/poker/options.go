package poker

import "fmt"

type Option func(poker *Poker)

// WithHand 通过绑定特定牌型的方式创建扑克玩法
func WithHand(pokerHand string, handle HandHandle) Option {
	return func(poker *Poker) {
		if _, exist := poker.pokerHand[pokerHand]; exist {
			panic(fmt.Errorf("same poker hand name: %s", pokerHand))
		}
		poker.pokerHand[pokerHand] = handle
		poker.pokerHandPriority = append(poker.pokerHandPriority, pokerHand)
	}
}

// WithPointValue 通过特定的扑克点数牌值创建扑克玩法
func WithPointValue(pointValues map[Point]int) Option {
	return func(poker *Poker) {
		poker.pointValue = pointValues
	}
}

// WithColorValue 通过特定的扑克花色牌值创建扑克玩法
func WithColorValue(colorValues map[Color]int) Option {
	return func(poker *Poker) {
		poker.colorValue = colorValues
	}
}

// WithPointSort 通过特定的扑克点数顺序创建扑克玩法
func WithPointSort(pointSort map[Point]int) Option {
	return func(poker *Poker) {
		for k, v := range pointSort {
			poker.pointSort[k] = v
		}
	}
}
