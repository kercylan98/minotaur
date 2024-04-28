package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"testing"
)

func TestChooseRandomSliceElementRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestChooseRandomSliceElementRepeatN_NonEmptySlice", []int{1, 2, 3}, 3},
		{"TestChooseRandomSliceElementRepeatN_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomSliceElementRepeatN(c.input, 3)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomIndexRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{
		{"TestChooseRandomIndexRepeatN_NonEmptySlice", []int{1, 2, 3}, 3},
		{"TestChooseRandomIndexRepeatN_EmptySlice", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomIndexRepeatN(c.input, 3)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomSliceElement(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{
		{"TestChooseRandomSliceElement_NonEmptySlice", []int{1, 2, 3}, map[int]bool{1: true, 2: true, 3: true}},
		{"TestChooseRandomSliceElement_EmptySlice", []int{}, map[int]bool{0: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomSliceElement(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomIndex(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{
		{"TestChooseRandomIndex_NonEmptySlice", []int{1, 2, 3}, map[int]bool{0: true, 1: true, 2: true}},
		{"TestChooseRandomIndex_EmptySlice", []int{}, map[int]bool{-1: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomIndex(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomSliceElementN(t *testing.T) {
	var cases = []struct {
		name  string
		input []int
	}{
		{"TestChooseRandomSliceElementN_NonEmptySlice", []int{1, 2, 3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ChooseRandomSliceElementN(c.input, 3)
			if !collection.AllInComparableSlice(actual, c.input) {
				t.Fatalf("%s failed, actual: %v, error: %s", c.name, actual, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomIndexN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{
		{"TestChooseRandomIndexN_NonEmptySlice", []int{1, 2, 3}, map[int]bool{0: true, 1: true, 2: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomIndexN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapKeyRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapKeyRepeatN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapKeyRepeatN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapValueRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapValueRepeatN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapValueRepeatN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapKeyAndValueRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapKeyAndValueRepeatN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapKeyAndValueRepeatN(c.input, 3)
			if !c.expected[result[1]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapKey(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapKey_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
		{"TestChooseRandomMapKey_EmptyMap", map[int]int{}, map[int]bool{0: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapKey(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapValue(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapValue_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
		{"TestChooseRandomMapValue_EmptyMap", map[int]int{}, map[int]bool{0: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapValue(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapValueN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapValueN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapValueN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapKeyAndValue(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapKeyAndValue_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
		{"TestChooseRandomMapKeyAndValue_EmptyMap", map[int]int{}, map[int]bool{0: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k, v := collection.ChooseRandomMapKeyAndValue(c.input)
			if !c.expected[k] || !c.expected[v] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, k, "the length of input is not equal")
			}
		})
	}
}

func TestChooseRandomMapKeyAndValueN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{"TestChooseRandomMapKeyAndValueN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			kvm := collection.ChooseRandomMapKeyAndValueN(c.input, 1)
			for k := range kvm {
				if !c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, k, "the length of input is not equal")
				}
			}
		})
	}
}
