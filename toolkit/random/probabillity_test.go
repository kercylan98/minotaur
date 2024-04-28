package random_test

import (
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
)

func TestProbabilitySlice(t *testing.T) {
	var awards = []int{1, 2, 3, 4, 5, 6, 7}
	var probability = []float64{0.1, 2, 0.1, 0.1, 0.1, 0.1, 0.1}

	for i := 0; i < 50; i++ {
		t.Log(random.ProbabilitySlice(func(data int) float64 {
			return probability[data-1]
		}, awards...))
	}
}
