package sher_test

import (
	"github.com/kercylan98/minotaur/utils/sher"
	"testing"
)

func TestDeduplicateSliceInPlace(t *testing.T) {
	var cases = []struct {
		s        []int
		expected []int
	}{
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, c := range cases {
		sher.DeduplicateSliceInPlace(&c.s)
		if len(c.s) != len(c.expected) {
			t.Errorf("DeduplicateSliceInPlace(%v) == %v, expected %v", c.s, c.s, c.expected)
		}
	}
}

func TestDeduplicateSlice(t *testing.T) {
	var cases = []struct {
		s        []int
		expected []int
	}{
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, c := range cases {
		sl := len(c.s)
		if r := sher.DeduplicateSlice(c.s); len(r) != len(c.expected) || len(c.s) != sl {
			t.Errorf("DeduplicateSlice(%v) == %v, expected %v", c.s, r, c.expected)
		}
	}
}

func TestDeduplicateSliceInPlaceWithCompare(t *testing.T) {
	var cases = []struct {
		s        []int
		expected []int
	}{
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, c := range cases {
		sher.DeduplicateSliceInPlaceWithCompare(&c.s, func(a, b int) bool {
			return a == b
		})
		if len(c.s) != len(c.expected) {
			t.Errorf("DeduplicateSliceInPlaceWithCompare(%v) == %v, expected %v", c.s, c.s, c.expected)
		}
	}
}

func TestDeduplicateSliceWithCompare(t *testing.T) {
	var cases = []struct {
		s        []int
		expected []int
	}{
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, c := range cases {
		sl := len(c.s)
		if r := sher.DeduplicateSliceWithCompare(c.s, func(a, b int) bool {
			return a == b
		}); len(r) != len(c.expected) || len(c.s) != sl {
			t.Errorf("DeduplicateSliceWithCompare(%v) == %v, expected %v", c.s, r, c.expected)
		}
	}
}
