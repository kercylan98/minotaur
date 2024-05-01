package toolkit_test

import (
	"github.com/kercylan98/minotaur/toolkit"
	"testing"
)

func BenchmarkCompareVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toolkit.CompareVersion("vfe2faf.d2ad5.dd3", "afe2faf.d2ad5.dd2")
	}
}

func BenchmarkOldVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toolkit.OldVersion("vfe2faf.d2ad5.dd3", "vda2faf.d2ad5.dd2")
	}
}
