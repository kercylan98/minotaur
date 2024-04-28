package super_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/super"
)

func ExampleRecoverTransform() {
	defer func() {
		if err := super.RecoverTransform(recover()); err != nil {
			fmt.Println(err)
		}
	}()
	panic("test")
	// Output:
	// test
}
