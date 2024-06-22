package eventstream_test

import (
	"github.com/kercylan98/minotaur/toolkit/eventstream"
	"testing"
)

func BenchmarkUnreliableSortStream_Publish(b *testing.B) {
	var es eventstream.UnreliableSortStream
	es.Subscribe(func(event eventstream.Event) {})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		es.Publish(nil)
	}
}

func BenchmarkUnreliableSortStream_Subscribe(b *testing.B) {
	var es eventstream.UnreliableSortStream
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		es.Subscribe(func(event eventstream.Event) {})
	}
}

func BenchmarkUnreliableSortStream_Unsubscribe(b *testing.B) {
	var es eventstream.UnreliableSortStream
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		es.Unsubscribe(es.Subscribe(func(event eventstream.Event) {}))
	}
}
