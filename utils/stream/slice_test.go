package stream_test

import (
	"github.com/kercylan98/minotaur/utils/stream"
	"testing"
)

func TestStream_Chunk(t *testing.T) {
	var s = stream.Slice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var chunks = s.Chunk(3)
	for _, chunk := range chunks {
		t.Log(chunk)
	}
}
