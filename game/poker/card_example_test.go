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

func ExampleCard_Equal() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorSpade)
	fmt.Println(card1.Equal(card2))

	// Output:
	// true
}

func ExampleCard_EqualColor() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.EqualColor(card2))

	// Output:
	// false
}

func ExampleCard_EqualPoint() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.EqualPoint(card2))

	// Output:
	// true
}

func ExampleCard_GetColor() {
	card := poker.NewCard(poker.PointA, poker.ColorSpade)
	fmt.Println(card.GetColor())

	// Output:
	// Spade
}

func ExampleCard_GetPoint() {
	card := poker.NewCard(poker.PointA, poker.ColorSpade)
	fmt.Println(card.GetPoint())

	// Output:
	// A
}

func ExampleCard_GetPointAndColor() {
	card := poker.NewCard(poker.PointA, poker.ColorSpade)
	fmt.Println(card.GetPointAndColor())

	// Output:
	// A Spade
}

func ExampleCard_MaxPoint() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.Point2, poker.ColorSpade)
	fmt.Println(card1.MaxPoint(card2))

	// Output:
	// (2 Spade)
}

func ExampleCard_MinPoint() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.Point2, poker.ColorSpade)
	fmt.Println(card1.MinPoint(card2))

	// Output:
	// (A Spade)
}

func ExampleCard_String() {
	card := poker.NewCard(poker.PointA, poker.ColorSpade)
	fmt.Println(card.String())

	// Output:
	// (A Spade)
}

func ExampleCard_MaxColor() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.MaxColor(card2))

	// Output:
	// (A Spade)
}

func ExampleCard_MinColor() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.MinColor(card2))

	// Output:
	// (A Heart)
}

func ExampleCard_Max() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.Max(card2))

	// Output:
	// (A Spade)
}

func ExampleCard_Min() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.Min(card2))

	// Output:
	// (A Heart)
}

func ExampleCard_IsJoker() {
	card := poker.NewCard(poker.PointA, poker.ColorSpade)
	fmt.Println(card.IsJoker())

	// Output:
	// false
}

func ExampleCard_CalcPointDifference() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.Point2, poker.ColorSpade)
	fmt.Println(card1.CalcPointDifference(card2))

	// Output:
	// -1
}

func ExampleCard_CalcColorDifference() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.CalcColorDifference(card2))

	// Output:
	// 1
}

func ExampleCard_CalcPointDifferenceAbs() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.Point2, poker.ColorSpade)
	fmt.Println(card1.CalcPointDifferenceAbs(card2))

	// Output:
	// 1
}

func ExampleCard_CalcColorDifferenceAbs() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.CalcColorDifferenceAbs(card2))

	// Output:
	// 1
}

func ExampleCard_IsNeighborPoint() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.Point2, poker.ColorSpade)
	fmt.Println(card1.IsNeighborPoint(card2))

	// Output:
	// true
}

func ExampleCard_IsNeighborColor() {
	card1 := poker.NewCard(poker.PointA, poker.ColorSpade)
	card2 := poker.NewCard(poker.PointA, poker.ColorHeart)
	fmt.Println(card1.IsNeighborColor(card2))

	// Output:
	// true
}
