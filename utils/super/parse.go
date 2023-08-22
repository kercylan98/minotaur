package super

import "strconv"

// StringToInt 字符串转换为整数
func StringToInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

// StringToFloat64 字符串转换为 float64
func StringToFloat64(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}
