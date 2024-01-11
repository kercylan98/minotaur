package collection

import "fmt"

func ExampleMappingFromSlice() {
	result := MappingFromSlice[[]int, []int]([]int{1, 2, 3}, func(value int) int {
		return value + 1
	})
	fmt.Println(result)
	// Output:
	// [2 3 4]
}

func ExampleMappingFromMap() {
	result := MappingFromMap[map[int]int, map[int]int](map[int]int{1: 1, 2: 2, 3: 3}, func(value int) int {
		return value + 1
	})
	fmt.Println(result)
	// Output:
	// map[1:2 2:3 3:4]
}
