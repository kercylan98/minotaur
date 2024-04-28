package collection

import "fmt"

func ExampleClearSlice() {
	slice := []int{1, 2, 3, 4, 5}
	ClearSlice(&slice)
	fmt.Println(slice)
	// Output:
	// []
}

func ExampleClearMap() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	ClearMap(m)
	fmt.Println(m)
	// Output:
	// map[]
}

func ExampleDropSliceByIndices() {
	slice := []int{1, 2, 3, 4, 5}
	DropSliceByIndices(&slice, 1, 3)
	fmt.Println(slice)
	// Output:
	// [1 3 5]
}

func ExampleDropSliceByCondition() {
	slice := []int{1, 2, 3, 4, 5}
	DropSliceByCondition(&slice, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(slice)
	// Output:
	// [1 3 5]
}

func ExampleDropSliceOverlappingElements() {
	slice := []int{1, 2, 3, 4, 5}
	DropSliceOverlappingElements(&slice, []int{1, 3, 5}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(slice)
	// Output:
	// [2 4]
}
