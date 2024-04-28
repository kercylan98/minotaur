package super_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/super"
)

func ExampleCompareVersion() {
	result := super.CompareVersion("1.2.3", "1.2.2")
	fmt.Println(result)
	// Output: 1
}

func ExampleOldVersion() {
	result := super.OldVersion("1.2.3", "1.2.2")
	fmt.Println(result)
	// Output: true
}
