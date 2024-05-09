package geometry

import (
	"strings"
)

// FloorPlan 平面图
type FloorPlan []string

// IsFree 检查位置是否为空格
func (fp FloorPlan) IsFree(position Vector2) bool {
	return fp.IsInBounds(position) && fp[int(position.GetY())][int(position.GetX())] == ' '
}

// IsInBounds 检查位置是否在边界内
func (fp FloorPlan) IsInBounds(position Vector2) bool {
	x, y := int(position.GetX()), int(position.GetY())
	return (0 <= x && x < len(fp[y])) && (0 <= y && y < len(fp))
}

// Put 设置平面图特定位置的字符
func (fp FloorPlan) Put(position Vector2, c rune) {
	x, y := int(position.GetX()), int(position.GetY())
	fp[y] = fp[y][:x] + string(c) + fp[y][x+1:]
}

// String 获取平面图结果
func (fp FloorPlan) String() string {
	var builder strings.Builder
	var last = len(fp) - 1
	for i, row := range fp {
		builder.WriteString(row)
		if i != last {
			builder.WriteByte('\n')
		}
	}
	return builder.String()
}
