package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math"
)

// Circle 圆形
type Circle[V generic.SignedNumber] struct {
	Shape[V]
	radius       V
	centroid     Point[V]
	initCentroid bool
}

// Radius 获取圆形半径
func (slf Circle[V]) Radius() V {
	if slf.radius > V(-0) {
		return slf.radius
	}
	for _, point := range slf.Points() {
		slf.radius = CalcDistanceWithPoint(slf.Centroid(), point)
		return slf.radius
	}
	panic("circle without any points")
}

// Centroid 获取圆形质心位置
func (slf Circle[V]) Centroid() Point[V] {
	if slf.initCentroid {
		return slf.centroid
	}
	slf.centroid = CalcRectangleCentroid(slf.Shape)
	slf.initCentroid = true
	return slf.centroid
}

// Overlap 与另一个圆是否发生重叠
func (slf Circle[V]) Overlap(circle Circle[V]) bool {
	return slf.CentroidDistance(circle) < slf.Radius()+circle.Radius()
}

// Area 获取圆形面积
func (slf Circle[V]) Area() V {
	return V(math.Pi * math.Pow(float64(slf.Radius()), 2))
}

// Length 获取圆的周长
func (slf Circle[V]) Length() V {
	return V(2 * math.Pi * float64(slf.Radius()))
}

// CentroidDistance 计算与另一个圆的质心距离
func (slf Circle[V]) CentroidDistance(circle Circle[V]) V {
	return CalcCircleCentroidDistance(slf, circle)
}

// NewCircle 通过传入圆的半径和需要的点数量，生成一个圆
func NewCircle[V generic.SignedNumber](radius V, points int) Circle[V] {
	angle := 2.0 * math.Pi / float64(points)
	var shape = make(Shape[V], points)
	for i := 0; i < points; i++ {
		curAngle := float64(i) * angle
		x := radius * V(math.Cos(curAngle))
		y := radius * V(math.Sin(curAngle))
		shape = append(shape, NewPoint(x, y))
	}
	return shape.ToCircle()
}

// CalcCircleCentroidDistance 计算两个圆质心距离
func CalcCircleCentroidDistance[V generic.SignedNumber](circle1, circle2 Circle[V]) V {
	return CalcDistanceWithPoint(circle1.Centroid(), circle2.Centroid())
}
