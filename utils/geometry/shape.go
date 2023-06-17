package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/slice"
	"math"
	"sort"
)

var (
	ShapeStringHasBorder = false // 控制 Shape.String 是否拥有边界
)

// Shape 通过多个点表示了一个形状
type Shape[V generic.Number] []Point[V]

// Points 获取这个形状的所有点
func (slf Shape[V]) Points() []Point[V] {
	return slf
}

// String 将该形状转换为可视化的字符串进行返回
func (slf Shape[V]) String() string {
	var result string
	left, right, top, bottom := GetShapeCoverageAreaWithCoordinateArray(slf.Points()...)
	width := right - left + 1
	height := bottom - top + 1
	if !ShapeStringHasBorder {
		for y := top; y < top+height; y++ {
			for x := left; x < left+width; x++ {
				exist := false
				for _, p := range slf {
					if x == p.GetX() && y == p.GetY() {
						exist = true
						break
					}
				}
				if exist {
					result += "X "
				} else {
					result += "# "
				}
			}
			result += "\r\n"
		}
	} else {
		for y := V(0); y < top+height; y++ {
			for x := V(0); x < left+width; x++ {
				exist := false
				for _, p := range slf {
					if x == p.GetX() && y == p.GetY() {
						exist = true
						break
					}
				}
				if exist {
					result += "X "
				} else {
					result += "# "
				}
			}
			result += "\r\n"
		}
	}
	return result
}

// ShapeSearch 获取该形状中包含的所有图形组合及其位置
//   - 需要注意的是，即便图形最终表示为相同的，但是只要位置组合顺序不同，那么也将被认定为一种图形组合
//   - [[1 0] [1 1] [1 2]] 和 [[1 1] [1 0] [1 2]] 可以被视为两个图形组合
//   - 返回的坐标为原始形状的坐标
//
// 可通过可选项对搜索结果进行过滤
func (slf Shape[V]) ShapeSearch(options ...ShapeSearchOption) (result []Shape[V]) {
	opt := &shapeSearchOptions{upperLimit: math.MaxInt}
	opt.directionCountUpper = map[Direction]int{}
	for _, d := range DirectionUDLR {
		opt.directionCountUpper[d] = math.MaxInt
	}

	for _, option := range options {
		option(opt)
	}

	var shapes []Shape[V]
	switch opt.sort {
	case 1:
		shapes = slf.getAllGraphicCompositionWithAsc(opt)
	case -1:
		shapes = slf.getAllGraphicCompositionWithDesc(opt)
	default:
		shapes = slf.getAllGraphicComposition(opt)
	}
	result = shapes

	if opt.deduplication {
		deduplication := make(map[V]struct{})
		w := V(len(slf.Points()))

		var notRepeat = make([]Shape[V], 0, len(result))
		for _, points := range result {
			count := len(points)
			if count < opt.lowerLimit || count > opt.upperLimit {
				continue
			}
			var match = true
			for _, point := range points {
				pos := point.GetPos(w)
				if _, exist := deduplication[pos]; exist {
					match = false
					break
				}
				deduplication[pos] = struct{}{}
			}
			if match {
				notRepeat = append(notRepeat, points)
			}
		}

		result = notRepeat
	} else {
		limit := make([]Shape[V], 0, len(result))
		for _, shape := range result {
			count := len(shape.Points())
			if count < opt.lowerLimit || count > opt.upperLimit {
				continue
			}
			limit = append(limit, shape)
		}
		result = limit
	}

	return
}

// getAllGraphicComposition 获取该形状中包含的所有图形组合及其位置
//   - 需要注意的是，即便图形最终表示为相同的，但是只要位置组合顺序不同，那么也将被认定为一种图形组合
//   - [[1 0] [1 1] [1 2]] 和 [[1 1] [1 0] [1 2]] 可以被视为两个图形组合
//   - 返回的坐标为原始形状的坐标
func (slf Shape[V]) getAllGraphicComposition(opt *shapeSearchOptions) (result []Shape[V]) {
	left, right, top, bottom := GetShapeCoverageAreaWithCoordinateArray(slf.Points()...)
	width := right - left + 1
	height := bottom - top + 1
	areaWidth := width + left
	areaHeight := height + top
	rectangleShape := GenerateShapeOnRectangle(slf.Points()...)
	records := make(map[V]struct{})

	var match = func(links Shape[V], directionCount map[Direction]int, count int) bool {
		match := true
		for _, direction := range DirectionUDLR {
			c := directionCount[direction]
			if c < opt.directionCountLower[direction] || c > opt.directionCountUpper[direction] {
				match = false
				break
			}
		}

		if opt.directionCount > 0 && len(directionCount) != opt.directionCount {
			match = false
		}

		if directionCount[GetOppositionDirection(opt.oppositionDirection)] > 0 {
			match = false
		}

		if opt.ra {
			match = false
			if directionCount[DirectionLeft] > 0 && directionCount[DirectionUp] > 0 && count == directionCount[DirectionLeft]+directionCount[DirectionUp] {
				match = true
			} else if directionCount[DirectionUp] > 0 && directionCount[DirectionRight] > 0 && count == directionCount[DirectionUp]+directionCount[DirectionRight] {
				match = true
			} else if directionCount[DirectionRight] > 0 && directionCount[DirectionDown] > 0 && count == directionCount[DirectionRight]+directionCount[DirectionDown] {
				match = true
			} else if directionCount[DirectionDown] > 0 && directionCount[DirectionLeft] > 0 && count == directionCount[DirectionDown]+directionCount[DirectionLeft] {
				match = true
			}
		}

		if match {
			result = append(result, links)
		}
		return match
	}

	// 通过每个点扩散图形
	for _, point := range slf.Points() {
		// 搜索四个方向
		var next = -1
		var directionPoint = point
		var links = Shape[V]{point}
		var directionCount = map[Direction]int{}
		var count = 0
		for i, directions := range [][]Direction{DirectionUDLR, DirectionLRUD} {
			var direction Direction
			next, direction = slice.NextLoop(directions, next)
			for {
				directionPoint = GetDirectionNextWithCoordinateArray(direction, directionPoint)
				if px, py := directionPoint.GetXY(); px < 0 || px >= areaWidth || py < 0 || py >= areaHeight {
					break
				}
				offset := directionPoint.GetOffset(-left, -top)
				if offset.Negative() {
					break
				}
				offsetPos := int(offset.GetPos(width))
				if offsetPos < 0 || offsetPos >= len(rectangleShape) || !rectangleShape[offsetPos].Data {
					break
				}
				links = append(links, directionPoint)
				directionCount[direction]++
				count++
				match(links, directionCount, count)
				pos := directionPoint.GetPos(areaWidth)
				if _, exist := records[pos]; !exist {
					result = append(result, Shape[V]{directionPoint})
					records[pos] = struct{}{}
				}

			}

			finish := false
			switch i {
			case 0:
				if direction == DirectionRight {
					finish = true
				}
			case 1:
				if direction == DirectionDown {
					finish = true
				}
			}
			if finish {
				break
			}
			directionPoint = point
		}

	}

	return result
}

// getAllGraphicCompositionWithAsc 通过升序的方式获取该形状中包含的所有图形组合及其位置
//   - 升序指标为图形包含的点数量
//   - 其余内容可参考 getAllGraphicComposition
func (slf Shape[V]) getAllGraphicCompositionWithAsc(opt *shapeSearchOptions) (result []Shape[V]) {
	result = slf.getAllGraphicComposition(opt)
	sort.Slice(result, func(i, j int) bool {
		return len(result[i].Points()) < len(result[j].Points())
	})
	return
}

// getAllGraphicCompositionWithDesc 通过降序的方式获取该形状中包含的所有图形组合及其位置
//   - 降序指标为图形包含的点数量
//   - 其余内容可参考 GetAllGraphicComposition
func (slf Shape[V]) getAllGraphicCompositionWithDesc(opt *shapeSearchOptions) (result []Shape[V]) {
	result = slf.getAllGraphicComposition(opt)
	sort.Slice(result, func(i, j int) bool {
		return len(result[i].Points()) > len(result[j].Points())
	})
	return
}

//
//// SearchNotRepeatFullRectangle 在一组二维坐标中从大到小搜索不重复的填充满的矩形
////   - 不重复指一个位置被使用后将不会被其他矩形使用
////   - 返回值表示了匹配的形状的左上角和右下角的点坐标
//func SearchNotRepeatFullRectangle(minWidth, minHeight int, xys ...[2]int) (result [][2][2]int) {
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	width := len(rectangleShape)
//	height := len(rectangleShape[0])
//	for x := 0; x < width; x++ {
//		for y := 0; y < height; y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	shapes := GetExpressibleRectangleBySize(width, height, minWidth, minHeight)
//	for _, s := range shapes {
//		x, y := 0, 0
//		for {
//			if x+s[0] >= width {
//				x = 0
//				y++
//			}
//			if y+s[1] >= height {
//				break
//			}
//			points := GetRectangleFullPoints(s[0]+1, s[1]+1)
//			find := 0
//			for _, point := range points {
//				px, py := CoordinateArrayToCoordinate(point)
//				ox, oy := px+x, py+y
//				if record[ox][oy] || !rectangleShape[ox][oy] {
//					find = 0
//					break
//				}
//				find++
//			}
//			if find == len(points) {
//				for _, point := range points {
//					px, py := CoordinateArrayToCoordinate(point)
//					record[px+x][py+y] = true
//				}
//				result = append(result, [2][2]int{
//					{x + left, y + top}, {x + left + s[0], y + top + s[1]},
//				})
//			}
//
//			x++
//		}
//	}
//
//	return result
//}
//
//// SearchContainFullRectangle 在一组二维坐标中查找是否存在填充满的矩形
//func SearchContainFullRectangle(minWidth, minHeight int, xys ...[2]int) bool {
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	width := len(rectangleShape)
//	height := len(rectangleShape[0])
//	for x := 0; x < width; x++ {
//		for y := 0; y < height; y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	shapes := GetExpressibleRectangleBySize(width, height, minWidth, minHeight)
//	for _, s := range shapes {
//		x, y := 0, 0
//		for {
//			if x+s[0] >= width {
//				x = 0
//				y++
//			}
//			if y+s[1] >= height {
//				break
//			}
//			points := GetRectangleFullPoints(s[0]+1, s[1]+1)
//			find := 0
//			for _, point := range points {
//				px, py := CoordinateArrayToCoordinate(point)
//				ox, oy := px+x, py+y
//				if record[ox][oy] || !rectangleShape[ox][oy] {
//					find = 0
//					break
//				}
//				find++
//			}
//			if find == len(points) {
//				for _, point := range points {
//					px, py := CoordinateArrayToCoordinate(point)
//					record[px+x][py+y] = true
//				}
//				return true
//			}
//
//			x++
//		}
//	}
//
//	return false
//}
