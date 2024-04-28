package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

func ExampleMappingFromSlice() {
	result := collection.MappingFromSlice([]int{1, 2, 3}, func(value int) int {
		return value + 1
	})
	fmt.Println(result)
	// Output:
	// [2 3 4]
}

func ExampleMappingFromMap() {
	result := collection.MappingFromMap(map[int]int{1: 1, 2: 2, 3: 3}, func(value int) int {
		return value + 1
	})
	fmt.Println(result)
	// Output:
	// map[1:2 2:3 3:4]
}
