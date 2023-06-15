package g2d

import "github.com/kercylan98/minotaur/utils/generic"

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

// GetAdjacentTranslatePos 获取一个连续位置的矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
func GetAdjacentTranslatePos[T any](matrix []T, width, pos int) (result []int) {
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

// GetAdjacentDiagonalsPos 获取一个连续位置的矩阵中，特定位置相邻的对角线最多四个方向的位置
func GetAdjacentDiagonalsPos[T any](matrix []T, width, pos int) (result []int) {
	size := len(matrix)
	currentRow := pos / width
	if topLeft := pos - width - 1; topLeft >= 0 && currentRow-1 == (topLeft/width) {
		result = append(result, topLeft)
	}
	if topRight := pos - width + 1; topRight >= 0 && currentRow-1 == (topRight/width) {
		result = append(result, topRight)
	}
	if bottomLeft := pos + width - 1; bottomLeft < size && currentRow+1 == (bottomLeft/width) {
		result = append(result, bottomLeft)
	}
	if bottomRight := pos + width + 1; bottomRight < size && currentRow+1 == (bottomRight/width) {
		result = append(result, bottomRight)
	}
	return
}

// GetAdjacentPos 获取一个连续位置的矩阵中，特定位置相邻的最多八个方向的位置
func GetAdjacentPos[T any](matrix []T, width, pos int) (result []int) {
	return append(GetAdjacentTranslatePos(matrix, width, pos), GetAdjacentDiagonalsPos(matrix, width, pos)...)
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

// PointOnLine 接受六个参数 top、bottom、left、right 和 x、y，分别表示一个矩形位置尺寸和一个点的横纵坐标，判断这个点是否在一条线段上。
//   - 首先计算点 (x, y) 与线段起点 (left, top) 之间的斜率即 (x - left) / (y - top)。
//   - 然后计算线段起点 (left, top) 与线段终点 (right, bottom) 之间的斜率，即 (right - left) / (bottom - top)。
//   - 如果这两个斜率等，那么点 (x, y) 就在这条线段上。为了避免除法可能导致的浮点数误差，我们可以将两个斜率的计算转换为乘法形式，即比较 (x - left) * (bottom - top) 是否等于 (right - left) * y - top)。
//   - 如果上述等式成立，说明点 (x, y) 在线段上，函数返回 true；否则，返回 false。
func PointOnLine[V generic.Number](top, bottom, left, right, x, y V) bool {
	return (x-left)*(bottom-top) == (right-left)*(y-top)
}
