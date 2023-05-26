package synchronization

type SliceOption[T any] func(slice *Slice[T])

func WithSliceLen[T any](len int) SliceOption[T] {
	return func(slice *Slice[T]) {
		slice.data = make([]T, len)
	}
}

func WithSliceCap[T any](cap int) SliceOption[T] {
	return func(slice *Slice[T]) {
		slice.data = make([]T, 0, cap)
	}
}

func WithSliceLenCap[T any](len, cap int) SliceOption[T] {
	return func(slice *Slice[T]) {
		slice.data = make([]T, len, cap)
	}
}
