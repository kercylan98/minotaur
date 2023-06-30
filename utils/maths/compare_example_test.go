package maths_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/maths"
)

func ExampleIsContinuity() {
	fmt.Println(maths.IsContinuity([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(maths.IsContinuity([]int{1, 2, 3, 5, 5, 6, 7}))
	// Output:
	// true
	// false
}
