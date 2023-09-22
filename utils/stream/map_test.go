package stream_test

import (
	"github.com/kercylan98/minotaur/utils/stream"
	"testing"
)

func TestMap_Chunk(t *testing.T) {
	var m = map[string]int{
		"a": 1,
		"b": 2,
	}
	t.Log(stream.WithMap(m).Chunk(1).Merge())
}
