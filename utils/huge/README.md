# Huge

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
|[NewFloat](#NewFloat)|创建一个 Float
|[NewFloatByString](#NewFloatByString)|通过字符串创建一个 Float
|[NewInt](#NewInt)|创建一个 Int 对象，该对象的值为 x


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Float](#struct_Float)|暂无描述...
|`STRUCT`|[Int](#struct_Int)|暂无描述...

</details>


***
## 详情信息
#### func NewFloat\[T generic.Number\](x T) *Float
<span id="NewFloat"></span>
> 创建一个 Float

***
#### func NewFloatByString(i string) *Float
<span id="NewFloatByString"></span>
> 通过字符串创建一个 Float
>   - 如果字符串不是一个合法的数字，则返回 0

***
#### func NewInt\[T generic.Basic\](x T) *Int
<span id="NewInt"></span>
> 创建一个 Int 对象，该对象的值为 x

**示例代码：**

该案例展示了 NewInt 对各种基本类型的支持及用法


```go

func ExampleNewInt() {
	fmt.Println(huge.NewInt("12345678900000000"))
	fmt.Println(huge.NewInt(1234567890))
	fmt.Println(huge.NewInt(true))
	fmt.Println(huge.NewInt(123.123))
	fmt.Println(huge.NewInt(byte(1)))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewInt(t *testing.T) {
	var cases = []struct {
		name string
		nil  bool
		in   int64
		mul  int64
		want string
	}{{name: "TestNewIntNegative", in: -1, want: "-1"}, {name: "TestNewIntZero", in: 0, want: "0"}, {name: "TestNewIntPositive", in: 1, want: "1"}, {name: "TestNewIntMax", in: 9223372036854775807, want: "9223372036854775807"}, {name: "TestNewIntMin", in: -9223372036854775808, want: "-9223372036854775808"}, {name: "TestNewIntMulNegative", in: -9223372036854775808, mul: 10000000, want: "-92233720368547758080000000"}, {name: "TestNewIntMulPositive", in: 9223372036854775807, mul: 10000000, want: "92233720368547758070000000"}, {name: "TestNewIntNil", nil: true, want: "0"}, {name: "TestNewIntNilMul", nil: true, mul: 10000000, want: "0"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got *huge.Int
			switch {
			case c.nil:
				if c.mul > 0 {
					got = huge.NewInt(0).MulInt64(c.mul)
				}
			case c.mul == 0:
				got = huge.NewInt(c.in)
			default:
				got = huge.NewInt(c.in).MulInt64(c.mul)
			}
			if s := got.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, got.String())
			} else {
				t.Log(s)
			}
		})
	}
	t.Run("TestNewIntFromString", func(t *testing.T) {
		if got := huge.NewInt("1234567890123456789012345678901234567890"); got.String() != "1234567890123456789012345678901234567890" {
			t.Fatalf("want: %s, got: %s", "1234567890123456789012345678901234567890", got.String())
		}
	})
	t.Run("TestNewIntFromInt", func(t *testing.T) {
		if got := huge.NewInt(1234567890); got.String() != "1234567890" {
			t.Fatalf("want: %s, got: %s", "1234567890", got.String())
		}
	})
	t.Run("TestNewIntFromBool", func(t *testing.T) {
		if got := huge.NewInt(true); got.String() != "1" {
			t.Fatalf("want: %s, got: %s", "1", got.String())
		}
	})
	t.Run("TestNewIntFromFloat", func(t *testing.T) {
		if got := huge.NewInt(1234567890.1234567890); got.String() != "1234567890" {
			t.Fatalf("want: %s, got: %s", "1234567890", got.String())
		}
	})
}

```


</details>


***
<span id="struct_Float"></span>
### Float `STRUCT`

```go
type Float big.Float
```
<span id="struct_Float_Copy"></span>

#### func (*Float) Copy()  *Float

***
<span id="struct_Float_Set"></span>

#### func (*Float) Set(i *Float)  *Float

***
<span id="struct_Float_IsZero"></span>

#### func (*Float) IsZero()  bool

***
<span id="struct_Float_ToBigFloat"></span>

#### func (*Float) ToBigFloat()  *big.Float

***
<span id="struct_Float_Cmp"></span>

#### func (*Float) Cmp(i *Float)  int
> 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0

***
<span id="struct_Float_GreaterThan"></span>

#### func (*Float) GreaterThan(i *Float)  bool
> 大于

***
<span id="struct_Float_GreaterThanOrEqualTo"></span>

#### func (*Float) GreaterThanOrEqualTo(i *Float)  bool
> 大于或等于

***
<span id="struct_Float_LessThan"></span>

#### func (*Float) LessThan(i *Float)  bool
> 小于

***
<span id="struct_Float_LessThanOrEqualTo"></span>

#### func (*Float) LessThanOrEqualTo(i *Float)  bool
> 小于或等于

***
<span id="struct_Float_EqualTo"></span>

#### func (*Float) EqualTo(i *Float)  bool
> 等于

***
<span id="struct_Float_Float64"></span>

#### func (*Float) Float64()  float64

***
<span id="struct_Float_String"></span>

#### func (*Float) String()  string

***
<span id="struct_Float_Add"></span>

#### func (*Float) Add(i *Float)  *Float

***
<span id="struct_Float_Sub"></span>

#### func (*Float) Sub(i *Float)  *Float

***
<span id="struct_Float_Mul"></span>

#### func (*Float) Mul(i *Float)  *Float

***
<span id="struct_Float_Div"></span>

#### func (*Float) Div(i *Float)  *Float

***
<span id="struct_Float_Sqrt"></span>

#### func (*Float) Sqrt()  *Float
> 平方根

***
<span id="struct_Float_Abs"></span>

#### func (*Float) Abs()  *Float
> 返回数字的绝对值

***
<span id="struct_Float_Sign"></span>

#### func (*Float) Sign()  int
> 返回数字的符号
>   - 1：正数
>   - 0：零
>   - -1：负数

***
<span id="struct_Float_IsPositive"></span>

#### func (*Float) IsPositive()  bool
> 是否为正数

***
<span id="struct_Float_IsNegative"></span>

#### func (*Float) IsNegative()  bool
> 是否为负数

***
<span id="struct_Int"></span>
### Int `STRUCT`

```go
type Int big.Int
```
<span id="struct_Int_Copy"></span>

#### func (*Int) Copy()  *Int
> 拷贝当前 Int 对象

**示例代码：**

```go

func ExampleInt_Copy() {
	var a = huge.NewInt(1234567890)
	var b = a.Copy().SetInt64(9876543210)
	fmt.Println(a)
	fmt.Println(b)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_Copy(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{{name: "TestIntCopyNegative", in: -1, want: "-1"}, {name: "TestIntCopyZero", in: 0, want: "0"}, {name: "TestIntCopyPositive", in: 1, want: "1"}, {name: "TestIntCopyMax", in: 9223372036854775807, want: "9223372036854775807"}, {name: "TestIntCopyMin", in: -9223372036854775808, want: "-9223372036854775808"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in = huge.NewInt(c.in)
			var got = in.Copy()
			if in.Int64() != c.in {
				t.Fatalf("want: %d, got: %d", c.in, in.Int64())
			}
			if s := got.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, got.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_Set"></span>

#### func (*Int) Set(i *Int)  *Int
> 设置当前 Int 对象的值为 i

**示例代码：**

```go

func ExampleInt_Set() {
	var a = huge.NewInt(1234567890)
	var b = huge.NewInt(9876543210)
	fmt.Println(a)
	a.Set(b)
	fmt.Println(a)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_Set(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{{name: "TestIntSetNegative", in: -1, want: "-1"}, {name: "TestIntSetZero", in: 0, want: "0"}, {name: "TestIntSetPositive", in: 1, want: "1"}, {name: "TestIntSetMax", in: 9223372036854775807, want: "9223372036854775807"}, {name: "TestIntSetMin", in: -9223372036854775808, want: "-9223372036854775808"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.Set(huge.NewInt(c.in))
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetString"></span>

#### func (*Int) SetString(i string)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetString(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{{name: "TestIntSetStringNegative", in: "-1", want: "-1"}, {name: "TestIntSetStringZero", in: "0", want: "0"}, {name: "TestIntSetStringPositive", in: "1", want: "1"}, {name: "TestIntSetStringMax", in: "9223372036854775807", want: "9223372036854775807"}, {name: "TestIntSetStringMin", in: "-9223372036854775808", want: "-9223372036854775808"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetString(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetInt"></span>

#### func (*Int) SetInt(i int)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetInt(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{{name: "TestIntSetIntNegative", in: -1, want: "-1"}, {name: "TestIntSetIntZero", in: 0, want: "0"}, {name: "TestIntSetIntPositive", in: 1, want: "1"}, {name: "TestIntSetIntMax", in: 9223372036854775807, want: "9223372036854775807"}, {name: "TestIntSetIntMin", in: -9223372036854775808, want: "-9223372036854775808"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetInt8"></span>

#### func (*Int) SetInt8(i int8)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetInt8(t *testing.T) {
	var cases = []struct {
		name string
		in   int8
		want string
	}{{name: "TestIntSetInt8Negative", in: -1, want: "-1"}, {name: "TestIntSetInt8Zero", in: 0, want: "0"}, {name: "TestIntSetInt8Positive", in: 1, want: "1"}, {name: "TestIntSetInt8Max", in: 127, want: "127"}, {name: "TestIntSetInt8Min", in: -128, want: "-128"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt8(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetInt16"></span>

#### func (*Int) SetInt16(i int16)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetInt16(t *testing.T) {
	var cases = []struct {
		name string
		in   int16
		want string
	}{{name: "TestIntSetInt16Negative", in: -1, want: "-1"}, {name: "TestIntSetInt16Zero", in: 0, want: "0"}, {name: "TestIntSetInt16Positive", in: 1, want: "1"}, {name: "TestIntSetInt16Max", in: 32767, want: "32767"}, {name: "TestIntSetInt16Min", in: -32768, want: "-32768"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt16(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetInt32"></span>

#### func (*Int) SetInt32(i int32)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetInt32(t *testing.T) {
	var cases = []struct {
		name string
		in   int32
		want string
	}{{name: "TestIntSetInt32Negative", in: -1, want: "-1"}, {name: "TestIntSetInt32Zero", in: 0, want: "0"}, {name: "TestIntSetInt32Positive", in: 1, want: "1"}, {name: "TestIntSetInt32Max", in: 2147483647, want: "2147483647"}, {name: "TestIntSetInt32Min", in: -2147483648, want: "-2147483648"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt32(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetInt64"></span>

#### func (*Int) SetInt64(i int64)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetInt64(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{{name: "TestIntSetInt64Negative", in: -1, want: "-1"}, {name: "TestIntSetInt64Zero", in: 0, want: "0"}, {name: "TestIntSetInt64Positive", in: 1, want: "1"}, {name: "TestIntSetInt64Max", in: 9223372036854775807, want: "9223372036854775807"}, {name: "TestIntSetInt64Min", in: -9223372036854775808, want: "-9223372036854775808"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetUint"></span>

#### func (*Int) SetUint(i uint)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetUint(t *testing.T) {
	var cases = []struct {
		name string
		in   uint64
		want string
	}{{name: "TestIntSetUintNegative", in: 0, want: "0"}, {name: "TestIntSetUintZero", in: 0, want: "0"}, {name: "TestIntSetUintPositive", in: 1, want: "1"}, {name: "TestIntSetUintMax", in: 18446744073709551615, want: "18446744073709551615"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetUint8"></span>

#### func (*Int) SetUint8(i uint8)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetUint8(t *testing.T) {
	var cases = []struct {
		name string
		in   uint8
		want string
	}{{name: "TestIntSetUint8Negative", in: 0, want: "0"}, {name: "TestIntSetUint8Zero", in: 0, want: "0"}, {name: "TestIntSetUint8Positive", in: 1, want: "1"}, {name: "TestIntSetUint8Max", in: 255, want: "255"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint8(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetUint16"></span>

#### func (*Int) SetUint16(i uint16)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetUint16(t *testing.T) {
	var cases = []struct {
		name string
		in   uint16
		want string
	}{{name: "TestIntSetUint16Negative", in: 0, want: "0"}, {name: "TestIntSetUint16Zero", in: 0, want: "0"}, {name: "TestIntSetUint16Positive", in: 1, want: "1"}, {name: "TestIntSetUint16Max", in: 65535, want: "65535"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint16(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetUint32"></span>

#### func (*Int) SetUint32(i uint32)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetUint32(t *testing.T) {
	var cases = []struct {
		name string
		in   uint32
		want string
	}{{name: "TestIntSetUint32Negative", in: 0, want: "0"}, {name: "TestIntSetUint32Zero", in: 0, want: "0"}, {name: "TestIntSetUint32Positive", in: 1, want: "1"}, {name: "TestIntSetUint32Max", in: 4294967295, want: "4294967295"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint32(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetUint64"></span>

#### func (*Int) SetUint64(i uint64)  *Int
> 设置当前 Int 对象的值为 i

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetUint64(t *testing.T) {
	var cases = []struct {
		name string
		in   uint64
		want string
	}{{name: "TestIntSetUint64Negative", in: 0, want: "0"}, {name: "TestIntSetUint64Zero", in: 0, want: "0"}, {name: "TestIntSetUint64Positive", in: 1, want: "1"}, {name: "TestIntSetUint64Max", in: 18446744073709551615, want: "18446744073709551615"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetFloat32"></span>

#### func (*Int) SetFloat32(i float32)  *Int
> 设置当前 Int 对象的值为 i 向下取整后的值

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetFloat32(t *testing.T) {
	var cases = []struct {
		name string
		in   float32
		want string
	}{{name: "TestIntSetFloat32Negative", in: -1.1, want: "-1"}, {name: "TestIntSetFloat32Zero", in: 0, want: "0"}, {name: "TestIntSetFloat32Positive", in: 1.1, want: "1"}, {name: "TestIntSetFloat32Max", in: 9223372036854775807, want: "9223372036854775807"}, {name: "TestIntSetFloat32Min", in: -9223372036854775808, want: "-9223372036854775808"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetFloat32(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetFloat64"></span>

#### func (*Int) SetFloat64(i float64)  *Int
> 设置当前 Int 对象的值为 i 向下取整后的值

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetFloat64(t *testing.T) {
	var cases = []struct {
		name string
		in   float64
		want string
	}{{name: "TestIntSetFloat64Negative", in: -1.1, want: "-1"}, {name: "TestIntSetFloat64Zero", in: 0, want: "0"}, {name: "TestIntSetFloat64Positive", in: 1.1, want: "1"}, {name: "TestIntSetFloat64Max", in: 9223372036854775807, want: "9223372036854775807"}, {name: "TestIntSetFloat64Min", in: -9223372036854775808, want: "-9223372036854775808"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetFloat64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_SetBool"></span>

#### func (*Int) SetBool(i bool)  *Int
> 设置当前 Int 对象的值为 i，当 i 为 true 时，值为 1，当 i 为 false 时，值为 0

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_SetBool(t *testing.T) {
	var cases = []struct {
		name string
		in   bool
		want string
	}{{name: "TestIntSetBoolFalse", in: false, want: "0"}, {name: "TestIntSetBoolTrue", in: true, want: "1"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetBool(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_IsZero"></span>

#### func (*Int) IsZero()  bool
> 判断当前 Int 对象的值是否为 0

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_IsZero(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want bool
	}{{name: "TestIntIsZeroNegative", in: -1, want: false}, {name: "TestIntIsZeroZero", in: 0, want: true}, {name: "TestIntIsZeroPositive", in: 1, want: false}, {name: "TestIntIsZeroMax", in: 9223372036854775807, want: false}, {name: "TestIntIsZeroMin", in: -9223372036854775808, want: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := huge.NewInt(c.in).IsZero(); got != c.want {
				t.Fatalf("want: %t, got: %t", c.want, got)
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_ToBigint"></span>

#### func (*Int) ToBigint()  *big.Int
> 转换为 *big.Int

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInt_ToBigint(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want *big.Int
	}{{name: "TestIntToBigintNegative", in: -1, want: big.NewInt(-1)}, {name: "TestIntToBigintZero", in: 0, want: big.NewInt(0)}, {name: "TestIntToBigintPositive", in: 1, want: big.NewInt(1)}, {name: "TestIntToBigintMax", in: 9223372036854775807, want: big.NewInt(9223372036854775807)}, {name: "TestIntToBigintMin", in: -9223372036854775808, want: big.NewInt(-9223372036854775808)}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := huge.NewInt(c.in).ToBigint(); got.Cmp(c.want) != 0 {
				t.Fatalf("want: %s, got: %s", c.want.String(), got.String())
			}
		})
	}
}

```


</details>


***
<span id="struct_Int_Cmp"></span>

#### func (*Int) Cmp(i *Int)  int
> 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0

***
<span id="struct_Int_GreaterThan"></span>

#### func (*Int) GreaterThan(i *Int)  bool
> 检查 slf 是否大于 i

***
<span id="struct_Int_GreaterThanOrEqualTo"></span>

#### func (*Int) GreaterThanOrEqualTo(i *Int)  bool
> 检查 slf 是否大于或等于 i

***
<span id="struct_Int_LessThan"></span>

#### func (*Int) LessThan(i *Int)  bool
> 检查 slf 是否小于 i

***
<span id="struct_Int_LessThanOrEqualTo"></span>

#### func (*Int) LessThanOrEqualTo(i *Int)  bool
> 检查 slf 是否小于或等于 i

***
<span id="struct_Int_EqualTo"></span>

#### func (*Int) EqualTo(i *Int)  bool
> 检查 slf 是否等于 i

***
<span id="struct_Int_Int64"></span>

#### func (*Int) Int64()  int64
> 转换为 int64 类型进行返回

***
<span id="struct_Int_String"></span>

#### func (*Int) String()  string
> 转换为 string 类型进行返回

***
<span id="struct_Int_Add"></span>

#### func (*Int) Add(i *Int)  *Int
> 使用 i 对 slf 进行加法运算，slf 的值会变为运算后的值。返回 slf

***
<span id="struct_Int_AddInt"></span>

#### func (*Int) AddInt(i int)  *Int

***
<span id="struct_Int_AddInt8"></span>

#### func (*Int) AddInt8(i int8)  *Int

***
<span id="struct_Int_AddInt16"></span>

#### func (*Int) AddInt16(i int16)  *Int

***
<span id="struct_Int_AddInt32"></span>

#### func (*Int) AddInt32(i int32)  *Int

***
<span id="struct_Int_AddInt64"></span>

#### func (*Int) AddInt64(i int64)  *Int

***
<span id="struct_Int_AddUint"></span>

#### func (*Int) AddUint(i uint)  *Int

***
<span id="struct_Int_AddUint8"></span>

#### func (*Int) AddUint8(i uint8)  *Int

***
<span id="struct_Int_AddUint16"></span>

#### func (*Int) AddUint16(i uint16)  *Int

***
<span id="struct_Int_AddUint32"></span>

#### func (*Int) AddUint32(i uint32)  *Int

***
<span id="struct_Int_AddUint64"></span>

#### func (*Int) AddUint64(i uint64)  *Int

***
<span id="struct_Int_Mul"></span>

#### func (*Int) Mul(i *Int)  *Int

***
<span id="struct_Int_MulInt"></span>

#### func (*Int) MulInt(i int)  *Int

***
<span id="struct_Int_MulInt8"></span>

#### func (*Int) MulInt8(i int8)  *Int

***
<span id="struct_Int_MulInt16"></span>

#### func (*Int) MulInt16(i int16)  *Int

***
<span id="struct_Int_MulInt32"></span>

#### func (*Int) MulInt32(i int32)  *Int

***
<span id="struct_Int_MulInt64"></span>

#### func (*Int) MulInt64(i int64)  *Int

***
<span id="struct_Int_MulUint"></span>

#### func (*Int) MulUint(i uint)  *Int

***
<span id="struct_Int_MulUint8"></span>

#### func (*Int) MulUint8(i uint8)  *Int

***
<span id="struct_Int_MulUint16"></span>

#### func (*Int) MulUint16(i uint16)  *Int

***
<span id="struct_Int_MulUint32"></span>

#### func (*Int) MulUint32(i uint32)  *Int

***
<span id="struct_Int_MulUint64"></span>

#### func (*Int) MulUint64(i uint64)  *Int

***
<span id="struct_Int_Sub"></span>

#### func (*Int) Sub(i *Int)  *Int

***
<span id="struct_Int_SubInt"></span>

#### func (*Int) SubInt(i int)  *Int

***
<span id="struct_Int_SubInt8"></span>

#### func (*Int) SubInt8(i int8)  *Int

***
<span id="struct_Int_SubInt16"></span>

#### func (*Int) SubInt16(i int16)  *Int

***
<span id="struct_Int_SubInt32"></span>

#### func (*Int) SubInt32(i int32)  *Int

***
<span id="struct_Int_SubInt64"></span>

#### func (*Int) SubInt64(i int64)  *Int

***
<span id="struct_Int_SubUint"></span>

#### func (*Int) SubUint(i uint)  *Int

***
<span id="struct_Int_SubUint8"></span>

#### func (*Int) SubUint8(i uint8)  *Int

***
<span id="struct_Int_SubUint16"></span>

#### func (*Int) SubUint16(i uint16)  *Int

***
<span id="struct_Int_SubUint32"></span>

#### func (*Int) SubUint32(i uint32)  *Int

***
<span id="struct_Int_SubUint64"></span>

#### func (*Int) SubUint64(i uint64)  *Int

***
<span id="struct_Int_Div"></span>

#### func (*Int) Div(i *Int)  *Int

***
<span id="struct_Int_DivInt"></span>

#### func (*Int) DivInt(i int)  *Int

***
<span id="struct_Int_DivInt8"></span>

#### func (*Int) DivInt8(i int8)  *Int

***
<span id="struct_Int_DivInt16"></span>

#### func (*Int) DivInt16(i int16)  *Int

***
<span id="struct_Int_DivInt32"></span>

#### func (*Int) DivInt32(i int32)  *Int

***
<span id="struct_Int_DivInt64"></span>

#### func (*Int) DivInt64(i int64)  *Int

***
<span id="struct_Int_DivUint"></span>

#### func (*Int) DivUint(i uint)  *Int

***
<span id="struct_Int_DivUint8"></span>

#### func (*Int) DivUint8(i uint8)  *Int

***
<span id="struct_Int_DivUint16"></span>

#### func (*Int) DivUint16(i uint16)  *Int

***
<span id="struct_Int_DivUint32"></span>

#### func (*Int) DivUint32(i uint32)  *Int

***
<span id="struct_Int_DivUint64"></span>

#### func (*Int) DivUint64(i uint64)  *Int

***
<span id="struct_Int_Mod"></span>

#### func (*Int) Mod(i *Int)  *Int

***
<span id="struct_Int_ModInt"></span>

#### func (*Int) ModInt(i int)  *Int

***
<span id="struct_Int_ModInt8"></span>

#### func (*Int) ModInt8(i int8)  *Int

***
<span id="struct_Int_ModInt16"></span>

#### func (*Int) ModInt16(i int16)  *Int

***
<span id="struct_Int_ModInt32"></span>

#### func (*Int) ModInt32(i int32)  *Int

***
<span id="struct_Int_ModInt64"></span>

#### func (*Int) ModInt64(i int64)  *Int

***
<span id="struct_Int_ModUint"></span>

#### func (*Int) ModUint(i uint)  *Int

***
<span id="struct_Int_ModUint8"></span>

#### func (*Int) ModUint8(i uint8)  *Int

***
<span id="struct_Int_ModUint16"></span>

#### func (*Int) ModUint16(i uint16)  *Int

***
<span id="struct_Int_ModUint32"></span>

#### func (*Int) ModUint32(i uint32)  *Int

***
<span id="struct_Int_ModUint64"></span>

#### func (*Int) ModUint64(i uint64)  *Int

***
<span id="struct_Int_Pow"></span>

#### func (*Int) Pow(i *Int)  *Int

***
<span id="struct_Int_PowInt"></span>

#### func (*Int) PowInt(i int)  *Int

***
<span id="struct_Int_PowInt8"></span>

#### func (*Int) PowInt8(i int8)  *Int

***
<span id="struct_Int_PowInt16"></span>

#### func (*Int) PowInt16(i int16)  *Int

***
<span id="struct_Int_PowInt32"></span>

#### func (*Int) PowInt32(i int32)  *Int

***
<span id="struct_Int_PowInt64"></span>

#### func (*Int) PowInt64(i int64)  *Int

***
<span id="struct_Int_PowUint"></span>

#### func (*Int) PowUint(i uint)  *Int

***
<span id="struct_Int_PowUint8"></span>

#### func (*Int) PowUint8(i uint8)  *Int

***
<span id="struct_Int_PowUint16"></span>

#### func (*Int) PowUint16(i uint16)  *Int

***
<span id="struct_Int_PowUint32"></span>

#### func (*Int) PowUint32(i uint32)  *Int

***
<span id="struct_Int_PowUint64"></span>

#### func (*Int) PowUint64(i uint64)  *Int

***
<span id="struct_Int_Lsh"></span>

#### func (*Int) Lsh(i int)  *Int
> 左移

***
<span id="struct_Int_Rsh"></span>

#### func (*Int) Rsh(i int)  *Int
> 右移

***
<span id="struct_Int_And"></span>

#### func (*Int) And(i *Int)  *Int
> 与

***
<span id="struct_Int_AndNot"></span>

#### func (*Int) AndNot(i *Int)  *Int
> 与非

***
<span id="struct_Int_Or"></span>

#### func (*Int) Or(i *Int)  *Int
> 或

***
<span id="struct_Int_Xor"></span>

#### func (*Int) Xor(i *Int)  *Int
> 异或

***
<span id="struct_Int_Not"></span>

#### func (*Int) Not()  *Int
> 非

***
<span id="struct_Int_Sqrt"></span>

#### func (*Int) Sqrt()  *Int
> 平方根

***
<span id="struct_Int_GCD"></span>

#### func (*Int) GCD(i *Int)  *Int
> 最大公约数

***
<span id="struct_Int_LCM"></span>

#### func (*Int) LCM(i *Int)  *Int
> 最小公倍数

***
<span id="struct_Int_ModInverse"></span>

#### func (*Int) ModInverse(i *Int)  *Int
> 模反元素

***
<span id="struct_Int_ModSqrt"></span>

#### func (*Int) ModSqrt(i *Int)  *Int
> 模平方根

***
<span id="struct_Int_BitLen"></span>

#### func (*Int) BitLen()  int
> 二进制长度

***
<span id="struct_Int_Bit"></span>

#### func (*Int) Bit(i int)  uint
> 二进制位

***
<span id="struct_Int_SetBit"></span>

#### func (*Int) SetBit(i int, v uint)  *Int
> 设置二进制位

***
<span id="struct_Int_Neg"></span>

#### func (*Int) Neg()  *Int
> 返回数字的相反数

***
<span id="struct_Int_Abs"></span>

#### func (*Int) Abs()  *Int
> 返回数字的绝对值

***
<span id="struct_Int_Sign"></span>

#### func (*Int) Sign()  int
> 返回数字的符号
>   - 1：正数
>   - 0：零
>   - -1：负数

***
<span id="struct_Int_IsPositive"></span>

#### func (*Int) IsPositive()  bool
> 是否为正数

***
<span id="struct_Int_IsNegative"></span>

#### func (*Int) IsNegative()  bool
> 是否为负数

***
<span id="struct_Int_IsEven"></span>

#### func (*Int) IsEven()  bool
> 是否为偶数

***
<span id="struct_Int_IsOdd"></span>

#### func (*Int) IsOdd()  bool
> 是否为奇数

***
<span id="struct_Int_ProportionalCalc"></span>

#### func (*Int) ProportionalCalc(proportional *Int, formula func (v *Int)  *Int)  *Int
> 比例计算，该函数会再 formula 返回值的基础上除以 proportional
>   - formula 为计算公式，该公式的参数为调用该函数的 Int 的拷贝

***
