package geometry_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
	"testing"
)

func TestGenerateCircle(t *testing.T) {
	fmt.Println(geometry.GenerateCircle[float64](5, 12))
}
