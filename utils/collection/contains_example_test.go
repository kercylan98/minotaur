package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

func ExampleInSlice() {
	result := collection.InSlice([]int{1, 2, 3}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleInComparableSlice() {
	result := collection.InComparableSlice([]int{1, 2, 3}, 2)
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllInSlice() {
	result := collection.AllInSlice([]int{1, 2, 3}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllInComparableSlice() {
	result := collection.AllInComparableSlice([]int{1, 2, 3}, []int{1, 2})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyInSlice() {
	result := collection.AnyInSlice([]int{1, 2, 3}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyInComparableSlice() {
	result := collection.AnyInComparableSlice([]int{1, 2, 3}, []int{1, 2})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleInSlices() {
	result := collection.InSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleInComparableSlices() {
	result := collection.InComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2)
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllInSlices() {
	result := collection.AllInSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllInComparableSlices() {
	result := collection.AllInComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyInSlices() {
	result := collection.AnyInSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyInComparableSlices() {
	result := collection.AnyInComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleInAllSlices() {
	result := collection.InAllSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// false
}

func ExampleInAllComparableSlices() {
	result := collection.InAllComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2)
	fmt.Println(result)
	// Output:
	// false
}

func ExampleAnyInAllSlices() {
	result := collection.AnyInAllSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// false
}

func ExampleAnyInAllComparableSlices() {
	result := collection.AnyInAllComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2})
	fmt.Println(result)
	// Output:
	// false
}

func ExampleKeyInMap() {
	result := collection.KeyInMap(map[string]int{"a": 1, "b": 2}, "a")
	fmt.Println(result)
	// Output:
	// true
}

func ExampleValueInMap() {
	result := collection.ValueInMap(map[string]int{"a": 1, "b": 2}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllKeyInMap() {
	result := collection.AllKeyInMap(map[string]int{"a": 1, "b": 2}, "a", "b")
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllValueInMap() {
	result := collection.AllValueInMap(map[string]int{"a": 1, "b": 2}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyKeyInMap() {
	result := collection.AnyKeyInMap(map[string]int{"a": 1, "b": 2}, "a", "b")
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyValueInMap() {
	result := collection.AnyValueInMap(map[string]int{"a": 1, "b": 2}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllKeyInMaps() {
	result := collection.AllKeyInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, "a", "b")
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAllValueInMaps() {
	result := collection.AllValueInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyKeyInMaps() {
	result := collection.AnyKeyInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, "a", "b")
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyValueInMaps() {
	result := collection.AnyValueInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleKeyInAllMaps() {
	result := collection.KeyInAllMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, "a")
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAnyKeyInAllMaps() {
	result := collection.AnyKeyInAllMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, []string{"a"})
	fmt.Println(result)
	// Output:
	// true
}
