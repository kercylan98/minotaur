package collection

import "github.com/kercylan98/minotaur/toolkit/constraints"

// SliceSum 获取切片 slice 的总和
//   - 它在一些场景中可以很好的消除冗余的代码，例如排行榜分数为多个成员聚合时，不需要单独建立循环来获取总和
func SliceSum[S ~[]V, V any, R constraints.Number](slice S, handler func(i int, value V) R) R {
	var sum R
	for i, v := range slice {
		sum += handler(i, v)
	}
	return sum
}

// MapSum 获取 map 的总和
//   - 它在一些场景中可以很好的消除冗余的代码，例如排行榜分数为多个成员聚合时，不需要单独建立循环来获取总和
func MapSum[M ~map[K]V, K comparable, V any, R constraints.Number](m M, handler func(key K, value V) R) R {
	var sum R
	for k, v := range m {
		sum += handler(k, v)
	}
	return sum
}
