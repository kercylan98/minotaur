# Stream

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

暂无介绍...


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewString](#NewString)|创建字符串流
|[NewStrings](#NewStrings)|创建字符串切片


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[String](#struct_String)|字符串流
|`STRUCT`|[Strings](#struct_Strings)|字符串切片

</details>


***
## 详情信息
#### func NewString\[S ~string\](s S) *String[S]
<span id="NewString"></span>
> 创建字符串流

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewString(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "hello"}, {name: "case2", in: "world", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in)
			if got.String() != c.want {
				t.Fatalf("NewString(%s) = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
#### func NewStrings\[S ~string\](s ...S) *Strings[S]
<span id="NewStrings"></span>
> 创建字符串切片

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewStrings(t *testing.T) {
	var cases = []struct {
		name string
		in   []string
		want []string
	}{{name: "empty", in: []string{}, want: []string{}}, {name: "one", in: []string{"a"}, want: []string{"a"}}, {name: "two", in: []string{"a", "b"}, want: []string{"a", "b"}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...)
			if got.Len() != len(c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String"></span>
### String `STRUCT`
字符串流
```go
type String[S ~string] struct {
	str S
}
```
<span id="struct_String_Elem"></span>

#### func (*String) Elem()  S
> 返回原始元素

***
<span id="struct_String_String"></span>

#### func (*String) String()  string
> 返回字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_String(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "hello"}, {name: "case2", in: "world", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).String()
			if got != c.want {
				t.Fatalf("String(%s).String() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Index"></span>

#### func (*String) Index(i int)  *String[S]
> 返回字符串指定位置的字符，当索引超出范围时将会触发 panic

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Index(t *testing.T) {
	var cases = []struct {
		name        string
		in          string
		i           int
		want        string
		shouldPanic bool
	}{{name: "case1", in: "hello", i: 0, want: "h", shouldPanic: false}, {name: "case2", in: "world", i: 2, want: "r", shouldPanic: false}, {name: "case3", in: "world", i: 5, want: "", shouldPanic: true}, {name: "case4", in: "world", i: -1, want: "", shouldPanic: true}}
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

```


</details>


***
<span id="struct_String_Range"></span>

#### func (*String) Range(start int, end int)  *String[S]
> 返回字符串指定范围的字符

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Range(t *testing.T) {
	var cases = []struct {
		name        string
		in          string
		start       int
		end         int
		want        string
		shouldPanic bool
	}{{name: "case1", in: "hello", start: 0, end: 2, want: "he", shouldPanic: false}, {name: "case2", in: "world", start: 2, end: 5, want: "rld", shouldPanic: false}, {name: "case3", in: "world", start: 5, end: 6, want: "", shouldPanic: true}, {name: "case4", in: "world", start: -1, end: 6, want: "", shouldPanic: true}}
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

```


</details>


***
<span id="struct_String_TrimSpace"></span>

#### func (*String) TrimSpace()  *String[S]
> 返回去除字符串首尾空白字符的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_TrimSpace(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: " hello ", want: "hello"}, {name: "case2", in: " world ", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TrimSpace()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TrimSpace() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Trim"></span>

#### func (*String) Trim(cs string)  *String[S]
> 返回去除字符串首尾指定字符的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Trim(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		cs   string
		want string
	}{{name: "case1", in: "hello", cs: "h", want: "ello"}, {name: "case2", in: "world", cs: "d", want: "worl"}, {name: "none", in: "world", cs: "", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Trim(c.cs)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Trim(%s) = %s; want %s", c.in, c.cs, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_TrimPrefix"></span>

#### func (*String) TrimPrefix(prefix string)  *String[S]
> 返回去除字符串前缀的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_TrimPrefix(t *testing.T) {
	var cases = []struct {
		name    string
		in      string
		prefix  string
		want    string
		isEqual bool
	}{{name: "case1", in: "hello", prefix: "h", want: "ello", isEqual: false}, {name: "case2", in: "world", prefix: "w", want: "orld", isEqual: false}, {name: "none", in: "world", prefix: "x", want: "world", isEqual: true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TrimPrefix(c.prefix)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TrimPrefix(%s) = %s; want %s", c.in, c.prefix, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_TrimSuffix"></span>

#### func (*String) TrimSuffix(suffix string)  *String[S]
> 返回去除字符串后缀的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_TrimSuffix(t *testing.T) {
	var cases = []struct {
		name    string
		in      string
		suffix  string
		want    string
		isEqual bool
	}{{name: "case1", in: "hello", suffix: "o", want: "hell", isEqual: false}, {name: "case2", in: "world", suffix: "d", want: "worl", isEqual: false}, {name: "none", in: "world", suffix: "x", want: "world", isEqual: true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TrimSuffix(c.suffix)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TrimSuffix(%s) = %s; want %s", c.in, c.suffix, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_ToUpper"></span>

#### func (*String) ToUpper()  *String[S]
> 返回字符串的大写形式

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_ToUpper(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "HELLO"}, {name: "case2", in: "world", want: "WORLD"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).ToUpper()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).ToUpper() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_ToLower"></span>

#### func (*String) ToLower()  *String[S]
> 返回字符串的小写形式

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_ToLower(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "HELLO", want: "hello"}, {name: "case2", in: "WORLD", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).ToLower()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).ToLower() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Equal"></span>

#### func (*String) Equal(ss S)  bool
> 返回字符串是否相等

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Equal(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want bool
	}{{name: "case1", in: "hello", ss: "hello", want: true}, {name: "case2", in: "world", ss: "world", want: true}, {name: "case3", in: "world", ss: "worldx", want: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Equal(c.ss)
			if got != c.want {
				t.Fatalf("NewString(%s).Equal(%s) = %t; want %t", c.in, c.ss, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_HasPrefix"></span>

#### func (*String) HasPrefix(prefix S)  bool
> 返回字符串是否包含指定前缀

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_HasPrefix(t *testing.T) {
	var cases = []struct {
		name   string
		in     string
		prefix string
		want   bool
	}{{name: "case1", in: "hello", prefix: "h", want: true}, {name: "case2", in: "world", prefix: "w", want: true}, {name: "case3", in: "world", prefix: "x", want: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).HasPrefix(c.prefix)
			if got != c.want {
				t.Fatalf("NewString(%s).HasPrefix(%s) = %t; want %t", c.in, c.prefix, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_HasSuffix"></span>

#### func (*String) HasSuffix(suffix S)  bool
> 返回字符串是否包含指定后缀

***
<span id="struct_String_Len"></span>

#### func (*String) Len()  int
> 返回字符串长度

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Len(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want int
	}{{name: "case1", in: "hello", want: 5}, {name: "case2", in: "world", want: 5}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Len()
			if got != c.want {
				t.Fatalf("NewString(%s).Len() = %d; want %d", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Contains"></span>

#### func (*String) Contains(sub S)  bool
> 返回字符串是否包含指定子串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Contains(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want bool
	}{{name: "case1", in: "hello", ss: "he", want: true}, {name: "case2", in: "world", ss: "or", want: true}, {name: "case3", in: "world", ss: "x", want: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Contains(c.ss)
			if got != c.want {
				t.Fatalf("NewString(%s).Contains(%s) = %t; want %t", c.in, c.ss, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Count"></span>

#### func (*String) Count(sub S)  int
> 返回字符串包含指定子串的次数

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Count(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want int
	}{{name: "case1", in: "hello", ss: "l", want: 2}, {name: "case2", in: "world", ss: "o", want: 1}, {name: "case3", in: "world", ss: "x", want: 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Count(c.ss)
			if got != c.want {
				t.Fatalf("NewString(%s).Count(%s) = %d; want %d", c.in, c.ss, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Repeat"></span>

#### func (*String) Repeat(count int)  *String[S]
> 返回重复 count 次的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Repeat(t *testing.T) {
	var cases = []struct {
		name  string
		in    string
		count int
		want  string
	}{{name: "case1", in: "hello", count: 2, want: "hellohello"}, {name: "case2", in: "world", count: 3, want: "worldworldworld"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Repeat(c.count)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Repeat(%d) = %s; want %s", c.in, c.count, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Replace"></span>

#### func (*String) Replace(old S, new S, n int)  *String[S]
> 返回替换指定子串后的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Replace(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		old  string
		new  string
		want string
	}{{name: "case1", in: "hello", old: "l", new: "x", want: "hexxo"}, {name: "case2", in: "world", old: "o", new: "x", want: "wxrld"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Replace(c.old, c.new, -1)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Replace(%s, %s) = %s; want %s", c.in, c.old, c.new, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_ReplaceAll"></span>

#### func (*String) ReplaceAll(old S, new S)  *String[S]
> 返回替换所有指定子串后的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_ReplaceAll(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		old  string
		new  string
		want string
	}{{name: "case1", in: "hello", old: "l", new: "x", want: "hexxo"}, {name: "case2", in: "world", old: "o", new: "x", want: "wxrld"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).ReplaceAll(c.old, c.new)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).ReplaceAll(%s, %s) = %s; want %s", c.in, c.old, c.new, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Append"></span>

#### func (*String) Append(ss S)  *String[S]
> 返回追加指定字符串后的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Append(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want string
	}{{name: "case1", in: "hello", ss: " world", want: "hello world"}, {name: "case2", in: "world", ss: " hello", want: "world hello"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Append(c.ss)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Append(%s) = %s; want %s", c.in, c.ss, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Prepend"></span>

#### func (*String) Prepend(ss S)  *String[S]
> 返回追加指定字符串后的字符串，追加的字符串在前

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Prepend(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		ss   string
		want string
	}{{name: "case1", in: "hello", ss: "world ", want: "world hello"}, {name: "case2", in: "world", ss: "hello ", want: "hello world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Prepend(c.ss)
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Prepend(%s) = %s; want %s", c.in, c.ss, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Clear"></span>

#### func (*String) Clear()  *String[S]
> 返回清空字符串后的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Clear(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: ""}, {name: "case2", in: "world", want: ""}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Clear()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Clear() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Reverse"></span>

#### func (*String) Reverse()  *String[S]
> 返回反转字符串后的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Reverse(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "olleh"}, {name: "case2", in: "world", want: "dlrow"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Reverse()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Reverse() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Queto"></span>

#### func (*String) Queto()  *String[S]
> 返回带引号的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Queto(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "\"hello\""}, {name: "case2", in: "world", want: "\"world\""}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Queto()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Queto() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_QuetoToASCII"></span>

#### func (*String) QuetoToASCII()  *String[S]
> 返回带引号的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_QuetoToASCII(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "\"hello\""}, {name: "case2", in: "world", want: "\"world\""}, {name: "case3", in: "你好", want: "\"\\u4f60\\u597d\""}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).QuetoToASCII()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).QuetoToASCII() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_FirstUpper"></span>

#### func (*String) FirstUpper()  *String[S]
> 返回首字母大写的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_FirstUpper(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "Hello"}, {name: "case2", in: "world", want: "World"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).FirstUpper()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).FirstUpper() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_FirstLower"></span>

#### func (*String) FirstLower()  *String[S]
> 返回首字母小写的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_FirstLower(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "Hello", want: "hello"}, {name: "case2", in: "World", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).FirstLower()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).FirstLower() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_SnakeCase"></span>

#### func (*String) SnakeCase()  *String[S]
> 返回蛇形命名的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_SnakeCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "HelloWorld", want: "hello_world"}, {name: "case2", in: "HelloWorldHello", want: "hello_world_hello"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).SnakeCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).SnakeCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_CamelCase"></span>

#### func (*String) CamelCase()  *String[S]
> 返回驼峰命名的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_CamelCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello_world", want: "helloWorld"}, {name: "case2", in: "hello_world_hello", want: "helloWorldHello"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).CamelCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).CamelCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_KebabCase"></span>

#### func (*String) KebabCase()  *String[S]
> 返回短横线命名的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_KebabCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "HelloWorld", want: "hello-world"}, {name: "case2", in: "HelloWorldHello", want: "hello-world-hello"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).KebabCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).KebabCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_TitleCase"></span>

#### func (*String) TitleCase()  *String[S]
> 返回标题命名的字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_TitleCase(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello_world", want: "HelloWorld"}, {name: "case2", in: "hello_world_hello", want: "HelloWorldHello"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).TitleCase()
			if got.String() != c.want {
				t.Fatalf("NewString(%s).TitleCase() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Bytes"></span>

#### func (*String) Bytes()  []byte
> 返回字符串的字节数组

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Bytes(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "hello"}, {name: "case2", in: "world", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Bytes()
			if string(got) != c.want {
				t.Fatalf("NewString(%s).Bytes() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Runes"></span>

#### func (*String) Runes()  []rune
> 返回字符串的字符数组

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Runes(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "hello"}, {name: "case2", in: "world", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Runes()
			if string(got) != c.want {
				t.Fatalf("NewString(%s).Runes() = %v; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Default"></span>

#### func (*String) Default(def S)  *String[S]
> 当字符串为空时设置默认值

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Default(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "", want: "default"}, {name: "case2", in: "world", want: "world"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Default("default")
			if got.String() != c.want {
				t.Fatalf("NewString(%s).Default() = %s; want %s", c.in, got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Handle"></span>

#### func (*String) Handle(f func ( S))  *String[S]
> 处理字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Handle(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "hello"}, {name: "case2", in: "world", want: "world"}}
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

```


</details>


***
<span id="struct_String_Update"></span>

#### func (*String) Update(f func ( S)  S)  *String[S]
> 更新字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Update(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "case1", in: "hello", want: "HELLO"}, {name: "case2", in: "world", want: "WORLD"}}
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

```


</details>


***
<span id="struct_String_Split"></span>

#### func (*String) Split(sep string)  *Strings[S]
> 返回字符串切片

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Split(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		sep  string
		want []string
	}{{name: "case1", in: "hello world", sep: " ", want: []string{"hello", "world"}}, {name: "case2", in: "hello,world", sep: ",", want: []string{"hello", "world"}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Split(c.sep)
			for i, v := range got.Elem() {
				if v != c.want[i] {
					t.Fatalf("NewString(%s).Split(%s) = %v; want %v", c.in, c.sep, got, c.want)
				}
			}
		})
	}
}

```


</details>


***
<span id="struct_String_SplitN"></span>

#### func (*String) SplitN(sep string, n int)  *Strings[S]
> 返回字符串切片

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_SplitN(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		sep  string
		n    int
		want []string
	}{{name: "case1", in: "hello world", sep: " ", n: 2, want: []string{"hello", "world"}}, {name: "case2", in: "hello,world", sep: ",", n: 2, want: []string{"hello", "world"}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).SplitN(c.sep, c.n)
			for i, v := range got.Elem() {
				if v != c.want[i] {
					t.Fatalf("NewString(%s).SplitN(%s, %d) = %v; want %v", c.in, c.sep, c.n, got, c.want)
				}
			}
		})
	}
}

```


</details>


***
<span id="struct_String_Batched"></span>

#### func (*String) Batched(size int)  *Strings[S]
> 将字符串按照指定长度分组，最后一组可能小于指定长度

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestString_Batched(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		size int
		want []string
	}{{name: "case1", in: "hello world", size: 5, want: []string{"hello", " worl", "d"}}, {name: "case2", in: "hello,world", size: 5, want: []string{"hello", ",worl", "d"}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewString(c.in).Batched(c.size)
			for i, v := range got.Elem() {
				if v != c.want[i] {
					t.Fatalf("NewString(%s).Batched(%d) = %v; want %v", c.in, c.size, got, c.want)
				}
			}
		})
	}
}

```


</details>


***
<span id="struct_Strings"></span>
### Strings `STRUCT`
字符串切片
```go
type Strings[S ~string] struct {
	s []S
}
```
<span id="struct_Strings_Elem"></span>

#### func (*Strings) Elem()  []S
> 返回原始元素

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestStrings_Elem(t *testing.T) {
	var cases = []struct {
		name string
		in   []string
		want []string
	}{{name: "empty", in: []string{}, want: []string{}}, {name: "one", in: []string{"a"}, want: []string{"a"}}, {name: "two", in: []string{"a", "b"}, want: []string{"a", "b"}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...).Elem()
			if len(got) != len(c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_Strings_Len"></span>

#### func (*Strings) Len()  int
> 返回切片长度

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestStrings_Len(t *testing.T) {
	var cases = []struct {
		name string
		in   []string
		want int
	}{{name: "empty", in: []string{}, want: 0}, {name: "one", in: []string{"a"}, want: 1}, {name: "two", in: []string{"a", "b"}, want: 2}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...)
			if got.Len() != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_Strings_Append"></span>

#### func (*Strings) Append(ss ...S)  *Strings[S]
> 添加字符串

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestStrings_Append(t *testing.T) {
	var cases = []struct {
		name   string
		in     []string
		append []string
		want   []string
	}{{name: "empty", in: []string{}, append: []string{"a"}, want: []string{"a"}}, {name: "one", in: []string{"a"}, append: []string{"b"}, want: []string{"a", "b"}}, {name: "two", in: []string{"a", "b"}, append: []string{"c"}, want: []string{"a", "b", "c"}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := stream.NewStrings(c.in...).Append(c.append...)
			if got.Len() != len(c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

```


</details>


***
<span id="struct_Strings_Join"></span>

#### func (*Strings) Join(sep S)  *String[S]
> 连接字符串

***
<span id="struct_Strings_Choice"></span>

#### func (*Strings) Choice(i int)  *String[S]
> 选择字符串

***
<span id="struct_Strings_Choices"></span>

#### func (*Strings) Choices(i ...int)  *Strings[S]
> 选择多个字符串

***
<span id="struct_Strings_ChoiceInRange"></span>

#### func (*Strings) ChoiceInRange(start int, end int)  *Strings[S]
> 选择范围内的字符串

***
<span id="struct_Strings_Remove"></span>

#### func (*Strings) Remove(i int)  *Strings[S]
> 移除字符串

***
<span id="struct_Strings_Removes"></span>

#### func (*Strings) Removes(i ...int)  *Strings[S]
> 移除多个字符串

***
<span id="struct_Strings_RemoveInRange"></span>

#### func (*Strings) RemoveInRange(start int, end int)  *Strings[S]
> 移除范围内的字符串

***
<span id="struct_Strings_Clear"></span>

#### func (*Strings) Clear()  *Strings[S]
> 清空字符串

***
<span id="struct_Strings_First"></span>

#### func (*Strings) First()  *String[S]
> 第一个字符串

***
<span id="struct_Strings_Last"></span>

#### func (*Strings) Last()  *String[S]
> 最后一个字符串

***
