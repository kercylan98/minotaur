package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/collection"
)

func ExampleTopologicalSort() {
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

	var sorted, err = collection.TopologicalSort(items, func(item Item) int {
		return item.ID
	}, func(item Item) []int {
		return item.Depends
	})

	if err != nil {
		return
	}

	for _, item := range sorted {
		fmt.Println(item.ID, "|", item.Depends)
	}
	// Output:
	// 1 | [2 3]
	// 2 | [4]
	// 3 | [4]
	// 4 | [5]
	// 5 | []
}
