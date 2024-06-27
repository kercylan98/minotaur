package queues_test

import (
	"github.com/kercylan98/minotaur/toolkit/queues"
	"testing"
)

func BenchmarkMPSC_Push(b *testing.B) {
	q := queues.NewMPSC()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkMPSC_Pop(b *testing.B) {
	q := queues.NewMPSC()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkMPSC_PushAndPop(b *testing.B) {
	q := queues.NewMPSC()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}
