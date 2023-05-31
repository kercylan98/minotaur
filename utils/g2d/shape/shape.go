package shape

func NewShape[Mark any](mark Mark, points ...Point) *Shape[Mark] {
	shape := &Shape[Mark]{
		maxX:   -1,
		maxY:   -1,
		points: map[int]map[int]Point{},
		mark:   mark,
	}
	shape.AddPoints(points...)
	return shape
}

// Shape 2D形状定义
type Shape[Mark any] struct {
	maxX   int
	maxY   int
	points map[int]map[int]Point
	mark   Mark
}

func (slf *Shape[Mark]) AddPoints(points ...Point) {
	for _, point := range points {
		slf.AddPoint(point)
	}
}

func (slf *Shape[Mark]) AddPoint(point Point) {
	x, y := point.GetXY()
	if x < 0 || y < 0 {
		panic("only positive integers are allowed for shape point positions")
	}
	if x > slf.maxX {
		slf.maxX = x
	}
	if y > slf.maxY {
		slf.maxY = y
	}
	ys, exist := slf.points[x]
	if !exist {
		ys = map[int]Point{}
		slf.points[x] = ys
	}
	ys[y] = point
}

func (slf *Shape[Mark]) GetPoints() []Point {
	var points []Point
	for _, m := range slf.points {
		for _, point := range m {
			points = append(points, point)
		}
	}
	return points
}

func (slf *Shape[Mark]) GetMaxX() int {
	return slf.maxX
}

func (slf *Shape[Mark]) GetMaxY() int {
	return slf.maxY
}

func (slf *Shape[Mark]) GetMaxXY() (int, int) {
	return slf.maxX, slf.maxY
}

func (slf *Shape[Mark]) GetMark() Mark {
	return slf.mark
}

func (slf *Shape[Mark]) String() string {
	var str string
	for y := 0; y <= slf.maxY; y++ {
		for x := 0; x <= slf.maxX; x++ {
			if _, exist := slf.points[x][y]; exist {
				str += "1"
			} else {
				str += "0"
			}
		}
		str += "\r\n"
	}
	return str
}
