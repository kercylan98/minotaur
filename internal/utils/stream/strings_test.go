package stream_test

import (
	"github.com/kercylan98/minotaur/utils/stream"
	"testing"
)

func TestNewStrings(t *testing.T) {
	var cases = []struct {
		name string
		in   []string
		want []string
	}{
		{name: "empty", in: []string{}, want: []string{}},
		{name: "one", in: []string{"a"}, want: []string{"a"}},
		{name: "two", in: []string{"a", "b"}, want: []string{"a", "b"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...)
			if got.Len() != len(c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestStrings_Elem(t *testing.T) {
	var cases = []struct {
		name string
		in   []string
		want []string
	}{
		{name: "empty", in: []string{}, want: []string{}},
		{name: "one", in: []string{"a"}, want: []string{"a"}},
		{name: "two", in: []string{"a", "b"}, want: []string{"a", "b"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...).Elem()
			if len(got) != len(c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestStrings_Len(t *testing.T) {
	var cases = []struct {
		name string
		in   []string
		want int
	}{
		{name: "empty", in: []string{}, want: 0},
		{name: "one", in: []string{"a"}, want: 1},
		{name: "two", in: []string{"a", "b"}, want: 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...)
			if got.Len() != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestStrings_Append(t *testing.T) {
	var cases = []struct {
		name   string
		in     []string
		append []string
		want   []string
	}{
		{name: "empty", in: []string{}, append: []string{"a"}, want: []string{"a"}},
		{name: "one", in: []string{"a"}, append: []string{"b"}, want: []string{"a", "b"}},
		{name: "two", in: []string{"a", "b"}, append: []string{"c"}, want: []string{"a", "b", "c"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...).Append(c.append...)
			if got.Len() != len(c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}
