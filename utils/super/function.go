package super

// Handle 执行 f 函数，如果 f 为 nil，则不执行
func Handle(f func()) {
	if f != nil {
		f()
	}
}

// HandleErr 执行 f 函数，如果 f 为 nil，则不执行
func HandleErr(f func() error) error {
	if f != nil {
		return f()
	}
	return nil
}

// HandleV 执行 f 函数，如果 f 为 nil，则不执行
func HandleV[V any](v V, f func(v V)) {
	if f != nil {
		f(v)
	}
}

// SafeExec 安全的执行函数，当 f 为 nil 时，不执行，当 f 执行出错时，不会 panic，而是转化为 error 进行返回
func SafeExec(f func()) (err error) {
	if f == nil {
		return
	}
	defer func() {
		err = RecoverTransform(recover())
	}()
	f()
	return
}
