package queues_test

import (
	"github.com/kercylan98/minotaur/toolkit/queues"
	"testing"
	"unsafe"
)

func BenchmarkLFQueue_Push(b *testing.B) {
	q := queues.NewLFQueue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(unsafe.Pointer(&i))
	}
}

func BenchmarkLFQueue_Pop(b *testing.B) {
	q := queues.NewLFQueue()
	for i := 0; i < b.N; i++ {
		q.Push(unsafe.Pointer(&i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkLFQueue_PushAndPop(b *testing.B) {
	q := queues.NewLFQueue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(unsafe.Pointer(&i))
		q.Pop()
	}
}

func BenchmarkLFQueue_BatchPop(b *testing.B) {
	q := queues.NewLFQueue()
	for i := 0; i < b.N; i++ {
		q.Push(unsafe.Pointer(&i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.BatchPop(1)
	}
}
