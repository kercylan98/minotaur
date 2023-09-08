package stream_test

import (
	"github.com/kercylan98/minotaur/utils/stream"
	"testing"
)

func TestStream(t *testing.T) {
	var s = stream.Slice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(s, s.
		Copy().
		Shuffle().
		Filter(true, func(index int, item int) bool {
			return item%2 == 0
		}).
		Zoom(20).
		Each(true, func(index int, item int) bool {
			t.Log(index, item)
			return false
		}).
		Chunk(3).
		EachT(func(index int, item stream.Slice[int]) bool {
			t.Log(item)
			return false
		}).
		Merge().
		FillBy(func(index int, value int) int {
			if value == 0 {
				return 999
			}
			return value
		}),
	)
}
