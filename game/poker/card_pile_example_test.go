package poker_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/poker"
)

func ExampleNewCardPile() {
	var pile = poker.NewCardPile(1,
		poker.WithCardPileExcludeCard(poker.NewCard(poker.PointBlackJoker, poker.ColorNone)),
	)

	fmt.Println(pile.Cards())

	// Output:
	// [(R None) (K Spade) (K Heart) (K Club) (K Diamond) (Q Spade) (Q Heart) (Q Club) (Q Diamond) (J Spade) (J Heart) (J Club) (J Diamond) (10 Spade) (10 Heart) (10 Club) (10 Diamond) (9 Spade) (9 Heart) (9 Club) (9 Diamond) (8 Spade) (8 Heart) (8 Club) (8 Diamond) (7 Spade) (7 Heart) (7 Club) (7 Diamond) (6 Spade) (6 Heart) (6 Club) (6 Diamond) (5 Spade) (5 Heart) (5 Club) (5 Diamond) (4 Spade) (4 Heart) (4 Club) (4 Diamond) (3 Spade) (3 Heart) (3 Club) (3 Diamond) (2 Spade) (2 Heart) (2 Club) (2 Diamond) (A Spade) (A Heart) (A Club) (A Diamond)]
}
