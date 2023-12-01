package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	t.Log(super.CompareVersion("1", "2"), -1)
	t.Log(super.CompareVersion("1", "vv2"), -1)
	t.Log(super.CompareVersion("1", "vv2.3.1"), -1)
	t.Log(super.CompareVersion("11", "vv2.3.1"), 1)
}
