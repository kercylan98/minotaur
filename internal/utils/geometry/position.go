package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/maths"
	"math"
)

// Point 表示了一个由 x、y 坐标组成的点
type Point[V generic.SignedNumber] [2]V

// NewPoint 创建一个由 x、y 坐标组成的点
func NewPoint[V generic.SignedNumber](x, y V) Point[V] {
	return Point[V]{x, y}
}

// GetX 返回该点的 x 坐标
func (slf Point[V]) GetX() V {
	return slf[0]
}

// GetY 返回该点的 y 坐标
func (slf Point[V]) GetY() V {
	return slf[1]
}

// GetXY 返回该点的 x、y 坐标
func (slf Point[V]) GetXY() (x, y V) {
	return slf[0], slf[1]
}

// GetPos 返回该点位于特定宽度的二维数组的顺序位置
func (slf Point[V]) GetPos(width V) V {
	return PointToPos(width, slf)
}

// GetOffset 获取偏移后的新坐标
func (slf Point[V]) GetOffset(x, y V) Point[V] {
	return NewPoint(slf.GetX()+x, slf.GetY()+y)
}

// Negative 返回该点是否是一个负数坐标
func (slf Point[V]) Negative() bool {
	return slf.GetX() < V(0) || slf.GetY() < V(0)
}

// OutOf 返回该点在特定宽高下是否越界f
func (slf Point[V]) OutOf(minWidth, minHeight, maxWidth, maxHeight V) bool {
	return slf.GetX() < minWidth || slf.GetY() < minHeight || slf.GetX() >= maxWidth || slf.GetY() >= maxHeight
}

// Equal 返回两个点是否相等
func (slf Point[V]) Equal(point Point[V]) bool {
	return slf.GetX() == point.GetX() && slf.GetY() == point.GetY()
}

// Copy 复制一个点位置
func (slf Point[V]) Copy() Point[V] {
	return PointCopy(slf)
}

// Add 得到加上 point 后的点
func (slf Point[V]) Add(point Point[V]) Point[V] {
	return slf.GetOffset(point.GetXY())
}

// Sub 得到减去 point 后的点
func (slf Point[V]) Sub(point Point[V]) Point[V] {
	return NewPoint(slf.GetX()-point.GetX(), slf.GetY()-point.GetY())
}

// Mul 得到乘以 point 后的点
func (slf Point[V]) Mul(point Point[V]) Point[V] {
	return NewPoint(slf.GetX()*point.GetX(), slf.GetY()*point.GetY())
}

// Div 得到除以 point 后的点
func (slf Point[V]) Div(point Point[V]) Point[V] {
	return NewPoint(slf.GetX()/point.GetX(), slf.GetY()/point.GetY())
}

// Abs 返回位置的绝对值
func (slf Point[V]) Abs() Point[V] {
	return NewPoint(V(math.Abs(float64(slf.GetX()))), V(math.Abs(float64(slf.GetY()))))
}

// Distance 返回两个点之间的距离
func (slf Point[V]) Distance(point Point[V]) float64 {
	return math.Sqrt(float64(slf.DistanceSquared(point)))
}

// DistanceSquared 返回两个点之间的距离的平方
func (slf Point[V]) DistanceSquared(point Point[V]) V {
	x, y := slf.GetXY()
	px, py := point.GetXY()
	return (x-px)*(x-px) + (y-py)*(y-py)
}

// Max 返回两个位置中每个维度的最大值组成的新的位置
func (slf Point[V]) Max(point Point[V]) Point[V] {
	x, y := slf.GetXY()
	px, py := point.GetXY()
	if px > x {
		x = px
	}
	if py > y {
		y = py
	}
	return NewPoint(x, y)
}

// Move 返回向特定角度移动特定距离后的新的位置，其中 angle 期待的角度范围是 -180~180
func (slf Point[V]) Move(angle, direction V) Point[V] {
	df := float64(direction)
	// 将角度转换为弧度
	radian := float64(angle) * (math.Pi / 180.0)

	// 计算新的坐标
	newX := float64(slf.GetX()) + df*math.Cos(radian)
	newY := float64(slf.GetY()) + df*math.Sin(radian)

	return NewPoint(V(newX), V(newY))
}

// Min 返回两个位置中每个维度的最小值组成的新的位置
func (slf Point[V]) Min(point Point[V]) Point[V] {
	x, y := slf.GetXY()
	px, py := point.GetXY()
	if px < x {
		x = px
	}
	if py < y {
		y = py
	}
	return NewPoint(x, y)
}

// NewPointCap 创建一个由 x、y 坐标组成的点，这个点具有一个数据容量
func NewPointCap[V generic.SignedNumber, D any](x, y V) PointCap[V, D] {
	return PointCap[V, D]{
		Point: NewPoint(x, y),
	}
}

// NewPointCapWithData 通过设置数据的方式创建一个由 x、y 坐标组成的点，这个点具有一个数据容量
func NewPointCapWithData[V generic.SignedNumber, D any](x, y V, data D) PointCap[V, D] {
	return PointCap[V, D]{
		Point: NewPoint(x, y),
		Data:  data,
	}
}

// NewPointCapWithPoint 通过设置数据的方式创建一个由已有坐标组成的点，这个点具有一个数据容量
func NewPointCapWithPoint[V generic.SignedNumber, D any](point Point[V], data D) PointCap[V, D] {
	return PointCap[V, D]{
		Point: point,
		Data:  data,
	}
}

// PointCap 表示了一个由 x、y 坐标组成的点，这个点具有一个数据容量
type PointCap[V generic.SignedNumber, D any] struct {
	Point[V]
	Data D
}

// GetData 获取数据
func (slf PointCap[V, D]) GetData() D {
	return slf.Data
}

// CoordinateToPoint 将坐标转换为x、y的坐标数组
func CoordinateToPoint[V generic.SignedNumber](x, y V) Point[V] {
	return [2]V{x, y}
}

// CoordinateToPos 将坐标转换为二维数组的顺序位置坐标
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func CoordinateToPos[V generic.SignedNumber](width, x, y V) V {
	return y*width + x
}

// PointToCoordinate 将坐标数组转换为x和y坐标
func PointToCoordinate[V generic.SignedNumber](position Point[V]) (x, y V) {
	return position[0], position[1]
}

// PointToPos 将坐标转换为二维数组的顺序位置
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func PointToPos[V generic.SignedNumber](width V, xy Point[V]) V {
	return CoordinateToPos(width, xy[0], xy[1])
}

// PosToCoordinate 通过宽度将一个二维数组的顺序位置转换为xy坐标
func PosToCoordinate[V generic.SignedNumber](width, pos V) (x, y V) {
	x = V(math.Mod(float64(pos), float64(width)))
	y = pos / width
	return x, y
}

// PosToPoint 通过宽度将一个二维数组的顺序位置转换为x、y的坐标数组
func PosToPoint[V generic.SignedNumber](width, pos V) Point[V] {
	return [2]V{V(math.Mod(float64(pos), float64(width))), pos / width}
}

// PosToCoordinateX 通过宽度将一个二维数组的顺序位置转换为X坐标
func PosToCoordinateX[V generic.SignedNumber](width, pos V) V {
	return V(math.Mod(float64(pos), float64(width)))
}

// PosToCoordinateY 通过宽度将一个二维数组的顺序位置转换为Y坐标
func PosToCoordinateY[V generic.SignedNumber](width, pos V) V {
	return pos / width
}

// PointCopy 复制一个坐标数组
func PointCopy[V generic.SignedNumber](point Point[V]) Point[V] {
	return NewPoint(point.GetXY())
}

// PointToPosWithMulti 将一组坐标转换为二维数组的顺序位置
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func PointToPosWithMulti[V generic.SignedNumber](width V, points ...Point[V]) []V {
	var result = make([]V, len(points), len(points))
	for i := 0; i < len(points); i++ {
		result[i] = PointToPos(width, points[i])
	}
	return result
}

// PosToPointWithMulti 将一组二维数组的顺序位置转换为一组数组坐标
func PosToPointWithMulti[V generic.SignedNumber](width V, positions ...V) []Point[V] {
	var result = make([]Point[V], len(positions))
	for i := 0; i < len(positions); i++ {
		result[i] = PosToPoint(width, positions[i])
	}
	return result
}

// PosSameRow 返回两个顺序位置在同一宽度是否位于同一行
func PosSameRow[V generic.SignedNumber](width, pos1, pos2 V) bool {
	return (pos1 / width) == (pos2 / width)
}

// DoublePointToCoordinate 将两个位置转换为 x1, y1, x2, y2 的坐标进行返回
func DoublePointToCoordinate[V generic.SignedNumber](point1, point2 Point[V]) (x1, y1, x2, y2 V) {
	return point1.GetX(), point1.GetY(), point2.GetX(), point2.GetY()
}

// CalcProjectionPoint 计算一个点到一条线段的最近点（即投影点）的。这个函数接收一个点和一条线段作为输入，线段由两个端点组成。
//   - 该函数的主要用于需要计算一个点到一条线段的最近点的情况下
func CalcProjectionPoint[V generic.SignedNumber](line LineSegment[V], point Point[V]) Point[V] {
	ax, ay, bx, by := DoublePointToCoordinate(line.GetStart(), line.GetEnd())
	ds := CalcDistanceSquared(ax, ay, bx, by)
	px, py := point.GetXY()
	clamp := maths.Clamp((px-ax)*(bx-ax)+(py-ay)*(by-ay)/ds, V(0), V(1))
	return NewPoint(ax+clamp*(bx-ax), ay+clamp*(by-ay))
}
