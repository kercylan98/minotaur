package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
)

// PointOnLineWithCoordinate 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
func PointOnLineWithCoordinate[V generic.Number](x1, y1, x2, y2, x, y V) bool {
	return (x-x1)*(y2-y1) == (x2-x1)*(y-y1)
}

//
//func PointOnLineWithPos[V generic.Number](width, pos1, pos2, pos V) bool {
//	x1, y1 := PosToCoordinate(width, pos1)
//	x2, y2 := PosToCoordinate(width, pos2)
//	return (x-x1)*(y2-y1) == (x2-x1)*(y-y1)
//}
//
//// PointOnSegment 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
////   - 与 PointOnLine 不同的是， PointOnSegment 中会判断线段及点的位置是否正确
//func PointOnSegment[V generic.Number](x1, y1, x2, y2, x, y V) bool {
//	return x >= x1 && x <= x2 && y >= y1 && y <= y2 && PointOnLine(x1, y1, x2, y2, x, y)
//}
