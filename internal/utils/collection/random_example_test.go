package collection_test

import (
	"fmt"
"github.com/kercylan98/minotaur/toolkit/collection"
)

func ExampleChooseRandomSliceElementRepeatN() {
	result := collection.ChooseRandomSliceElementRepeatN([]int{1}, 10)
	fmt.Println(result)
	// Output:
	// [1 1 1 1 1 1 1 1 1 1]
}

func ExampleChooseRandomIndexRepeatN() {
	result := collection.ChooseRandomIndexRepeatN([]int{1}, 10)
	fmt.Println(result)
	// Output:
	// [0 0 0 0 0 0 0 0 0 0]
}

func ExampleChooseRandomSliceElement() {
	result := collection.ChooseRandomSliceElement([]int{1})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleChooseRandomIndex() {
	result := collection.ChooseRandomIndex([]int{1})
	fmt.Println(result)
	// Output:
	// 0
}

func ExampleChooseRandomSliceElementN() {
	result := collection.ChooseRandomSliceElementN([]int{1}, 1)
	fmt.Println(result)
	// Output:
	// [1]
}

func ExampleChooseRandomIndexN() {
	result := collection.ChooseRandomIndexN([]int{1}, 1)
	fmt.Println(result)
	// Output:
	// [0]
}

func ExampleChooseRandomMapKeyRepeatN() {
	result := collection.ChooseRandomMapKeyRepeatN(map[int]int{1: 1}, 1)
	fmt.Println(result)
	// Output:
	// [1]
}

func ExampleChooseRandomMapValueRepeatN() {
	result := collection.ChooseRandomMapValueRepeatN(map[int]int{1: 1}, 1)
	fmt.Println(result)
	// Output:
	// [1]
}

func ExampleChooseRandomMapKeyAndValueRepeatN() {
	result := collection.ChooseRandomMapKeyAndValueRepeatN(map[int]int{1: 1}, 1)
	fmt.Println(result)
	// Output:
	// map[1:1]
}

func ExampleChooseRandomMapKey() {
	result := collection.ChooseRandomMapKey(map[int]int{1: 1})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleChooseRandomMapValue() {
	result := collection.ChooseRandomMapValue(map[int]int{1: 1})
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleChooseRandomMapKeyN() {
	result := collection.ChooseRandomMapKeyN(map[int]int{1: 1}, 1)
	fmt.Println(result)
	// Output:
	// [1]
}

func ExampleChooseRandomMapValueN() {
	result := collection.ChooseRandomMapValueN(map[int]int{1: 1}, 1)
	fmt.Println(result)
	// Output:
	// [1]
}

func ExampleChooseRandomMapKeyAndValue() {
	k, v := collection.ChooseRandomMapKeyAndValue(map[int]int{1: 1})
	fmt.Println(k, v)
	// Output:
	// 1 1
}

func ExampleChooseRandomMapKeyAndValueN() {
	result := collection.ChooseRandomMapKeyAndValueN(map[int]int{1: 1}, 1)
	fmt.Println(result)
	// Output:
	// map[1:1]
}
