package super

import "reflect"

// Matcher 匹配器
type Matcher[Value, Result any] struct {
	value Value
	r     Result
	d     bool
}

// Match 匹配
func Match[Value, Result any](value Value) *Matcher[Value, Result] {
	return &Matcher[Value, Result]{
		value: value,
	}
}

// Case 匹配
func (slf *Matcher[Value, Result]) Case(value Value, result Result) *Matcher[Value, Result] {
	if !slf.d && reflect.DeepEqual(slf.value, value) {
		slf.r = result
		slf.d = true
	}
	return slf
}

// Default 默认
func (slf *Matcher[Value, Result]) Default(value Result) Result {
	if slf.d {
		return slf.r
	}
	return value
}
