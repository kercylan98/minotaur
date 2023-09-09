package super_test

import (
	"errors"
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestGetError(t *testing.T) {
	var ErrNotFound = errors.New("not found")
	var _ = super.RegErrorRef(100, "test error", ErrNotFound)
	t.Log(super.GetError(ErrNotFound))
}
