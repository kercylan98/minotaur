package geometry

import (
	"strings"
)

// FloorPlan 平面图
type FloorPlan []string

// IsFree 检查位置是否为空格
func (slf FloorPlan) IsFree(point Point[int]) bool {
	return slf.IsInBounds(point) && slf[point.GetY()][point.GetX()] == ' '
}

// IsInBounds 检查位置是否在边界内
func (slf FloorPlan) IsInBounds(point Point[int]) bool {
	return (0 <= point.GetX() && point.GetX() < len(slf[point.GetY()])) && (0 <= point.GetY() && point.GetY() < len(slf))
}

// Put 设置平面图特定位置的字符
func (slf FloorPlan) Put(point Point[int], c rune) {
	slf[point.GetY()] = slf[point.GetY()][:point.GetX()] + string(c) + slf[point.GetY()][point.GetX()+1:]
}

// String 获取平面图结果
func (slf FloorPlan) String() string {
	var result string
	for _, row := range slf {
		result += row + "\n"
	}
	return strings.TrimSuffix(result, "\n")
}
