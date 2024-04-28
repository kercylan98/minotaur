package reflects

import (
	"fmt"
	"reflect"
)

// WrapperFunc 包装函数
func WrapperFunc[Func any](f any, wrapper func(call func([]reflect.Value) []reflect.Value) func(args []reflect.Value) []reflect.Value) (wf Func, err error) {
	tof := reflect.TypeOf(f)
	if tof.Kind() != reflect.Func {
		return wf, fmt.Errorf("f is not a function, got %v", tof.String())
	}
	return reflect.MakeFunc(tof, wrapper(func(args []reflect.Value) []reflect.Value {
		return reflect.ValueOf(f).Call(args)
	})).Interface().(Func), nil
}

// WrapperFuncBefore2After 包装函数，前置函数执行前，后置函数执行后
func WrapperFuncBefore2After[Func any](f Func, before, after func()) (wf Func, err error) {
	vof := reflect.ValueOf(f)
	tof := vof.Type()
	if tof.Kind() != reflect.Func {
		return wf, fmt.Errorf("f is not a function, got %v", tof.String())
	}
	wrapped := reflect.MakeFunc(tof, func(args []reflect.Value) []reflect.Value {
		if before != nil {
			before()
		}
		result := vof.Call(args)
		if after != nil {
			after()
		}
		return result
	})

	return wrapped.Interface().(Func), nil
}

// WrapperFuncBefore 包装函数，前置函数执行前
func WrapperFuncBefore[Func any](f Func, before func()) (wf Func, err error) {
	return WrapperFuncBefore2After(f, before, nil)
}

// WrapperFuncAfter 包装函数，后置函数执行后
func WrapperFuncAfter[Func any](f Func, after func()) (wf Func, err error) {
	return WrapperFuncBefore2After(f, nil, after)
}
