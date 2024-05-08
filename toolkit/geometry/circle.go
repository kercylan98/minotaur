package geometry

import "math"

// NewCircle 创建一个圆
func NewCircle(center Vector2, radius float64) Circle {
	if len(center) != 2 {
		panic("vector size mismatch")
	}
	return Circle{
		Center: center,
		Radius: radius,
	}
}

type Circle struct {
	Center Vector2
	Radius float64
}

// GetCenter 获取圆心
func (c Circle) GetCenter() Vector2 {
	if len(c.Center) != 2 {
		panic("vector size mismatch")
	}
	return c.Center
}

// GetRadius 获取半径
func (c Circle) GetRadius() float64 {
	return c.Radius
}

// GetDiameter 获取直径
func (c Circle) GetDiameter() float64 {
	return c.GetRadius() * 2
}

// GetCircumference 获取周长
func (c Circle) GetCircumference() float64 {
	return 2 * math.Pi * c.GetRadius()
}

// GetArea 获取面积
func (c Circle) GetArea() float64 {
	return math.Pi * c.GetRadius() * c.GetRadius()
}

// Contains 判断点是否在圆内
func (c Circle) Contains(point Vector2) bool {
	if len(point) != 2 {
		panic("vector size mismatch")
	}
	return c.GetCenter().Sub(point).Length() <= c.GetRadius()
}

// Intersect 判断两个圆是否相交
func (c Circle) Intersect(c2 Circle) bool {
	return c.GetCenter().Sub(c2.GetCenter()).Length() <= c.GetRadius()+c2.GetRadius()
}

// GetIntersectionPoints 获取两个圆的交点
func (c Circle) GetIntersectionPoints(c2 Circle) (Vector2, Vector2) {
	if !c.Intersect(c2) {
		panic("circles do not intersect")
	}
	d := c.GetCenter().Sub(c2.GetCenter()).Length()
	a := (c.GetRadius()*c.GetRadius() - c2.GetRadius()*c2.GetRadius() + d*d) / (2 * d)
	h := math.Sqrt(c.GetRadius()*c.GetRadius() - a*a)
	p2 := c.GetCenter().Add(c2.GetCenter().Sub(c.GetCenter()).Mul(a / d))
	x3 := p2.GetX() + h*(c.GetCenter().GetY()-c2.GetCenter().GetY())/d
	y3 := p2.GetY() + h*(c2.GetCenter().GetX()-c.GetCenter().GetX())/d
	x4 := p2.GetX() - h*(c.GetCenter().GetY()-c2.GetCenter().GetY())/d
	y4 := p2.GetY() - h*(c2.GetCenter().GetX()-c.GetCenter().GetX())/d
	return NewVector(x3, y3), NewVector(x4, y4)
}
