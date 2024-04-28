package generic

import "reflect"

// IsNil 检查指定的值是否为 nil
func IsNil[V any](v V) bool {
	return reflect.ValueOf(v).IsNil()
}

// IsAllNil 检查指定的值是否全部为 nil
func IsAllNil[V any](v ...V) bool {
	for _, v := range v {
		if !IsNil(v) {
			return false
		}
	}
	return true
}

// IsHasNil 检查指定的值是否存在 nil
func IsHasNil[V any](v ...V) bool {
	for _, v := range v {
		if IsNil(v) {
			return true
		}
	}
	return false
}
