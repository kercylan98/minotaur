package toolkit

// NotNil 返回一个非 nil 的值，如果 a 和 b 都为 nil，则抛出异常
func NotNil(a, b any) any {
	if a != nil && b != nil {
		panic("a and b are not nil")
	}
	if a != nil {
		return a
	}
	return b
}
