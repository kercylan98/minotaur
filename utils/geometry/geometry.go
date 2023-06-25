package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math"
)

// Direction 方向
type Direction uint8

const (
	DirectionUnknown = Direction(iota) // 未知
	DirectionUp                        // 上方
	DirectionDown                      // 下方
	DirectionLeft                      // 左方
	DirectionRight                     // 右方
)

var (
	DirectionUDLR = []Direction{DirectionUp, DirectionDown, DirectionLeft, DirectionRight} // 上下左右四个方向的数组
	DirectionLRUD = []Direction{DirectionLeft, DirectionRight, DirectionUp, DirectionDown} // 左右上下四个方向的数组
)

// GetOppositionDirection 获取特定方向的对立方向
func GetOppositionDirection(direction Direction) Direction {
	switch direction {
	case DirectionUp:
		return DirectionDown
	case DirectionDown:
		return DirectionUp
	case DirectionLeft:
		return DirectionRight
	case DirectionRight:
		return DirectionLeft
	}
	return DirectionUnknown
}

// GetDirectionNextWithCoordinate 获取特定方向上的下一个坐标
func GetDirectionNextWithCoordinate[V generic.SignedNumber](direction Direction, x, y V) (nx, ny V) {
	switch direction {
	case DirectionUp:
		nx, ny = x, y-1
	case DirectionDown:
		nx, ny = x, y+1
	case DirectionLeft:
		nx, ny = x-1, y
	case DirectionRight:
		nx, ny = x+1, y
	default:
		panic("unexplained direction")
	}
	return
}

// GetDirectionNextWithPoint 获取特定方向上的下一个坐标
func GetDirectionNextWithPoint[V generic.SignedNumber](direction Direction, point Point[V]) Point[V] {
	x, y := point.GetXY()
	switch direction {
	case DirectionUp:
		return NewPoint(x, y-1)
	case DirectionDown:
		return NewPoint(x, y+1)
	case DirectionLeft:
		return NewPoint(x-1, y)
	case DirectionRight:
		return NewPoint(x+1, y)
	default:
		panic("unexplained direction")
	}
}

// GetDirectionNextWithPos 获取位置在特定宽度和特定方向上的下一个位置
//   - 需要注意的是，在左右方向时，当下一个位置不在矩形区域内时，将会返回上一行的末位置或下一行的首位置
func GetDirectionNextWithPos[V generic.SignedNumber](direction Direction, width, pos V) V {
	switch direction {
	case DirectionUp:
		return pos - width
	case DirectionDown:
		return pos + width
	case DirectionLeft:
		return pos - 1
	case DirectionRight:
		return pos + 1
	default:
		panic("unexplained direction")
	}
}

// CalcDirection 计算点2位于点1的方向
func CalcDirection[V generic.SignedNumber](x1, y1, x2, y2 V) Direction {
	var oneEighty = 180
	var fortyFive = 45
	var oneThirtyFive = 135
	var twoTwentyFive = 225
	var threeFifteen = 315
	var end = 360
	var start = 0
	angle := CalcAngle(x1, y1, x2, y2) + V(oneEighty)
	if angle > V(oneThirtyFive) && angle <= V(twoTwentyFive) {
		return DirectionRight
	} else if (angle > V(threeFifteen) && angle <= V(end)) || (angle >= V(start) && angle <= V(fortyFive)) {
		return DirectionLeft
	} else if angle > V(twoTwentyFive) && angle <= V(threeFifteen) {
		return DirectionUp
	} else if angle > V(fortyFive) && angle <= V(oneThirtyFive) {
		return DirectionDown
	}
	return DirectionUnknown
}

// CalcDistanceWithCoordinate 计算两点之间的距离
func CalcDistanceWithCoordinate[V generic.SignedNumber](x1, y1, x2, y2 V) V {
	return V(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

// CalcDistanceWithPoint 计算两点之间的距离
func CalcDistanceWithPoint[V generic.SignedNumber](point1, point2 Point[V]) V {
	return CalcDistanceWithCoordinate(DoublePointToCoordinate(point1, point2))
}

// CalcDistanceSquared 计算两点之间的平方距离
//   - 这个函数的主要用途是在需要计算两点之间距离的情况下，但不需要得到实际的距离值，而只需要比较距离大小。因为平方根运算相对较为耗时，所以在只需要比较大小的情况下，通常会使用平方距离。
func CalcDistanceSquared[V generic.SignedNumber](x1, y1, x2, y2 V) V {
	dx, dy := x2-x1, y2-y1
	return dx*dx + dy*dy
}

// CalcAngle 计算点2位于点1之间的角度
func CalcAngle[V generic.SignedNumber](x1, y1, x2, y2 V) V {
	return V(math.Atan2(float64(y2-y1), float64(x2-x1)) * 180 / math.Pi)
}

// CalcNewCoordinate 根据给定的x、y坐标、角度和距离计算新的坐标
func CalcNewCoordinate[V generic.SignedNumber](x, y, angle, distance V) (newX, newY V) {
	// 将角度转换为弧度
	var pi = math.Pi
	var dividend = 180.0
	radians := angle * V(pi) / V(dividend)

	// 计算新的坐标
	newX = x + distance*V(math.Cos(float64(radians)))
	newY = y + distance*V(math.Sin(float64(radians)))

	return newX, newY
}

// CalcAngleDifference 计算两个角度之间的最小角度差
func CalcAngleDifference[V generic.SignedNumber](angleA, angleB V) V {
	pi := math.Pi
	t := angleA - angleB
	a := t + V(pi)
	b := V(pi) * 2
	t = V(math.Floor(float64(a/b))) * b
	t -= V(pi)
	return t
}
