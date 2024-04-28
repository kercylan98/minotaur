package maths_test

import (
	"github.com/kercylan98/minotaur/utils/maths"
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
)

func TestMakeLastDigitsZero(t *testing.T) {
	for i := 0; i < 20; i++ {
		n := float64(random.Int64(100, 999999))
		t.Log(n, 3, maths.MakeLastDigitsZero(n, 3))
	}
}
