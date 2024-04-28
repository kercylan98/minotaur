package convert

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"strconv"
)

// BooleanToInt 将 bool 转换为 int 类型
func BooleanToInt[I constraints.Int](b bool) I {
	if b {
		return 1
	}
	return 0
}

// BooleanToString 将 bool 转换为 string 类型
func BooleanToString(b bool) string {
	return strconv.FormatBool(b)
}

// BooleanToByte 将 bool 转换为 byte 类型
func BooleanToByte(b bool) byte {
	return BooleanToInt[byte](b)
}

// BooleanToRune 将 bool 转换为 rune 类型
func BooleanToRune(b bool) rune {
	return BooleanToInt[rune](b)
}

// BooleanToFloat32 将 bool 转换为 float32 类型
func BooleanToFloat32(b bool) float32 {
	if b {
		return 1
	}
	return 0
}

// BooleanToFloat64 将 bool 转换为 float64 类型
func BooleanToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
