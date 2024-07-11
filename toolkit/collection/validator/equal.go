package validator

func Equal[T any](comparator func(a, b T) bool) Validator[T] {
	return &equal[T]{
		comparator: comparator,
	}
}

type equal[T any] struct {
	comparator func(a, b T) bool
}

func (e *equal[T]) Evaluate(entries []T) bool {
	for i := 1; i < len(entries); i++ {
		if !e.comparator(entries[i], entries[i-1]) {
			return false
		}
	}
	return true
}
