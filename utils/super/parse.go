package super

import "strconv"

// StringToInt 字符串转换为整数
func StringToInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}
