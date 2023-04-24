package g2d

import "math"

// CalcDistance 计算两点之间的距离
func CalcDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// CalcAngle 计算点2位于点1之间的角度
func CalcAngle(x1, y1, x2, y2 float64) float64 {
	return math.Atan2(y2-y1, x2-x1) * 180 / math.Pi
}

// CalculateNewCoordinate 根据给定的x、y坐标、角度和距离计算新的坐标
func CalculateNewCoordinate(x, y, angle, distance float64) (newX, newY float64) {
	// 将角度转换为弧度
	radians := angle * math.Pi / 180.0

	// 计算新的坐标
	newX = x + distance*math.Cos(radians)
	newY = y + distance*math.Sin(radians)

	return newX, newY
}
