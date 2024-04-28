package collection_test

import (
	"github.com/kercylan98/minotaur/utils/collection"
	"testing"
)

func TestDeduplicateSliceInPlace(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "TestDeduplicateSliceInPlace_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "TestDeduplicateSliceInPlace_Empty", input: []int{}, expected: []int{}},
		{name: "TestDeduplicateSliceInPlace_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DeduplicateSliceInPlace(&c.input)
			if len(c.input) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, c.input)
			}
			for i := 0; i < len(c.input); i++ {
				av, ev := c.input[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
			}
		})
	}
}

func TestDeduplicateSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "TestDeduplicateSlice_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "TestDeduplicateSlice_Empty", input: []int{}, expected: []int{}},
		{name: "TestDeduplicateSlice_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.DeduplicateSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for i := 0; i < len(actual); i++ {
				av, ev := actual[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestDeduplicateSliceInPlaceWithCompare(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "TestDeduplicateSliceInPlaceWithCompare_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "TestDeduplicateSliceInPlaceWithCompare_Empty", input: []int{}, expected: []int{}},
		{name: "TestDeduplicateSliceInPlaceWithCompare_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DeduplicateSliceInPlaceWithCompare(&c.input, func(a, b int) bool {
				return a == b
			})
			if len(c.input) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, c.input)
			}
			for i := 0; i < len(c.input); i++ {
				av, ev := c.input[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
			}
		})
	}
}

func TestDeduplicateSliceWithCompare(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "TestDeduplicateSliceWithCompare_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "TestDeduplicateSliceWithCompare_Empty", input: []int{}, expected: []int{}},
		{name: "TestDeduplicateSliceWithCompare_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.DeduplicateSliceWithCompare(c.input, func(a, b int) bool {
				return a == b
			})
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for i := 0; i < len(actual); i++ {
				av, ev := actual[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}
