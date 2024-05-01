package convert

import "unsafe"

// BytesToStringByZeroCopy 以零拷贝的方式将 []byte 转换为 string
func BytesToStringByZeroCopy(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
