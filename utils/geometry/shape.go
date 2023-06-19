package geometry

import (
	"bytes"
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/slice"
	"math"
	"sort"
	"strings"
)

var (
	ShapeStringHasBorder = false // 控制 Shape.String 是否拥有边界
)

// NewShape 通过多个点生成一个形状进行返回
func NewShape[V generic.SignedNumber](points ...Point[V]) Shape[V] {
	return points
}

// NewShapeWithString 通过字符串将指定 rune 转换为点位置生成形状进行返回
//   - 每个点的顺序从上到下，从左到右
func NewShapeWithString[V generic.SignedNumber](rows []string, point rune) (shape Shape[V]) {
	var width int
	for _, row := range rows {
		length := len(row)
		if length > width {
			width = length
		}
	}
	for y := 0; y < len(rows); y++ {
		runes := bytes.Runes([]byte(rows[y]))
		for x := 0; x < width; x++ {
			if x >= len(runes) {
				break
			}
			if point == runes[x] {
				shape = append(shape, NewPoint(V(x), V(y)))
			}
		}
	}

	return shape
}

// Shape 通过多个点表示了一个形状
type Shape[V generic.SignedNumber] []Point[V]

// Points 获取这个形状的所有点
func (slf Shape[V]) Points() []Point[V] {
	return slf
}

// PointCount 获取这个形状的点数量
func (slf Shape[V]) PointCount() int {
	return len(slf)
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
					if int(x) == int(p.GetX()) && int(y) == int(p.GetY()) {
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
			result = result[:len(result)-1]
			result += "\n"
		}
	} else {
		if left < 0 {
			left += -left
			right += -left
		}
		if top < 0 {
			top += -top
			bottom += -top
		}
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
			result = result[:len(result)-1]
			result += "\n"
		}
	}
	return fmt.Sprintf("%v\n%s", slf.Points(), strings.TrimSuffix(result, "\n"))
}

// ShapeSearch 获取该形状中包含的所有图形组合及其位置
//   - 需要注意的是，即便图形最终表示为相同的，但是只要位置组合顺序不同，那么也将被认定为一种图形组合
//   - [[1 0] [1 1] [1 2]] 和 [[1 1] [1 0] [1 2]] 可以被视为两个图形组合
//   - 返回的坐标为原始形状的坐标
//
// 可通过可选项对搜索结果进行过滤
func (slf Shape[V]) ShapeSearch(options ...ShapeSearchOption) (result []Shape[V]) {
	opt := newShapeSearchOptions()
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
	rectangleShape := GenerateShapeOnRectangleWithCoordinate(slf.Points()...)
	records := make(map[V]struct{})

	var match = func(links Shape[V], directionCount map[Direction]int, count int) bool {
		if opt.rectangle {
			return false
		}
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

	if opt.rectangle {
		l, r, t, b := GetShapeCoverageAreaWithCoordinateArray(slf.Points()...)
		rs := GenerateShapeOnRectangleWithCoordinate(slf.Points()...)
		w := r - l + 1
		h := b - t + 1
		shapes := GetExpressibleRectangleBySize(w, h, V(opt.rectangleMinWidth), V(opt.rectangleMinHeight))
		for _, s := range shapes {
			x, y := 0, 0
			for {
				if V(x)+s.GetX() >= w {
					x = 0
					y++
				}
				if V(y)+s.GetY() >= h {
					break
				}
				points := GetRectangleFullPoints(s[0]+1, s[1]+1)
				find := 0
				for _, point := range points {
					px, py := CoordinateArrayToCoordinate(point)
					ox, oy := px+V(x), py+V(y)
					if !rs[int(ox)][int(oy)] {
						find = 0
						break
					}
					find++
				}
				if find == len(points) {
					sw := s.GetX() + 1
					sh := s.GetY() + 1
					if !(sw < V(opt.rectangleMinWidth) || sw > V(opt.rectangleMaxWidth) || sh < V(opt.rectangleMinHeight) || sh > V(opt.rectangleMaxHeight)) {
						result = append(result, points)
					}
				}

				x++
			}
		}
	} else {
		for _, point := range slf.Points() {
			// 搜索四个方向
			var next = -1
			var directionPoint = point
			var links = Shape[V]{}
			var linkRecord = map[V]struct{}{}
			var directionCount = map[Direction]int{}
			var count = 0
			for i, directions := range [][]Direction{DirectionUDLR, DirectionLRUD} {
				var direction Direction
				for {
					next, direction = slice.NextLoop(directions, next)
					for {
						directionPoint = GetDirectionNextWithCoordinateArray(direction, directionPoint)
						if px, py := directionPoint.GetXY(); px < 0 || px >= areaWidth || py < 0 || py >= areaHeight {
							break
						}
						offset := directionPoint.GetOffset(-left, -top)
						if offset.OutOf(V(0), V(0), width, height) || !rectangleShape[int(offset.GetX())][int(offset.GetY())] {
							break
						}
						recordPos := directionPoint.GetPos(areaWidth)
						if _, exist := linkRecord[recordPos]; !exist {
							linkRecord[recordPos] = struct{}{}
							links = append(links, directionPoint)
							directionCount[direction]++
							count++
							match(links, directionCount, count)
						}

						pos := directionPoint.GetPos(areaWidth)
						if _, exist := records[pos]; !exist && opt.directionCount < 1 {
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
