package slice

import "math/rand"

// Shuffle 随机打乱切片
//   - 该函数会改变传入的切片，如果不希望改变原切片，需要在函数调用之前手动复制一份或者使用 ShuffleCopy 函数
func Shuffle[T any](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}

// ShuffleCopy 返回随机打乱后的切片
//   - 该函数不会改变原切片
func ShuffleCopy[T any](collection []T) []T {
	if len(collection) == 0 {
		return nil
	}

	result := make([]T, len(collection))
	perm := rand.Perm(len(collection))

	for i, randIndex := range perm {
		result[i] = collection[randIndex]
	}

	return result
}
