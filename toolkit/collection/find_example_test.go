package collection_test

import (
	"fmt"
"github.com/kercylan98/minotaur/toolkit/collection"
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

func ExampleFindInSlice() {
	_, result := collection.FindInSlice([]int{1, 2, 3}, func(v int) bool {
		return v == 2
	})
	fmt.Println(result)
	// Output:
	// 2
}

func ExampleFindIndexInSlice() {
	result := collection.FindIndexInSlice([]int{1, 2, 3}, func(v int) bool {
		return v == 2
	})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindInComparableSlice() {
	index, result := collection.FindInComparableSlice([]int{1, 2, 3}, 2)
	fmt.Println(index, result)
	// Output:
	// 1 2
}

func ExampleFindIndexInComparableSlice() {
	result := collection.FindIndexInComparableSlice([]int{1, 2, 3}, 2)
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindMinimumInComparableSlice() {
	result := collection.FindMinimumInComparableSlice([]int{1, 2, 3})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindMinimumInSlice() {
	result := collection.FindMinimumInSlice([]int{1, 2, 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindMaximumInComparableSlice() {
	result := collection.FindMaximumInComparableSlice([]int{1, 2, 3})
	fmt.Println(result)
	// Output:
	// 3
}

func ExampleFindMaximumInSlice() {
	result := collection.FindMaximumInSlice([]int{1, 2, 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
	// Output:
	// 3
}

func ExampleFindMin2MaxInComparableSlice() {
	minimum, maximum := collection.FindMin2MaxInComparableSlice([]int{1, 2, 3})
	fmt.Println(minimum, maximum)
	// Output:
	// 1 3
}

func ExampleFindMin2MaxInSlice() {
	minimum, maximum := collection.FindMin2MaxInSlice([]int{1, 2, 3}, func(v int) int {
		return v
	})
	fmt.Println(minimum, maximum)
	// Output:
	// 1 3
}

func ExampleFindMinFromComparableMap() {
	result := collection.FindMinFromComparableMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindMinFromMap() {
	result := collection.FindMinFromMap(map[int]int{1: 1, 2: 2, 3: 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindMaxFromComparableMap() {
	result := collection.FindMaxFromComparableMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(result)
	// Output:
	// 3
}

func ExampleFindMaxFromMap() {
	result := collection.FindMaxFromMap(map[int]int{1: 1, 2: 2, 3: 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
	// Output:
	// 3
}

func ExampleFindMin2MaxFromComparableMap() {
	minimum, maximum := collection.FindMin2MaxFromComparableMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(minimum, maximum)
	// Output:
	// 1 3
}

func ExampleFindMin2MaxFromMap() {
	minimum, maximum := collection.FindMin2MaxFromMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(minimum, maximum)
	// Output:
	// 1 3
}
