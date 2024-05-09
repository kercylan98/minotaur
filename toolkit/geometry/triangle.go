package geometry

// CalcTriangleAreaTwice 计算三角形面积的两倍
func CalcTriangleAreaTwice(a, b, c Vector2) float64 {
	ax := b.GetX() - a.GetX()
	ay := b.GetY() - a.GetY()
	bx := c.GetX() - a.GetX()
	by := c.GetY() - a.GetY()
	return bx*ay - ax*by
}
