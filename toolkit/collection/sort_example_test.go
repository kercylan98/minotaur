package collection_test

import (
	"fmt"
"github.com/kercylan98/minotaur/toolkit/collection"
"sort"
)

func ExampleDescBy() {
	var slice = []int{1, 2, 3}
	sort.Slice(slice, func(i, j int) bool {
		return collection.DescBy(slice[i], slice[j])
	})
	fmt.Println(slice)
	// Output:
	// [3 2 1]
}

func ExampleAscBy() {
	var slice = []int{1, 2, 3}
	sort.Slice(slice, func(i, j int) bool {
		return collection.AscBy(slice[i], slice[j])
	})
	fmt.Println(slice)
	// Output:
	// [1 2 3]
}

func ExampleDesc() {
	var slice = []int{1, 2, 3}
	collection.Desc(&slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(slice)
	// Output:
	// [3 2 1]
}

func ExampleDescByClone() {
	var slice = []int{1, 2, 3}
	result := collection.DescByClone(slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(result)
	// Output:
	// [3 2 1]
}

func ExampleAsc() {
	var slice = []int{1, 2, 3}
	collection.Asc(&slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(slice)
	// Output:
	// [1 2 3]
}

func ExampleAscByClone() {
	var slice = []int{1, 2, 3}
	result := collection.AscByClone(slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleShuffle() {
	var slice = []int{1, 2, 3}
	collection.Shuffle(&slice)
	fmt.Println(len(slice))
	// Output:
	// 3
}

func ExampleShuffleByClone() {
	var slice = []int{1, 2, 3}
	result := collection.ShuffleByClone(slice)
	fmt.Println(len(result))
	// Output:
	// 3
}
