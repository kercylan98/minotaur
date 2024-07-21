package convert

import (
	"strconv"
	"strings"
)

// StringToBytes 转换 string 为 []byte 类型
func StringToBytes(str string) []byte {
	return []byte(str)
}

// StringToRune 转换 string 为 []rune 类型
func StringToRune(str string) []rune {
	return []rune(str)
}

// StringToBoolean 转换 string 为 bool 类型
func StringToBoolean(str string) bool {
	result, _ := strconv.ParseBool(strings.ToLower(str))
	return result
}

// StringToFloat32 转换 string 为 float32 类型
func StringToFloat32(str string) float32 {
	result, _ := strconv.ParseFloat(str, 32)
	return float32(result)
}

// StringToFloat64 转换 string 为 float64 类型
func StringToFloat64(str string) float64 {
	result, _ := strconv.ParseFloat(str, 64)
	return result
}

// StringToInt 转换 string 为 int 类型
func StringToInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}

// StringToInt8 转换 string 为 int8 类型
func StringToInt8(str string) int8 {
	result, _ := strconv.ParseInt(str, 10, 8)
	return int8(result)
}

// StringToInt16 转换 string 为 int16 类型
func StringToInt16(str string) int16 {
	result, _ := strconv.ParseInt(str, 10, 16)
	return int16(result)
}

// StringToInt32 转换 string 为 int32 类型
func StringToInt32(str string) int32 {
	result, _ := strconv.ParseInt(str, 10, 32)
	return int32(result)
}

// StringToInt64 转换 string 为 int64 类型
func StringToInt64(str string) int64 {
	result, _ := strconv.ParseInt(str, 10, 64)
	return result
}

// StringToUint 转换 string 为 uint 类型
func StringToUint(str string) uint {
	result, _ := strconv.ParseUint(str, 10, 0)
	return uint(result)
}

// StringToUint8 转换 string 为 uint8 类型
func StringToUint8(str string) uint8 {
	result, _ := strconv.ParseUint(str, 10, 8)
	return uint8(result)
}

// StringToUint16 转换 string 为 uint16 类型
func StringToUint16(str string) uint16 {
	result, _ := strconv.ParseUint(str, 10, 16)
	return uint16(result)
}

// StringToUint32 转换 string 为 uint32 类型
func StringToUint32(str string) uint32 {
	result, _ := strconv.ParseUint(str, 10, 32)
	return uint32(result)
}

// StringToUint64 转换 string 为 uint64 类型
func StringToUint64(str string) uint64 {
	result, _ := strconv.ParseUint(str, 10, 64)
	return result
}
