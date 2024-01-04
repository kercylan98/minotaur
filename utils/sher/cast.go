package sher

// SliceCastToAny 将切片转换为任意类型的切片
func SliceCastToAny[S ~[]V, V any](s S) []any {
	var r = make([]any, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}
