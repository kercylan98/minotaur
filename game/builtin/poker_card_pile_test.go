package builtin

import (
	"fmt"
	"testing"
)

func TestNewPokerCardPile(t *testing.T) {
	pile := NewPokerCardPile(2)
	_ = pile.PullTop()
	pile.Reset()
	fmt.Println(pile)
}
