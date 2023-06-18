package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math"
)

// GenerateCircle 通过传入圆的半径和需要的点数量，生成一个圆
func GenerateCircle[V generic.SignedNumber](radius V, points int) Shape[V] {
	angle := 2.0 * math.Pi / float64(points)
	var shape = make(Shape[V], points)
	for i := 0; i < points; i++ {
		curAngle := float64(i) * angle
		x := radius * V(math.Cos(curAngle))
		y := radius * V(math.Sin(curAngle))

		shape = append(shape, NewPoint(x, y))
	}
	return shape
}
