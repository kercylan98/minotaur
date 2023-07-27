package super_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/super"
	"go.uber.org/atomic"
	"testing"
)

func TestConvert(t *testing.T) {
	type B struct {
		nocmp [0]func()
		v     atomic.Value
	}
	var a = atomic.NewString("hello")
	var b = super.Convert[*atomic.String, *B](a)
	fmt.Println(b.v.Load())
}
