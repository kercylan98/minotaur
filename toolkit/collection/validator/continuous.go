package validator

import (
	"sort"
)

// Continuous 校验是否为连续的值，如果传入了 valueSort，则会按照 valueSort 的顺序进行排序，否则按照值的大小进行排序
func Continuous[T any](minCount, maxCount int, valueGetter func(t T) int, valueSort ...int) Validator[T] {
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

func (s *continuous[T]) Evaluate(entries []T) bool {
	if len(entries) < s.minCount || len(entries) > s.maxCount {
		return false
	}

	sortValues := make(map[int]int)
	for i, s := range s.valueSort {
		sortValues[s] = i
	}

	values := make([]int, len(entries))
	for i, entry := range entries {
		value := s.valueGetter(entry)
		values[i] = value
		if _, exist := sortValues[value]; !exist {
			sortValues[value] = value
		}
	}

	sort.Slice(values, func(i, j int) bool {
		return sortValues[values[i]] < sortValues[values[j]]
	})

	for i := 1; i < len(values); i++ {
		if values[i] != values[i-1]+1 {
			return false
		}
	}

	return true
}
