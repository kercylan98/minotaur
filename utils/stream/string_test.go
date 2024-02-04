package stream_test

import (
	"github.com/kercylan98/minotaur/utils/stream"
	"testing"
)

func TestNewString(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "hello"},
		{name: "case2", in: "world", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in)
			if got.String() != c.want {
				t.Fatalf("NewString(%s) = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_String(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "hello"},
		{name: "case2", in: "world", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).String()
			if got != c.want {
				t.Fatalf("String(%s).String() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Index(t *testing.T) {
	var cases = []struct {
		name        string
		in          string
		i           int
		want        string
		shouldPanic bool
	}{
		{name: "case1", in: "hello", i: 0, want: "h", shouldPanic: false},
		{name: "case2", in: "world", i: 2, want: "r", shouldPanic: false},
		{name: "case3", in: "world", i: 5, want: "", shouldPanic: true},
		{name: "case4", in: "world", i: -1, want: "", shouldPanic: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != c.shouldPanic {
					t.Fatalf("NewString(%s).Index(%d) should panic", c.in, c.i)
				}
			}()
			got := stream.NewString(c.in).Index(c.i)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Index(%d) = %s; want %s", c.in, c.i, got, c.want)
			}
		})
	}
}

func TestString_Range(t *testing.T) {
	var cases = []struct {
		name        string
		in          string
		start       int
		end         int
		want        string
		shouldPanic bool
	}{
		{name: "case1", in: "hello", start: 0, end: 2, want: "he", shouldPanic: false},
		{name: "case2", in: "world", start: 2, end: 5, want: "rld", shouldPanic: false},
		{name: "case3", in: "world", start: 5, end: 6, want: "", shouldPanic: true},
		{name: "case4", in: "world", start: -1, end: 6, want: "", shouldPanic: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != c.shouldPanic {
					t.Fatalf("NewString(%s).Range(%d, %d) should panic", c.in, c.start, c.end)
				}
			}()
			got := stream.NewString(c.in).Range(c.start, c.end)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Range(%d, %d) = %s; want %s", c.in, c.start, c.end, got, c.want)
			}
		})
	}
}

func TestString_TrimSpace(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: " hello ", want: "hello"},
		{name: "case2", in: " world ", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TrimSpace()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TrimSpace() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Trim(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		cs   string
		want string
	}{
		{name: "case1", in: "hello", cs: "h", want: "ello"},
		{name: "case2", in: "world", cs: "d", want: "worl"},
		{name: "none", in: "world", cs: "", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Trim(c.cs)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Trim(%s) = %s; want %s", c.in, c.cs, got, c.want)
			}
		})
	}
}

func TestString_TrimPrefix(t *testing.T) {
	var cases = []struct {
		name    string
		in      string
		prefix  string
		want    string
		isEqual bool
	}{
		{name: "case1", in: "hello", prefix: "h", want: "ello", isEqual: false},
		{name: "case2", in: "world", prefix: "w", want: "orld", isEqual: false},
		{name: "none", in: "world", prefix: "x", want: "world", isEqual: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TrimPrefix(c.prefix)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TrimPrefix(%s) = %s; want %s", c.in, c.prefix, got, c.want)
			}
		})
	}
}

func TestString_TrimSuffix(t *testing.T) {
	var cases = []struct {
		name    string
		in      string
		suffix  string
		want    string
		isEqual bool
	}{
		{name: "case1", in: "hello", suffix: "o", want: "hell", isEqual: false},
		{name: "case2", in: "world", suffix: "d", want: "worl", isEqual: false},
		{name: "none", in: "world", suffix: "x", want: "world", isEqual: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TrimSuffix(c.suffix)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TrimSuffix(%s) = %s; want %s", c.in, c.suffix, got, c.want)
			}
		})
	}
}

func TestString_ToUpper(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "HELLO"},
		{name: "case2", in: "world", want: "WORLD"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).ToUpper()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).ToUpper() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_ToLower(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "HELLO", want: "hello"},
		{name: "case2", in: "WORLD", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).ToLower()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).ToLower() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Equal(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want bool
	}{
		{name: "case1", in: "hello", ss: "hello", want: true},
		{name: "case2", in: "world", ss: "world", want: true},
		{name: "case3", in: "world", ss: "worldx", want: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Equal(c.ss)
			if got != c.want {
				t.Fatalf("NewString(%s).Equal(%s) = %t; want %t", c.in, c.ss, got, c.want)
			}
		})
	}
}

func TestString_HasPrefix(t *testing.T) {
	var cases = []struct {
		name   string
		in     string
		prefix string
		want   bool
	}{
		{name: "case1", in: "hello", prefix: "h", want: true},
		{name: "case2", in: "world", prefix: "w", want: true},
		{name: "case3", in: "world", prefix: "x", want: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).HasPrefix(c.prefix)
			if got != c.want {
				t.Fatalf("NewString(%s).HasPrefix(%s) = %t; want %t", c.in, c.prefix, got, c.want)
			}
		})
	}
}

func TestString_Len(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want int
	}{
		{name: "case1", in: "hello", want: 5},
		{name: "case2", in: "world", want: 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Len()
			if got != c.want {
				t.Fatalf("NewString(%s).Len() = %d; want %d", c.in, got, c.want)
			}
		})
	}
}

func TestString_Contains(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want bool
	}{
		{name: "case1", in: "hello", ss: "he", want: true},
		{name: "case2", in: "world", ss: "or", want: true},
		{name: "case3", in: "world", ss: "x", want: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Contains(c.ss)
			if got != c.want {
				t.Fatalf("NewString(%s).Contains(%s) = %t; want %t", c.in, c.ss, got, c.want)
			}
		})
	}
}

func TestString_Count(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want int
	}{
		{name: "case1", in: "hello", ss: "l", want: 2},
		{name: "case2", in: "world", ss: "o", want: 1},
		{name: "case3", in: "world", ss: "x", want: 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Count(c.ss)
			if got != c.want {
				t.Fatalf("NewString(%s).Count(%s) = %d; want %d", c.in, c.ss, got, c.want)
			}
		})
	}
}

func TestString_Repeat(t *testing.T) {
	var cases = []struct {
		name  string
		in    string
		count int
		want  string
	}{
		{name: "case1", in: "hello", count: 2, want: "hellohello"},
		{name: "case2", in: "world", count: 3, want: "worldworldworld"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Repeat(c.count)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Repeat(%d) = %s; want %s", c.in, c.count, got, c.want)
			}
		})
	}
}

func TestString_Replace(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		old  string
		new  string
		want string
	}{
		{name: "case1", in: "hello", old: "l", new: "x", want: "hexxo"},
		{name: "case2", in: "world", old: "o", new: "x", want: "wxrld"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Replace(c.old, c.new, -1)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Replace(%s, %s) = %s; want %s", c.in, c.old, c.new, got, c.want)
			}
		})
	}
}

func TestString_ReplaceAll(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		old  string
		new  string
		want string
	}{
		{name: "case1", in: "hello", old: "l", new: "x", want: "hexxo"},
		{name: "case2", in: "world", old: "o", new: "x", want: "wxrld"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).ReplaceAll(c.old, c.new)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).ReplaceAll(%s, %s) = %s; want %s", c.in, c.old, c.new, got, c.want)
			}
		})
	}
}

func TestString_Append(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want string
	}{
		{name: "case1", in: "hello", ss: " world", want: "hello world"},
		{name: "case2", in: "world", ss: " hello", want: "world hello"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Append(c.ss)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Append(%s) = %s; want %s", c.in, c.ss, got, c.want)
			}
		})
	}
}

func TestString_Prepend(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want string
	}{
		{name: "case1", in: "hello", ss: "world ", want: "world hello"},
		{name: "case2", in: "world", ss: "hello ", want: "hello world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Prepend(c.ss)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Prepend(%s) = %s; want %s", c.in, c.ss, got, c.want)
			}
		})
	}
}

func TestString_Clear(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: ""},
		{name: "case2", in: "world", want: ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Clear()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Clear() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Reverse(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "olleh"},
		{name: "case2", in: "world", want: "dlrow"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Reverse()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Reverse() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Queto(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "\"hello\""},
		{name: "case2", in: "world", want: "\"world\""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Queto()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Queto() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_QuetoToASCII(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "\"hello\""},
		{name: "case2", in: "world", want: "\"world\""},
		{name: "case3", in: "你好", want: "\"\\u4f60\\u597d\""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).QuetoToASCII()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).QuetoToASCII() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_FirstUpper(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "Hello"},
		{name: "case2", in: "world", want: "World"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).FirstUpper()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).FirstUpper() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_FirstLower(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "Hello", want: "hello"},
		{name: "case2", in: "World", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).FirstLower()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).FirstLower() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_SnakeCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "HelloWorld", want: "hello_world"},
		{name: "case2", in: "HelloWorldHello", want: "hello_world_hello"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).SnakeCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).SnakeCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_CamelCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello_world", want: "helloWorld"},
		{name: "case2", in: "hello_world_hello", want: "helloWorldHello"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).CamelCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).CamelCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_KebabCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "HelloWorld", want: "hello-world"},
		{name: "case2", in: "HelloWorldHello", want: "hello-world-hello"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).KebabCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).KebabCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_TitleCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello_world", want: "HelloWorld"},
		{name: "case2", in: "hello_world_hello", want: "HelloWorldHello"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TitleCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TitleCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Bytes(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "hello"},
		{name: "case2", in: "world", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Bytes()
			if string(got) != c.want {
				t.Fatalf("NewString(%s).Bytes() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Runes(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "hello"},
		{name: "case2", in: "world", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Runes()
			if string(got) != c.want {
				t.Fatalf("NewString(%s).Runes() = %v; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Default(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "", want: "default"},
		{name: "case2", in: "world", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Default("default")
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Default() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Handle(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "hello"},
		{name: "case2", in: "world", want: "world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var w string
			got := stream.NewString(c.in).Handle(func(s string) {
				w = s
			})
			if w != c.want {
				t.Fatalf("NewString(%s).Handle() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Update(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "case1", in: "hello", want: "HELLO"},
		{name: "case2", in: "world", want: "WORLD"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Update(func(s string) string {
				return stream.NewString(s).ToUpper().String()
			})
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Update() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

func TestString_Split(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		sep  string
		want []string
	}{
		{name: "case1", in: "hello world", sep: " ", want: []string{"hello", "world"}},
		{name: "case2", in: "hello,world", sep: ",", want: []string{"hello", "world"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Split(c.sep)
			for i, v := range got {
				if v != c.want[i] {
					t.Fatalf("NewString(%s).Split(%s) = %v; want %v", c.in, c.sep, got, c.want)
				}
			}
		})
	}
}

func TestString_SplitN(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		sep  string
		n    int
		want []string
	}{
		{name: "case1", in: "hello world", sep: " ", n: 2, want: []string{"hello", "world"}},
		{name: "case2", in: "hello,world", sep: ",", n: 2, want: []string{"hello", "world"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).SplitN(c.sep, c.n)
			for i, v := range got {
				if v != c.want[i] {
					t.Fatalf("NewString(%s).SplitN(%s, %d) = %v; want %v", c.in, c.sep, c.n, got, c.want)
				}
			}
		})
	}
}

func TestString_Batched(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		size int
		want []string
	}{
		{name: "case1", in: "hello world", size: 5, want: []string{"hello", " worl", "d"}},
		{name: "case2", in: "hello,world", size: 5, want: []string{"hello", ",worl", "d"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Batched(c.size)
			for i, v := range got {
				if v != c.want[i] {
					t.Fatalf("NewString(%s).Batched(%d) = %v; want %v", c.in, c.size, got, c.want)
				}
			}
		})
	}
}
