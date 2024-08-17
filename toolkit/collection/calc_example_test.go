package collection_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/collection"
)

func ExampleSliceSum() {
	type Score struct {
		Name  string
		Score int64
	}

	var scores = []Score{
		{"A", 100},
		{"B", 200},
		{"C", 300},
		{"D", 400},
		{"E", 500},
	}

	sum := collection.SliceSum(scores, func(i int, value Score) int64 {
		return value.Score
	})
	fmt.Println(sum)

	// Output: 1500
}

func ExampleMapSum() {
	type Score struct {
		Name  string
		Score int64
	}

	var scores = map[string]Score{
		"A": {"A", 100},
		"B": {"B", 200},
		"C": {"C", 300},
		"D": {"D", 400},
		"E": {"E", 500},
	}

	sum := collection.MapSum(scores, func(key string, value Score) int64 {
		return value.Score
	})
	fmt.Println(sum)

	// Output: 1500
}
