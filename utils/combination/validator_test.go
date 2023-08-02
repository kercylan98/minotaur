package combination_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/combination"
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

type Card struct {
	Point string
	Color string
}

func TestValidator_Validate(t *testing.T) {
	v := combination.NewValidator[*Card](
		combination.WithValidatorHandleContinuous[*Card, int](func(item *Card) int {
			switch item.Point {
			case "A":
				return 1
			case "2", "3", "4", "5", "6", "7", "8", "9", "10":
				return super.StringToInt(item.Point)
			case "J":
				return 11
			case "Q":
				return 12
			case "K":
				return 13
			}
			return -1
		}),
		combination.WithValidatorHandleLength[*Card](3),
	)

	cards := []*Card{
		{Point: "2", Color: "Spade"},
		{Point: "4", Color: "Heart"},
		{Point: "3", Color: "Diamond"},
	}

	fmt.Println(v.Validate(cards))
}
