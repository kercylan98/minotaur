package buffer_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/buffer"
	"testing"
)

func TestUnbounded_Get(t *testing.T) {
	ub := buffer.NewUnbounded[int]()
	for i := 0; i < 100; i++ {
		ub.Put(i + 1)
		fmt.Println(<-ub.Get())
		//<-ub.Get()
		ub.Load()
	}
}
