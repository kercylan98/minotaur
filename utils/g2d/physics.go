package g2d

// SlideDropXY 侧滑掉落特定位置成员
func SlideDropXY[T any](matrix [][]T, x, y int, isStop func(start T, x, y int, data T) bool) (dropX, dropY int) {
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

// SlideDropX 侧滑掉落一整列
//   - 返回每一个成员的y轴新坐标
func SlideDropX[T any](matrix [][]T, x int, isStop func(start T, x, y int, data T) bool) (result []int, change bool) {
	result = make([]int, len(matrix[x]))
	for y := len(matrix[x]) - 1; y >= 0; y-- {
		_, dropY := SlideDropXY(matrix, x, y, isStop)
		result[y] = dropY
		if y != dropY {
			change = true
		}
	}
	return
}

// SlideDropY 侧滑掉落一整行
//   - 返回每一个成员的x轴新坐标
func SlideDropY[T any](matrix [][]T, y int, isStop func(start T, x, y int, data T) bool) (result []int, change bool) {
	result = make([]int, len(matrix))
	for x := 0; x < len(matrix); x++ {
		dropX, _ := SlideDropXY(matrix, x, y, isStop)
		result[x] = dropX
		if x != dropX {
			change = true
		}
	}
	return
}

// SlideDrop 侧滑掉落整个矩阵
func SlideDrop[T any](matrix [][]T, isStop func(start T, x, y int, data T) bool) (result [][][2]int, change bool) {
	result = make([][][2]int, len(matrix))
	for x := 0; x < len(matrix); x++ {
		ys := make([][2]int, len(matrix[x]))
		var dropYs []int
		dropYs, change = SlideDropX(matrix, x, isStop)
		for y, positionY := range dropYs {
			ys[y] = PositionToArray(x, positionY)
		}
		result[x] = ys
	}
	return
}

// VerticalDropXY 垂直掉落特定位置成员
func VerticalDropXY[T any](matrix [][]T, x, y int, isStop func(start T, x, y int, data T) bool) (dropX, dropY int) {
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

// VerticalDropX 垂直掉落一整列
//   - 返回每一个成员的y轴新坐标
func VerticalDropX[T any](matrix [][]T, x int, isStop func(start T, x, y int, data T) bool) (result []int, change bool) {
	result = make([]int, len(matrix[x]))
	for y := len(matrix[x]) - 1; y >= 0; y-- {
		_, dropY := VerticalDropXY(matrix, x, y, isStop)
		result[y] = dropY
		if y != dropY {
			change = true
		}
	}
	return
}

// VerticalDropY 垂直掉落一整行
//   - 返回每一个成员的x轴新坐标
func VerticalDropY[T any](matrix [][]T, y int, isStop func(start T, x, y int, data T) bool) (result []int, change bool) {
	result = make([]int, len(matrix))
	for x := 0; x < len(matrix); x++ {
		dropX, _ := VerticalDropXY(matrix, x, y, isStop)
		result[x] = dropX
		if x != dropX {
			change = true
		}
	}
	return
}

// VerticalDrop 垂直掉落整个矩阵
func VerticalDrop[T any](matrix [][]T, isStop func(start T, x, y int, data T) bool) (result [][][2]int, change bool) {
	result = make([][][2]int, len(matrix))
	for x := 0; x < len(matrix); x++ {
		ys := make([][2]int, len(matrix[x]))
		var dropYs []int
		dropYs, change = VerticalDropX(matrix, x, isStop)
		for y, positionY := range dropYs {
			ys[y] = PositionToArray(x, positionY)
		}
		result[x] = ys
	}
	return
}
