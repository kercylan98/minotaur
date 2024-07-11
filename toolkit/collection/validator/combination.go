package validator

func Combination[T any](countCombination []int, getter func(T) int) Validator[T] {
	return &combination[T]{
		getter:           getter,
		countCombination: countCombination,
	}
}

type combination[T any] struct {
	getter           func(T) int
	countCombination []int
}

func (c *combination[T]) Evaluate(entries []T) bool {
	if len(entries) == 0 {
		return false
	}

	counts := make(map[int]int)
	for _, entry := range entries {
		counts[c.getter(entry)]++
	}

	reversion := make(map[int]struct{})
	for _, v := range counts {
		reversion[v] = struct{}{}
	}

	for _, c := range c.countCombination {
		if _, exists := reversion[c]; !exists {
			return false
		}
	}

	return true
}
