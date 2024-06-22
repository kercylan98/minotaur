package eventstream_test

import (
	"github.com/kercylan98/minotaur/toolkit/eventstream"
	"testing"
)

func TestUnreliableSortStream(t *testing.T) {
	var es eventstream.UnreliableSortStream
	var counter int
	es.Subscribe(func(event eventstream.Event) {
		t.Log(event)
		counter++
	})

	es.Subscribe(func(event eventstream.Event) {
		t.Log(event)
		counter++
	})

	es.Publish("Hello, World!")

	if counter != 2 {
		t.Error("counter should be 2")
	}
}
