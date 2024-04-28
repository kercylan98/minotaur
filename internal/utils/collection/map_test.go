package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"testing"
)

func TestMappingFromSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"TestMappingFromSlice_NonEmptySlice", []int{1, 2, 3}, []int{2, 3, 4}},
		{"TestMappingFromSlice_EmptySlice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MappingFromSlice[[]int, []int](c.input, func(value int) int {
				return value + 1
			})
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

func TestMappingFromMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]int
	}{
		{"TestMappingFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 2, 2: 3, 3: 4}},
		{"TestMappingFromMap_EmptyMap", map[int]int{}, map[int]int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MappingFromMap[map[int]int, map[int]int](c.input, func(value int) int {
				return value + 1
			})
			if len(result) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
			for k, v := range result {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the value of input is not equal")
				}
			}
		})
	}
}
