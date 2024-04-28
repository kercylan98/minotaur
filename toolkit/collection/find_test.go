package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
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

func TestFindInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindInSlice_NonEmptySlice", []int{1, 2, 3}, 2},
		{"TestFindInSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, actual := collection.FindInSlice(c.input, func(v int) bool {
				return v == 2
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}

func TestFindIndexInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindIndexInSlice_NonEmptySlice", []int{1, 2, 3}, 1},
		{"TestFindIndexInSlice_EmptySlice", []int{}, -1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindIndexInSlice(c.input, func(v int) bool {
				return v == 2
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the index of input is not equal")
			}
		})
	}
}

func TestFindInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 2},
		{"TestFindInComparableSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, actual := collection.FindInComparableSlice(c.input, 2)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}

func TestFindIndexInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindIndexInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 1},
		{"TestFindIndexInComparableSlice_EmptySlice", []int{}, -1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindIndexInComparableSlice(c.input, 2)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the index of input is not equal")
			}
		})
	}
}

func TestFindMinimumInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindMinimumInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 1},
		{"TestFindMinimumInComparableSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinimumInComparableSlice(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

func TestFindMinimumInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindMinimumInSlice_NonEmptySlice", []int{1, 2, 3}, 1},
		{"TestFindMinimumInSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinimumInSlice(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

func TestFindMaximumInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindMaximumInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 3},
		{"TestFindMaximumInComparableSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaximumInComparableSlice(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

func TestFindMaximumInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestFindMaximumInSlice_NonEmptySlice", []int{1, 2, 3}, 3},
		{"TestFindMaximumInSlice_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaximumInSlice(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

func TestFindMin2MaxInComparableSlice(t *testing.T) {
	var cases = []struct {
		name        string
		input       []int
		expectedMin int
		expectedMax int
	}{
		{"TestFindMin2MaxInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 1, 3},
		{"TestFindMin2MaxInComparableSlice_EmptySlice", []int{}, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxInComparableSlice(c.input)
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}

func TestFindMin2MaxInSlice(t *testing.T) {
	var cases = []struct {
		name        string
		input       []int
		expectedMin int
		expectedMax int
	}{
		{"TestFindMin2MaxInSlice_NonEmptySlice", []int{1, 2, 3}, 1, 3},
		{"TestFindMin2MaxInSlice_EmptySlice", []int{}, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxInSlice(c.input, func(v int) int {
				return v
			})
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}

func TestFindMinFromComparableMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{
		{"TestFindMinFromComparableMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1},
		{"TestFindMinFromComparableMap_EmptyMap", map[int]int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinFromComparableMap(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

func TestFindMinFromMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{
		{"TestFindMinFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1},
		{"TestFindMinFromMap_EmptyMap", map[int]int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinFromMap(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

func TestFindMaxFromComparableMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{
		{"TestFindMaxFromComparableMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 3},
		{"TestFindMaxFromComparableMap_EmptyMap", map[int]int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaxFromComparableMap(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

func TestFindMaxFromMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{
		{"TestFindMaxFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 3},
		{"TestFindMaxFromMap_EmptyMap", map[int]int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaxFromMap(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

func TestFindMin2MaxFromComparableMap(t *testing.T) {
	var cases = []struct {
		name        string
		input       map[int]int
		expectedMin int
		expectedMax int
	}{
		{"TestFindMin2MaxFromComparableMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, 3},
		{"TestFindMin2MaxFromComparableMap_EmptyMap", map[int]int{}, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxFromComparableMap(c.input)
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}

func TestFindMin2MaxFromMap(t *testing.T) {
	var cases = []struct {
		name        string
		input       map[int]int
		expectedMin int
		expectedMax int
	}{
		{"TestFindMin2MaxFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, 3},
		{"TestFindMin2MaxFromMap_EmptyMap", map[int]int{}, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxFromMap(c.input)
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}
