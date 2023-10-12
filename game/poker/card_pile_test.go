package poker_test

import (
	"github.com/kercylan98/minotaur/game/poker"
	"testing"
)

type Card struct {
	Guid  int64
	Point int32
	Color int32
}

func (slf *Card) GetGuid() int64 {
	return slf.Guid
}

func (slf *Card) GetPoint() int32 {
	return slf.Point
}

func (slf *Card) GetColor() int32 {
	return slf.Color
}

func TestCardPile_PullTop(t *testing.T) {
	var pile = poker.NewCardPile[int32, int32, *Card](6,
		[2]int32{14, 15},
		[13]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		[4]int32{1, 2, 3, 4},
		func(guid int64, point int32, color int32) *Card {
			return &Card{Guid: guid, Point: point, Color: color}
		},
	)

	pile.Shuffle()
	for i := 0; i < 10; i++ {
		t.Log(pile.PullTop())
	}
}
