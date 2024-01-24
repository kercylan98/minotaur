# Generic

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

generic 目的在于提供一组基于泛型的用于处理通用功能的函数和数据结构。该包旨在简化通用功能的实现，并提供一致的接口和易于使用的功能。
主要特性：
  - 通用功能：generic 包支持处理各种通用功能，如数据结构操作、算法实现和常用工具等。您可以使用这些功能来解决各种通用问题，并提高代码的复用性和可维护性。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[IsNil](#IsNil)|检查指定的值是否为 nil
|[IsAllNil](#IsAllNil)|检查指定的值是否全部为 nil
|[IsHasNil](#IsHasNil)|检查指定的值是否存在 nil


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`INTERFACE`|[IdR](#struct_IdR)|暂无描述...
|`INTERFACE`|[IDR](#struct_IDR)|暂无描述...
|`INTERFACE`|[IdW](#struct_IdW)|暂无描述...
|`INTERFACE`|[IDW](#struct_IDW)|暂无描述...
|`INTERFACE`|[IdR2W](#struct_IdR2W)|暂无描述...
|`INTERFACE`|[IDR2W](#struct_IDR2W)|暂无描述...
|`INTERFACE`|[Ordered](#struct_Ordered)|可排序类型
|`INTERFACE`|[Number](#struct_Number)|数字类型
|`INTERFACE`|[SignedNumber](#struct_SignedNumber)|有符号数字类型
|`INTERFACE`|[Integer](#struct_Integer)|整数类型
|`INTERFACE`|[Signed](#struct_Signed)|有符号整数类型
|`INTERFACE`|[Unsigned](#struct_Unsigned)|无符号整数类型
|`INTERFACE`|[UnsignedNumber](#struct_UnsignedNumber)|无符号数字类型
|`INTERFACE`|[Float](#struct_Float)|浮点类型
|`INTERFACE`|[Basic](#struct_Basic)|基本类型

</details>


***
## 详情信息
#### func IsNil\[V any\](v V) bool
<span id="IsNil"></span>
> 检查指定的值是否为 nil

***
#### func IsAllNil\[V any\](v ...V) bool
<span id="IsAllNil"></span>
> 检查指定的值是否全部为 nil

***
#### func IsHasNil\[V any\](v ...V) bool
<span id="IsHasNil"></span>
> 检查指定的值是否存在 nil

***
<span id="struct_IdR"></span>
### IdR `INTERFACE`

```go
type IdR[ID comparable] interface {
	GetId() ID
}
```
<span id="struct_IDR"></span>
### IDR `INTERFACE`

```go
type IDR[ID comparable] interface {
	GetID() ID
}
```
<span id="struct_IdW"></span>
### IdW `INTERFACE`

```go
type IdW[ID comparable] interface {
	SetId(id ID)
}
```
<span id="struct_IDW"></span>
### IDW `INTERFACE`

```go
type IDW[ID comparable] interface {
	SetID(id ID)
}
```
<span id="struct_IdR2W"></span>
### IdR2W `INTERFACE`

```go
type IdR2W[ID comparable] interface {
	IdR[ID]
	IdW[ID]
}
```
<span id="struct_IDR2W"></span>
### IDR2W `INTERFACE`

```go
type IDR2W[ID comparable] interface {
	IDR[ID]
	IDW[ID]
}
```
<span id="struct_Ordered"></span>
### Ordered `INTERFACE`
可排序类型
```go
type Ordered interface {
	Integer | Float | ~string
}
```
<span id="struct_Number"></span>
### Number `INTERFACE`
数字类型
```go
type Number interface {
	Integer | Float
}
```
<span id="struct_SignedNumber"></span>
### SignedNumber `INTERFACE`
有符号数字类型
```go
type SignedNumber interface {
	Signed | Float
}
```
<span id="struct_Integer"></span>
### Integer `INTERFACE`
整数类型
```go
type Integer interface {
	Signed | Unsigned
}
```
<span id="struct_Signed"></span>
### Signed `INTERFACE`
有符号整数类型
```go
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
```
<span id="struct_Unsigned"></span>
### Unsigned `INTERFACE`
无符号整数类型
```go
type Unsigned interface {
	UnsignedNumber | ~uintptr
}
```
<span id="struct_UnsignedNumber"></span>
### UnsignedNumber `INTERFACE`
无符号数字类型
```go
type UnsignedNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
```
<span id="struct_Float"></span>
### Float `INTERFACE`
浮点类型
```go
type Float interface {
	~float32 | ~float64
}
```
<span id="struct_Basic"></span>
### Basic `INTERFACE`
基本类型
```go
type Basic interface {
	Signed | Unsigned | Float | ~string | ~bool | ~byte
}
```
