package collection_test

import (
	"github.com/kercylan98/minotaur/toolkit/collection"
	"testing"
)

func TestSliceSum(t *testing.T) {
	var tests = []struct {
		name string
		args []int
		want int
	}{
		{name: "test1", args: []int{1, 2, 3, 4, 5}, want: 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collection.SliceSum(tt.args, func(i int, value int) int {
				return value
			}); got != tt.want {
				t.Errorf("SliceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapSum(t *testing.T) {
	var tests = []struct {
		name string
		args map[int]int
		want int
	}{
		{name: "test1", args: map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}, want: 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collection.MapSum(tt.args, func(key int, value int) int {
				return value
			}); got != tt.want {
				t.Errorf("MapSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
