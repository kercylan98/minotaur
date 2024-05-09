package geometry

import "github.com/kercylan98/minotaur/toolkit/constraints"

// NewPoint 创建一个点
func NewPoint[V constraints.Number](x, y V) Point {
	return NewVector2(x, y)
}

// NewPosition 创建一个位置
func NewPosition[V constraints.Number](x, y V) Position {
	return NewVector2(x, y)
}

// Point 由二维向量组成的点
type Point = Vector2

// Position 位置，等同于 Point
type Position = Point
