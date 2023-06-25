package poker_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/poker"
)

func ExampleNewCard() {
	card := poker.NewCard(poker.PointA, poker.ColorSpade)
	fmt.Println(card)

	// Output:
	// (A Spade)
}
