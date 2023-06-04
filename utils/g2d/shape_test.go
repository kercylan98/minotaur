package g2d

import (
	"fmt"
	"testing"
)

func TestGetShapeCoverageArea(t *testing.T) {
	for _, xy := range GetExpressibleRectangleBySize(2, 3, 2, 2) {
		for y := 0; y < xy[1]+1; y++ {
			for x := 0; x < xy[0]+1; x++ {
				fmt.Print("0", " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func TestGetExpressibleRectangleBySize(t *testing.T) {
	for _, xy := range GetExpressibleRectangleBySize(3, 3, 2, 2) {
		for y := 0; y < xy[1]+1; y++ {
			for x := 0; x < xy[0]+1; x++ {
				fmt.Print("0", " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
