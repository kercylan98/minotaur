package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

func ExampleLoopSlice() {
	var result []int
	collection.LoopSlice([]int{1, 2, 3, 4, 5}, func(i int, val int) bool {
		result = append(result, val)
		if uint(i) == 1 {
			return false
		}
		return true
	})
	fmt.Println(result)
	// Output: [1 2]
}

func ExampleReverseLoopSlice() {
	var result []int
	collection.ReverseLoopSlice([]int{1, 2, 3, 4, 5}, func(i int, val int) bool {
		result = append(result, val)
		if uint(i) == 1 {
			return false
		}
		return true
	})
	fmt.Println(result)
	// Output: [5 4 3 2]
}

func ExampleLoopMap() {
	var result []int
	collection.LoopMap(map[string]int{"a": 1, "b": 2, "c": 3}, func(i int, key string, val int) bool {
		result = append(result, val)
		return true
	})
	fmt.Println(collection.AllInComparableSlice(result, []int{1, 2, 3}))
	// Output:
	// true
}

func ExampleLoopMapByOrderedKeyAsc() {
	var result []int
	collection.LoopMapByOrderedKeyAsc(map[string]int{"a": 1, "b": 2, "c": 3}, func(i int, key string, val int) bool {
		result = append(result, val)
		return true
	})
	fmt.Println(collection.AllInComparableSlice(result, []int{1, 2, 3}))
	// Output:
	// true
}

func ExampleLoopMapByOrderedKeyDesc() {
	var result []int
	collection.LoopMapByOrderedKeyDesc(map[string]int{"a": 1, "b": 2, "c": 3}, func(i int, key string, val int) bool {
		result = append(result, val)
		return true
	})
	fmt.Println(collection.AllInComparableSlice(result, []int{3, 2, 1}))
	// Output:
	// true
}

func ExampleLoopMapByOrderedValueAsc() {
	var result []int
	collection.LoopMapByOrderedValueAsc(map[string]int{"a": 1, "b": 2, "c": 3}, func(i int, key string, val int) bool {
		result = append(result, val)
		return true
	})
	fmt.Println(collection.AllInComparableSlice(result, []int{1, 2, 3}))
	// Output:
	// true
}

func ExampleLoopMapByOrderedValueDesc() {
	var result []int
	collection.LoopMapByOrderedValueDesc(map[string]int{"a": 1, "b": 2, "c": 3}, func(i int, key string, val int) bool {
		result = append(result, val)
		return true
	})
	fmt.Println(collection.AllInComparableSlice(result, []int{3, 2, 1}))
	// Output:
	// true
}

func ExampleLoopMapByKeyGetterAsc() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result []int
	collection.LoopMapByKeyGetterAsc(
		m,
		func(k string) int {
			return m[k]
		},
		func(i int, key string, val int) bool {
			result = append(result, val)
			return true
		},
	)
	fmt.Println(collection.AllInComparableSlice(result, []int{1, 2, 3}))
	// Output:
	// true
}

func ExampleLoopMapByKeyGetterDesc() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result []int
	collection.LoopMapByKeyGetterDesc(
		m,
		func(k string) int {
			return m[k]
		},
		func(i int, key string, val int) bool {
			result = append(result, val)
			return true
		},
	)
	fmt.Println(collection.AllInComparableSlice(result, []int{3, 2, 1}))
	// Output:
	// true
}

func ExampleLoopMapByValueGetterAsc() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result []int
	collection.LoopMapByValueGetterAsc(
		m,
		func(v int) int {
			return v
		},
		func(i int, key string, val int) bool {
			result = append(result, val)
			return true
		},
	)
	fmt.Println(collection.AllInComparableSlice(result, []int{1, 2, 3}))
	// Output:
	// true
}

func ExampleLoopMapByValueGetterDesc() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result []int
	collection.LoopMapByValueGetterDesc(
		m,
		func(v int) int {
			return v
		},
		func(i int, key string, val int) bool {
			result = append(result, val)
			return true
		},
	)
	fmt.Println(collection.AllInComparableSlice(result, []int{3, 2, 1}))
	// Output:
	// true
}
