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
