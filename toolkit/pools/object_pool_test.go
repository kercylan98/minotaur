package pools_test

import (
	"github.com/kercylan98/minotaur/toolkit/pools"
	"testing"
)

func TestNewObjectPool(t *testing.T) {
	var cases = []struct {
		name        string
		generator   func() *map[string]int
		releaser    func(data *map[string]int)
		shouldPanic bool
	}{
		{name: "TestNewObjectPool_NilGenerator", generator: nil, releaser: func(data *map[string]int) {}, shouldPanic: true},
		{name: "TestNewObjectPool_NilReleaser", generator: func() *map[string]int { return &map[string]int{} }, releaser: nil, shouldPanic: true},
		{name: "TestNewObjectPool_NilGeneratorAndReleaser", generator: nil, releaser: nil, shouldPanic: true},
		{name: "TestNewObjectPool_Normal", generator: func() *map[string]int { return &map[string]int{} }, releaser: func(data *map[string]int) {}, shouldPanic: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if err := recover(); c.shouldPanic && err == nil {
					t.Error("TestNewObjectPool should panic")
				}
			}()
			_ = pools.NewObjectPool[map[string]int](c.generator, c.releaser)
		})
	}
}

func TestObjectPool_Get(t *testing.T) {
	var cases = []struct {
		name      string
		generator func() *map[string]int
		releaser  func(data *map[string]int)
	}{
		{
			name: "TestObjectPool_Get_Normal",
			generator: func() *map[string]int {
				return &map[string]int{}
			},
			releaser: func(data *map[string]int) {
				for k := range *data {
					delete(*data, k)
				}
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pool := pools.NewObjectPool[map[string]int](c.generator, c.releaser)
			if actual := pool.Get(); len(*actual) != 0 {
				t.Error("TestObjectPool_Get failed")
			}
		})
	}
}

func TestObjectPool_Release(t *testing.T) {
	var cases = []struct {
		name      string
		generator func() *map[string]int
		releaser  func(data *map[string]int)
	}{
		{
			name: "TestObjectPool_Release_Normal",
			generator: func() *map[string]int {
				return &map[string]int{}
			},
			releaser: func(data *map[string]int) {
				for k := range *data {
					delete(*data, k)
				}
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pool := pools.NewObjectPool[map[string]int](c.generator, c.releaser)
			msg := pool.Get()
			m := *msg
			m["test"] = 1
			pool.Put(msg)
			if len(m) != 0 {
				t.Error("TestObjectPool_Release failed")
			}
		})
	}
}
