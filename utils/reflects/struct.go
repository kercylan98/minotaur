package reflects

import (
	"fmt"
	"reflect"
	"unsafe"
)

// DeepCopy 深拷贝
func DeepCopy[T any](src T) T {
	vof := reflect.Indirect(reflect.ValueOf(src))
	tof := vof.Type()
	fmt.Println(tof)

	return src
}

// GetPtrUnExportFiled 获取指针类型的未导出字段
func GetPtrUnExportFiled(s reflect.Value, filedIndex int) reflect.Value {
	v := s.Elem().Field(filedIndex)
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

// SetPtrUnExportFiled 设置指针类型的未导出字段
func SetPtrUnExportFiled(s reflect.Value, filedIndex int, val reflect.Value) {
	v := GetPtrUnExportFiled(s, filedIndex)
	v.Set(val)
}

// Copy 拷贝
func Copy(s reflect.Value) reflect.Value {
	return reflect.NewAt(s.Type(), unsafe.Pointer(s.UnsafeAddr())).Elem()
}

// GetPointer 获取指针
func GetPointer[T any](src T) reflect.Value {
	s := reflect.TypeOf(src)
	if s.Kind() == reflect.Pointer {
		return reflect.ValueOf(src)
	}
	return reflect.ValueOf(&src)
}
