package g2d

func NewDropper[T any](matrix [][]T, setPosition func(x, y, newX, newY int), allowDrop, allowCross, isObstacle func(data T) bool) *Dropper[T] {
	return &Dropper[T]{
		matrix:      matrix,
		width:       len(matrix) + 1,
		height:      len(matrix[0]),
		setPosition: setPosition,
		allowDrop:   allowDrop,
		allowCross:  allowCross,
		isObstacle:  isObstacle,
	}
}

// Dropper 掉落器
type Dropper[T any] struct {
	matrix                            [][]T
	width, height                     int
	setPosition                       func(x, y, newX, newY int)
	allowDrop, allowCross, isObstacle func(data T) bool
}

func (slf *Dropper[T]) Drop() {

	for x := 0; x < slf.width; x++ {
		for y := slf.height - 1; y >= 0; y-- {
			data := slf.matrix[x][y]
			if slf.allowDrop(data) {

			}
		}
	}
}
