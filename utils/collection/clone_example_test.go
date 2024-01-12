package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

func ExampleCloneSlice() {
	var slice = []int{1, 2, 3}
	var result = collection.CloneSlice(slice)
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleCloneMap() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMap(m)
	fmt.Println(len(result))
	// Output:
	// 3
}

func ExampleCloneSliceN() {
	var slice = []int{1, 2, 3}
	var result = collection.CloneSliceN(slice, 2)
	fmt.Println(len(result))
	// Output:
	// 2
}

func ExampleCloneMapN() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMapN(m, 2)
	fmt.Println(len(result))
	// Output:
	// 2
}

func ExampleCloneSlices() {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 3}
	var result = collection.CloneSlices(slice1, slice2)
	fmt.Println(len(result))
	// Output:
	// 2
}

func ExampleCloneMaps() {
	var m1 = map[int]int{1: 1, 2: 2, 3: 3}
	var m2 = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMaps(m1, m2)
	fmt.Println(len(result))
	// Output:
	// 2
}
