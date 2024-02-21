package geometry

import "github.com/kercylan98/minotaur/utils/generic"

// GetAdjacentTranslatePos 获取一个连续位置的矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
func GetAdjacentTranslatePos[T any, P generic.SignedNumber](matrix []T, width, pos P) (result []P) {
	wf, pf := float64(width), float64(pos)
	size := float64(len(matrix))
	currentRow := pf / wf
	if up := -wf; up >= 0 {
		result = append(result, P(up))
	}
	if down := pf + wf; down < size {
		result = append(result, P(down))
	}
	if left := pf - 1; left >= 0 && currentRow == (left/wf) {
		result = append(result, P(left))
	}
	if right := pf + 1; right < size && currentRow == (right/wf) {
		result = append(result, P(right))
	}
	return
}

// GetAdjacentTranslateCoordinateXY 获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
func GetAdjacentTranslateCoordinateXY[T any, P generic.SignedNumber](matrix [][]T, x, y P) (result []Point[P]) {
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
func GetAdjacentTranslateCoordinateYX[T any, P generic.SignedNumber](matrix [][]T, x, y P) (result []Point[P]) {
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
func GetAdjacentDiagonalsPos[T any, P generic.SignedNumber](matrix []T, width, pos P) (result []P) {
	size := float64(len(matrix))
	wf, pf := float64(width), float64(pos)
	currentRow := pf / wf
	if topLeft := pf - wf - 1; topLeft >= 0 && currentRow-1 == (topLeft/wf) {
		result = append(result, P(topLeft))
	}
	if topRight := pf - wf + 1; topRight >= 0 && currentRow-1 == (topRight/wf) {
		result = append(result, P(topRight))
	}
	if bottomLeft := pf + wf - 1; bottomLeft < size && currentRow+1 == (bottomLeft/wf) {
		result = append(result, P(bottomLeft))
	}
	if bottomRight := pf + wf + 1; bottomRight < size && currentRow+1 == (bottomRight/wf) {
		result = append(result, P(bottomRight))
	}
	return
}

// GetAdjacentDiagonalsCoordinateXY 获取一个基于 x、y 的二维矩阵中，特定位置相邻的对角线最多四个方向的位置
func GetAdjacentDiagonalsCoordinateXY[T any, P generic.SignedNumber](matrix [][]T, x, y P) (result []Point[P]) {
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
func GetAdjacentDiagonalsCoordinateYX[T any, P generic.SignedNumber](matrix [][]T, x, y P) (result []Point[P]) {
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
func GetAdjacentPos[T any, P generic.SignedNumber](matrix []T, width, pos P) (result []P) {
	return append(GetAdjacentTranslatePos(matrix, width, pos), GetAdjacentDiagonalsPos(matrix, width, pos)...)
}

// GetAdjacentCoordinateXY 获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多八个方向的位置
func GetAdjacentCoordinateXY[T any, P generic.SignedNumber](matrix [][]T, x, y P) (result []Point[P]) {
	return append(GetAdjacentTranslateCoordinateXY(matrix, x, y), GetAdjacentDiagonalsCoordinateXY(matrix, x, y)...)
}

// GetAdjacentCoordinateYX 获取一个基于 yx 的二维矩阵中，特定位置相邻的最多八个方向的位置
func GetAdjacentCoordinateYX[T any, P generic.SignedNumber](matrix [][]T, x, y P) (result []Point[P]) {
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

// GetShapeCoverageAreaWithPoint 通过传入的一组坐标 points 计算一个图形覆盖的矩形范围
func GetShapeCoverageAreaWithPoint[V generic.SignedNumber](points ...Point[V]) (left, right, top, bottom V) {
	hasLeft, hasTop := false, false
	for _, xy := range points {
		x, y := PointToCoordinate(xy)
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

// GetShapeCoverageAreaWithPos 通过传入的一组坐标 positions 计算一个图形覆盖的矩形范围
func GetShapeCoverageAreaWithPos[V generic.SignedNumber](width V, positions ...V) (left, right, top, bottom V) {
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

// CoverageAreaBoundless 将一个图形覆盖矩形范围设置为无边的
//   - 无边化表示会将多余的部分进行裁剪，例如图形左边从 2 开始的时候，那么左边将会被裁剪到从 0 开始
func CoverageAreaBoundless[V generic.SignedNumber](l, r, t, b V) (left, right, top, bottom V) {
	differentX := 0 - l
	differentY := 0 - t
	left = l + differentX
	right = r + differentX
	top = t + differentY
	bottom = b + differentY
	return
}

// GenerateShapeOnRectangle 生成一组二维坐标的形状
//   - 这个形状将被在一个刚好能容纳形状的矩形中表示
//   - 为 true 的位置表示了形状的每一个点
func GenerateShapeOnRectangle[V generic.SignedNumber](points ...Point[V]) (result []PointCap[V, bool]) {
	left, r, top, b := GetShapeCoverageAreaWithPoint(points...)
	_, right, _, bottom := CoverageAreaBoundless(left, r, top, b)
	w, h := right+1, bottom+1
	result = make([]PointCap[V, bool], int(w*h))
	for _, xy := range points {
		x, y := xy.GetXY()
		sx := x - (r - right)
		sy := y - (b - bottom)
		pos := CoordinateToPos(w, sx, sy)
		pointCap := &result[int(pos)]
		pointCap.Point[0] = sx
		pointCap.Point[1] = sy
		pointCap.Data = true
	}
	for pos, pointCap := range result {
		if !pointCap.Data {
			pointCap := &result[pos]
			sx, sy := PosToCoordinate(w, V(pos))
			pointCap.Point[0] = sx
			pointCap.Point[1] = sy
		}
	}
	return
}

// GenerateShapeOnRectangleWithCoordinate 生成一组二维坐标的形状
//   - 这个形状将被在一个刚好能容纳形状的矩形中表示
//   - 为 true 的位置表示了形状的每一个点
func GenerateShapeOnRectangleWithCoordinate[V generic.SignedNumber](points ...Point[V]) (result [][]bool) {
	left, r, top, b := GetShapeCoverageAreaWithPoint(points...)
	_, right, _, bottom := CoverageAreaBoundless(left, r, top, b)
	w, h := right+1, bottom+1
	result = make([][]bool, int(w))
	for x := V(0); x < w; x++ {
		result[int(x)] = make([]bool, int(h))
	}
	for _, xy := range points {
		x, y := xy.GetXY()
		sx := x - (r - right)
		sy := y - (b - bottom)
		result[int(sx)][int(sy)] = true
	}
	return
}

// GetExpressibleRectangleBySize 获取一个宽高可表达的所有特定尺寸以上的矩形形状
//   - 返回值表示了每一个矩形右下角的x,y位置（左上角始终为0, 0）
//   - 矩形尺寸由大到小
func GetExpressibleRectangleBySize[V generic.SignedNumber](width, height, minWidth, minHeight V) (result []Point[V]) {
	sourceWidth := width
	if width == 0 || height == 0 {
		return nil
	}
	if width < minWidth || height < minHeight {
		return nil
	}
	width--
	height--
	for {
		rightBottom := NewPoint(width, height)
		result = append(result, rightBottom)
		if width == 0 && height == 0 || (width < minWidth && height < minHeight) {
			return
		}
		if width == height {
			width--
		} else if width < height {
			if width+1 == sourceWidth {
				height--
			} else {
				width++
				height--
			}
		} else if width > height {
			width--
		}
	}
}

// GetExpressibleRectangle 获取一个宽高可表达的所有矩形形状
//   - 返回值表示了每一个矩形右下角的x,y位置（左上角始终为0, 0）
//   - 矩形尺寸由大到小
func GetExpressibleRectangle[V generic.SignedNumber](width, height V) (result []Point[V]) {
	return GetExpressibleRectangleBySize(width, height, 1, 1)
}

// GetRectangleFullPointsByXY 通过开始结束坐标获取一个矩形包含的所有点
//   - 例如 1,1 到 2,2 的矩形结果为 1,1 2,1 1,2 2,2
func GetRectangleFullPointsByXY[V generic.SignedNumber](startX, startY, endX, endY V) (result []Point[V]) {
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			result = append(result, NewPoint(x, y))
		}
	}
	return
}

// GetRectangleFullPoints 获取一个矩形填充满后包含的所有点
func GetRectangleFullPoints[V generic.SignedNumber](width, height V) (result []Point[V]) {
	for x := V(0); x < width; x++ {
		for y := V(0); y < height; y++ {
			result = append(result, NewPoint(x, y))
		}
	}
	return
}

// GetRectangleFullPos 获取一个矩形填充满后包含的所有位置
func GetRectangleFullPos[V generic.SignedNumber](width, height V) (result []V) {
	size := int(width * height)
	result = make([]V, 0, size)
	for pos := 0; pos < size; pos++ {
		result[pos] = V(pos)
	}
	return
}

// CalcRectangleCentroid 计算矩形质心
//   - 非多边形质心计算，仅为顶点的平均值 - 该区域中多边形因子的适当质心
func CalcRectangleCentroid[V generic.SignedNumber](shape Shape[V]) Point[V] {
	var x, y float64
	length := float64(shape.PointCount())
	for _, point := range shape.Points() {
		x += float64(point.GetX())
		y += float64(point.GetY())
	}
	x /= length
	y /= length
	return NewPoint(V(x), V(x))
}
