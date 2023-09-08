package slice

// Each 根据传入的 abort 遍历 slice，如果 iterator 返回值与 abort 相同，则停止遍历
func Each[V any](abort bool, slice []V, iterator func(index int, item V) bool) {
	for i := 0; i < len(slice); i++ {
		if iterator(i, slice[i]) == abort {
			break
		}
	}
}

// EachT 与 Each 的功能相同，但是 abort 被默认为 true
func EachT[V any](slice []V, iterator func(index int, item V) bool) {
	Each(true, slice, iterator)
}

// EachF 与 Each 的功能相同，但是 abort 被默认为 false
func EachF[V any](slice []V, iterator func(index int, item V) bool) {
	Each(false, slice, iterator)
}

// EachReverse 根据传入的 abort 从后往前遍历 slice，如果 iterator 返回值与 abort 相同，则停止遍历
func EachReverse[V any](abort bool, slice []V, iterator func(index int, item V) bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if iterator(i, slice[i]) == abort {
			break
		}
	}
}

// EachReverseT 与 EachReverse 的功能相同，但是 abort 被默认为 true
func EachReverseT[V any](slice []V, iterator func(index int, item V) bool) {
	EachReverse(true, slice, iterator)
}

// EachReverseF 与 EachReverse 的功能相同，但是 abort 被默认为 false
func EachReverseF[V any](slice []V, iterator func(index int, item V) bool) {
	EachReverse(false, slice, iterator)
}

// EachResult 根据传入的 abort 遍历 slice，得到遍历的结果，如果 iterator 返回值中的 bool 值与 abort 相同，则停止遍历，并返回当前已积累的结果
func EachResult[V any, R any](abort bool, slice []V, iterator func(index int, item V) (R, bool)) []R {
	var result []R
	for i := 0; i < len(slice); i++ {
		r, ok := iterator(i, slice[i])
		result = append(result, r)
		if ok == abort {
			break
		}
	}
	return result
}

// EachResultT 与 EachResult 的功能相同，但是 abort 被默认为 true
func EachResultT[V any, R any](slice []V, iterator func(index int, item V) (R, bool)) []R {
	return EachResult(true, slice, iterator)
}

// EachResultF 与 EachResult 的功能相同，但是 abort 被默认为 false
func EachResultF[V any, R any](slice []V, iterator func(index int, item V) (R, bool)) []R {
	return EachResult(false, slice, iterator)
}

// EachResultReverse 根据传入的 abort 从后往前遍历 slice，得到遍历的结果，如果 iterator 返回值中的 bool 值与 abort 相同，则停止遍历，并返回当前已积累的结果
func EachResultReverse[V any, R any](abort bool, slice []V, iterator func(index int, item V) (R, bool)) []R {
	var result []R
	for i := len(slice) - 1; i >= 0; i-- {
		r, ok := iterator(i, slice[i])
		result = append(result, r)
		if ok == abort {
			break
		}
	}
	return result
}

// EachResultReverseT 与 EachResultReverse 的功能相同，但是 abort 被默认为 true
func EachResultReverseT[V any, R any](slice []V, iterator func(index int, item V) (R, bool)) []R {
	return EachResultReverse(true, slice, iterator)
}

// EachResultReverseF 与 EachResultReverse 的功能相同，但是 abort 被默认为 false
func EachResultReverseF[V any, R any](slice []V, iterator func(index int, item V) (R, bool)) []R {
	return EachResultReverse(false, slice, iterator)
}
