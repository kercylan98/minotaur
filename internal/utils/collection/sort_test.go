package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"sort"
"testing"
)

func TestDescBy(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestDescBy_NonEmptySlice", []int{1, 2, 3}, []int{3, 2, 1}},
		{"TestDescBy_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Slice(c.input, func(i, j int) bool {
				return collection.DescBy(c.input[i], c.input[j])
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

func TestAscBy(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestAscBy_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"TestAscBy_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Slice(c.input, func(i, j int) bool {
				return collection.AscBy(c.input[i], c.input[j])
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

func TestDesc(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestDesc_NonEmptySlice", []int{1, 2, 3}, []int{3, 2, 1}},
		{"TestDesc_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.Desc(&c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

func TestDescByClone(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestDescByClone_NonEmptySlice", []int{1, 2, 3}, []int{3, 2, 1}},
		{"TestDescByClone_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.DescByClone(c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range result {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, result)
				}
			}
		})
	}
}

func TestAsc(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestAsc_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"TestAsc_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.Asc(&c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

func TestAscByClone(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestAscByClone_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"TestAscByClone_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.AscByClone(c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range result {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, result)
				}
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestShuffle_NonEmptySlice", []int{1, 2, 3}, 3},
		{"TestShuffle_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.Shuffle(&c.input)
			if len(c.input) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
			}
		})
	}
}

func TestShuffleByClone(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestShuffleByClone_NonEmptySlice", []int{1, 2, 3}, 3},
		{"TestShuffleByClone_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ShuffleByClone(c.input)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, result)
			}
		})
	}
}
