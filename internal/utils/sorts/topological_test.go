package sorts_test

import (
	"github.com/kercylan98/minotaur/utils/sorts"
	"testing"
)

func TestTopological(t *testing.T) {
	type Item struct {
		ID      int
		Depends []int
	}

	var items = []Item{
		{ID: 2, Depends: []int{4}},
		{ID: 1, Depends: []int{2, 3}},
		{ID: 3, Depends: []int{4}},
		{ID: 4, Depends: []int{5}},
		{ID: 5, Depends: []int{}},
	}

	var sorted, err = sorts.Topological(items, func(item Item) int {
		return item.ID
	}, func(item Item) []int {
		return item.Depends
	})

	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range sorted {
		t.Log(item.ID, "|", item.Depends)
	}
}
