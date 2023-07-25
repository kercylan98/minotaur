package poker_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/poker"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
	"time"
)

type Card[P, C generic.Number] struct {
	guid  int64
	point P
	color C
}

func (slf *Card[P, C]) GetGuid() int64 {
	return slf.guid
}

func (slf *Card[P, C]) GetPoint() P {
	return slf.point
}

func (slf *Card[P, C]) GetColor() C {
	return slf.color
}

func TestMatcher_Group(t *testing.T) {

	evaluate := func(cards []*Card[int, int]) int64 {
		score := int64(0)
		for _, card := range cards {
			if card.point == 1 {
				score += 14
				continue
			}
			score += int64(card.GetPoint())
		}
		return score
	}

	matcher := poker.NewMatcher[int, int, *Card[int, int]]()
	//matcher.RegType("三条", evaluate,
	//	poker.WithMatcherNCarryMSingle[*Card](3, 2))
	matcher.RegType("皇家同花顺", evaluate,
		poker.WithMatcherFlush[int, int, *Card[int, int]](),
		poker.WithMatcherContinuityPointOrder[int, int, *Card[int, int]](map[int]int{1: 14}),
		poker.WithMatcherLength[int, int, *Card[int, int]](5),
	).RegType("同花顺", evaluate,
		poker.WithMatcherFlush[int, int, *Card[int, int]](),
		poker.WithMatcherContinuityPointOrder[int, int, *Card[int, int]](map[int]int{1: 14}),
		poker.WithMatcherLeastLength[int, int, *Card[int, int]](3),
	).RegType("四条", evaluate,
		poker.WithMatcherTieCount[int, int, *Card[int, int]](4),
	).RegType("葫芦", evaluate,
		poker.WithMatcherNCarryM[int, int, *Card[int, int]](3, 2),
	).RegType("顺子", evaluate,
		poker.WithMatcherContinuityPointOrder[int, int, *Card[int, int]](map[int]int{1: 14}),
		poker.WithMatcherLength[int, int, *Card[int, int]](5),
	).RegType("三条", evaluate,
		poker.WithMatcherNCarryMSingle[int, int, *Card[int, int]](3, 2),
	).RegType("两对", evaluate,
		poker.WithMatcherTieCountNum[int, int, *Card[int, int]](2, 2),
	).RegType("一对", evaluate,
		poker.WithMatcherTieCount[int, int, *Card[int, int]](2),
	).RegType("高牌", evaluate,
		poker.WithMatcherTieCount[int, int, *Card[int, int]](1),
	)

	var pub = []*Card[int, int]{
		{point: 4, color: 3},
		{point: 5, color: 2},
		{point: 6, color: 1},
		{point: 6, color: 2},
		{point: 13, color: 2},
	}

	var pri = []*Card[int, int]{
		{point: 1, color: 1},
		{point: 1, color: 2},
		{point: 4, color: 3},
		{point: 5, color: 4},
	}

	var start = time.Now()
	var usePub, usePri = slice.LimitedCombinations(pub, 3, 3), slice.LimitedCombinations(pri, 2, 2)

	var topResult []*Card[int, int]
	var topScore int64
	var topName string
	var source []*Card[int, int]
	for _, handCards := range usePri {
		for _, pubCards := range usePub {
			cards := append(handCards, pubCards...)
			name, result := matcher.Group(cards)
			score := evaluate(result)
			if score > topScore || topResult == nil {
				topScore = score
				topResult = result
				topName = name
				source = cards
			}
		}
	}

	fmt.Println("time:", time.Since(start))
	fmt.Println("result:", topName)
	for _, card := range topResult {
		fmt.Println(fmt.Sprintf("Point: %d Color: %d", card.GetPoint(), card.GetColor()))
	}
	fmt.Println("source:", topScore)
	for _, card := range source {
		fmt.Println(fmt.Sprintf("Point: %d Color: %d", card.GetPoint(), card.GetColor()))
	}
}
