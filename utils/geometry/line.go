package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
)

// PointOnLineWithCoordinate 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
func PointOnLineWithCoordinate[V generic.Number](x1, y1, x2, y2, x, y V) bool {
	return (x-x1)*(y2-y1) == (x2-x1)*(y-y1)
}

// PointOnLineWithPos 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
func PointOnLineWithPos[V generic.Number](width, pos1, pos2, pos V) bool {
	x1, y1 := PosToCoordinate(width, pos1)
	x2, y2 := PosToCoordinate(width, pos2)
	x, y := PosToCoordinate(width, pos)
	return PointOnLineWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnLineWithCoordinateArray 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
func PointOnLineWithCoordinateArray[V generic.Number](point1, point2, point Point[V]) bool {
	x1, y1 := point1.GetXY()
	x2, y2 := point2.GetXY()
	x, y := point.GetXY()
	return PointOnLineWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnSegmentWithCoordinate 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
//   - 与 PointOnLineWithCoordinate 不同的是， PointOnSegmentWithCoordinate 中会判断线段及点的位置是否正确
func PointOnSegmentWithCoordinate[V generic.Number](x1, y1, x2, y2, x, y V) bool {
	return x >= x1 && x <= x2 && y >= y1 && y <= y2 && PointOnLineWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnSegmentWithPos 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
//   - 与 PointOnLineWithPos 不同的是， PointOnSegmentWithPos 中会判断线段及点的位置是否正确
func PointOnSegmentWithPos[V generic.Number](width, pos1, pos2, pos V) bool {
	x1, y1 := PosToCoordinate(width, pos1)
	x2, y2 := PosToCoordinate(width, pos2)
	x, y := PosToCoordinate(width, pos)
	return x >= x1 && x <= x2 && y >= y1 && y <= y2 && PointOnLineWithCoordinate(x1, y1, x2, y2, x, y)
}

// PointOnSegmentWithCoordinateArray 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
//   - 与 PointOnLineWithCoordinateArray 不同的是， PointOnSegmentWithCoordinateArray 中会判断线段及点的位置是否正确
func PointOnSegmentWithCoordinateArray[V generic.Number](point1, point2, point Point[V]) bool {
	x1, y1 := point1.GetXY()
	x2, y2 := point2.GetXY()
	x, y := point.GetXY()
	return x >= x1 && x <= x2 && y >= y1 && y <= y2 && PointOnLineWithCoordinate(x1, y1, x2, y2, x, y)
}
