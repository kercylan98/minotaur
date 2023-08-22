package super

import (
	"strconv"
	"strings"
)

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

// StringToBool 字符串转换为 bool
func StringToBool(value string) bool {
	result, _ := strconv.ParseBool(strings.ToLower(value))
	return result
}

// StringToUint64 字符串转换为 uint64
func StringToUint64(value string) uint64 {
	result, _ := strconv.ParseUint(value, 10, 64)
	return result
}

// StringToUint 字符串转换为 uint
func StringToUint(value string) uint {
	result, _ := strconv.ParseUint(value, 10, 64)
	return uint(result)
}

// StringToFloat32 字符串转换为 float32
func StringToFloat32(value string) float32 {
	result, _ := strconv.ParseFloat(value, 32)
	return float32(result)
}

// StringToInt64 字符串转换为 int64
func StringToInt64(value string) int64 {
	result, _ := strconv.ParseInt(value, 10, 64)
	return result
}

// StringToUint32 字符串转换为 uint32
func StringToUint32(value string) uint32 {
	result, _ := strconv.ParseUint(value, 10, 32)
	return uint32(result)
}

// StringToInt32 字符串转换为 int32
func StringToInt32(value string) int32 {
	result, _ := strconv.ParseInt(value, 10, 32)
	return int32(result)
}

// StringToUint16 字符串转换为 uint16
func StringToUint16(value string) uint16 {
	result, _ := strconv.ParseUint(value, 10, 16)
	return uint16(result)
}

// StringToInt16 字符串转换为 int16
func StringToInt16(value string) int16 {
	result, _ := strconv.ParseInt(value, 10, 16)
	return int16(result)
}

// StringToUint8 字符串转换为 uint8
func StringToUint8(value string) uint8 {
	result, _ := strconv.ParseUint(value, 10, 8)
	return uint8(result)
}

// StringToInt8 字符串转换为 int8
func StringToInt8(value string) int8 {
	result, _ := strconv.ParseInt(value, 10, 8)
	return int8(result)
}

// StringToByte 字符串转换为 byte
func StringToByte(value string) byte {
	result, _ := strconv.ParseUint(value, 10, 8)
	return byte(result)
}

// StringToRune 字符串转换为 rune
func StringToRune(value string) rune {
	result, _ := strconv.ParseInt(value, 10, 32)
	return rune(result)
}

// IntToString 整数转换为字符串
func IntToString(value int) string {
	return strconv.Itoa(value)
}

// Float64ToString float64 转换为字符串
func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// BoolToString bool 转换为字符串
func BoolToString(value bool) string {
	return strconv.FormatBool(value)
}

// Uint64ToString uint64 转换为字符串
func Uint64ToString(value uint64) string {
	return strconv.FormatUint(value, 10)
}

// UintToString uint 转换为字符串
func UintToString(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Float32ToString float32 转换为字符串
func Float32ToString(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', -1, 32)
}

// Int64ToString int64 转换为字符串
func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

// Uint32ToString uint32 转换为字符串
func Uint32ToString(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Int32ToString int32 转换为字符串
func Int32ToString(value int32) string {
	return strconv.FormatInt(int64(value), 10)
}

// Uint16ToString uint16 转换为字符串
func Uint16ToString(value uint16) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Int16ToString int16 转换为字符串
func Int16ToString(value int16) string {
	return strconv.FormatInt(int64(value), 10)
}

// Uint8ToString uint8 转换为字符串
func Uint8ToString(value uint8) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Int8ToString int8 转换为字符串
func Int8ToString(value int8) string {
	return strconv.FormatInt(int64(value), 10)
}

// ByteToString byte 转换为字符串
func ByteToString(value byte) string {
	return strconv.FormatUint(uint64(value), 10)
}

// RuneToString rune 转换为字符串
func RuneToString(value rune) string {
	return strconv.FormatInt(int64(value), 10)
}

// StringToSlice 字符串转换为切片
func StringToSlice(value string) []string {
	return strings.Split(value, "")
}

// SliceToString 切片转换为字符串
func SliceToString(value []string) string {
	return strings.Join(value, "")
}
