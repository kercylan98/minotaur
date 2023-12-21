package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestBitSet_Set(t *testing.T) {
	bs := super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	bs.Set(11)
	bs.Set(12)
	bs.Set(13)
	t.Log(bs)
}

func TestBitSet_Del(t *testing.T) {
	bs := super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	bs.Del(11)
	bs.Del(12)
	bs.Del(13)
	bs.Del(10)
	t.Log(bs)
}

func TestBitSet_Shrink(t *testing.T) {
	bs := super.NewBitSet(63)
	t.Log(bs.Cap())
	bs.Set(200)
	t.Log(bs.Cap())
	bs.Del(200)
	bs.Shrink()
	t.Log(bs.Cap())
}
