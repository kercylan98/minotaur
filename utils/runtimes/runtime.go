package runtimes

import (
	"runtime"
)

// CurrentRunningFuncName 获取正在运行的函数名
func CurrentRunningFuncName(skip ...int) string {
	pc := make([]uintptr, 1)
	var s = 2
	if len(skip) > 0 {
		s += skip[0]
	}
	runtime.Callers(s, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
