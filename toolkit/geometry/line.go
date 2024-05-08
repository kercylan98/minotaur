package geometry

import "math"

type Line [2]Vector2

// NewLine 创建一条直线
func NewLine(start, end Vector2) Line {
	if len(start) != 2 || len(end) != 2 {
		panic("vector size mismatch")
	}
	return Line{start, end}
}

// GetStart 获取起点
func (l Line) GetStart() Vector2 {
	return l[0]
}

// GetEnd 获取终点
func (l Line) GetEnd() Vector2 {
	return l[1]
}

// GetLength 获取长度
func (l Line) GetLength() float64 {
	return l.GetStart().Sub(l.GetEnd()).Length()
}

// Contains 判断点是否在直线上
func (l Line) Contains(point Vector2) bool {
	if len(point) != 2 {
		panic("vector size mismatch")
	}

	v1 := l.GetEnd().Sub(l.GetStart()) // 直线的向量
	v2 := point.Sub(l.GetStart())      // 直线起点到给定点的向量

	// 计算交叉积
	crossProduct := v1[0]*v2[1] - v1[1]*v2[0]

	// 数值误差
	epsilon := 1e-9
	return math.Abs(crossProduct) < epsilon
}

// Intersect 判断两条直线是否相交
func (l Line) Intersect(l2 Line) bool {
	// 两条直线的向量
	v1 := l.GetEnd().Sub(l.GetStart())
	v2 := l2.GetEnd().Sub(l2.GetStart())

	// 计算交叉积
	crossProduct1 := v1[0]*v2[1] - v1[1]*v2[0]

	// 两条直线的起点到另一条直线起点的向量
	v3 := l2.GetStart().Sub(l.GetStart())
	v4 := l2.GetEnd().Sub(l.GetStart())

	// 计算交叉积
	crossProduct2 := v1[0]*v3[1] - v1[1]*v3[0]
	crossProduct3 := v1[0]*v4[1] - v1[1]*v4[0]

	return crossProduct1 != 0 && crossProduct2*crossProduct3 < 0
}

// IntersectCircle 判断直线是否与圆相交
func (l Line) IntersectCircle(c Circle) bool {
	// 直线的向量
	v1 := l.GetEnd().Sub(l.GetStart())

	// 直线起点到圆心的向量
	v2 := c.GetCenter().Sub(l.GetStart())

	// 计算直线到圆心的投影长度
	projection := v1.Dot(v2) / v1.Length()

	// 计算直线到圆心的距离
	distance := v2.Length()

	return distance <= c.GetRadius() && projection >= 0 && projection <= v1.Length()
}
