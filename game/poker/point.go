package poker

import "strconv"

const (
	PointA          Point = 1
	Point2          Point = 2
	Point3          Point = 3
	Point4          Point = 4
	Point5          Point = 5
	Point6          Point = 6
	Point7          Point = 7
	Point8          Point = 8
	Point9          Point = 9
	Point10         Point = 10
	PointJ          Point = 11
	PointQ          Point = 12
	PointK          Point = 13
	PointBlackJoker Point = 14
	PointRedJoker   Point = 15
)

var defaultPointSort = map[Point]int{
	PointA:          int(PointA),
	Point2:          int(Point2),
	Point3:          int(Point3),
	Point4:          int(Point4),
	Point5:          int(Point5),
	Point6:          int(Point6),
	Point7:          int(Point7),
	Point8:          int(Point8),
	Point9:          int(Point9),
	Point10:         int(Point10),
	PointJ:          int(PointJ),
	PointQ:          int(PointQ),
	PointK:          int(PointK),
	PointBlackJoker: int(PointBlackJoker),
	PointRedJoker:   int(PointRedJoker),
}

// Point 扑克点数
type Point int

func (slf Point) String() string {
	var str string
	switch slf {
	case PointA:
		str = "A"
	case PointJ:
		str = "J"
	case PointQ:
		str = "Q"
	case PointK:
		str = "K"
	case PointBlackJoker:
		str = "B"
	case PointRedJoker:
		str = "R"
	default:
		str = strconv.Itoa(int(slf))
	}
	return str
}
