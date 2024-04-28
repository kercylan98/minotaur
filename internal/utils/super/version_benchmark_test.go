package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func BenchmarkCompareVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		super.CompareVersion("vfe2faf.d2ad5.dd3", "afe2faf.d2ad5.dd2")
	}
}

func BenchmarkOldVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		super.OldVersion("vfe2faf.d2ad5.dd3", "vda2faf.d2ad5.dd2")
	}
}
