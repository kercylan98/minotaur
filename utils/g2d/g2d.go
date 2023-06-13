package g2d

// CoordinateToCoordinateArray 将坐标转换为x、y的坐标数组
func CoordinateToCoordinateArray(x, y int) [2]int {
	return [2]int{x, y}
}

// CoordinateArrayToCoordinate 将坐标数组转换为x和y坐标
func CoordinateArrayToCoordinate(position [2]int) (x, y int) {
	return position[0], position[1]
}

// CoordinateArrayClone 克隆一个坐标数组
func CoordinateArrayClone(position [2]int) [2]int {
	return [2]int{position[0], position[1]}
}

// GetAdjacentCoordinates 获取一个矩阵中，特定位置相邻的最多四个方向的坐标
func GetAdjacentCoordinates[T any](matrix [][]T, x, y int) (result [][2]int) {
	width, height := len(matrix), len(matrix[0])
	if tx := x - 1; tx >= 0 {
		result = append(result, CoordinateToCoordinateArray(tx, y))
	}
	if tx := x + 1; tx < width {
		result = append(result, CoordinateToCoordinateArray(tx, y))
	}
	if ty := y - 1; ty >= 0 {
		result = append(result, CoordinateToCoordinateArray(x, ty))
	}
	if ty := y + 1; ty < height {
		result = append(result, CoordinateToCoordinateArray(x, ty))
	}
	return
}

// GetAdjacentCoordinatesWithPos 获取一个连续位置的矩阵中，特定位置相邻的最多四个方向的坐标
func GetAdjacentCoordinatesWithPos[T any](matrix []T, width, pos int) (result []int) {
	size := len(matrix)
	currentRow := pos / width
	if up := pos - width; up >= 0 {
		result = append(result, up)
	}
	if down := pos + width; down < size {
		result = append(result, down)
	}
	if left := pos - 1; left >= 0 && currentRow == (left/width) {
		result = append(result, left)
	}
	if right := pos + 1; right < size && currentRow == (right/width) {
		result = append(result, right)
	}
	return
}

// CoordinateToPos 将坐标转换为二维数组的顺序位置坐标
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func CoordinateToPos(width, x, y int) int {
	return y*width + x
}

// CoordinateArrayToPos 将坐标转换为二维数组的顺序位置
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func CoordinateArrayToPos(width int, xy [2]int) int {
	return CoordinateToPos(width, xy[0], xy[1])
}

// CoordinateArrayToPosWithMulti 将一组坐标转换为二维数组的顺序位置
//   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值
func CoordinateArrayToPosWithMulti(width int, xys ...[2]int) []int {
	var result = make([]int, len(xys), len(xys))
	for i := 0; i < len(xys); i++ {
		result[i] = CoordinateArrayToPos(width, xys[i])
	}
	return result
}

// PosToCoordinate 通过宽度将一个二维数组的顺序位置转换为xy坐标
func PosToCoordinate(width, pos int) (x, y int) {
	x = pos % width
	y = pos / width
	return x, y
}

// PosToCoordinateX 通过宽度将一个二维数组的顺序位置转换为X坐标
func PosToCoordinateX(width, pos int) int {
	return pos % width
}

// PosToCoordinateY 通过宽度将一个二维数组的顺序位置转换为Y坐标
func PosToCoordinateY(width, pos int) int {
	return pos / width
}

// MatrixToPosMatrix 将二维矩阵转换为顺序的二维矩阵
func MatrixToPosMatrix[V any](matrix [][]V) (width int, posMatrix []V) {
	width = len(matrix)
	height := len(matrix[0])
	posMatrix = make([]V, width*height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			posMatrix[CoordinateToPos(width, x, y)] = matrix[x][y]
		}
	}
	return
}
