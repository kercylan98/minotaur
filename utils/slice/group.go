package slice

// GroupBy 返回分组后的切片
func GroupBy[T any, K comparable](collection []T, fn func(T) K) map[K][]T {
	if len(collection) == 0 {
		return nil
	}

	result := make(map[K][]T, len(collection))

	for _, item := range collection {
		key := fn(item)
		result[key] = append(result[key], item)
	}

	return result
}
