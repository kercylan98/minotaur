package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
	"reflect"
)

func ExampleConvertSliceToAny() {
	result := collection.ConvertSliceToAny([]int{1, 2, 3})
	fmt.Println(reflect.TypeOf(result).String(), len(result))
	// Output:
	// []interface {} 3
}

func ExampleConvertSliceToIndexMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToIndexMap(slice)
	for i, v := range slice {
		fmt.Println(result[i], v)
	}
	// Output:
	// 1 1
	// 2 2
	// 3 3
}

func ExampleConvertSliceToIndexOnlyMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToIndexOnlyMap(slice)
	expected := map[int]bool{0: true, 1: true, 2: true}
	for k := range result {
		fmt.Println(expected[k])
	}
	// Output:
	// true
	// true
	// true
}

func ExampleConvertSliceToMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToMap(slice)
	fmt.Println(collection.AllKeyInMap(result, slice...))
	// Output:
	// true
}

func ExampleConvertSliceToBoolMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToBoolMap(slice)
	for _, v := range slice {
		fmt.Println(v, result[v])
	}
	// Output:
	// 1 true
	// 2 true
	// 3 true
}

func ExampleConvertMapKeysToSlice() {
	result := collection.ConvertMapKeysToSlice(map[int]int{1: 1, 2: 2, 3: 3})
	for i, v := range result {
		fmt.Println(i, v)
	}
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleConvertMapValuesToSlice() {
	result := collection.ConvertMapValuesToSlice(map[int]int{1: 1, 2: 2, 3: 3})
	expected := map[int]bool{1: true, 2: true, 3: true}
	for _, v := range result {
		fmt.Println(expected[v])
	}
	// Output:
	// true
	// true
	// true
}

func ExampleInvertMap() {
	result := collection.InvertMap(map[int]string{1: "a", 2: "b", 3: "c"})
	fmt.Println(collection.AllKeyInMap(result, "a", "b", "c"))
	// Output:
	// true
}

func ExampleConvertMapValuesToBool() {
	result := collection.ConvertMapValuesToBool(map[int]int{1: 1})
	fmt.Println(result)
	// Output:
	// map[1:true]
}

func ExampleReverseSlice() {
	var s = []int{1, 2, 3}
	collection.ReverseSlice(&s)
	fmt.Println(s)
	// Output:
	// [3 2 1]
}
