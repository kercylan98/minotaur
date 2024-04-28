package buffer_test

import (
	"testing"
)

func TestNewRing(t *testing.T) {
	ring := buffer.NewRing[int]()
	for i := 0; i < 100; i++ {
		ring.Write(i)
		t.Log(ring.Read())
	}
}
