package super

import (
	"unsafe"
)

// StringToBytes 以零拷贝的方式将字符串转换为字节切片
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString 以零拷贝的方式将字节切片转换为字符串
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

// Convert 以零拷贝的方式将一个对象转换为另一个对象
//   - 两个对象字段必须完全一致
//   - 该函数可以绕过私有字段的访问限制
func Convert[A, B any](src A) B {
	return *(*B)(unsafe.Pointer(&src))
}
