package g2d

// SlideDrop 侧滑掉落特定位置成员
func SlideDrop[T any](matrix [][]T, x, y int, isStop func(start T, x, y int, data T) bool) (dropX, dropY int) {
	width, height := len(matrix), len(matrix[0])
	start := matrix[x][y]
	var offsetX, offsetY = -1, 1
	for {
		targetX, targetY := x+offsetX, y+offsetY
		if targetX < 0 || targetY == height || isStop(start, targetX, targetY, matrix[targetX][targetY]) {
			dropX, dropY = targetX+1, targetY-1
			break
		}
		offsetX--
		offsetY++
	}
	offsetX, offsetY = 1, 1
	for dropX != x && dropY != y {
		targetX, targetY := x+offsetX, y+offsetY
		if targetX == width || targetY == height || isStop(start, targetX, targetY, matrix[targetX][targetY]) {
			dropX, dropY = targetX-1, targetY-1
			break
		}
		offsetX++
		offsetY++
	}
	return
}

// VerticalDrop 垂直掉落特定位置成员
func VerticalDrop[T any](matrix [][]T, x, y int, isStop func(start T, x, y int, data T) bool) (dropX, dropY int) {
	height := len(matrix[0])
	start := matrix[x][y]
	var offsetY = 1
	for {
		testY := y + offsetY
		if testY == height || isStop(start, x, testY, matrix[x][testY]) {
			return x, testY - 1
		}
		offsetY++
	}
}
