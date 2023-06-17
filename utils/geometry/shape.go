package geometry

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/slice"
	"math"
	"sort"
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

	// 通过每个点扩散图形
	for _, point := range slf.Points() {
		// 搜索四个方向
		var next = -1
		var directionPoint = point
		var links = Shape[V]{point}
		for {
			var direction Direction
			next, direction = slice.NextLoop(Directions, next)
			for {
				directionPoint = GetDirectionNextWithCoordinateArray(direction, directionPoint)
				if px, py := directionPoint.GetXY(); px < 0 || px >= areaWidth || py < 0 || py >= areaHeight {
					break
				}
				if offsetPos := int(CoordinateArrayToPos(width, directionPoint.GetOffset(-left, -top))); offsetPos < 0 || offsetPos >= len(rectangleShape) || !rectangleShape[offsetPos].Data {
					break
				}
				links = append(links, directionPoint)
				pos := directionPoint.GetPos(areaWidth)
				if _, exist := records[pos]; !exist {
					result = append(result, Shape[V]{directionPoint})
					records[pos] = struct{}{}
				}
			}
			if direction == DirectionRight {
				break
			}
			directionPoint = point
		}
		result = append(result, links)
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
//// SearchNotRepeatCross 在一组二维坐标中从大到小搜索不重复交叉（十字）线
////   - 不重复指一个位置被使用后将不会被其他交叉线（十字）使用
//func SearchNotRepeatCross[V generic.Number](findHandle func(findCount map[Direction]int, nextDirection func(direction Direction), stop func()), points []Point[V]) (result [][]Point[V]) {
//	left, right, top, bottom := GetShapeCoverageAreaWithCoordinateArray(points...)
//	width := right - left + 1
//	height := bottom - top + 1
//	size := width * height
//	rectangleShape := GenerateShapeOnRectangle(points...)
//	record := map[V]map[V]bool{}
//	for x := V(0); x < width; x++ {
//		for y := V(0); y < height; y++ {
//			record[x] = map[V]bool{}
//		}
//	}
//
//	var findCount = map[Direction]int{}
//
//	for _, point := range points {
//
//		var next = -1
//		for {
//			var direction Direction
//			next, direction = slice.NextLoop(Directions, next)
//			nextPoint := point
//			for {
//				nextPoint = GetDirectionNextWithCoordinateArray(direction, point)
//				nextPos := nextPoint.GetPos(width)
//				if nextPos < 0 || nextPos >= size {
//					break
//				}
//				if rectangleShape[int(nextPos)].Data {
//					findCount[direction]++
//					var goToNextDirection bool
//					var stop bool
//					findHandle(findCount, func(direction Direction) {
//						switch direction {
//						case DirectionUp:
//							next = -1
//						case DirectionDown:
//							next = 0
//						case DirectionLeft:
//							next = 1
//						case DirectionRight:
//							next = 2
//						}
//						goToNextDirection = true
//					}, func() {
//						stop = true
//					})
//					if stop {
//						return
//					}
//					if goToNextDirection {
//						break
//					}
//				} else {
//					break
//				}
//			}
//		}
//
//		for _, direction := range Directions {
//			for {
//				nextPoint := GetDirectionNextWithCoordinateArray(direction, point)
//				nextPos := nextPoint.GetPos(width)
//				if nextPos < 0 || nextPos >= size {
//					break
//				}
//				if rectangleShape[int(nextPos)].Data {
//					findCount[direction]++
//				} else {
//					break
//				}
//			}
//
//			// 十字至少各边需要长度1
//			totalCount := hash.Sum(findCount)
//			if totalCount < 4 {
//				continue
//			}
//		}
//	}
//
//	for _, xy := range xys {
//		var points []Point[V]
//		var find = map[int]bool{}
//		x, y := xy.GetXY()
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[int(CoordinateToPos(width, sx, y))].Data {
//				break
//			}
//			find[1] = true
//			points = append(points, NewPoint(sx+left, y+top))
//		}
//		if !find[1] {
//			continue
//		}
//		for sx := x + 1; sx < V(len(rectangleShape)); sx++ {
//			if !rectangleShape[int(CoordinateToPos(width, sx, y))].Data {
//				break
//			}
//			find[2] = true
//			points = append(points, NewPoint(sx+left, y+top))
//		}
//		if !find[2] {
//			continue
//		}
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[int(CoordinateToPos(width, x, sy))].Data {
//				break
//			}
//			find[3] = true
//			points = append(points, NewPoint(x+left, sy+top))
//		}
//		if !find[3] {
//			continue
//		}
//		for sy := y + 1; sy < V(len(rectangleShape)); sy++ {
//			if !rectangleShape[int(CoordinateToPos(width, x, sy))].Data {
//				break
//			}
//			find[4] = true
//			points = append(points, NewPoint(x+left, sy+top))
//		}
//		if !find[4] {
//			continue
//		}
//		result = append(result, append(points, NewPoint(x+left, y+top)))
//	}
//
//	sort.Slice(result, func(i, j int) bool {
//		return len(result[i]) > len(result[j])
//	})
//
//	var notRepeat [][]Point[V]
//	for _, points := range result {
//		var match = true
//		for _, point := range points {
//			x, y := CoordinateArrayToCoordinate(point)
//			x = x + (0 - left)
//			y = y + (0 - top)
//			if record[x][y] {
//				match = false
//				break
//			}
//			record[x][y] = true
//		}
//		if match {
//			notRepeat = append(notRepeat, points)
//		}
//	}
//
//	return notRepeat
//}
//
//// SearchContainCross 在一组二维坐标中查找是否存在交叉（十字）线
//func SearchContainCross(xys ...[2]int) bool {
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	for x := 0; x < len(rectangleShape); x++ {
//		for y := 0; y < len(rectangleShape[0]); y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	for _, xy := range xys {
//		var points [][2]int
//		var find = map[int]bool{}
//		x, y := CoordinateArrayToCoordinate(xy)
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[1] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		if !find[1] {
//			continue
//		}
//		for sx := x + 1; sx < len(rectangleShape); sx++ {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[2] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		if !find[2] {
//			continue
//		}
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[3] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if !find[3] {
//			continue
//		}
//		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[4] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if !find[4] {
//			continue
//		}
//		return true
//	}
//
//	return false
//}
//
//// SearchNotRepeatStraightLine 在一组二维坐标中从大到小搜索不重复的直线
////   - 最低需要长度为3
//func SearchNotRepeatStraightLine(minLength int, xys ...[2]int) (result [][][2]int) {
//	if minLength < 3 {
//		return nil
//	}
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	for x := 0; x < len(rectangleShape); x++ {
//		for y := 0; y < len(rectangleShape[0]); y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	for _, xy := range xys {
//		var points [][2]int
//		var find = map[int]bool{}
//		x, y := CoordinateArrayToCoordinate(xy)
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[1] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		for sx := x + 1; sx < len(rectangleShape); sx++ {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[2] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		if len(find) == 0 {
//			points = nil
//		} else if len(points) >= minLength-1 {
//			goto end
//		} else {
//			points = nil
//		}
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[3] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[4] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if !find[3] && !find[4] {
//			continue
//		}
//	end:
//		{
//			if len(points) < minLength-1 {
//				continue
//			}
//			result = append(result, append(points, [2]int{x + left, y + top}))
//		}
//	}
//
//	sort.Slice(result, func(i, j int) bool {
//		return len(result[i]) > len(result[j])
//	})
//
//	var notRepeat [][][2]int
//	for _, points := range result {
//		var match = true
//		for _, point := range points {
//			x, y := CoordinateArrayToCoordinate(point)
//			x = x + (0 - left)
//			y = y + (0 - top)
//			if record[x][y] {
//				match = false
//				break
//			}
//			record[x][y] = true
//		}
//		if match {
//			notRepeat = append(notRepeat, points)
//		}
//	}
//
//	return notRepeat
//}
//
//// SearchContainStraightLine 在一组二维坐标中查找是否存在直线
//func SearchContainStraightLine(minLength int, xys ...[2]int) bool {
//	if minLength < 3 {
//		return false
//	}
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	for x := 0; x < len(rectangleShape); x++ {
//		for y := 0; y < len(rectangleShape[0]); y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	for _, xy := range xys {
//		var points [][2]int
//		var find = map[int]bool{}
//		x, y := CoordinateArrayToCoordinate(xy)
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[1] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		for sx := x + 1; sx < len(rectangleShape); sx++ {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[2] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		if len(find) == 0 {
//			points = nil
//		} else if len(points) >= minLength-1 {
//			goto end
//		} else {
//			points = nil
//		}
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[3] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[4] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if !find[3] && !find[4] {
//			continue
//		}
//	end:
//		{
//			if len(points) < minLength-1 {
//				continue
//			}
//			return true
//		}
//	}
//
//	return false
//}
//
//// SearchNotRepeatT 在一组二维坐标中从大到小搜索不重复T型（T）线
//func SearchNotRepeatT(minLength int, xys ...[2]int) (result [][][2]int) {
//	if minLength < 4 {
//		return nil
//	}
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	for x := 0; x < len(rectangleShape); x++ {
//		for y := 0; y < len(rectangleShape[0]); y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	for _, xy := range xys {
//		var points [][2]int
//		var find = map[int]bool{}
//		x, y := CoordinateArrayToCoordinate(xy)
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[1] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		for sx := x + 1; sx < len(rectangleShape); sx++ {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[2] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[3] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[4] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if len(find) != 3 || len(points) < minLength {
//			continue
//		}
//		result = append(result, append(points, [2]int{x + left, y + top}))
//	}
//
//	sort.Slice(result, func(i, j int) bool {
//		return len(result[i]) > len(result[j])
//	})
//
//	var notRepeat [][][2]int
//	for _, points := range result {
//		var match = true
//		for _, point := range points {
//			x, y := CoordinateArrayToCoordinate(point)
//			x = x + (0 - left)
//			y = y + (0 - top)
//			if record[x][y] {
//				match = false
//				break
//			}
//			record[x][y] = true
//		}
//		if match {
//			notRepeat = append(notRepeat, points)
//		}
//	}
//
//	return notRepeat
//}
//
//// SearchContainT 在一组二维坐标中查找是否存在T型（T）线
//func SearchContainT(minLength int, xys ...[2]int) bool {
//	if minLength < 4 {
//		return false
//	}
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	for x := 0; x < len(rectangleShape); x++ {
//		for y := 0; y < len(rectangleShape[0]); y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	for _, xy := range xys {
//		var points [][2]int
//		var find = map[int]bool{}
//		x, y := CoordinateArrayToCoordinate(xy)
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[1] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		for sx := x + 1; sx < len(rectangleShape); sx++ {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[2] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[3] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[4] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if len(find) != 3 || len(points) < minLength-1 {
//			continue
//		}
//		return true
//	}
//
//	return false
//}
//
//// SearchNotRepeatRightAngle 在一组二维坐标中从大到小搜索不重复的直角（L）线
//func SearchNotRepeatRightAngle(minLength int, xys ...[2]int) (result [][][2]int) {
//	if minLength < 3 {
//		return nil
//	}
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	for x := 0; x < len(rectangleShape); x++ {
//		for y := 0; y < len(rectangleShape[0]); y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	for _, xy := range xys {
//		var points [][2]int
//		var find = map[int]bool{}
//		x, y := CoordinateArrayToCoordinate(xy)
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[1] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		if find[1] {
//			goto up
//		}
//		for sx := x + 1; sx < len(rectangleShape); sx++ {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[2] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//	up:
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[3] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if find[3] {
//			goto end
//		}
//		// down
//		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[4] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if !find[4] {
//			continue
//		}
//	end:
//		{
//			if len(find) != 2 || len(points) < minLength-1 {
//				continue
//			}
//			result = append(result, append(points, [2]int{x + left, y + top}))
//		}
//	}
//
//	sort.Slice(result, func(i, j int) bool {
//		return len(result[i]) > len(result[j])
//	})
//
//	var notRepeat [][][2]int
//	for _, points := range result {
//		var match = true
//		for _, point := range points {
//			x, y := CoordinateArrayToCoordinate(point)
//			x = x + (0 - left)
//			y = y + (0 - top)
//			if record[x][y] {
//				match = false
//				break
//			}
//			record[x][y] = true
//		}
//		if match {
//			notRepeat = append(notRepeat, points)
//		}
//	}
//
//	return notRepeat
//}
//
//// SearchContainRightAngle 在一组二维坐标中查找是否存在直角（L）线
//func SearchContainRightAngle(minLength int, xys ...[2]int) bool {
//	if minLength < 3 {
//		return false
//	}
//	left, _, top, _ := GetShapeCoverageArea(xys...)
//	rectangleShape := GenerateShape(xys...)
//	record := map[int]map[int]bool{}
//	for x := 0; x < len(rectangleShape); x++ {
//		for y := 0; y < len(rectangleShape[0]); y++ {
//			record[x] = map[int]bool{}
//		}
//	}
//
//	for _, xy := range xys {
//		var points [][2]int
//		var find = map[int]bool{}
//		x, y := CoordinateArrayToCoordinate(xy)
//		x = x + (0 - left)
//		y = y + (0 - top)
//		// 搜索四个方向
//		for sx := x - 1; sx >= 0; sx-- {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[1] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//		if find[1] {
//			goto up
//		}
//		for sx := x + 1; sx < len(rectangleShape); sx++ {
//			if !rectangleShape[sx][y] {
//				break
//			}
//			find[2] = true
//			points = append(points, [2]int{sx + left, y + top})
//		}
//	up:
//		for sy := y - 1; sy >= 0; sy-- {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[3] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if find[3] {
//			goto end
//		}
//		// down
//		for sy := y + 1; sy < len(rectangleShape[0]); sy++ {
//			if !rectangleShape[x][sy] {
//				break
//			}
//			find[4] = true
//			points = append(points, [2]int{x + left, sy + top})
//		}
//		if !find[4] {
//			continue
//		}
//	end:
//		{
//			if len(find) != 2 || len(points) < minLength-1 {
//				continue
//			}
//			return true
//		}
//	}
//
//	return false
//}
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
