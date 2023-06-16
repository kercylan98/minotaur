package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math"
)

// Point 表示了一个由 x、y 坐标组成的点
type Point[V generic.Number] [2]V

// NewPoint 创建一个由 x、y 坐标组成的点
func NewPoint[V generic.Number](x, y V) Point[V] {
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
	return CoordinateArrayToPos(width, slf)
}

// Copy 复制一个点位置
func (slf Point[V]) Copy() Point[V] {
	return CoordinateArrayCopy(slf)
}

// NewPointCap 创建一个由 x、y 坐标组成的点，这个点具有一个数据容量
func NewPointCap[V generic.Number, D any](x, y V) PointCap[V, D] {
	return PointCap[V, D]{
		Point: NewPoint(x, y),
	}
}

// NewPointCapWithData 通过设置数据的方式创建一个由 x、y 坐标组成的点，这个点具有一个数据容量
func NewPointCapWithData[V generic.Number, D any](x, y V, data D) PointCap[V, D] {
	return PointCap[V, D]{
		Point: NewPoint(x, y),
		Data:  data,
	}
}

// PointCap 表示了一个由 x、y 坐标组成的点，这个点具有一个数据容量
type PointCap[V generic.Number, D any] struct {
	Point[V]
	Data D
}

// GetData 获取数据
func (slf PointCap[V, D]) GetData() D {
	return slf.Data
}

// CoordinateToCoordinateArray 将坐标转换为x、y的坐标数组
func CoordinateToCoordinateArray[V generic.Number](x, y V) Point[V] {
	return [2]V{x, y}
}

// CoordinateToPos 将坐标转换为二维数组的顺序位置坐标
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func CoordinateToPos[V generic.Number](width, x, y V) V {
	return y*width + x
}

// CoordinateArrayToCoordinate 将坐标数组转换为x和y坐标
func CoordinateArrayToCoordinate[V generic.Number](position Point[V]) (x, y V) {
	return position[0], position[1]
}

// CoordinateArrayToPos 将坐标转换为二维数组的顺序位置
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func CoordinateArrayToPos[V generic.Number](width V, xy Point[V]) V {
	return CoordinateToPos(width, xy[0], xy[1])
}

// PosToCoordinate 通过宽度将一个二维数组的顺序位置转换为xy坐标
func PosToCoordinate[V generic.Number](width, pos V) (x, y V) {

	x = V(math.Mod(float64(pos), float64(width)))
	y = pos / width
	return x, y
}

// PosToCoordinateArray 通过宽度将一个二维数组的顺序位置转换为x、y的坐标数组
func PosToCoordinateArray[V generic.Number](width, pos V) Point[V] {
	return [2]V{V(math.Mod(float64(pos), float64(width))), pos / width}
}

// PosToCoordinateX 通过宽度将一个二维数组的顺序位置转换为X坐标
func PosToCoordinateX[V generic.Number](width, pos V) V {
	return V(math.Mod(float64(pos), float64(width)))
}

// PosToCoordinateY 通过宽度将一个二维数组的顺序位置转换为Y坐标
func PosToCoordinateY[V generic.Number](width, pos V) V {
	return pos / width
}

// CoordinateArrayCopy 复制一个坐标数组
func CoordinateArrayCopy[V generic.Number](position Point[V]) Point[V] {
	return NewPoint(position[0], position[1])
}

// CoordinateArrayToPosWithMulti 将一组坐标转换为二维数组的顺序位置
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func CoordinateArrayToPosWithMulti[V generic.Number](width V, xys ...Point[V]) []V {
	var result = make([]V, len(xys), len(xys))
	for i := 0; i < len(xys); i++ {
		result[i] = CoordinateArrayToPos(width, xys[i])
	}
	return result
}

// PosToCoordinateArrayWithMulti 将一组二维数组的顺序位置转换为一组数组坐标
func PosToCoordinateArrayWithMulti[V generic.Number](width V, positions ...V) []Point[V] {
	var result = make([]Point[V], len(positions))
	for i := 0; i < len(positions); i++ {
		result[i] = PosToCoordinateArray(width, positions[i])
	}
	return result
}
