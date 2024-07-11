package validator

func Or[T any](validators ...Validator[T]) Validator[T] {
	return &or[T]{
		validators: validators,
	}
}

type or[T any] struct {
	validators []Validator[T]
}

func (a *or[T]) Evaluate(entries []T) bool {
	for _, validator := range a.validators {
		if validator.Evaluate(entries) {
			return true
		}
	}
	return false
}
