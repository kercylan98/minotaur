package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestMask(t *testing.T) {
	mask := super.Mask(1, 2, 3, 5)

	t.Log(mask.Matches(super.Mask(1, 2, 3, 5)))
	t.Log(mask.Matches(super.Mask(1, 2, 3, 4, 5)))
	t.Log(mask.Matches(super.Mask(1, 2, 5)))
	t.Log(mask.Contains(super.Mask(1, 2, 5)))
	t.Log(mask)
}
