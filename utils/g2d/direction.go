package g2d

type Direction uint8

const (
	DirectionUp = Direction(iota)
	DirectionDown
	DirectionLeft
	DirectionRight
)

// CalcDirection 计算点2位于点1的方向
func CalcDirection(x1, y1, x2, y2 float64) Direction {
	angle := CalcAngle(x1, y1, x2, y2)
	if angle > -45 && angle < 45 {
		return DirectionRight
	} else if angle > 135 && angle < -135 {
		return DirectionLeft
	} else if angle > 45 && angle < 135 {
		return DirectionUp
	} else if angle > -135 && angle < -45 {
		return DirectionDown
	}
	return 0
}
