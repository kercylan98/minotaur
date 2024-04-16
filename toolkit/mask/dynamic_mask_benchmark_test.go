package mask_test

import (
	"github.com/kercylan98/minotaur/toolkit/mask"
	"testing"
)

func BenchmarkDynamicMask_Set(b *testing.B) {

	b.Run("Set", func(b *testing.B) {
		var m mask.DynamicMask
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Set(i)
		}
	})

	b.Run("Set64", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var m mask.DynamicMask
			for j := 0; j < 64; j++ {
				m.Set(j)
			}
		}
	})

	b.Run("Set128", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var m mask.DynamicMask
			for j := 0; j < 128; j++ {
				m.Set(j)
			}
		}
	})

	b.Run("Set256", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var m mask.DynamicMask
			for j := 0; j < 256; j++ {
				m.Set(j)
			}
		}
	})

}

func BenchmarkDynamicMask_Bits(b *testing.B) {
	b.Run("Bits", func(b *testing.B) {
		var m mask.DynamicMask
		for i := 0; i < 1000; i++ {
			m.Set(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Bits()
		}
	})

	b.Run("Bits64", func(b *testing.B) {

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var m mask.DynamicMask
			for j := 0; j < 64; j++ {
				m.Set(j)
			}
			m.Bits()
		}
	})
}

func BenchmarkDynamicMask_Has(b *testing.B) {
	b.Run("Has", func(b *testing.B) {
		var m mask.DynamicMask
		for j := 0; j < 64; j++ {
			m.Set(j)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Has(i)
		}
	})
}

func BenchmarkDynamicMask_HasAny(b *testing.B) {
	b.Run("HasAny", func(b *testing.B) {
		var m mask.DynamicMask
		for j := 0; j < 64; j++ {
			m.Set(j)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.HasAny([]int{1, 2, 3})
		}
	})
}

func BenchmarkDynamicMask_Equal(b *testing.B) {
	b.Run("Equal", func(b *testing.B) {
		var m mask.DynamicMask
		for j := 0; j < 64; j++ {
			m.Set(j)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Equal(m)
		}
	})
}

func BenchmarkDynamicMask_EqualBits(b *testing.B) {
	b.Run("EqualBits", func(b *testing.B) {
		var m mask.DynamicMask
		for j := 0; j < 3; j++ {
			m.Set(j)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.EqualBits([]int{1, 2, 3})
		}
	})
}

func BenchmarkDynamicMask_Clear(b *testing.B) {
	b.Run("Clear", func(b *testing.B) {
		var m mask.DynamicMask
		for j := 0; j < 64; j++ {
			m.Set(j)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Clear()
		}
	})
}

func BenchmarkDynamicMask_Clone(b *testing.B) {
	b.Run("Clone", func(b *testing.B) {
		var m mask.DynamicMask
		for j := 0; j < 64; j++ {
			m.Set(j)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Clone()
		}
	})
}

func BenchmarkDynamicMask_Del(b *testing.B) {
	b.Run("Del", func(b *testing.B) {
		var m mask.DynamicMask
		for j := 0; j < 64; j++ {
			m.Set(j)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Del(i)
		}
	})
}

func BenchmarkDynamicMask_SetAndDel(b *testing.B) {
	b.Run("SetAndDel", func(b *testing.B) {
		var m mask.DynamicMask
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Set(i)
			m.Del(i)
		}
	})
}
