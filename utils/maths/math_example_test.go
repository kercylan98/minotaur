package maths_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/maths"
)

func ExampleToContinuous() {
	var nums = []int{1, 2, 3, 5, 6, 7, 9, 10, 11}
	var continuous = maths.ToContinuous(nums)

	fmt.Println(nums)
	fmt.Println(continuous)

	// Output:
	// [1 2 3 5 6 7 9 10 11]
	// map[1:1 2:2 3:3 4:5 5:6 6:7 7:9 8:10 9:11]
}
