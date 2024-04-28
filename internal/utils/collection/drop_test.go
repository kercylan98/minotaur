package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"testing"
)

func TestClearSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestClearSlice_NonEmptySlice", []int{1, 2, 3}, []int{}},
		{"TestClearSlice_EmptySlice", []int{}, []int{}},
		{"TestClearSlice_NilSlice", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.ClearSlice(&c.input)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the inputV of input is not equal")
				}
			}
		})
	}
}

func TestClearMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]int
	}{
		{"TestClearMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{}},
		{"TestClearMap_EmptyMap", map[int]int{}, map[int]int{}},
		{"TestClearMap_NilMap", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.ClearMap(c.input)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the length of map is not equal")
			}
			for k, v := range c.input {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the inputV of map is not equal")
				}
			}
		})
	}
}

func TestDropSliceByIndices(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		indices  []int
		expected []int
	}{
		{"TestDropSliceByIndices_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{1, 3}, []int{1, 3, 5}},
		{"TestDropSliceByIndices_EmptySlice", []int{}, []int{1, 3}, []int{}},
		{"TestDropSliceByIndices_NilSlice", nil, []int{1, 3}, nil},
		{"TestDropSliceByIndices_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{}, []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DropSliceByIndices(&c.input, c.indices...)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the inputV of input is not equal")
				}
			}
		})
	}
}

func TestDropSliceByCondition(t *testing.T) {
	var cases = []struct {
		name      string
		input     []int
		condition func(v int) bool
		expected  []int
	}{
		{"TestDropSliceByCondition_NonEmptySlice", []int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 }, []int{1, 3, 5}},
		{"TestDropSliceByCondition_EmptySlice", []int{}, func(v int) bool { return v%2 == 0 }, []int{}},
		{"TestDropSliceByCondition_NilSlice", nil, func(v int) bool { return v%2 == 0 }, nil},
		{"TestDropSliceByCondition_NonEmptySlice", []int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 1 }, []int{2, 4}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DropSliceByCondition(&c.input, c.condition)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the inputV of input is not equal")
				}
			}
		})
	}
}

func TestDropSliceOverlappingElements(t *testing.T) {
	var cases = []struct {
		name               string
		input              []int
		anotherSlice       []int
		comparisonHandler  collection.ComparisonHandler[int]
		expected           []int
		expectedAnother    []int
		expectedComparison []int
	}{
		{"TestDropSliceOverlappingElements_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}, func(v1, v2 int) bool { return v1 == v2 }, []int{1, 2}, []int{6, 7}, []int{3, 4, 5}},
		{"TestDropSliceOverlappingElements_EmptySlice", []int{}, []int{3, 4, 5, 6, 7}, func(v1, v2 int) bool { return v1 == v2 }, []int{}, []int{3, 4, 5, 6, 7}, []int{}},
		{"TestDropSliceOverlappingElements_NilSlice", nil, []int{3, 4, 5, 6, 7}, func(v1, v2 int) bool { return v1 == v2 }, nil, []int{3, 4, 5, 6, 7}, nil},
		{"TestDropSliceOverlappingElements_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{}, func(v1, v2 int) bool { return v1 == v2 }, []int{1, 2, 3, 4, 5}, []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DropSliceOverlappingElements(&c.input, c.anotherSlice, c.comparisonHandler)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the inputV of input is not equal")
				}
			}
		})
	}
}
