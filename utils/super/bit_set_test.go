package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestNewBitSet(t *testing.T) {
	var cases = []struct {
		name        string
		in          []int
		shouldPanic bool
	}{
		{name: "normal", in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "empty", in: make([]int, 0)},
		{name: "nil", in: nil},
		{name: "negative", in: []int{-1, -2}, shouldPanic: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !c.shouldPanic {
					t.Fatalf("panic: %v", r)
				}
			}()
			bs := super.NewBitSet(c.in...)
			t.Log(bs)
		})
	}
}

func TestBitSet_Set(t *testing.T) {
	var cases = []struct {
		name        string
		in          []int
		shouldPanic bool
	}{
		{name: "normal", in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "empty", in: make([]int, 0)},
		{name: "nil", in: nil},
		{name: "negative", in: []int{-1, -2}, shouldPanic: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !c.shouldPanic {
					t.Fatalf("panic: %v", r)
				}
			}()
			bs := super.NewBitSet[int]()
			for _, bit := range c.in {
				bs.Set(bit)
			}
			for _, bit := range c.in {
				if !bs.Has(bit) {
					t.Fatalf("bit %v not set", bit)
				}
			}
		})
	}
}

func TestBitSet_Del(t *testing.T) {
	var cases = []struct {
		name        string
		in          []int
		shouldPanic bool
	}{
		{name: "normal", in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "empty", in: make([]int, 0)},
		{name: "nil", in: nil},
		{name: "negative", in: []int{-1, -2}, shouldPanic: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !c.shouldPanic {
					t.Fatalf("panic: %v", r)
				}
			}()
			bs := super.NewBitSet[int]()
			for _, bit := range c.in {
				bs.Set(bit)
			}
			for _, bit := range c.in {
				bs.Del(bit)
			}
			for _, bit := range c.in {
				if bs.Has(bit) {
					t.Fatalf("bit %v not del", bit)
				}
			}
		})
	}
}

func TestBitSet_Shrink(t *testing.T) {
	var cases = []struct {
		name string
		in   []int
	}{
		{name: "normal", in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "empty", in: make([]int, 0)},
		{name: "nil", in: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bs := super.NewBitSet(c.in...)
			for _, v := range c.in {
				bs.Del(v)
			}
			bs.Shrink()
			if bs.Cap() != 0 {
				t.Fatalf("cap %v != 0", bs.Cap())
			}
		})
	}
}
