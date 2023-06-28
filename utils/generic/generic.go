package generic

import "reflect"

// IsNil 检查指定的值是否为 nil
func IsNil[V any](v V) bool {
	return reflect.ValueOf(v).IsNil()
}
