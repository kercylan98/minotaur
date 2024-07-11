package validator

func And[T any](validators ...Validator[T])  Validator[T]{
	return &and[T]{
		validators:validators,
	}
}

type and[T any] struct {
	validators []Validator[T]
}

func (a *and[T]) Evaluate(entries []T) bool {
	for _, validator := range a.validators {
		if !validator.Evaluate(entries) {
			return false
		}
	}
	return true
}

