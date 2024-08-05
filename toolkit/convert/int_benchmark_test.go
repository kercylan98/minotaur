package convert_test

import (
	"github.com/kercylan98/minotaur/toolkit/convert"
	"testing"
)

func BenchmarkUint64ToString(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		_ = convert.Uint64ToString(i)
	}
	b.StopTimer()
}

func BenchmarkFastUint64ToString(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		_ = convert.FastUint64ToString(i)
	}
}
