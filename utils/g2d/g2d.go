package g2d

// PositionToArray 将坐标转换为x、y的数组
func PositionToArray(x, y int) [2]int {
	return [2]int{x, y}
}

// PositionArrayToXY 将坐标数组转换为x和y坐标
func PositionArrayToXY(position [2]int) (x, y int) {
	return position[0], position[1]
}

// CoverageAreaBoundless 将一个图形覆盖范围设置为无边的
//   - 例如一个图形的left和top从2开始，那么将被转换到从0开始
func CoverageAreaBoundless(l, r, t, b int) (left, right, top, bottom int) {
	differentX := 0 - l
	differentY := 0 - t
	left = l + differentX
	right = r + differentX
	top = t + differentY
	bottom = b + differentY
	return
}

// GetShapeCoverageArea 获取一个图形覆盖的范围
func GetShapeCoverageArea(xys ...[2]int) (left, right, top, bottom int) {
	left, top = -1, -1
	for _, xy := range xys {
		x, y := PositionArrayToXY(xy)
		if x < left || left == -1 {
			left = x
		}
		if x > right {
			right = x
		}
		if y < top || top == -1 {
			top = y
		}
		if y > bottom {
			bottom = y
		}
	}
	return
}
