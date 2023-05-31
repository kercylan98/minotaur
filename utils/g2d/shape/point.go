package shape

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func NewPointWithArray(arr [2]int) Point {
	return Point{arr[0], arr[1]}
}

func NewPointWithArrays(arrays ...[2]int) []Point {
	var points = make([]Point, len(arrays), len(arrays))
	for i, arr := range arrays {
		points[i] = NewPointWithArray(arr)
	}
	return points
}

type Point [2]int

func (slf Point) GetX() int {
	return slf[0]
}

func (slf Point) GetY() int {
	return slf[1]
}

func (slf Point) GetXY() (int, int) {
	return slf[0], slf[1]
}
