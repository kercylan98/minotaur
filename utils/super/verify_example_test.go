package super_test

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/utils/super"
)

func ExampleVerify() {
	var getId = func() int { return 1 }
	var n *super.VerifyHandle[int]

	super.Verify(func(err error) {
		fmt.Println(err)
	}).Case(getId() == 1, errors.New("id can't be 1")).
		Do()

	super.Verify(func(err error) {
		fmt.Println(err)
	}).PreCase(func() bool {
		return n == nil
	}, errors.New("n can't be nil"), func(verify *super.VerifyHandle[error]) bool {
		return verify.Do()
	})

	// Output:
	// id can't be 1
	// n can't be nil
}
