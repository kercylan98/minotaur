package task

import "time"

// Cond 创建任务条件
func Cond(k, v any) Condition {
	return map[any]any{k: v}
}

// Condition 任务条件
type Condition map[any]any

// Cond 创建任务条件
func (slf Condition) Cond(k, v any) Condition {
	slf[k] = v
	return slf
}

// GetString 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetString(key any) string {
	v, _ := slf[key].(string)
	return v
}

// GetInt 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetInt(key any) int {
	v, _ := slf[key].(int)
	return v
}

// GetInt8 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetInt8(key any) int8 {
	v, _ := slf[key].(int8)
	return v
}

// GetInt16 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetInt16(key any) int16 {
	v, _ := slf[key].(int16)
	return v
}

// GetInt32 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetInt32(key any) int32 {
	v, _ := slf[key].(int32)
	return v
}

// GetInt64 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetInt64(key any) int64 {
	v, _ := slf[key].(int64)
	return v
}

// GetUint 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetUint(key any) uint {
	v, _ := slf[key].(uint)
	return v
}

// GetUint8 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetUint8(key any) uint8 {
	v, _ := slf[key].(uint8)
	return v
}

// GetUint16 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetUint16(key any) uint16 {
	v, _ := slf[key].(uint16)
	return v
}

// GetUint32 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetUint32(key any) uint32 {
	v, _ := slf[key].(uint32)
	return v
}

// GetUint64 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetUint64(key any) uint64 {
	v, _ := slf[key].(uint64)
	return v
}

// GetFloat32 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetFloat32(key any) float32 {
	v, _ := slf[key].(float32)
	return v
}

// GetFloat64 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetFloat64(key any) float64 {
	v, _ := slf[key].(float64)
	return v
}

// GetBool 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetBool(key any) bool {
	v, _ := slf[key].(bool)
	return v
}

// GetTime 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetTime(key any) time.Time {
	v, _ := slf[key].(time.Time)
	return v
}

// GetDuration 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetDuration(key any) time.Duration {
	v, _ := slf[key].(time.Duration)
	return v
}

// GetByte 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetByte(key any) byte {
	v, _ := slf[key].(byte)
	return v
}

// GetBytes 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetBytes(key any) []byte {
	v, _ := slf[key].([]byte)
	return v
}

// GetRune 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetRune(key any) rune {
	v, _ := slf[key].(rune)
	return v
}

// GetRunes 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetRunes(key any) []rune {
	v, _ := slf[key].([]rune)
	return v
}

// GetAny 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
func (slf Condition) GetAny(key any) any {
	return slf[key]
}
