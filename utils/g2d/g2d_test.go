package g2d

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
	"testing"
)

func TestPositionIntToXY(t *testing.T) {
	pos := geometry.CoordinateToPos(9, 7, 8)
	fmt.Println(pos)
	fmt.Println(geometry.PosToCoordinate(9, pos))

	fmt.Println(geometry.CoordinateToPos(65000, 61411, 158266))
	fmt.Println(geometry.PosToCoordinate(65000, 10287351411))

}
