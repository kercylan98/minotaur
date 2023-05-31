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
