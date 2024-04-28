package buffer_test

import (
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"testing"
)

func TestRingUnbounded_Write2Read(t *testing.T) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	for i := 0; i < 100; i++ {
		ring.Write(i)
	}
	t.Log("write done")
	for i := 0; i < 100; i++ {
		t.Log(<-ring.Read())
	}
	t.Log("read done")
}

func TestRingUnbounded_Close(t *testing.T) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	for i := 0; i < 100; i++ {
		ring.Write(i)
	}
	t.Log("write done")
	ring.Close()
	t.Log("close done")
	for v := range ring.Read() {
		ring.Write(v)
		t.Log(v)
	}
	t.Log("read done")
}
