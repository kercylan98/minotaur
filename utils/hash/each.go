package hash

// Each 根据传入的 abort 遍历 m，如果 iterator 返回值与 abort 相同，则停止遍历
func Each[K comparable, V any](abort bool, m map[K]V, iterator func(i int, key K, item V) bool) {
	i := 0
	for k, v := range m {
		if iterator(i, k, v) == abort {
			break
		}
		i++
	}
}

// EachT 与 Each 的功能相同，但是 abort 被默认为 true
func EachT[K comparable, V any](m map[K]V, iterator func(i int, key K, item V) bool) {
	Each(true, m, iterator)
}

// EachF 与 Each 的功能相同，但是 abort 被默认为 false
func EachF[K comparable, V any](m map[K]V, iterator func(i int, key K, item V) bool) {
	Each(false, m, iterator)
}

// EachResult 根据传入的 abort 遍历 m，得到遍历的结果，如果 iterator 返回值中的 bool 值与 abort 相同，则停止遍历，并返回当前已积累的结果
func EachResult[K comparable, V any, R any](abort bool, m map[K]V, iterator func(i int, key K, item V) (R, bool)) []R {
	var result []R
	i := 0
	for k, v := range m {
		r, ok := iterator(i, k, v)
		result = append(result, r)
		if ok == abort {
			break
		}
		i++
	}
	return result
}

// EachResultT 与 EachResult 的功能相同，但是 abort 被默认为 true
func EachResultT[K comparable, V any, R any](m map[K]V, iterator func(i int, key K, item V) (R, bool)) []R {
	return EachResult(true, m, iterator)
}

// EachResultF 与 EachResult 的功能相同，但是 abort 被默认为 false
func EachResultF[K comparable, V any, R any](m map[K]V, iterator func(i int, key K, item V) (R, bool)) []R {
	return EachResult(false, m, iterator)
}
