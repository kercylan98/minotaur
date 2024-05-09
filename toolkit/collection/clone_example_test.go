package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/collection"
)

// 在该示例中，将 slice 克隆后将会得到一个新的 slice result，而 result 和 slice 将不会有任何关联，但是如果 slice 中的元素是引用类型，那么 result 中的元素将会和 slice 中的元素指向同一个地址
//   - 示例中的结果将会输出 [1 2 3]
func ExampleCloneSlice() {
	var slice = []int{1, 2, 3}
	var result = collection.CloneSlice(slice)
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

// 在该示例中，将 map 克隆后将会得到一个新的 map result，而 result 和 map 将不会有任何关联，但是如果 map 中的元素是引用类型，那么 result 中的元素将会和 map 中的元素指向同一个地址
//   - 示例中的结果将会输出 3
func ExampleCloneMap() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMap(m)
	fmt.Println(len(result))
	// Output:
	// 3
}

// 通过将 slice 克隆为 2 个新的 slice，将会得到一个新的 slice result，而 result 和 slice 将不会有任何关联，但是如果 slice 中的元素是引用类型，那么 result 中的元素将会和 slice 中的元素指向同一个地址
//   - result 的结果为 [[1 2 3] [1 2 3]]
//   - 示例中的结果将会输出 2
func ExampleCloneSliceN() {
	var slice = []int{1, 2, 3}
	var result = collection.CloneSliceN(slice, 2)
	fmt.Println(len(result))
	// Output:
	// 2
}

// 通过将 map 克隆为 2 个新的 map，将会得到一个新的 map result，而 result 和 map 将不会有任何关联，但是如果 map 中的元素是引用类型，那么 result 中的元素将会和 map 中的元素指向同一个地址
//   - result 的结果为 [map[1:1 2:2 3:3] map[1:1 2:2 3:3]] `无序的 Key-value 对`
//   - 示例中的结果将会输出 2
func ExampleCloneMapN() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMapN(m, 2)
	fmt.Println(len(result))
	// Output:
	// 2
}

// 通过将多个 slice 克隆为 2 个新的 slice，将会得到一个新的 slice result，而 result 和 slice 将不会有任何关联，但是如果 slice 中的元素是引用类型，那么 result 中的元素将会和 slice 中的元素指向同一个地址
//   - result 的结果为 [[1 2 3] [1 2 3]]
func ExampleCloneSlices() {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 3}
	var result = collection.CloneSlices(slice1, slice2)
	fmt.Println(len(result))
	// Output:
	// 2
}

// 通过将多个 map 克隆为 2 个新的 map，将会得到一个新的 map result，而 result 和 map 将不会有任何关联，但是如果 map 中的元素是引用类型，那么 result 中的元素将会和 map 中的元素指向同一个地址
//   - result 的结果为 [map[1:1 2:2 3:3] map[1:1 2:2 3:3]] `无序的 Key-value 对`
func ExampleCloneMaps() {
	var m1 = map[int]int{1: 1, 2: 2, 3: 3}
	var m2 = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMaps(m1, m2)
	fmt.Println(len(result))
	// Output:
	// 2
}
