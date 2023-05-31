package shape

func NewShape() *Shape {
	shape := &Shape{
		maxX:   -1,
		maxY:   -1,
		points: map[int]map[int]Point{},
	}
	return shape
}

// Shape 2D形状定义
type Shape struct {
	maxX   int
	maxY   int
	points map[int]map[int]Point
}

func (slf *Shape) AddPoints(points ...Point) {
	for _, point := range points {
		slf.AddPoint(point)
	}
}

func (slf *Shape) AddPoint(point Point) {
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

func (slf *Shape) GetPoints() []Point {
	var points []Point
	for _, m := range slf.points {
		for _, point := range m {
			points = append(points, point)
		}
	}
	return points
}

func (slf *Shape) GetMaxX() int {
	return slf.maxX
}

func (slf *Shape) GetMaxY() int {
	return slf.maxY
}

func (slf *Shape) GetMaxXY() (int, int) {
	return slf.maxX, slf.maxY
}

func (slf *Shape) String() string {
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
