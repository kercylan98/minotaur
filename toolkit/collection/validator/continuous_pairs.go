package validator

import "sort"

// ContinuousPairs 创建一个验证器，用于检查给定的切片是否包含指定数量的连续对。
// -  minCount 和 maxCount 指定了连续对的最小和最大数量。
// -  pair 指定了需要检查的连续对的数量。
// -  valueGetter 是一个函数，用于从类型 T 的元素中获取用于比较的整数值。
// -  valueSort 是一个可选参数，用于指定值的排序规则。如果不提供，则按照值的自然顺序进行排序。
func ContinuousPairs[T any](minCount, maxCount, pair int, valueGetter func(t T) int, valueSort ...int) Validator[T] {
	return &continuousPairs[T]{
		pair:        pair,
		minCount:    minCount,
		maxCount:    maxCount,
		valueGetter: valueGetter,
		valueSort:   valueSort,
	}
}

type continuousPairs[T any] struct {
	minCount    int
	maxCount    int
	valueGetter func(t T) int
	valueSort   []int
	pair        int
}

func (p *continuousPairs[T]) Evaluate(entries []T) bool {
	// 检查条目的数量是否满足最小和最大成对数量的要求
	if len(entries) < p.minCount*p.pair || len(entries) > p.maxCount*p.pair {
		return false
	}

	sortValues := make(map[int]int)
	for i, s := range p.valueSort {
		sortValues[s] = i
	}

	type same struct {
		value  int
		values []int
	}
	var sames = make(map[int]*same)

	for _, entry := range entries {
		value := p.valueGetter(entry)
		target, exists := sames[value]
		if !exists {
			target = &same{value: value}
			sames[value] = target
		}
		target.values = append(target.values, value)
		if _, exist := sortValues[value]; !exist {
			sortValues[value] = value
		}
	}

	samesSlice := make([]*same, 0, len(sames))
	for _, s := range sames {
		samesSlice = append(samesSlice, s)
	}

	sort.Slice(samesSlice, func(i, j int) bool {
		return sortValues[samesSlice[i].value] < sortValues[samesSlice[j].value]
	})

	for i := 1; i < len(samesSlice); i++ {
		if samesSlice[i].value != samesSlice[i-1].value+1 {
			return false
		}
	}

	return true
}
