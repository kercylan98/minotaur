package g2d

// PositionToArray 将坐标转换为x、y的数组
func PositionToArray(x, y int) [2]int {
	return [2]int{x, y}
}

// PositionArrayToXY 将坐标数组转换为x和y坐标
func PositionArrayToXY(position [2]int) (x, y int) {
	return position[0], position[1]
}

func PositionClone(position [2]int) [2]int {
	return [2]int{position[0], position[1]}
}
