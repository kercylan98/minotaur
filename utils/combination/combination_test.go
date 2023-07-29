package combination_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/combination"
	"testing"
)

type Poker struct {
	Point int
	Color int
}

func TestCombination_Best(t *testing.T) {
	combine := combination.NewCombination(combination.WithCombinationEvaluation(func(items []*Poker) float64 {
		var total float64
		for _, item := range items {
			total += float64(item.Point)
		}
		return total
	}))

	combine.NewMatcher("炸弹",
		combination.WithMatcherSame[*Poker, int](4, func(item *Poker) int {
			return item.Point
		}),
	).NewMatcher("三带一",
		combination.WithMatcherNCarryM[*Poker, int](3, 1, func(item *Poker) int {
			return item.Point
		}),
	)

	var cards = []*Poker{
		{Point: 2, Color: 1},
		{Point: 2, Color: 2},
		{Point: 2, Color: 3},
		{Point: 3, Color: 4},
		{Point: 4, Color: 1},
		{Point: 4, Color: 2},
		{Point: 5, Color: 3},
		{Point: 6, Color: 4},
		{Point: 7, Color: 1},
	}

	name, result := combine.Best(cards)
	fmt.Println("best:", name)
	for _, item := range result {
		fmt.Println(item)
	}
}
