package collection_test

import (
	"github.com/kercylan98/minotaur/toolkit/collection"
	"reflect"
	"testing"
)

func TestConvertSliceToBatches(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		batch    int
		expected [][]int
	}{
		{name: "TestConvertSliceToBatches_NonEmpty", input: []int{1, 2, 3}, batch: 2, expected: [][]int{{1, 2}, {3}}},
		{name: "TestConvertSliceToBatches_Empty", input: []int{}, batch: 2, expected: nil},
		{name: "TestConvertSliceToBatches_Nil", input: nil, batch: 2, expected: nil},
		{name: "TestConvertSliceToBatches_NonPositive", input: []int{1, 2, 3}, batch: 0, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToBatches(c.input, c.batch)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for i := 0; i < len(actual); i++ {
				av, ev := actual[i], c.expected[i]
				if len(av) != len(ev) {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
				for j := 0; j < len(av); j++ {
					aj, ej := av[j], ev[j]
					if reflect.TypeOf(aj).Kind() != reflect.TypeOf(ej).Kind() {
						t.Errorf("expected: %v, actual: %v", c.expected, actual)
					}
					if aj != ej {
						t.Errorf("expected: %v, actual: %v", c.expected, actual)
					}
				}
			}
		})
	}
}

func TestConvertMapKeysToBatches(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		batch    int
		expected [][]int
	}{
		{name: "TestConvertMapKeysToBatches_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, batch: 2, expected: [][]int{{1, 2}, {3}}},
		{name: "TestConvertMapKeysToBatches_Empty", input: map[int]int{}, batch: 2, expected: nil},
		{name: "TestConvertMapKeysToBatches_Nil", input: nil, batch: 2, expected: nil},
		{name: "TestConvertMapKeysToBatches_NonPositive", input: map[int]int{1: 1, 2: 2, 3: 3}, batch: 0, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapKeysToBatches(c.input, c.batch)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestConvertMapValuesToBatches(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		batch    int
		expected [][]int
	}{
		{name: "TestConvertMapValuesToBatches_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, batch: 2, expected: [][]int{{1, 2}, {3}}},
		{name: "TestConvertMapValuesToBatches_Empty", input: map[int]int{}, batch: 2, expected: nil},
		{name: "TestConvertMapValuesToBatches_Nil", input: nil, batch: 2, expected: nil},
		{name: "TestConvertMapValuesToBatches_NonPositive", input: map[int]int{1: 1, 2: 2, 3: 3}, batch: 0, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapValuesToBatches(c.input, c.batch)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestConvertSliceToAny(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []interface{}
	}{
		{name: "TestConvertSliceToAny_NonEmpty", input: []int{1, 2, 3}, expected: []any{1, 2, 3}},
		{name: "TestConvertSliceToAny_Empty", input: []int{}, expected: []any{}},
		{name: "TestConvertSliceToAny_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToAny(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for i := 0; i < len(actual); i++ {
				av, ev := actual[i], c.expected[i]
				if reflect.TypeOf(av).Kind() != reflect.TypeOf(ev).Kind() {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestConvertSliceToIndexMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]int
	}{
		{name: "TestConvertSliceToIndexMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]int{0: 1, 1: 2, 2: 3}},
		{name: "TestConvertSliceToIndexMap_Empty", input: []int{}, expected: map[int]int{}},
		{name: "TestConvertSliceToIndexMap_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToIndexMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestConvertSliceToIndexOnlyMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]struct{}
	}{
		{name: "TestConvertSliceToIndexOnlyMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]struct{}{0: {}, 1: {}, 2: {}}},
		{name: "TestConvertSliceToIndexOnlyMap_Empty", input: []int{}, expected: map[int]struct{}{}},
		{name: "TestConvertSliceToIndexOnlyMap_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToIndexOnlyMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if _, ok := c.expected[k]; !ok {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
				if v != struct{}{} {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestConvertSliceToMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]struct{}
	}{
		{name: "TestConvertSliceToMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]struct{}{1: {}, 2: {}, 3: {}}},
		{name: "TestConvertSliceToMap_Empty", input: []int{}, expected: map[int]struct{}{}},
		{name: "TestConvertSliceToMap_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if _, ok := c.expected[k]; !ok {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
				if v != struct{}{} {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestConvertSliceToBoolMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{
		{name: "TestConvertSliceToBoolMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]bool{1: true, 2: true, 3: true}},
		{name: "TestConvertSliceToBoolMap_Empty", input: []int{}, expected: map[int]bool{}},
		{name: "TestConvertSliceToBoolMap_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToBoolMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestConvertMapKeysToSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected []int
	}{
		{name: "TestConvertMapKeysToSlice_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, expected: []int{1, 2, 3}},
		{name: "TestConvertMapKeysToSlice_Empty", input: map[int]int{}, expected: []int{}},
		{name: "TestConvertMapKeysToSlice_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapKeysToSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			var matchCount = 0
			for _, av := range actual {
				for _, ev := range c.expected {
					if av == ev {
						matchCount++
					}
				}
			}
			if matchCount != len(actual) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestConvertMapValuesToSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected []int
	}{
		{name: "TestConvertMapValuesToSlice_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, expected: []int{1, 2, 3}},
		{name: "TestConvertMapValuesToSlice_Empty", input: map[int]int{}, expected: []int{}},
		{name: "TestConvertMapValuesToSlice_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapValuesToSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			var matchCount = 0
			for _, av := range actual {
				for _, ev := range c.expected {
					if av == ev {
						matchCount++
					}
				}
			}
			if matchCount != len(actual) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestInvertMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]string
		expected map[string]int
	}{
		{name: "TestInvertMap_NonEmpty", input: map[int]string{1: "1", 2: "2", 3: "3"}, expected: map[string]int{"1": 1, "2": 2, "3": 3}},
		{name: "TestInvertMap_Empty", input: map[int]string{}, expected: map[string]int{}},
		{name: "TestInvertMap_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.InvertMap[map[int]string, map[string]int](c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestConvertMapValuesToBool(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{
		{name: "TestConvertMapValuesToBool_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, expected: map[int]bool{1: true, 2: true, 3: true}},
		{name: "TestConvertMapValuesToBool_Empty", input: map[int]int{}, expected: map[int]bool{}},
		{name: "TestConvertMapValuesToBool_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapValuesToBool[map[int]int, map[int]bool](c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

func TestReverseSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "TestReverseSlice_NonEmpty", input: []int{1, 2, 3}, expected: []int{3, 2, 1}},
		{name: "TestReverseSlice_Empty", input: []int{}, expected: []int{}},
		{name: "TestReverseSlice_Nil", input: nil, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.ReverseSlice(&c.input)
			if len(c.input) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, c.input)
			}
			for i := 0; i < len(c.input); i++ {
				av, ev := c.input[i], c.expected[i]
				if reflect.TypeOf(av).Kind() != reflect.TypeOf(ev).Kind() {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
			}
		})
	}
}
