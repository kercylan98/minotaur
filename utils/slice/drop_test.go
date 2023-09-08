package slice_test

import (
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
)

func TestDrop(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	t.Log(s, slice.Drop(1, 3, s))
}
