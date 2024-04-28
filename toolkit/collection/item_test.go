package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"testing"
)

func TestSwapSlice(t *testing.T) {
	var cases = []struct {
		name   string
		slice  []int
		i      int
		j      int
		expect []int
	}{
		{"TestSwapSliceNonEmpty", []int{1, 2, 3}, 0, 1, []int{2, 1, 3}},
		{"TestSwapSliceEmpty", []int{}, 0, 0, []int{}},
		{"TestSwapSliceIndexOutOfBound", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
		{"TestSwapSliceNil", nil, 0, 0, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.SwapSlice(&c.slice, c.i, c.j)
			for i, v := range c.slice {
				if v != c.expect[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expect, c.slice, "the slice is not equal")
				}
			}
		})
	}
}
