package slice

// FillBy 用指定的值填充切片
//   - slice: 待填充的切片
func FillBy[V any](slice []V, fn func(index int, value V) V) []V {
	for i, v := range slice {
		slice[i] = fn(i, v)
	}
	return slice
}

// FillByCopy 与 FillBy 的功能相同，但是不会改变原切片，而是返回一个新的切片
func FillByCopy[V any](slice []V, fn func(index int, value V) V) []V {
	var s = make([]V, len(slice))
	copy(s, slice)
	return FillBy(s, fn)
}

// FillUntil 用指定的值填充切片，如果 fn 返回的 bool 值与 abort 相同，则停止填充
//   - abort: 填充中止条件
//   - slice: 待填充的切片
func FillUntil[V any](abort bool, slice []V, fn func(index int, value V) (V, bool)) []V {
	for i, v := range slice {
		if value, b := fn(i, v); b == abort {
			break
		} else {
			slice[i] = value
		}
	}
	return slice
}

// FillUntilCopy 与 FillUntil 的功能相同，但不会改变原切片，而是返回一个新的切片
func FillUntilCopy[V any](abort bool, slice []V, fn func(index int, value V) (V, bool)) []V {
	var s = make([]V, len(slice))
	copy(s, slice)
	return FillUntil(abort, s, fn)
}

// FillUntilT 是 FillUntil 的简化版本，其中 abort 参数为 true
func FillUntilT[V any](slice []V, fn func(index int, value V) (V, bool)) []V {
	return FillUntil(true, slice, fn)
}

// FillUntilF 是 FillUntil 的简化版本，其中 abort 参数为 false
func FillUntilF[V any](slice []V, fn func(index int, value V) (V, bool)) []V {
	return FillUntil(false, slice, fn)
}

// FillUntilTCopy 是 FillUntilCopy 的简化版本，其中 abort 参数为 true
func FillUntilTCopy[V any](slice []V, fn func(index int, value V) (V, bool)) []V {
	return FillUntilCopy(true, slice, fn)
}

// FillUntilFCopy 是 FillUntilCopy 的简化版本，其中 abort 参数为 false
func FillUntilFCopy[V any](slice []V, fn func(index int, value V) (V, bool)) []V {
	return FillUntilCopy(false, slice, fn)
}
