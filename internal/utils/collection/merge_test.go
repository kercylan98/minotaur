package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"testing"
)

func TestMergeSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestMergeSlice_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"TestMergeSlice_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MergeSlice(c.input...)
			if len(result) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
			for i := 0; i < len(result); i++ {
				if result[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the value of input is not equal")
				}
			}
		})
	}
}

func TestMergeSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{"TestMergeSlices_NonEmptySlice", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"TestMergeSlices_EmptySlice", [][]int{{}, {}}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MergeSlices(c.input...)
			if len(result) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
			for i := 0; i < len(result); i++ {
				if result[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the value of input is not equal")
				}
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	var cases = []struct {
		name     string
		input    []map[int]int
		expected int
	}{
		{"TestMergeMaps_NonEmptyMap", []map[int]int{{1: 1, 2: 2, 3: 3}, {4: 4, 5: 5, 6: 6}}, 6},
		{"TestMergeMaps_EmptyMap", []map[int]int{{}, {}}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MergeMaps(c.input...)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestMergeMapsWithSkip(t *testing.T) {
	var cases = []struct {
		name     string
		input    []map[int]int
		expected int
	}{
		{"TestMergeMapsWithSkip_NonEmptyMap", []map[int]int{{1: 1}, {1: 2}}, 1},
		{"TestMergeMapsWithSkip_EmptyMap", []map[int]int{{}, {}}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MergeMapsWithSkip(c.input...)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}
