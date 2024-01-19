package huge_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/huge"
)

// 该案例展示了 NewInt 对各种基本类型的支持及用法
func ExampleNewInt() {
	fmt.Println(huge.NewInt("12345678900000000"))
	fmt.Println(huge.NewInt(1234567890))
	fmt.Println(huge.NewInt(true))
	fmt.Println(huge.NewInt(123.123))
	fmt.Println(huge.NewInt(byte(1)))
	// Output:
	// 12345678900000000
	// 1234567890
	// 1
	// 123
	// 1
}

func ExampleInt_Copy() {
	var a = huge.NewInt(1234567890)
	var b = a.Copy().SetInt64(9876543210)
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// 1234567890
	// 9876543210
}

func ExampleInt_Set() {
	var a = huge.NewInt(1234567890)
	var b = huge.NewInt(9876543210)
	fmt.Println(a)
	a.Set(b)
	fmt.Println(a)
	// Output:
	// 1234567890
	// 9876543210
}
