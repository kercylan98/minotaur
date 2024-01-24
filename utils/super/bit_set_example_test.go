package super_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/super"
)

func ExampleNewBitSet() {
	var bs = super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	bs.Set(10)
	fmt.Println(bs.Bits())
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExampleBitSet_Set() {
	var bs = super.NewBitSet[int]()
	bs.Set(10)
	fmt.Println(bs.Bits())
	// Output:
	// [10]
}

func ExampleBitSet_Del() {
	var bs = super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	bs.Del(1)
	fmt.Println(bs.Bits())
	// Output:
	// [2 3 4 5 6 7 8 9]
}

func ExampleBitSet_Shrink() {
	var bs = super.NewBitSet(111, 222, 333, 444)
	fmt.Println(bs.Cap())
	bs.Del(444)
	fmt.Println(bs.Cap())
	bs.Shrink()
	fmt.Println(bs.Cap())
	// Output:
	// 448
	// 448
	// 384
}

func ExampleBitSet_Cap() {
	var bs = super.NewBitSet(63)
	fmt.Println(bs.Cap())
	// Output:
	// 64
}

func ExampleBitSet_Has() {
	var bs = super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(bs.Has(1))
	fmt.Println(bs.Has(10))
	// Output:
	// true
	// false
}

func ExampleBitSet_Clear() {
	var bs = super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	bs.Clear()
	fmt.Println(bs.Bits())
	// Output:
	// []
}

func ExampleBitSet_Len() {
	var bs = super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(bs.Len())
	// Output:
	// 9
}
