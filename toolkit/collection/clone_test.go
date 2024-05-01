package collection_test

import (
	"github.com/kercylan98/minotaur/toolkit/collection"
	"testing"
)

func TestCloneSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestCloneSlice_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"TestCloneSlice_EmptySlice", []int{}, []int{}},
		{"TestCloneSlice_NilSlice", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of input is not equal")
				}
			}
		})
	}
}

func TestCloneMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]int
	}{
		{"TestCloneMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 2: 2, 3: 3}},
		{"TestCloneMap_EmptyMap", map[int]int{}, map[int]int{}},
		{"TestCloneMap_NilMap", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneMap(c.input)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of map is not equal")
				}
			}
		})
	}
}

func TestCloneSliceN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputN   int
		expected [][]int
	}{
		{"TestCloneSliceN_NonEmptySlice", []int{1, 2, 3}, 2, [][]int{{1, 2, 3}, {1, 2, 3}}},
		{"TestCloneSliceN_EmptySlice", []int{}, 2, [][]int{{}, {}}},
		{"TestCloneSliceN_NilSlice", nil, 2, nil},
		{"TestCloneSliceN_NonEmptySlice_ZeroN", []int{1, 2, 3}, 0, [][]int{}},
		{"TestCloneSliceN_EmptySlice_ZeroN", []int{}, 0, [][]int{}},
		{"TestCloneSliceN_NilSlice_ZeroN", nil, 0, nil},
		{"TestCloneSliceN_NonEmptySlice_NegativeN", []int{1, 2, 3}, -1, [][]int{}},
		{"TestCloneSliceN_EmptySlice_NegativeN", []int{}, -1, [][]int{}},
		{"TestCloneSliceN_NilSlice_NegativeN", nil, -1, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneSliceN(c.input, c.inputN)
			if actual == nil {
				if c.expected != nil {
					t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, c.inputN, actual, "after clone, the expected is nil")
				}
				return
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, c.inputN, actual, "after clone, the inputV of input is not equal")
					}
				}
			}
		})
	}
}

func TestCloneMapN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		inputN   int
		expected []map[int]int
	}{
		{"TestCloneMapN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 2, []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}},
		{"TestCloneMapN_EmptyMap", map[int]int{}, 2, []map[int]int{{}, {}}},
		{"TestCloneMapN_NilMap", nil, 2, nil},
		{"TestCloneMapN_NonEmptyMap_ZeroN", map[int]int{1: 1, 2: 2, 3: 3}, 0, []map[int]int{}},
		{"TestCloneMapN_EmptyMap_ZeroN", map[int]int{}, 0, []map[int]int{}},
		{"TestCloneMapN_NilMap_ZeroN", nil, 0, nil},
		{"TestCloneMapN_NonEmptyMap_NegativeN", map[int]int{1: 1, 2: 2, 3: 3}, -1, []map[int]int{}},
		{"TestCloneMapN_EmptyMap_NegativeN", map[int]int{}, -1, []map[int]int{}},
		{"TestCloneMapN_NilMap_NegativeN", nil, -1, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneMapN(c.input, c.inputN)
			if actual == nil {
				if c.expected != nil {
					t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, actual, c.inputN, "after clone, the expected is nil")
				}
				return
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, actual, c.inputN, "after clone, the inputV of map is not equal")
					}
				}
			}
		})
	}
}

func TestCloneSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{"TestCloneSlices_NonEmptySlices", [][]int{{1, 2, 3}, {1, 2, 3}}, [][]int{{1, 2, 3}, {1, 2, 3}}},
		{"TestCloneSlices_EmptySlices", [][]int{{}, {}}, [][]int{{}, {}}},
		{"TestCloneSlices_NilSlices", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneSlices(c.input...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of input is not equal")
					}
				}
			}
		})
	}
}

func TestCloneMaps(t *testing.T) {
	var cases = []struct {
		name     string
		input    []map[int]int
		expected []map[int]int
	}{
		{"TestCloneMaps_NonEmptyMaps", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}, []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}},
		{"TestCloneMaps_EmptyMaps", []map[int]int{{}, {}}, []map[int]int{{}, {}}},
		{"TestCloneMaps_NilMaps", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneMaps(c.input...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of maps is not equal")
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of maps is not equal")
					}
				}
			}
		})
	}
}
