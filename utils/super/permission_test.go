package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestNewPermission(t *testing.T) {
	const (
		Read = 1 << iota
		Write
		Execute
	)

	p := super.NewPermission[int, int]()
	p.AddPermission(1, Read, Write)
	t.Log(p.HasPermission(1, Read))
	t.Log(p.HasPermission(1, Write))
	p.SetPermission(2, Read|Write)
	t.Log(p.HasPermission(2, Read))
	t.Log(p.HasPermission(2, Execute))
	p.SetPermission(2, Execute)
	t.Log(p.HasPermission(2, Execute))
	t.Log(p.HasPermission(2, Read))
	t.Log(p.HasPermission(2, Write))
	p.RemovePermission(2, Execute)
	t.Log(p.HasPermission(2, Execute))
}
