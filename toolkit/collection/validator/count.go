package validator

func Count[T any](n int) Validator[T] {
	return &count[T]{count: n}
}

type count[T any] struct {
	count int
}

func (c *count[T]) Evaluate(entries []T) bool {
	return len(entries) == c.count
}
