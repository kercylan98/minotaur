package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

func ExampleFindLoopedNextInSlice() {
	next, v := collection.FindLoopedNextInSlice([]int{1, 2, 3}, 1)
	fmt.Println(next, v)
	// Output:
	// 2 3
}

func ExampleFindLoopedPrevInSlice() {
	prev, v := collection.FindLoopedPrevInSlice([]int{1, 2, 3}, 1)
	fmt.Println(prev, v)
	// Output:
	// 0 1
}

func ExampleFindCombinationsInSliceByRange() {
	result := collection.FindCombinationsInSliceByRange([]int{1, 2, 3}, 1, 2)
	fmt.Println(len(result))
	// Output:
	// 6
}

func ExampleFindFirstOrDefaultInSlice() {
	result := collection.FindFirstOrDefaultInSlice([]int{1, 2, 3}, 0)
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindOrDefaultInSlice() {
	result := collection.FindOrDefaultInSlice([]int{1, 2, 3}, 0, func(v int) bool {
		return v == 2
	})
	fmt.Println(result)
	// Output:
	// 2
}

func ExampleFindOrDefaultInComparableSlice() {
	result := collection.FindOrDefaultInComparableSlice([]int{1, 2, 3}, 2, 0)
	fmt.Println(result)
	// Output:
	// 2
}
