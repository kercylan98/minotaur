package slice_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
)

func TestPriority_Append(t *testing.T) {
	var s = slice.NewPriority[string]()
	s.Append("name_1", 2)
	s.Append("name_2", 1)
	fmt.Println(s)
}
