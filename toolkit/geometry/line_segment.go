package geometry

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/maths"
	"math"
	"sort"
)

// NewLineSegment 创建一个线段
func NewLineSegment(points ...Point) LineSegment {
	AssertLineSegmentValid(points)
	return points
}

// LineSegment 由至少两个点组成的线段
type LineSegment []Point

// GetLength 获取线段的长度
func (l LineSegment) GetLength() float64 {
	var length float64
	for i := 0; i < len(l)-1; i++ {
		length += l[i].Distance2D(l[i+1])
	}
	return length
}

// GetDirection 获取线段的方向
func (l LineSegment) GetDirection() Vector2 {
	return l[0].Sub(l[1]).Normalize()
}

// GetMidpoint 获取线段的中点
func (l LineSegment) GetMidpoint() Point {
	return l[0].Add(l[1]).Div(2)
}

// IsPointOnSegment 判断一个给定的点是否在该线段上，允许一定的误差来包容浮点数计算的误差，但不允许点超出线段的范围
func (l LineSegment) IsPointOnSegment(point Point) bool {
	d1 := point.Distance2D(l[0])
	d2 := point.Distance2D(l[1])
	length := l.GetLength()

	// 确保距离之和接近线段长度
	isClose := math.Abs((d1+d2)-length) < 1e-9

	if !isClose {
		return false
	}

	// 确保点的坐标在两个端点的范围内
	inXRange := (point.GetX() >= math.Min(l[0].GetX(), l[1].GetX())) && (point.GetX() <= math.Max(l[0].GetX(), l[1].GetX()))
	inYRange := (point.GetY() >= math.Min(l[0].GetY(), l[1].GetY())) && (point.GetY() <= math.Max(l[0].GetY(), l[1].GetY()))

	return inXRange && inYRange
}

// ClosestPoint 计算一个点到该线段的最近的点
func (l LineSegment) ClosestPoint(point Point) Point {
	ax, ay := l[0].GetXY()
	bx, by := l[1].GetXY()
	ds := l[0].DistanceSquared2D(l[1])
	px, py := point.GetXY()
	clamp := maths.Clamp((px-ax)*(bx-ax)+(py-ay)*(by-ay)/ds, 0, 1)
	return NewPoint(ax+clamp*(bx-ax), ay+clamp*(by-ay))
}

// CalcLineSegmentPointProjection 计算点在线段上的投影
//   - 根据向量的投影公式计算点在线段上的投影，然后将点投影到线段上，得到投影点
//   - 参考：https://en.wikipedia.org/wiki/Vector_projection
func CalcLineSegmentPointProjection(line LineSegment, point Point) Point {
	lineDir := line.GetDirection()
	pointDir := line[0].Sub(point)
	pointProjection := lineDir.Mul(pointDir.Dot(lineDir))
	return line[0].Sub(pointProjection)
}

// CalcLineSegmentDistanceToPoint 计算点到线段的距禂
//   - 如果点在线段上，则距离为 0
//   - 如果点在线段的延长线上，则距离为点到线段两个端点的最小距离
//   - 否则，计算点到线段的投影点，然后计算点到投影点的距离
//   - 参考：https://en.wikipedia.org/wiki/Distance_from_a_point_to_a_line
func CalcLineSegmentDistanceToPoint(line LineSegment, point Point) float64 {
	if line.IsPointOnSegment(point) {
		return 0
	}

	projection := CalcLineSegmentPointProjection(line, point)
	if line.IsPointOnSegment(projection) {
		return point.Distance2D(projection)
	}

	return math.Min(point.Distance2D(line[0]), point.Distance2D(line[1]))
}

// CalcLineSegmentPointProjectionDistance 计算点到线段的投影点以及点到投影点的距离
func CalcLineSegmentPointProjectionDistance(line LineSegment, point Point) (projection Point, distance float64) {
	projection = CalcLineSegmentPointProjection(line, point)
	distance = point.Distance2D(projection)
	return
}

// CalcLineSegmentCollinearWithEpsilon 计算两条线段是否共线，允许一定的误差来包容浮点数计算的误差
func CalcLineSegmentCollinearWithEpsilon[T constraints.Number](line1, line2 LineSegment, epsilon T) bool {
	area1 := CalcTriangleAreaTwice(line1[0], line1[1], line2[0])
	area2 := CalcTriangleAreaTwice(line1[0], line1[1], line2[1])
	e := float64(epsilon)
	return math.Abs(area1-0) <= e && math.Abs(area2-0) <= e
}

// CalcLineSegmentOverlap 通过对点进行排序来检查两条共线线段是否重叠，返回重叠线段
func CalcLineSegmentOverlap(line1, line2 LineSegment) (overlap LineSegment, isOverlap bool) {
	type PointData struct {
		Point
		state bool
	}

	l1ps, l1pe := PointData{Point: NewPoint(line1[0].GetXY()), state: true}, PointData{Point: NewPoint(line1[1].GetXY()), state: true}
	l2ps, l2pe := PointData{Point: NewPoint(line2[0].GetXY()), state: false}, PointData{Point: NewPoint(line2[1].GetXY()), state: false}

	var shapes = [][]PointData{
		{l1ps, l1pe, l1ps},
		{l1ps, l1pe, l1pe},
		{l2ps, l2pe, l2ps},
		{l2ps, l2pe, l2pe},
	}
	sort.Slice(shapes, func(i, j int) bool {
		a, b := shapes[i], shapes[j]
		if a[2].GetX() < b[2].GetX() {
			return true
		} else if a[2].GetX() > b[2].GetX() {
			return false
		} else {
			return a[2].GetY() < b[2].GetY()
		}
	})

	notOverlap := shapes[1][0].state == shapes[2][0].state
	singlePointOverlap := shapes[1][2].Equal(shapes[2][2].Point)
	if notOverlap || singlePointOverlap {
		return overlap, false
	}
	return NewLineSegment(shapes[1][2].Point, shapes[2][2].Point), true
}
