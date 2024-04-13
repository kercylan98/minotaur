package ident_test

import (
	"github.com/kercylan98/minotaur/toolkit/ident"
	"testing"
)

func BenchmarkGenerateOrderedUniqueIdentStringWithUInt64_10(b *testing.B) {
	var ids = []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ident.GenerateOrderedUniqueIdentStringWithUInt64(ids...)
	}
}

func BenchmarkGenerateOrderedUniqueIdentStringWithUInt64_100(b *testing.B) {
	var ids = make([]uint64, 100)
	for i := 0; i < 100; i++ {
		ids[i] = uint64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ident.GenerateOrderedUniqueIdentStringWithUInt64(ids...)
	}
}
