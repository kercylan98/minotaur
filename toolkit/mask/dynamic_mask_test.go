package mask_test

import (
	"github.com/kercylan98/minotaur/toolkit/mask"
	"testing"
)

func TestDynamicMask_Set(t *testing.T) {
	var m mask.DynamicMask

	m.Set(1)
	m.Set(2)
	m.Set(3)

	bits := m.Bits()
	t.Log(bits)
}

func TestDynamicMask_Del(t *testing.T) {
	var m mask.DynamicMask

	m.Set(1)
	m.Set(2)
	m.Del(1)

	bits := m.Bits()
	t.Log(bits)
}

func TestDynamicMask_Has(t *testing.T) {
	var m mask.DynamicMask

	m.Set(1)
	m.Set(2)

	has := m.Has(1)
	t.Log(has)
}

func TestDynamicMask_Bits(t *testing.T) {
	var m mask.DynamicMask

	m.Set(1)
	m.Set(2)

	bits := m.Bits()
	t.Log(bits)
}

func TestDynamicMask_Clear(t *testing.T) {
	var m mask.DynamicMask

	m.Set(1)
	m.Set(2)
	m.Clear()

	bits := m.Bits()
	t.Log(bits)
}

func TestDynamicMask_Clone(t *testing.T) {
	var m mask.DynamicMask

	m.Set(1)
	m.Set(2)

	clone := m.Clone()
	clone.Set(3)
	t.Log(m.Bits())
	t.Log(clone.Bits())
}
