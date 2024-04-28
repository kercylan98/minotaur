package collection_test

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"testing"
)

func TestLoopSlice(t *testing.T) {
	var cases = []struct {
		name       string
		in         []int
		out        []int
		breakIndex uint
	}{
		{"TestLoopSlice_Part", []int{1, 2, 3, 4, 5}, []int{1, 2}, 2},
		{"TestLoopSlice_All", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 0},
		{"TestLoopSlice_Empty", []int{}, []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []int
			collection.LoopSlice(c.in, func(i int, val int) bool {
				result = append(result, val)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopSlice(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestReverseLoopSlice(t *testing.T) {
	var cases = []struct {
		name       string
		in         []int
		out        []int
		breakIndex uint
	}{
		{"TestReverseLoopSlice_Part", []int{1, 2, 3, 4, 5}, []int{5, 4}, 2},
		{"TestReverseLoopSlice_All", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, 0},
		{"TestReverseLoopSlice_Empty", []int{}, []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []int
			collection.ReverseLoopSlice(c.in, func(i int, val int) bool {
				result = append(result, val)
				if c.breakIndex != 0 && uint(i) == uint(len(c.in))-c.breakIndex {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("ReverseLoopSlice(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMap(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        map[int]string
		breakIndex uint
	}{
		{"TestLoopMap_Part", map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2"}, 2},
		{"TestLoopMap_All", map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2", 3: "3"}, 0},
		{"TestLoopMap_Empty", map[int]string{}, map[int]string{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result = make(map[int]string)
			collection.LoopMap(c.in, func(i int, key int, val string) bool {
				result[key] = val
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableMap(result, c.out) {
				t.Errorf("LoopMap(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByOrderedKeyAsc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []int
		breakIndex uint
	}{
		{"TestLoopMapByOrderedKeyAsc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []int{1, 2}, 2},
		{"TestLoopMapByOrderedKeyAsc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []int{1, 2, 3}, 0},
		{"TestLoopMapByOrderedKeyAsc_Empty", map[int]string{}, []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []int
			collection.LoopMapByOrderedKeyAsc(c.in, func(i int, key int, val string) bool {
				result = append(result, key)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByOrderedKeyAsc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByOrderedKeyDesc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []int
		breakIndex uint
	}{
		{"TestLoopMapByOrderedKeyDesc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []int{3, 2}, 2},
		{"TestLoopMapByOrderedKeyDesc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []int{3, 2, 1}, 0},
		{"TestLoopMapByOrderedKeyDesc_Empty", map[int]string{}, []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []int
			collection.LoopMapByOrderedKeyDesc(c.in, func(i int, key int, val string) bool {
				result = append(result, key)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByOrderedKeyDesc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByOrderedValueAsc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []string
		breakIndex uint
	}{
		{"TestLoopMapByOrderedValueAsc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"1", "2"}, 2},
		{"TestLoopMapByOrderedValueAsc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"1", "2", "3"}, 0},
		{"TestLoopMapByOrderedValueAsc_Empty", map[int]string{}, []string{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []string
			collection.LoopMapByOrderedValueAsc(c.in, func(i int, key int, val string) bool {
				result = append(result, val)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByOrderedValueAsc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByOrderedValueDesc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []string
		breakIndex uint
	}{
		{"TestLoopMapByOrderedValueDesc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"3", "2"}, 2},
		{"TestLoopMapByOrderedValueDesc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"3", "2", "1"}, 0},
		{"TestLoopMapByOrderedValueDesc_Empty", map[int]string{}, []string{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []string
			collection.LoopMapByOrderedValueDesc(c.in, func(i int, key int, val string) bool {
				result = append(result, val)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByOrderedValueDesc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByKeyGetterAsc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []int
		breakIndex uint
	}{
		{"TestLoopMapByKeyGetterAsc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []int{1, 2}, 2},
		{"TestLoopMapByKeyGetterAsc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []int{1, 2, 3}, 0},
		{"TestLoopMapByKeyGetterAsc_Empty", map[int]string{}, []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []int
			collection.LoopMapByKeyGetterAsc(c.in, func(key int) int {
				return key
			}, func(i int, key int, val string) bool {
				result = append(result, key)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByKeyGetterAsc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByKeyGetterDesc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []int
		breakIndex uint
	}{
		{"TestLoopMapByKeyGetterDesc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []int{3, 2}, 2},
		{"TestLoopMapByKeyGetterDesc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []int{3, 2, 1}, 0},
		{"TestLoopMapByKeyGetterDesc_Empty", map[int]string{}, []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []int
			collection.LoopMapByKeyGetterDesc(c.in, func(key int) int {
				return key
			}, func(i int, key int, val string) bool {
				result = append(result, key)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByKeyGetterDesc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByValueGetterAsc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []string
		breakIndex uint
	}{
		{"TestLoopMapByValueGetterAsc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"1", "2"}, 2},
		{"TestLoopMapByValueGetterAsc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"1", "2", "3"}, 0},
		{"TestLoopMapByValueGetterAsc_Empty", map[int]string{}, []string{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []string
			collection.LoopMapByValueGetterAsc(c.in, func(val string) string {
				return val
			}, func(i int, key int, val string) bool {
				result = append(result, val)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByValueGetterAsc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}

func TestLoopMapByValueGetterDesc(t *testing.T) {
	var cases = []struct {
		name       string
		in         map[int]string
		out        []string
		breakIndex uint
	}{
		{"TestLoopMapByValueGetterDesc_Part", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"3", "2"}, 2},
		{"TestLoopMapByValueGetterDesc_All", map[int]string{1: "1", 2: "2", 3: "3"}, []string{"3", "2", "1"}, 0},
		{"TestLoopMapByValueGetterDesc_Empty", map[int]string{}, []string{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []string
			collection.LoopMapByValueGetterDesc(c.in, func(val string) string {
				return val
			}, func(i int, key int, val string) bool {
				result = append(result, val)
				if c.breakIndex != 0 && uint(i) == c.breakIndex-1 {
					return false
				}
				return true
			})
			if !collection.EqualComparableSlice(result, c.out) {
				t.Errorf("LoopMapByValueGetterDesc(%v) got %v, want %v", c.in, result, c.out)
			}
		})
	}
}
