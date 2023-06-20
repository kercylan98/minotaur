package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/maths"
	"sort"
)

// NewLineSegment 创建一根线段
func NewLineSegment[V generic.SignedNumber](start, end Point[V]) LineSegment[V] {
	if start.Equal(end) {
		panic("two points of the line segment are the same")
	}
	return LineSegment[V]{start, end}
}

// NewLineSegmentCap 创建一根包含数据的线段
func NewLineSegmentCap[V generic.SignedNumber, Data any](start, end Point[V], data Data) LineSegmentCap[V, Data] {
	return LineSegmentCap[V, Data]{NewLineSegment(start, end), data}
}

// NewLineSegmentCapWithLine 通过已有线段创建一根包含数据的线段
func NewLineSegmentCapWithLine[V generic.SignedNumber, Data any](line LineSegment[V], data Data) LineSegmentCap[V, Data] {
	return LineSegmentCap[V, Data]{line, data}
}

// LineSegment 通过两个点表示一根线段
type LineSegment[V generic.SignedNumber] [2]Point[V]

// LineSegmentCap 可以包含一份额外数据的线段
type LineSegmentCap[V generic.SignedNumber, Data any] struct {
	LineSegment[V]
	Data Data
}

func (slf *LineSegmentCap[V, Data]) GetData() Data {
	return slf.Data
}

// GetPoints 获取该线段的两个点
func (slf LineSegment[V]) GetPoints() [2]Point[V] {
	return slf
}

// GetStart 获取该线段的开始位置
func (slf LineSegment[V]) GetStart() Point[V] {
	return slf[0]
}

// GetEnd 获取该线段的结束位置
func (slf LineSegment[V]) GetEnd() Point[V] {
	return slf[1]
}

// GetLength 获取该线段的长度
func (slf LineSegment[V]) GetLength() V {
	return CalcDistanceWithCoordinate(DoublePointToCoordinate(slf.GetStart(), slf.GetEnd()))
}

// PointOnLineSegmentWithCoordinate 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
func PointOnLineSegmentWithCoordinate[V generic.SignedNumber](x1, y1, x2, y2, x, y V) bool {
	return (x-x1)*(y2-y1) == (x2-x1)*(y-y1)
}

// PointOnLineSegmentWithPos 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
func PointOnLineSegmentWithPos[V generic.SignedNumber](width, pos1, pos2, pos V) bool {
	x1, y1 := PosToCoordinate(width, pos1)
	x2, y2 := PosToCoordinate(width, pos2)
	x, y := PosToCoordinate(width, pos)
	return PointOnLineSegmentWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnLineSegmentWithPoint 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
func PointOnLineSegmentWithPoint[V generic.SignedNumber](point1, point2, point Point[V]) bool {
	x1, y1 := point1.GetXY()
	x2, y2 := point2.GetXY()
	x, y := point.GetXY()
	return PointOnLineSegmentWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnLineSegmentWithCoordinateInBounds 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
//   - 与 PointOnLineSegmentWithCoordinate 不同的是， PointOnLineSegmentWithCoordinateInBounds 中会判断线段及点的位置是否正确
func PointOnLineSegmentWithCoordinateInBounds[V generic.SignedNumber](x1, y1, x2, y2, x, y V) bool {
	return x >= x1 && x <= x2 && y >= y1 && y <= y2 && PointOnLineSegmentWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnLineSegmentWithPosInBounds 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
//   - 与 PointOnLineSegmentWithPos 不同的是， PointOnLineSegmentWithPosInBounds 中会判断线段及点的位置是否正确
func PointOnLineSegmentWithPosInBounds[V generic.SignedNumber](width, pos1, pos2, pos V) bool {
	x1, y1 := PosToCoordinate(width, pos1)
	x2, y2 := PosToCoordinate(width, pos2)
	x, y := PosToCoordinate(width, pos)
	return x >= x1 && x <= x2 && y >= y1 && y <= y2 && PointOnLineSegmentWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnLineSegmentWithPointInBounds 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
//   - 与 PointOnLineSegmentWithPoint 不同的是， PointOnLineSegmentWithPointInBounds 中会判断线段及点的位置是否正确
func PointOnLineSegmentWithPointInBounds[V generic.SignedNumber](point1, point2, point Point[V]) bool {
	x1, y1 := point1.GetXY()
	x2, y2 := point2.GetXY()
	x, y := point.GetXY()
	return x >= x1 && x <= x2 && y >= y1 && y <= y2 && PointOnLineSegmentWithCoordinate(x1, y1, x2, y2, x, y)
}

// CalcLineSegmentIsCollinear 检查两条线段在一个误差内是否共线
//   - 共线是指两条线段在同一直线上，即它们的延长线可以重合
func CalcLineSegmentIsCollinear[V generic.SignedNumber](line1, line2 LineSegment[V], tolerance V) bool {
	area1 := CalcTriangleTwiceArea(line1.GetStart(), line1.GetEnd(), line2.GetStart())
	area2 := CalcTriangleTwiceArea(line1.GetStart(), line1.GetEnd(), line2.GetEnd())
	return maths.Tolerance(area1, 0, tolerance) && maths.Tolerance(area2, 0, tolerance)
}

// CalcLineSegmentIsOverlap 通过对点进行排序来检查两条共线线段是否重叠，返回重叠线段
func CalcLineSegmentIsOverlap[V generic.SignedNumber](line1, line2 LineSegment[V]) (line LineSegment[V], overlap bool) {
	l1ps, l1pe := NewPointCapWithPoint(line1.GetStart(), true), NewPointCapWithPoint(line1.GetEnd(), true)
	l2ps, l2pe := NewPointCapWithPoint(line2.GetStart(), false), NewPointCapWithPoint(line2.GetEnd(), false)
	var shapes = [][]PointCap[V, bool]{
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

	notOverlap := shapes[1][0].GetData() == shapes[2][0].GetData()
	singlePointOverlap := shapes[1][2].Equal(shapes[2][2].Point)
	if notOverlap || singlePointOverlap {
		return line, false
	}
	return NewLineSegment(shapes[1][2].Point, shapes[2][2].Point), true
}
