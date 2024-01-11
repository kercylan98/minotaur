package collection

import "fmt"

func ExampleDeduplicateSliceInPlace() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	DeduplicateSliceInPlace(&slice)
	fmt.Println(slice)
	// Output:
	// [1 2 3 4 5]
}

func ExampleDeduplicateSlice() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	fmt.Println(DeduplicateSlice(slice))
	// Output:
	// [1 2 3 4 5]
}

func ExampleDeduplicateSliceInPlaceWithCompare() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	DeduplicateSliceInPlaceWithCompare(&slice, func(a, b int) bool {
		return a == b
	})
	fmt.Println(slice)
	// Output:
	// [1 2 3 4 5]
}

func ExampleDeduplicateSliceWithCompare() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	fmt.Println(DeduplicateSliceWithCompare(slice, func(a, b int) bool {
		return a == b
	}))
	// Output:
	// [1 2 3 4 5]
}
