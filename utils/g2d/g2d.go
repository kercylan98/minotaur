package g2d

// PositionToArray 将坐标转换为x、y的数组
func PositionToArray(x, y int) [2]int {
	return [2]int{x, y}
}

// PositionArrayToXY 将坐标数组转换为x和y坐标
func PositionArrayToXY(position [2]int) (x, y int) {
	return position[0], position[1]
}

// PositionClone 克隆一个坐标数组
func PositionClone(position [2]int) [2]int {
	return [2]int{position[0], position[1]}
}

// GetAdjacentPositions 获取一个矩阵中，特定位置相邻的最多四个方向的位置
func GetAdjacentPositions[T any](matrix [][]T, x, y int) (result [][2]int) {
	width, height := len(matrix), len(matrix[0])
	if tx := x - 1; tx >= 0 {
		result = append(result, PositionToArray(tx, y))
	}
	if tx := x + 1; tx < width {
		result = append(result, PositionToArray(tx, y))
	}
	if ty := y - 1; ty >= 0 {
		result = append(result, PositionToArray(x, ty))
	}
	if ty := y + 1; ty < height {
		result = append(result, PositionToArray(x, ty))
	}
	return
}

// GetAdjacentPositionsWithContinuousPosition 获取一个连续位置的矩阵中，特定位置相邻的最多四个方向的位置
func GetAdjacentPositionsWithContinuousPosition[T any](matrix []T, width, pos int) (result []int) {
	size := len(matrix)
	if up := pos - width; up >= 0 {
		result = append(result, up)
	}
	if down := pos + width; down < size {
		result = append(result, down)
	}
	if left := pos - 1; pos >= 0 {
		result = append(result, left)
	}
	if right := pos + 1; right < size {
		result = append(result, right)
	}
	return
}

// PositionToInt 将坐标转换为二维数组的顺序位置
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func PositionToInt(width, x, y int) int {
	return y*width + x
}

// PositionToIntWithArray 将坐标转换为数字
func PositionToIntWithArray(width int, xy [2]int) int {
	return PositionToInt(width, xy[0], xy[1])
}

// PositionsToIntWithArray 将一组坐标转换为数字
func PositionsToIntWithArray(width int, xys ...[2]int) []int {
	var result = make([]int, len(xys), len(xys))
	for i := 0; i < len(xys); i++ {
		result[i] = PositionToIntWithArray(width, xys[i])
	}
	return result
}

// PositionIntToXY 通过宽度将一个二维数组的顺序位置转换为xy坐标
func PositionIntToXY(width, pos int) (x, y int) {
	x = pos % width
	y = pos / width
	return x, y
}

// PositionIntGetX 通过宽度将一个二维数组的顺序位置转换为X坐标
func PositionIntGetX(width, pos int) int {
	return pos % width
}

// PositionIntGetY 通过宽度将一个二维数组的顺序位置转换为Y坐标
func PositionIntGetY(width, pos int) int {
	return pos / width
}
