package g2d

import (
	"github.com/kercylan98/minotaur/utils/g2d/shape"
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
	record := map[int]map[int]bool{}
	width := len(matrix)
	height := len(matrix[0])
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

// SearchNotRepeatFullRectangle 在一组二维坐标中从大到小搜索不重复的填充满的矩形
//   - 不重复指一个位置被使用后将不会被其他矩形使用
//   - 返回值表示了匹配的形状的左上角和右下角的点坐标
func SearchNotRepeatFullRectangle(xys ...[2]int) (result [][2][2]int) {
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

	shapes := GetExpressibleRectangleBySize(width, height, 2, 2)
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
	_, right, _, bottom := CoverageAreaBoundless(GetShapeCoverageArea(xys...))
	w, h := right+1, bottom+1
	m := make([][]bool, w)
	for x := 0; x < w; x++ {
		m[x] = make([]bool, h)
	}
	for _, xy := range xys {
		x, y := PositionArrayToXY(xy)
		m[x][y] = true
	}
	return m
}
