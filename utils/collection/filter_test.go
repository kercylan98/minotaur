package collection_test

import (
	"github.com/kercylan98/minotaur/utils/collection"
	"testing"
)

func TestFilterOutByIndices(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		indices  []int
		expected []int
	}{
		{"TestFilterOutByIndices_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{1, 3}, []int{1, 3, 5}},
		{"TestFilterOutByIndices_EmptySlice", []int{}, []int{1, 3}, []int{}},
		{"TestFilterOutByIndices_NilSlice", nil, []int{1, 3}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByIndices(c.input, c.indices...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of input is not equal")
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of input is not equal")
				}
			}
		})
	}
}

func TestFilterOutByCondition(t *testing.T) {
	var cases = []struct {
		name      string
		input     []int
		condition func(int) bool
		expected  []int
	}{
		{"TestFilterOutByCondition_NonEmptySlice", []int{1, 2, 3, 4, 5}, func(v int) bool {
			return v%2 == 0
		}, []int{1, 3, 5}},
		{"TestFilterOutByCondition_EmptySlice", []int{}, func(v int) bool {
			return v%2 == 0
		}, []int{}},
		{"TestFilterOutByCondition_NilSlice", nil, func(v int) bool {
			return v%2 == 0
		}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByCondition(c.input, c.condition)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of input is not equal")
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of input is not equal")
				}
			}
		})
	}
}

func TestFilterOutByKey(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		key      int
		expected map[int]int
	}{
		{"TestFilterOutByKey_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, map[int]int{2: 2, 3: 3}},
		{"TestFilterOutByKey_EmptyMap", map[int]int{}, 1, map[int]int{}},
		{"TestFilterOutByKey_NilMap", nil, 1, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByKey(c.input, c.key)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

func TestFilterOutByValue(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		value    int
		expected map[int]int
	}{
		{"TestFilterOutByValue_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, map[int]int{2: 2, 3: 3}},
		{"TestFilterOutByValue_EmptyMap", map[int]int{}, 1, map[int]int{}},
		{"TestFilterOutByValue_NilMap", nil, 1, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByValue(c.input, c.value, func(source, target int) bool {
				return source == target
			})
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

func TestFilterOutByKeys(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		keys     []int
		expected map[int]int
	}{
		{"TestFilterOutByKeys_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, []int{1, 3}, map[int]int{2: 2}},
		{"TestFilterOutByKeys_EmptyMap", map[int]int{}, []int{1, 3}, map[int]int{}},
		{"TestFilterOutByKeys_NilMap", nil, []int{1, 3}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByKeys(c.input, c.keys...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

func TestFilterOutByValues(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		values   []int
		expected map[int]int
	}{
		{"TestFilterOutByValues_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, []int{1, 3}, map[int]int{2: 2}},
		{"TestFilterOutByValues_EmptyMap", map[int]int{}, []int{1, 3}, map[int]int{}},
		{"TestFilterOutByValues_NilMap", nil, []int{1, 3}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByValues(c.input, c.values, func(source, target int) bool {
				return source == target
			})
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

func TestFilterOutByMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		filter   map[int]int
		expected map[int]int
	}{
		{"TestFilterOutByMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 3: 3}, map[int]int{2: 2}},
		{"TestFilterOutByMap_EmptyMap", map[int]int{}, map[int]int{1: 1, 3: 3}, map[int]int{}},
		{"TestFilterOutByMap_NilMap", nil, map[int]int{1: 1, 3: 3}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByMap(c.input, func(k int, v int) bool {
				return c.filter[k] == v
			})
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}
