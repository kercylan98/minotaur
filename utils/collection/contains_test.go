package collection_test

import (
	"github.com/kercylan98/minotaur/utils/collection"
	"testing"
)

var intComparisonHandler = func(source, target int) bool {
	return source == target
}

func TestInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   int
		expected bool
	}{
		{"TestInSlice_NonEmptySliceIn", []int{1, 2, 3}, 1, true},
		{"TestInSlice_NonEmptySliceNotIn", []int{1, 2, 3}, 4, false},
		{"TestInSlice_EmptySlice", []int{}, 1, false},
		{"TestInSlice_NilSlice", nil, 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InSlice(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   int
		expected bool
	}{
		{"TestInComparableSlice_NonEmptySliceIn", []int{1, 2, 3}, 1, true},
		{"TestInComparableSlice_NonEmptySliceNotIn", []int{1, 2, 3}, 4, false},
		{"TestInComparableSlice_EmptySlice", []int{}, 1, false},
		{"TestInComparableSlice_NilSlice", nil, 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InComparableSlice(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAllInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{
		{"TestAllInSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true},
		{"TestAllInSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, false},
		{"TestAllInSlice_EmptySlice", []int{}, []int{1, 2}, false},
		{"TestAllInSlice_NilSlice", nil, []int{1, 2}, false},
		{"TestAllInSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, true},
		{"TestAllInSlice_NilValueSlice", []int{1, 2, 3}, nil, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInSlice(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAllInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{
		{"TestAllInComparableSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true},
		{"TestAllInComparableSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, false},
		{"TestAllInComparableSlice_EmptySlice", []int{}, []int{1, 2}, false},
		{"TestAllInComparableSlice_NilSlice", nil, []int{1, 2}, false},
		{"TestAllInComparableSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, true},
		{"TestAllInComparableSlice_NilValueSlice", []int{1, 2, 3}, nil, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInComparableSlice(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestAnyInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{
		{"TestAnyInSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true},
		{"TestAnyInSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, true},
		{"TestAnyInSlice_EmptySlice", []int{}, []int{1, 2}, false},
		{"TestAnyInSlice_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyInSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, false},
		{"TestAnyInSlice_NilValueSlice", []int{1, 2, 3}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInSlice(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAnyInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{
		{"TestAnyInComparableSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true},
		{"TestAnyInComparableSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, true},
		{"TestAnyInComparableSlice_EmptySlice", []int{}, []int{1, 2}, false},
		{"TestAnyInComparableSlice_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyInComparableSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, false},
		{"TestAnyInComparableSlice_NilValueSlice", []int{1, 2, 3}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInComparableSlice(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestInSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		inputV   int
		expected bool
	}{
		{"TestInSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, 1, true},
		{"TestInSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false},
		{"TestInSlices_EmptySlice", [][]int{{}, {}}, 1, false},
		{"TestInSlices_NilSlice", nil, 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InSlices(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestInComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		inputV   int
		expected bool
	}{
		{"TestInComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, 1, true},
		{"TestInComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false},
		{"TestInComparableSlices_EmptySlice", [][]int{{}, {}}, 1, false},
		{"TestInComparableSlices_NilSlice", nil, 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InComparableSlices(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

func TestAllInSlices(t *testing.T) {
	var cases = []struct {
		name        string
		input       [][]int
		inputValues []int
		expected    bool
	}{
		{"TestAllInSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true},
		{"TestAllInSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false},
		{"TestAllInSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false},
		{"TestAllInSlices_NilSlice", nil, []int{1, 2}, false},
		{"TestAllInSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, true},
		{"TestAllInSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInSlices(c.input, c.inputValues, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAllInComparableSlices(t *testing.T) {
	var cases = []struct {
		name        string
		input       [][]int
		inputValues []int
		expected    bool
	}{
		{"TestAllInComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true},
		{"TestAllInComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false},
		{"TestAllInComparableSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false},
		{"TestAllInComparableSlices_NilSlice", nil, []int{1, 2}, false},
		{"TestAllInComparableSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, true},
		{"TestAllInComparableSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInComparableSlices(c.input, c.inputValues)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

func TestAnyInSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{
		{"TestAnyInSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true},
		{"TestAnyInSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, true},
		{"TestAnyInSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false},
		{"TestAnyInSlices_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyInSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false},
		{"TestAnyInSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInSlices(c.slices, c.values, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAnyInComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{
		{"TestAnyInComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true},
		{"TestAnyInComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, true},
		{"TestAnyInComparableSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false},
		{"TestAnyInComparableSlices_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyInComparableSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false},
		{"TestAnyInComparableSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInComparableSlices(c.slices, c.values)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

func TestInAllSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		value    int
		expected bool
	}{
		{"TestInAllSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, 1, true},
		{"TestInAllSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false},
		{"TestInAllSlices_EmptySlice", [][]int{{}, {}}, 1, false},
		{"TestInAllSlices_NilSlice", nil, 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InAllSlices(c.slices, c.value, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestInAllComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		value    int
		expected bool
	}{
		{"TestInAllComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, 1, true},
		{"TestInAllComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false},
		{"TestInAllComparableSlices_EmptySlice", [][]int{{}, {}}, 1, false},
		{"TestInAllComparableSlices_NilSlice", nil, 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InAllComparableSlices(c.slices, c.value)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

func TestAnyInAllSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{
		{"TestAnyInAllSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, []int{1, 2}, true},
		{"TestAnyInAllSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false},
		{"TestAnyInAllSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false},
		{"TestAnyInAllSlices_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyInAllSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false},
		{"TestAnyInAllSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInAllSlices(c.slices, c.values, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAnyInAllComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{
		{"TestAnyInAllComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, []int{1, 2}, true},
		{"TestAnyInAllComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false},
		{"TestAnyInAllComparableSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false},
		{"TestAnyInAllComparableSlices_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyInAllComparableSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false},
		{"TestAnyInAllComparableSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInAllComparableSlices(c.slices, c.values)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

func TestKeyInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		key      int
		expected bool
	}{
		{"TestKeyInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, 1, true},
		{"TestKeyInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, 3, false},
		{"TestKeyInMap_EmptySlice", map[int]int{}, 1, false},
		{"TestKeyInMap_NilSlice", nil, 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.KeyInMap(c.m, c.key)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestValueInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		value    int
		handler  collection.ComparisonHandler[int]
		expected bool
	}{
		{"TestValueInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, 1, intComparisonHandler, true},
		{"TestValueInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, 3, intComparisonHandler, false},
		{"TestValueInMap_EmptySlice", map[int]int{}, 1, intComparisonHandler, false},
		{"TestValueInMap_NilSlice", nil, 1, intComparisonHandler, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.ValueInMap(c.m, c.value, c.handler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAllKeyInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		keys     []int
		expected bool
	}{
		{"TestAllKeyInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true},
		{"TestAllKeyInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, false},
		{"TestAllKeyInMap_EmptySlice", map[int]int{}, []int{1, 2}, false},
		{"TestAllKeyInMap_NilSlice", nil, []int{1, 2}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllKeyInMap(c.m, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAllValueInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		values   []int
		expected bool
	}{
		{"TestAllValueInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true},
		{"TestAllValueInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, false},
		{"TestAllValueInMap_EmptySlice", map[int]int{}, []int{1, 2}, false},
		{"TestAllValueInMap_NilSlice", nil, []int{1, 2}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllValueInMap(c.m, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

func TestAnyKeyInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		keys     []int
		expected bool
	}{
		{"TestAnyKeyInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true},
		{"TestAnyKeyInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, true},
		{"TestAnyKeyInMap_EmptySlice", map[int]int{}, []int{1, 2}, false},
		{"TestAnyKeyInMap_NilSlice", nil, []int{1, 2}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyKeyInMap(c.m, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestAnyValueInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		values   []int
		expected bool
	}{
		{"TestAnyValueInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true},
		{"TestAnyValueInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, true},
		{"TestAnyValueInMap_EmptySlice", map[int]int{}, []int{1, 2}, false},
		{"TestAnyValueInMap_NilSlice", nil, []int{1, 2}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyValueInMap(c.m, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestAllKeyInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		keys     []int
		expected bool
	}{
		{"TestAllKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false},
		{"TestAllKeyInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, false},
		{"TestAllKeyInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false},
		{"TestAllKeyInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false},
		{"TestAllKeyInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false},
		{"TestAllKeyInMaps_NilSlice", nil, []int{1, 2}, false},
		{"TestAllKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllKeyInMaps(c.maps, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestAllValueInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		values   []int
		expected bool
	}{
		{"TestAllValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false},
		{"TestAllValueInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, false},
		{"TestAllValueInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false},
		{"TestAllValueInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false},
		{"TestAllValueInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false},
		{"TestAllValueInMaps_NilSlice", nil, []int{1, 2}, false},
		{"TestAllValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllValueInMaps(c.maps, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestAnyKeyInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		keys     []int
		expected bool
	}{
		{"TestAnyKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, true},
		{"TestAnyKeyInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, true},
		{"TestAnyKeyInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, true},
		{"TestAnyKeyInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false},
		{"TestAnyKeyInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false},
		{"TestAnyKeyInMaps_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyKeyInMaps(c.maps, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestAnyValueInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		values   []int
		expected bool
	}{
		{"TestAnyValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false},
		{"TestAnyValueInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, true},
		{"TestAnyValueInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false},
		{"TestAnyValueInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false},
		{"TestAnyValueInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false},
		{"TestAnyValueInMaps_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyValueInMaps(c.maps, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestKeyInAllMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		key      int
		expected bool
	}{
		{"TestKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, 1, false},
		{"TestKeyInAllMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, 3, false},
		{"TestKeyInAllMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, 1, false},
		{"TestKeyInAllMaps_NilSlice", []map[int]int{{}, {}}, 1, false},
		{"TestKeyInAllMaps_EmptySlice", []map[int]int{}, 1, false},
		{"TestKeyInAllMaps_NilSlice", nil, 1, false},
		{"TestKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, 1, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.KeyInAllMaps(c.maps, c.key)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

func TestAnyKeyInAllMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		keys     []int
		expected bool
	}{
		{"TestAnyKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false},
		{"TestAnyKeyInAllMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, true},
		{"TestAnyKeyInAllMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false},
		{"TestAnyKeyInAllMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false},
		{"TestAnyKeyInAllMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false},
		{"TestAnyKeyInAllMaps_NilSlice", nil, []int{1, 2}, false},
		{"TestAnyKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyKeyInAllMaps(c.maps, c.keys)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, actual)
			}
		})
	}
}
