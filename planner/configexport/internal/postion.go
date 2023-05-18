package internal

func NewPosition(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

type Position struct {
	X int
	Y int
}
