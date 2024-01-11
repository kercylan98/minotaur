package collection_test

import (
	"github.com/kercylan98/minotaur/utils/collection"
	"testing"
)

func TestFindLoopedNextInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		i        int
		expected int
	}{
		{"TestFindLoopedNextInSlice_NonEmptySlice", []int{1, 2, 3}, 1, 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, _ := collection.FindLoopedNextInSlice(c.input, c.i)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the next index of input is not equal")
			}
		})
	}
}

func TestFindLoopedPrevInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		i        int
		expected int
	}{
		{"TestFindLoopedPrevInSlice_NonEmptySlice", []int{1, 2, 3}, 1, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, _ := collection.FindLoopedPrevInSlice(c.input, c.i)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the prev index of input is not equal")
			}
		})
	}
}

func TestFindCombinationsInSliceByRange(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		minSize  int
		maxSize  int
		expected [][]int
	}{
		{"TestFindCombinationsInSliceByRange_NonEmptySlice", []int{1, 2, 3}, 1, 2, [][]int{{1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindCombinationsInSliceByRange(c.input, c.minSize, c.maxSize)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the length of input is not equal")
			}
		})
	}
}

func TestFindFirstOrDefaultInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindFirstOrDefaultInSlice_NonEmptySlice", []int{1, 2, 3}, 1},
		{"TestFindFirstOrDefaultInSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindFirstOrDefaultInSlice(c.input, 0)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the first element of input is not equal")
			}
		})
	}
}

func TestFindOrDefaultInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindOrDefaultInSlice_NonEmptySlice", []int{1, 2, 3}, 2},
		{"TestFindOrDefaultInSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindOrDefaultInSlice(c.input, 0, func(v int) bool {
				return v == 2
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}

func TestFindOrDefaultInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindOrDefaultInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 2},
		{"TestFindOrDefaultInComparableSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindOrDefaultInComparableSlice(c.input, 2, 0)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}
