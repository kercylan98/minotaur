package toolkit

// If 三元运算，如果 express 为真则返回 t，否则返回 f
//   - 该函数需要注意所有参数必须已经分配内存，否则极易发生 panic
func If[V any](express bool, t, f V) V {
	if express {
		return t
	}
	return f
}
