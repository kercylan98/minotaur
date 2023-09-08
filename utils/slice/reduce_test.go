package slice_test

import (
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
)

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	sum := slice.Reduce(0, s, func(index int, item int, current int) int {
		return current + item
	})
	t.Log(sum)
	if sum != 15 {
		t.Error("Reduce failed")
	}
}
