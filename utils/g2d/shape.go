package g2d

import (
	"github.com/kercylan98/minotaur/utils/g2d/shape"
	"sort"
)

type MatrixShapeSearchResult[Mark any] struct {
	Shape  *shape.Shape[Mark]
	Points []shape.Point
}

// MatrixShapeSearchWithYX 二维矩阵形状搜索
func MatrixShapeSearchWithYX[T any, Mark any](matrix [][]T, shapes []*shape.Shape[Mark], checkMatchHandle func(val T) bool) []MatrixShapeSearchResult[Mark] {
	record := map[int]map[int]bool{}
	width := len(matrix[0])
	height := len(matrix)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			record[x] = map[int]bool{y: true}
		}
	}

	var result []MatrixShapeSearchResult[Mark]

	for _, s := range shapes {
		points := s.GetPoints()
		x, y := 0, 0
		mx, my := s.GetMaxXY()
		for {
			if x+mx >= width {
				x = 0
				y++
			}
			if y+my >= height {
				break
			}
			var count int
			for _, point := range points {
				px, py := point.GetXY()
				px, py = px+x, py+y
				if record[px][py] {
					break
				}
				if checkMatchHandle(matrix[py][px]) {
					count++
				} else {
					break
				}
			}
			if count == len(points) {
				target := MatrixShapeSearchResult[Mark]{
					Shape: s,
				}
				for _, point := range points {
					px, py := point.GetXY()
					px, py = px+x, py+y
					ys, exist := record[px]
					if !exist {
						ys = map[int]bool{}
						record[px] = ys
					}
					ys[py] = true
					target.Points = append(target.Points, shape.NewPoint(px, py))
				}
				result = append(result, target)
			}
			x++
		}
	}

	return result
}

// MatrixShapeSearchWithXY 二维矩阵形状搜索
func MatrixShapeSearchWithXY[T any, Mark any](matrix [][]T, shapes []*shape.Shape[Mark], checkMatchHandle func(val T) bool) []MatrixShapeSearchResult[Mark] {
	width := len(matrix)
	height := len(matrix[0])
	record := map[int]map[int]bool{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			record[x] = map[int]bool{y: true}
		}
	}

	var result []MatrixShapeSearchResult[Mark]

	for _, s := range shapes {
		points := s.GetPoints()
		x, y := 0, 0
		mx, my := s.GetMaxXY()
		for {
			if x+mx >= width {
				x = 0
				y++
			}
			if y+my >= height {
				break
			}
			var count int
			for _, point := range points {
				px, py := point.GetXY()
				px, py = px+x, py+y
				if record[px][py] {
					break
				}
				if checkMatchHandle(matrix[px][py]) {
					count++
				} else {
					break
				}
			}
			if count == len(points) {
				target := MatrixShapeSearchResult[Mark]{
					Shape: s,
				}
				for _, point := range points {
					px, py := point.GetXY()
					px, py = px+x, py+y
					ys, exist := record[px]
					if !exist {
						ys = map[int]bool{}
						record[px] = ys
					}
					ys[py] = true
					target.Points = append(target.Points, shape.NewPoint(px, py))
				}
				result = append(result, target)
			}
			x++
		}
	}

	return result
}

// SearchNotRepeatCross 在一组二维坐标中从大到小搜索不重复交叉（十字）线
//   - 不重复指一个位置被使用后将不会被其他交叉线（十字）使用
func SearchNotRepeatCross(xys ...[2]int) (result [][][2]int) {
	left, _, top, _ := GetShapeCoverageArea(xys...)
	rectangleShape := GenerateShape(xys...)
	record := map[int]map[int]bool{}
	for x := 0; x < len(rectangleShape); x++ {
		for y := 0; y < len(rectangleShape[0]); y++ {
			record[x] = map[int]bool{}
		}
	}

	for _, xy := range xys {
		var points [][2]int
		var find = map[int]bool{}
		x, y := PositionArrayToXY(xy)
		x = x + (0 - left)
		y = y + (0 - top)
		// 搜索四个方向
		for sx := x - 1; sx >= 0; sx-- {
			if !rectangleShape[sx][y] {
				break
			}
			find[1] = true
			points = append(points, [2]int{sx + left, y + top})
		}
		if !find[1] {
			continue
		}
		for sx := x + 1; sx < len(rectangleShape); sx++ {
			if !rectangleShape[sx][y] {
				break
			}
			find[2] = true
			points = append(points, [2]int{sx + left, y + top})
		}
		if !find[2] {
			continue
		}
		for sy := y - 1; sy >= 0; sy-- {
			if !rectangleShape[x][sy] {
				break
			}
			find[3] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		if !find[3] {
			continue
		}
		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
			if !rectangleShape[x][sy] {
				break
			}
			find[4] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		if !find[4] {
			continue
		}
		result = append(result, append(points, [2]int{x + left, y + top}))
	}

	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})

	var notRepeat [][][2]int
	for _, points := range result {
		var match = true
		for _, point := range points {
			x, y := PositionArrayToXY(point)
			x = x + (0 - left)
			y = y + (0 - top)
			if record[x][y] {
				match = false
				break
			}
			record[x][y] = true
		}
		if match {
			notRepeat = append(notRepeat, points)
		}
	}

	return notRepeat
}

// SearchNotRepeatStraightLine 在一组二维坐标中从大到小搜索不重复的直线
//   - 最低需要长度为3
func SearchNotRepeatStraightLine(minLength int, xys ...[2]int) (result [][][2]int) {
	if minLength < 3 {
		return nil
	}
	left, _, top, _ := GetShapeCoverageArea(xys...)
	rectangleShape := GenerateShape(xys...)
	record := map[int]map[int]bool{}
	for x := 0; x < len(rectangleShape); x++ {
		for y := 0; y < len(rectangleShape[0]); y++ {
			record[x] = map[int]bool{}
		}
	}

	for _, xy := range xys {
		var points [][2]int
		var find = map[int]bool{}
		x, y := PositionArrayToXY(xy)
		x = x + (0 - left)
		y = y + (0 - top)
		// 搜索四个方向
		for sx := x - 1; sx >= 0; sx-- {
			if !rectangleShape[sx][y] {
				break
			}
			find[1] = true
			points = append(points, [2]int{sx + left, y + top})
		}
		if !find[1] {
			goto up
		}
		for sx := x + 1; sx < len(rectangleShape); sx++ {
			if !rectangleShape[sx][y] {
				break
			}
			find[2] = true
			points = append(points, [2]int{sx + left, y + top})
		}
		if !find[2] {
			points = nil
		}
	up:
		for sy := y - 1; sy >= 0; sy-- {
			if !rectangleShape[x][sy] {
				break
			}
			find[3] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		if !find[3] {
			continue
		}
		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
			if !rectangleShape[x][sy] {
				break
			}
			find[4] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		if !find[4] {
			continue
		}
		if len(find) != 2 {
			continue
		}
		result = append(result, append(points, [2]int{x + left, y + top}))
	}

	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})

	var notRepeat [][][2]int
	for _, points := range result {
		if len(points) < minLength {
			continue
		}
		var match = true
		for _, point := range points {
			x, y := PositionArrayToXY(point)
			x = x + (0 - left)
			y = y + (0 - top)
			if record[x][y] {
				match = false
				break
			}
			record[x][y] = true
		}
		if match {
			notRepeat = append(notRepeat, points)
		}
	}

	return notRepeat
}

// SearchNotRepeatT 在一组二维坐标中从大到小搜索不重复T型（T）线
func SearchNotRepeatT(minLength int, xys ...[2]int) (result [][][2]int) {
	if minLength < 4 {
		return nil
	}
	left, _, top, _ := GetShapeCoverageArea(xys...)
	rectangleShape := GenerateShape(xys...)
	record := map[int]map[int]bool{}
	for x := 0; x < len(rectangleShape); x++ {
		for y := 0; y < len(rectangleShape[0]); y++ {
			record[x] = map[int]bool{}
		}
	}

	for _, xy := range xys {
		var points [][2]int
		var find = map[int]bool{}
		x, y := PositionArrayToXY(xy)
		x = x + (0 - left)
		y = y + (0 - top)
		// 搜索四个方向
		for sx := x - 1; sx >= 0; sx-- {
			if !rectangleShape[sx][y] {
				break
			}
			find[1] = true
			points = append(points, [2]int{sx + left, y + top})
		}
		for sx := x + 1; sx < len(rectangleShape); sx++ {
			if !rectangleShape[sx][y] {
				break
			}
			find[2] = true
			points = append(points, [2]int{sx + left, y + top})
		}
		for sy := y - 1; sy >= 0; sy-- {
			if !rectangleShape[x][sy] {
				break
			}
			find[3] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
			if !rectangleShape[x][sy] {
				break
			}
			find[4] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		if len(find) != 3 {
			continue
		}
		result = append(result, append(points, [2]int{x + left, y + top}))
	}

	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})

	var notRepeat [][][2]int
	for _, points := range result {
		if len(points) < minLength {
			continue
		}
		var match = true
		for _, point := range points {
			x, y := PositionArrayToXY(point)
			x = x + (0 - left)
			y = y + (0 - top)
			if record[x][y] {
				match = false
				break
			}
			record[x][y] = true
		}
		if match {
			notRepeat = append(notRepeat, points)
		}
	}

	return notRepeat
}

// SearchNotRepeatRightAngle 在一组二维坐标中从大到小搜索不重复的直角（L）线
func SearchNotRepeatRightAngle(minLength int, xys ...[2]int) (result [][][2]int) {
	if minLength < 3 {
		return nil
	}
	left, _, top, _ := GetShapeCoverageArea(xys...)
	rectangleShape := GenerateShape(xys...)
	record := map[int]map[int]bool{}
	for x := 0; x < len(rectangleShape); x++ {
		for y := 0; y < len(rectangleShape[0]); y++ {
			record[x] = map[int]bool{}
		}
	}

	for _, xy := range xys {
		var points [][2]int
		var find = map[int]bool{}
		x, y := PositionArrayToXY(xy)
		x = x + (0 - left)
		y = y + (0 - top)
		// 搜索四个方向
		for sx := x - 1; sx >= 0; sx-- {
			if !rectangleShape[sx][y] {
				break
			}
			find[1] = true
			points = append(points, [2]int{sx + left, y + top})
		}
		if find[1] {
			goto up
		}
		for sx := x + 1; sx < len(rectangleShape); sx++ {
			if !rectangleShape[sx][y] {
				break
			}
			find[2] = true
			points = append(points, [2]int{sx + left, y + top})
		}
	up:
		for sy := y - 1; sy >= 0; sy-- {
			if !rectangleShape[x][sy] {
				break
			}
			find[3] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		if find[3] {
			goto end
		}
		// down
		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
			if !rectangleShape[x][sy] {
				break
			}
			find[4] = true
			points = append(points, [2]int{x + left, sy + top})
		}
		if !find[4] {
			continue
		}
	end:
		{
			if len(find) != 2 {
				continue
			}
			result = append(result, append(points, [2]int{x + left, y + top}))
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})

	var notRepeat [][][2]int
	for _, points := range result {
		if len(points) < minLength {
			continue
		}
		var match = true
		for _, point := range points {
			x, y := PositionArrayToXY(point)
			x = x + (0 - left)
			y = y + (0 - top)
			if record[x][y] {
				match = false
				break
			}
			record[x][y] = true
		}
		if match {
			notRepeat = append(notRepeat, points)
		}
	}

	return notRepeat
}

// SearchNotRepeatFullRectangle 在一组二维坐标中从大到小搜索不重复的填充满的矩形
//   - 不重复指一个位置被使用后将不会被其他矩形使用
//   - 返回值表示了匹配的形状的左上角和右下角的点坐标
func SearchNotRepeatFullRectangle(minWidth, minHeight int, xys ...[2]int) (result [][2][2]int) {
	left, _, top, _ := GetShapeCoverageArea(xys...)
	rectangleShape := GenerateShape(xys...)
	record := map[int]map[int]bool{}
	width := len(rectangleShape)
	height := len(rectangleShape[0])
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			record[x] = map[int]bool{}
		}
	}

	shapes := GetExpressibleRectangleBySize(width, height, minWidth, minHeight)
	for _, s := range shapes {
		x, y := 0, 0
		for {
			if x+s[0] >= width {
				x = 0
				y++
			}
			if y+s[1] >= height {
				break
			}
			points := GetRectangleFullPoints(s[0]+1, s[1]+1)
			find := 0
			for _, point := range points {
				px, py := PositionArrayToXY(point)
				ox, oy := px+x, py+y
				if record[ox][oy] || !rectangleShape[ox][oy] {
					find = 0
					break
				}
				find++
			}
			if find == len(points) {
				for _, point := range points {
					px, py := PositionArrayToXY(point)
					record[px+x][py+y] = true
				}
				result = append(result, [2][2]int{
					{x + left, y + top}, {x + left + s[0], y + top + s[1]},
				})
			}

			x++
		}
	}

	return result
}

// GetRectangleFullPoints 获取一个矩形包含的所有点
func GetRectangleFullPoints(width, height int) (result [][2]int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			result = append(result, [2]int{x, y})
		}
	}
	return
}

// GetRectangleFullPointsByXY 通过开始结束坐标获取一个矩形包含的所有点
//   - 例如 1,1 到 2,2 的矩形结果为 1,1 2,1 1,2 2,2
func GetRectangleFullPointsByXY(startX, startY, endX, endY int) (result [][2]int) {
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			result = append(result, [2]int{x, y})
		}
	}
	return
}

// GetExpressibleRectangle 获取一个宽高可表达的所有矩形形状
//   - 返回值表示了每一个矩形右下角的x,y位置（左上角始终为0, 0）
//   - 矩形尺寸由大到小
func GetExpressibleRectangle(width, height int) (result [][2]int) {
	return GetExpressibleRectangleBySize(width, height, 1, 1)
}

// GetExpressibleRectangleBySize 获取一个宽高可表达的所有特定尺寸以上的矩形形状
//   - 返回值表示了每一个矩形右下角的x,y位置（左上角始终为0, 0）
//   - 矩形尺寸由大到小
func GetExpressibleRectangleBySize(width, height, minWidth, minHeight int) (result [][2]int) {
	if width == 0 || height == 0 {
		return nil
	}
	if width < minWidth || height < minHeight {
		return nil
	}
	width--
	height--
	for {
		rightBottom := [2]int{width, height}
		result = append(result, rightBottom)
		if width == 0 && height == 0 || (width < minWidth && height < minHeight) {
			return
		}
		if width == height {
			width--
		} else if width < height {
			width++
			height--
		} else if width > height {
			width--
		}
	}
}

// GenerateShape 生成一组二维坐标的形状
//   - 这个形状将被在一个刚好能容纳形状的矩形中表示
//   - 为true的位置表示了形状的每一个点
func GenerateShape(xys ...[2]int) [][]bool {
	left, r, top, b := GetShapeCoverageArea(xys...)
	_, right, _, bottom := CoverageAreaBoundless(left, r, top, b)
	w, h := right+1, bottom+1
	m := make([][]bool, w)
	for x := 0; x < w; x++ {
		m[x] = make([]bool, h)
	}
	for _, xy := range xys {
		x, y := PositionArrayToXY(xy)
		m[x-(r-right)][y-(b-bottom)] = true
	}
	return m
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
