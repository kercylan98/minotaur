package combiner

import (
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/collection/validator"
)

func Continuous[T any](minCount, maxCount int, valueGetter func(t T) int, valueSort ...int) Combiner[T] {
	return &continuous[T]{
		minCount:    minCount,
		maxCount:    maxCount,
		valueGetter: valueGetter,
		valueSort:   valueSort,
	}
}

type continuous[T any] struct {
	minCount    int
	maxCount    int
	valueGetter func(t T) int
	valueSort   []int
}

func (s *continuous[T]) Evaluate(entries []T) [][]T {
	cs := collection.FindCombinationsInSliceByRange(entries, s.minCount, s.maxCount)
	result := make([][]T, 0)
	for _, c := range cs {
		if validator.Continuous(s.minCount, s.maxCount, s.valueGetter, s.valueSort...).Evaluate(c) {
			result = append(result, c)
		}
	}

	return result
}
