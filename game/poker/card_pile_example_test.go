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
	// [(R None) (K Diamond) (K Club) (K Heart) (K Spade) (Q Diamond) (Q Club) (Q Heart) (Q Spade) (J Diamond) (J Club) (J Heart) (J Spade) (10 Diamond) (10 Club) (10 Heart) (10 Spade) (9 Diamond) (9 Club) (9 Heart) (9 Spade) (8 Diamond) (8 Club) (8 Heart) (8 Spade) (7 Diamond) (7 Club) (7 Heart) (7 Spade) (6 Diamond) (6 Club) (6 Heart) (6 Spade) (5 Diamond) (5 Club) (5 Heart) (5 Spade) (4 Diamond) (4 Club) (4 Heart) (4 Spade) (3 Diamond) (3 Club) (3 Heart) (3 Spade) (2 Diamond) (2 Club) (2 Heart) (2 Spade) (A Diamond) (A Club) (A Heart) (A Spade)]
}
