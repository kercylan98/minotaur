package geometry

import "github.com/kercylan98/minotaur/utils/generic"

// GetAdjacentTranslatePos 获取一个连续位置的矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
func GetAdjacentTranslatePos[T any, P generic.Number](matrix []T, width, pos P) (result []P) {
	size := P(len(matrix))
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

// GetAdjacentTranslateCoordinateXY 获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
func GetAdjacentTranslateCoordinateXY[T any, P generic.Number](matrix [][]T, x, y P) (result []Point[P]) {
	width := P(len(matrix))
	height := P(len(matrix[0]))
	if up := y - 1; up >= 0 {
		result = append(result, NewPoint(x, up))
	}
	if down := y + 1; down < height {
		result = append(result, NewPoint(x, down))
	}
	if left := x - 1; left >= 0 {
		result = append(result, NewPoint(left, y))
	}
	if right := x + 1; right < width {
		result = append(result, NewPoint(right, y))
	}
	return
}

// GetAdjacentTranslateCoordinateYX 获取一个基于 y、x 的二维矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
func GetAdjacentTranslateCoordinateYX[T any, P generic.Number](matrix [][]T, x, y P) (result []Point[P]) {
	width := P(len(matrix[0]))
	height := P(len(matrix))
	if up := y - 1; up >= 0 {
		result = append(result, NewPoint(x, up))
	}
	if down := y + 1; down < height {
		result = append(result, NewPoint(x, down))
	}
	if left := x - 1; left >= 0 {
		result = append(result, NewPoint(left, y))
	}
	if right := x + 1; right < width {
		result = append(result, NewPoint(right, y))
	}
	return
}

// GetAdjacentDiagonalsPos 获取一个连续位置的矩阵中，特定位置相邻的对角线最多四个方向的位置
func GetAdjacentDiagonalsPos[T any, P generic.Number](matrix []T, width, pos P) (result []P) {
	size := P(len(matrix))
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

// GetAdjacentDiagonalsCoordinateXY 获取一个基于 x、y 的二维矩阵中，特定位置相邻的对角线最多四个方向的位置
func GetAdjacentDiagonalsCoordinateXY[T any, P generic.Number](matrix [][]T, x, y P) (result []Point[P]) {
	width := P(len(matrix[0]))
	height := P(len(matrix))
	if nx, ny := x-1, y-1; nx >= 0 && ny >= 0 {
		result = append(result, NewPoint(nx, ny))
	}
	if nx, ny := x+1, y-1; nx < width && ny >= 0 {
		result = append(result, NewPoint(nx, ny))
	}
	if nx, ny := x-1, y+1; nx >= 0 && ny < height {
		result = append(result, NewPoint(nx, ny))
	}
	if nx, ny := x+1, y+1; nx < width && ny < height {
		result = append(result, NewPoint(nx, ny))
	}
	return
}

// GetAdjacentDiagonalsCoordinateYX 获取一个基于 tx 的二维矩阵中，特定位置相邻的对角线最多四个方向的位置
func GetAdjacentDiagonalsCoordinateYX[T any, P generic.Number](matrix [][]T, x, y P) (result []Point[P]) {
	width := P(len(matrix))
	height := P(len(matrix[0]))
	if nx, ny := x-1, y-1; nx >= 0 && ny >= 0 {
		result = append(result, NewPoint(nx, ny))
	}
	if nx, ny := x+1, y-1; nx < width && ny >= 0 {
		result = append(result, NewPoint(nx, ny))
	}
	if nx, ny := x-1, y+1; nx >= 0 && ny < height {
		result = append(result, NewPoint(nx, ny))
	}
	if nx, ny := x+1, y+1; nx < width && ny < height {
		result = append(result, NewPoint(nx, ny))
	}
	return
}

// GetAdjacentPos 获取一个连续位置的矩阵中，特定位置相邻的最多八个方向的位置
func GetAdjacentPos[T any, P generic.Number](matrix []T, width, pos P) (result []P) {
	return append(GetAdjacentTranslatePos(matrix, width, pos), GetAdjacentDiagonalsPos(matrix, width, pos)...)
}

// GetAdjacentCoordinateXY 获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多八个方向的位置
func GetAdjacentCoordinateXY[T any, P generic.Number](matrix [][]T, x, y P) (result []Point[P]) {
	return append(GetAdjacentTranslateCoordinateXY(matrix, x, y), GetAdjacentDiagonalsCoordinateXY(matrix, x, y)...)
}

// GetAdjacentCoordinateYX 获取一个基于 yx 的二维矩阵中，特定位置相邻的最多八个方向的位置
func GetAdjacentCoordinateYX[T any, P generic.Number](matrix [][]T, x, y P) (result []Point[P]) {
	return append(GetAdjacentTranslateCoordinateYX(matrix, x, y), GetAdjacentDiagonalsCoordinateYX(matrix, x, y)...)
}

// CoordinateMatrixToPosMatrix 将二维矩阵转换为顺序的二维矩阵
func CoordinateMatrixToPosMatrix[V any](matrix [][]V) (width int, posMatrix []V) {
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

// GetShapeCoverageAreaWithCoordinateArray 通过传入的一组坐标 xys 计算一个图形覆盖的矩形范围
func GetShapeCoverageAreaWithCoordinateArray[V generic.Number](xys ...Point[V]) (left, right, top, bottom V) {
	hasLeft, hasTop := false, false
	for _, xy := range xys {
		x, y := CoordinateArrayToCoordinate(xy)
		if x < left || !hasLeft {
			hasLeft = true
			left = x
		}
		if x > right {
			right = x
		}
		if y < top || !hasTop {
			hasTop = true
			top = y
		}
		if y > bottom {
			bottom = y
		}
	}
	return
}

// GetShapeCoverageAreaWithPos 通过传入的一组坐标 xys 计算一个图形覆盖的矩形范围
func GetShapeCoverageAreaWithPos[V generic.Number](width V, positions ...V) (left, right, top, bottom V) {
	hasLeft, hasTop := false, false
	for _, pos := range positions {
		x, y := PosToCoordinate(width, pos)
		if x < left || !hasLeft {
			hasLeft = true
			left = x
		}
		if x > right {
			right = x
		}
		if y < top || !hasTop {
			hasTop = true
			top = y
		}
		if y > bottom {
			bottom = y
		}
	}
	return
}
