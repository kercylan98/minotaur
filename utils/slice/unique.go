package slice

// Unique 返回去重后的切片
func Unique[T comparable](collection []T) []T {
	if len(collection) == 0 {
		return nil
	}

	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// UniqueBy 返回去重后的切片
func UniqueBy[T any](collection []T, fn func(T) any) []T {
	if len(collection) == 0 {
		return nil
	}

	result := make([]T, 0, len(collection))
	seen := make(map[any]struct{}, len(collection))

	for _, item := range collection {
		key := fn(item)
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}
