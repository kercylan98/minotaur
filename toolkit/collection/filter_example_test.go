package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

func ExampleFilterOutByIndices() {
	var slice = []int{1, 2, 3, 4, 5}
	var result = collection.FilterOutByIndices(slice, 1, 3)
	fmt.Println(result)
	// Output:
	// [1 3 5]
}

func ExampleFilterOutByCondition() {
	var slice = []int{1, 2, 3, 4, 5}
	var result = collection.FilterOutByCondition(slice, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(result)
	// Output:
	// [1 3 5]
}

func ExampleFilterOutByKey() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByKey(m, "b")
	fmt.Println(result)
	// Output:
	// map[a:1 c:3]
}

func ExampleFilterOutByValue() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByValue(m, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(len(result))
	// Output:
	// 2
}

func ExampleFilterOutByKeys() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByKeys(m, "a", "c")
	fmt.Println(result)
	// Output:
	// map[b:2]
}

func ExampleFilterOutByValues() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByValues(m, []int{1}, func(source, target int) bool {
		return source == target
	})
	for i, s := range []string{"a", "b", "c"} {
		fmt.Println(i, result[s])
	}
	// Output:
	// 0 0
	// 1 2
	// 2 3
}

func ExampleFilterOutByMap() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByMap(m, func(k string, v int) bool {
		return k == "a" || v == 3
	})
	fmt.Println(result)
	// Output:
	// map[b:2]
}
