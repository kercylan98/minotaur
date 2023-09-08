package slice_test

import (
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
)

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s = slice.Filter(true, s, func(index int, item int) bool {
		return item%2 == 0
	})
	if len(s) != 2 {
		t.Error("Filter failed")
	}
	if s[0] != 2 {
		t.Error("Filter failed")
	}
	if s[1] != 4 {
		t.Error("Filter failed")
	}
}

func TestFilterCopy(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	cp := slice.FilterCopy(true, s, func(index int, item int) bool {
		return item%2 == 0
	})
	if len(s) != 5 {
		t.Error("FilterCopy failed")
	} else {
		t.Log(s, cp)
	}
}
