package geometry_test

import (
	"github.com/kercylan98/minotaur/utils/geometry"
	"testing"
)

func TestSimpleCircle_RandomSubCircle(t *testing.T) {
	for i := 0; i < 10; i++ {
		sc := geometry.NewSimpleCircle(10, geometry.NewPoint(0, 0))
		sub := sc.RandomCircleWithinParent(8)

		t.Log(sc)
		t.Log(sub)
		t.Log(sc.CentroidDistance(sub))
	}
}
