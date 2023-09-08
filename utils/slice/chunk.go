package slice

// Chunk 返回分块后的切片
func Chunk[T any](collection []T, size int) [][]T {
	if len(collection) == 0 {
		return nil
	}

	if size < 1 {
		panic("size must be greater than 0")
	}

	result := make([][]T, 0, (len(collection)+size-1)/size)
	for size < len(collection) {
		collection, result = collection[size:], append(result, collection[0:size])
	}
	return append(result, collection)
}
