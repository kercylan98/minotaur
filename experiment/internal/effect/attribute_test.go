package effect_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/effect"
	"testing"
)

func TestAttributesGetAndSet(t *testing.T) {
	var hp = int32(1)
	var attrs = effect.NewAttributes()
	attrs.Set(hp, attrs.Get(hp).AddInt(1000))

	t.Log(attrs.Get(hp))
}
