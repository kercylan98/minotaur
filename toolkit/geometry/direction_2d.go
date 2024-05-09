package geometry

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"math"
)

// Direction2D 二维方向
type Direction2D = Vector2

var (
	direction2DUnknown   = NewVector2(0, 0)   // 未知
	direction2DUp        = NewVector2(0, 1)   // 上
	direction2DDown      = NewVector2(0, -1)  // 下
	direction2DLeft      = NewVector2(-1, 0)  // 左
	direction2DRight     = NewVector2(1, 0)   // 右
	direction2DLeftUp    = NewVector2(-1, 1)  // 左上
	direction2DRightUp   = NewVector2(1, 1)   // 右上
	direction2DLeftDown  = NewVector2(-1, -1) // 左下
	direction2DRightDown = NewVector2(1, -1)  // 右下

	direction2D4 = []Direction2D{direction2DUp, direction2DDown, direction2DLeft, direction2DRight} // 上下左右四个方向的数组                                                                 // 上下左右四个方向的数组
	direction2D8 = []Direction2D{                                                                   // 上下左右、左上、右上、左下、右下八个方向的数组
		direction2DUp, direction2DDown, direction2DLeft, direction2DRight,
		direction2DLeftUp, direction2DRightUp, direction2DLeftDown, direction2DRightDown,
	}
)

// Direction2D4 上下左右四个方向
func Direction2D4() []Direction2D {
	return direction2D4
}

// Direction2D8 上下左右、左上、右上、左下、右下八个方向
func Direction2D8() []Direction2D {
	return direction2D8
}

// Direction2DUnknown 获取未知方向
func Direction2DUnknown() Direction2D {
	return direction2DUnknown
}

// Direction2DUp 获取上方向
func Direction2DUp() Direction2D {
	return direction2DUp
}

// Direction2DDown 获取下方向
func Direction2DDown() Direction2D {
	return direction2DDown
}

// Direction2DLeft 获取左方向
func Direction2DLeft() Direction2D {
	return direction2DLeft
}

// Direction2DRight 获取右方向
func Direction2DRight() Direction2D {
	return direction2DRight
}

// Direction2DLeftUp 获取左上方向
func Direction2DLeftUp() Direction2D {
	return direction2DLeftUp
}

// Direction2DRightUp 获取右上方向
func Direction2DRightUp() Direction2D {
	return direction2DRightUp
}

// Direction2DLeftDown 获取左下方向
func Direction2DLeftDown() Direction2D {
	return direction2DLeftDown
}

// Direction2DRightDown 获取右下方向
func Direction2DRightDown() Direction2D {
	return direction2DRightDown
}

// CalcOppositionDirection2D 计算二维方向的反方向
func CalcOppositionDirection2D(direction Direction2D) Direction2D {
	switch {
	case direction.Equal(direction2DUp):
		return direction2DDown
	case direction.Equal(direction2DDown):
		return direction2DUp
	case direction.Equal(direction2DLeft):
		return direction2DRight
	case direction.Equal(direction2DRight):
		return direction2DLeft
	case direction.Equal(direction2DLeftUp):
		return direction2DRightDown
	case direction.Equal(direction2DRightUp):
		return direction2DLeftDown
	case direction.Equal(direction2DLeftDown):
		return direction2DRightUp
	case direction.Equal(direction2DRightDown):
		return direction2DLeftUp
	default:
		return direction2DUnknown
	}
}

// CalcOffsetInDirection2D 计算特定方向上按照指定距离偏移后的坐标
func CalcOffsetInDirection2D[T constraints.Number](vector Vector2, direction Direction2D, offset T) Vector2 {
	return vector.Add(direction.Mul(float64(offset)))
}

// CalcDirection2DWithAngle 通过角度计算二维方向
func CalcDirection2DWithAngle[T constraints.Number](angle T) Direction2D {
	angleFloat := float64(angle)
	return NewVector2(math.Cos(angleFloat), math.Sin(angleFloat))
}

// CalcAngleWithDirection2D 计算二维方向的角度
func CalcAngleWithDirection2D(direction Direction2D) float64 {
	return math.Atan2(direction.GetY(), direction.GetX())
}
