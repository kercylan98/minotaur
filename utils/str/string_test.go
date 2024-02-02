package str_test

import (
	"github.com/kercylan98/minotaur/utils/str"
	"testing"
)

func TestString_ToLower(t *testing.T) {
	var s str.String = "HELLO"
	s.ToLower()
	t.Log(s)
}
