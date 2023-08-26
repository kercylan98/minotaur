package super

import "reflect"

// IsNumber 判断是否为数字
func IsNumber(v any) bool {
	kind := reflect.Indirect(reflect.ValueOf(v)).Kind()
	return kind >= reflect.Int && kind <= reflect.Float64
}
