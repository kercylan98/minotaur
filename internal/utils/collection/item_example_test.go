package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

func ExampleSwapSlice() {
	var s = []int{1, 2, 3}
	collection.SwapSlice(&s, 0, 1)
	fmt.Println(s)
	// Output:
	// [2 1 3]
}
