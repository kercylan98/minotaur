package g2d

import (
	"fmt"
	"testing"
)

func TestPositionIntToXY(t *testing.T) {
	pos := PositionToInt(9, 7, 8)
	fmt.Println(pos)
	fmt.Println(PositionIntToXY(9, pos))

	fmt.Println(PositionToInt(65000, 61411, 158266))
	fmt.Println(PositionIntToXY(65000, 10287351411))

}
