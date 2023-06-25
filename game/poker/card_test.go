package poker_test

import (
	"github.com/kercylan98/minotaur/game/poker"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCard_GetPoint(t *testing.T) {
	Convey("TestCard_GetPoint", t, func() {
		card := poker.NewCard(poker.PointA, poker.ColorSpade)
		So(card.GetPoint(), ShouldEqual, poker.PointA)
	})
}

func TestCard_GetColor(t *testing.T) {
	Convey("TestCard_GetColor", t, func() {
		card := poker.NewCard(poker.PointA, poker.ColorSpade)
		So(card.GetColor(), ShouldEqual, poker.ColorSpade)
	})
}

func TestCard_GetPointAndColor(t *testing.T) {
	Convey("TestCard_GetPointAndColor", t, func() {
		card := poker.NewCard(poker.PointA, poker.ColorSpade)
		point, color := card.GetPointAndColor()
		So(point, ShouldEqual, poker.PointA)
		So(color, ShouldEqual, poker.ColorSpade)
	})
}

func TestCard_EqualPoint(t *testing.T) {
	Convey("TestCard_EqualPoint", t, func() {
		card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
		card2 := poker.NewCard(poker.PointA, poker.ColorSpade)
		card3 := poker.NewCard(poker.Point2, poker.ColorSpade)
		So(card1.EqualPoint(card2), ShouldEqual, true)
		So(card2.EqualPoint(card3), ShouldEqual, false)
	})
}

func TestCard_EqualColor(t *testing.T) {
	Convey("TestCard_EqualColor", t, func() {
		card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
		card2 := poker.NewCard(poker.PointA, poker.ColorSpade)
		card3 := poker.NewCard(poker.PointA, poker.ColorHeart)
		So(card1.EqualColor(card2), ShouldEqual, true)
		So(card2.EqualColor(card3), ShouldEqual, false)
	})
}

func TestCard_Equal(t *testing.T) {
	Convey("TestCard_Equal", t, func() {
		card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
		card2 := poker.NewCard(poker.PointA, poker.ColorSpade)
		card3 := poker.NewCard(poker.Point2, poker.ColorHeart)
		So(card1.Equal(card2), ShouldEqual, true)
		So(card2.Equal(card3), ShouldEqual, false)
	})
}
