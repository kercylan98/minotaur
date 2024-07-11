package validator

// Pairs 校验所有值是否都是成对的，pair 表示每个值期望的对数
func Pairs[T any](pair int, valueGetter func(T) int) Validator[T] {
	return &pairs[T]{
		pair:        pair,
		valueGetter: valueGetter,
	}
}

type pairs[T any] struct {
	pair        int
	valueGetter func(T) int
}

func (p *pairs[T]) Evaluate(entries []T) bool {
	if len(entries)%p.pair != 0 {
		return false
	}

	var valueMap = make(map[int]int)
	for _, entry := range entries {
		value := p.valueGetter(entry)
		valueMap[value]++
		if valueMap[value] > p.pair {
			return false
		}
	}

	for _, i := range valueMap {
		if i != p.pair {
			return false
		}
	}

	return true
}
