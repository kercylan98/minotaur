package geometry

import (
	"fmt"
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := [2]int{1, 1}
	fmt.Println(PointToPos(9, p))
}
