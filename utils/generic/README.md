# Generic

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/generic)
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
|`INTERFACE`|[IdR](#idr)|暂无描述...
|`INTERFACE`|[IDR](#idr)|暂无描述...
|`INTERFACE`|[IdW](#idw)|暂无描述...
|`INTERFACE`|[IDW](#idw)|暂无描述...
|`INTERFACE`|[IdR2W](#idr2w)|暂无描述...
|`INTERFACE`|[IDR2W](#idr2w)|暂无描述...
|`INTERFACE`|[Ordered](#ordered)|可排序类型
|`INTERFACE`|[Number](#number)|数字类型
|`INTERFACE`|[SignedNumber](#signednumber)|有符号数字类型
|`INTERFACE`|[Integer](#integer)|整数类型
|`INTERFACE`|[Signed](#signed)|有符号整数类型
|`INTERFACE`|[Unsigned](#unsigned)|无符号整数类型
|`INTERFACE`|[UnsignedNumber](#unsignednumber)|无符号数字类型
|`INTERFACE`|[Float](#float)|浮点类型
|`INTERFACE`|[Basic](#basic)|基本类型

</details>


***
## 详情信息
#### func IsNil(v V)  bool
<span id="IsNil"></span>
> 检查指定的值是否为 nil

***
#### func IsAllNil(v ...V)  bool
<span id="IsAllNil"></span>
> 检查指定的值是否全部为 nil

***
#### func IsHasNil(v ...V)  bool
<span id="IsHasNil"></span>
> 检查指定的值是否存在 nil

***
### IdR `INTERFACE`

```go
type IdR[ID comparable] interface {
	GetId() ID
}
```
### IDR `INTERFACE`

```go
type IDR[ID comparable] interface {
	GetID() ID
}
```
### IdW `INTERFACE`

```go
type IdW[ID comparable] interface {
	SetId(id ID)
}
```
### IDW `INTERFACE`

```go
type IDW[ID comparable] interface {
	SetID(id ID)
}
```
### IdR2W `INTERFACE`

```go
type IdR2W[ID comparable] interface {
	IdR[ID]
	IdW[ID]
}
```
### IDR2W `INTERFACE`

```go
type IDR2W[ID comparable] interface {
	IDR[ID]
	IDW[ID]
}
```
### Ordered `INTERFACE`
可排序类型
```go
type Ordered interface {
}
```
### Number `INTERFACE`
数字类型
```go
type Number interface {
}
```
### SignedNumber `INTERFACE`
有符号数字类型
```go
type SignedNumber interface {
}
```
### Integer `INTERFACE`
整数类型
```go
type Integer interface {
}
```
### Signed `INTERFACE`
有符号整数类型
```go
type Signed interface {
}
```
### Unsigned `INTERFACE`
无符号整数类型
```go
type Unsigned interface {
}
```
### UnsignedNumber `INTERFACE`
无符号数字类型
```go
type UnsignedNumber interface {
}
```
### Float `INTERFACE`
浮点类型
```go
type Float interface {
}
```
### Basic `INTERFACE`
基本类型
```go
type Basic interface {
}
```
