package collection_test

import (
	"fmt"
"github.com/kercylan98/minotaur/toolkit/collection"
)

func ExampleMergeSlice() {
	fmt.Println(collection.MergeSlice(1, 2, 3))
	// Output:
	// [1 2 3]
}

func ExampleMergeSlices() {
	fmt.Println(
		collection.MergeSlices(
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		),
	)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleMergeMaps() {
	m := collection.MergeMaps(
		map[int]int{1: 1, 2: 2, 3: 3},
		map[int]int{4: 4, 5: 5, 6: 6},
	)
	fmt.Println(len(m))
	// Output:
	// 6
}

func ExampleMergeMapsWithSkip() {
	m := collection.MergeMapsWithSkip(
		map[int]int{1: 1},
		map[int]int{1: 2},
	)
	fmt.Println(m[1])
	// Output:
	// 1
}
