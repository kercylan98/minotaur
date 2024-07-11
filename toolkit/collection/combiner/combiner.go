package combiner

type Combiner[T any] interface {
	Evaluate(entries []T) [][]T
}
