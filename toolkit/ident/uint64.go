package ident

import "unsafe"

// GenerateOrderedUniqueIdentStringWithUInt64 从提供的 uint64 值生成一个唯一标识字符串
//   - 该函数生成的字符串是有序的，即使参数的顺序不同，生成的字符串也是相同的
func GenerateOrderedUniqueIdentStringWithUInt64(is ...uint64) string {
	switch len(is) {
	case 0:
		return ""
	case 1:
		v := is[0]
		b := make([]byte, 8)
		for i := 0; i < 8; i++ {
			b[i] = byte(v >> (8 * uint(i)))
		}
		return string(b)
	case 2:
		v1 := is[0]
		v2 := is[1]
		b := make([]byte, 16)
		for i := 0; i < 8; i++ {
			b[i] = byte(v1 >> (8 * uint(i)))
			b[i+8] = byte(v2 >> (8 * uint(i)))
		}
		return string(b)
	default:
		result := make([]byte, len(is)*8)
		ptr := unsafe.Pointer(&result[0])
		for _, v := range is {
			*(*uint64)(ptr) = v
			ptr = unsafe.Pointer(uintptr(ptr) + 8)
		}
		return *(*string)(unsafe.Pointer(&result))
	}
}
