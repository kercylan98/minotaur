package toolkit_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit"
)

func ExampleCompareVersion() {
	result := toolkit.CompareVersion("1.2.3", "1.2.2")
	fmt.Println(result)
	// Output: 1
}

func ExampleOldVersion() {
	result := toolkit.OldVersion("1.2.3", "1.2.2")
	fmt.Println(result)
	// Output: true
}
