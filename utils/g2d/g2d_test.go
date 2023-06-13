package g2d

import (
	"fmt"
	"testing"
)

func TestPositionIntToXY(t *testing.T) {
	pos := CoordinateToPos(9, 7, 8)
	fmt.Println(pos)
	fmt.Println(PosToCoordinate(9, pos))

	fmt.Println(CoordinateToPos(65000, 61411, 158266))
	fmt.Println(PosToCoordinate(65000, 10287351411))

}
