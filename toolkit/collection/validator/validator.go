package validator

type Validator[T any] interface {
	Evaluate(entries []T) bool
}
