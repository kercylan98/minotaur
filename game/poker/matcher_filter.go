package poker

import "github.com/kercylan98/minotaur/utils/generic"

type MatcherFilter[P, C generic.Number, T Card[P, C]] struct {
	evaluate func([]T) int64
	handles  []func(cards []T) [][]T
	asc      bool
}

func (slf *MatcherFilter[P, C, T]) AddHandle(handle func(cards []T) [][]T) {
	slf.handles = append(slf.handles, handle)
}

func (slf *MatcherFilter[P, C, T]) group(cards []T) []T {
	var bestCombination = cards

	for _, handle := range slf.handles {
		var bestScore int64
		filteredCombinations := handle(bestCombination)
		if len(filteredCombinations) == 0 {
			return []T{}
		}
		for _, combination := range filteredCombinations {
			score := slf.evaluate(combination)
			if slf.asc {
				if score < bestScore || bestCombination == nil {
					bestCombination = combination
					bestScore = score
				}
			} else {
				if score > bestScore || bestCombination == nil {
					bestCombination = combination
					bestScore = score
				}
			}

		}
	}

	return bestCombination
}
