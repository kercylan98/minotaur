package geometry

// NewLine 创建一条直线
func NewLine(point Point, slope float64) Line {
	AssertPointValid(point)
	return Line{point, slope}
}

// Line 直线是由一个点和斜率定义的
type Line struct {
	point Point   // 直线上的一个点
	slope float64 // 斜率
}

// IsPointOn 判断点是否在直线上
func (l Line) IsPointOn(point Point) bool {
	AssertPointValid(l.point, point)
	return l.slope*(point[0]-l.point[0]) == point[1]-l.point[1]
}

// IsPointOnOrAbove 判断点是否在直线上或者在直线上方
func (l Line) IsPointOnOrAbove(point Point) bool {
	AssertPointValid(l.point, point)
	return l.slope*(point[0]-l.point[0]) >= point[1]-l.point[1]
}

// IsPointOnOrBelow 判断点是否在直线上或者在直线下方
func (l Line) IsPointOnOrBelow(point Point) bool {
	AssertPointValid(l.point, point)
	return l.slope*(point[0]-l.point[0]) <= point[1]-l.point[1]
}

// IsPointAbove 判断点是否在直线上方
func (l Line) IsPointAbove(point Point) bool {
	AssertPointValid(l.point, point)
	return l.slope*(point[0]-l.point[0]) > point[1]-l.point[1]
}

// IsPointBelow 判断点是否在直线下方
func (l Line) IsPointBelow(point Point) bool {
	AssertPointValid(l.point, point)
	return l.slope*(point[0]-l.point[0]) < point[1]-l.point[1]
}

// IsPointLeft 判断点是否在直线左侧
func (l Line) IsPointLeft(point Point) bool {
	AssertPointValid(l.point, point)
	return l.slope*(point[0]-l.point[0]) < point[1]-l.point[1]
}

// IsPointRight 判断点是否在直线右侧
func (l Line) IsPointRight(point Point) bool {
	AssertPointValid(l.point, point)
	return l.slope*(point[0]-l.point[0]) > point[1]-l.point[1]
}
