package slice_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
)

func TestLimitedCombinations(t *testing.T) {
	c := slice.LimitedCombinations([]int{1, 2, 3, 4, 5}, 3, 3)
	for _, v := range c {
		fmt.Println(v)
	}
}
