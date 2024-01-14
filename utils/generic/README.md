# Generic

generic 目的在于提供一组基于泛型的用于处理通用功能的函数和数据结构。该包旨在简化通用功能的实现，并提供一致的接口和易于使用的功能。
主要特性：
  - 通用功能：generic 包支持处理各种通用功能，如数据结构操作、算法实现和常用工具等。您可以使用这些功能来解决各种通用问题，并提高代码的复用性和可维护性。

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/generic)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[IsNil](#IsNil)|检查指定的值是否为 nil
|[IsAllNil](#IsAllNil)|检查指定的值是否全部为 nil
|[IsHasNil](#IsHasNil)|检查指定的值是否存在 nil


> 结构体定义

|结构体|描述
|:--|:--
|[IdR](#idr)|暂无描述...
|[IDR](#idr)|暂无描述...
|[IdW](#idw)|暂无描述...
|[IDW](#idw)|暂无描述...
|[IdR2W](#idr2w)|暂无描述...
|[IDR2W](#idr2w)|暂无描述...
|[Ordered](#ordered)|可排序类型
|[Number](#number)|数字类型
|[SignedNumber](#signednumber)|有符号数字类型
|[Integer](#integer)|整数类型
|[Signed](#signed)|有符号整数类型
|[Unsigned](#unsigned)|无符号整数类型
|[UnsignedNumber](#unsignednumber)|无符号数字类型
|[Float](#float)|浮点类型
|[Basic](#basic)|基本类型

</details>


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
### IdR

```go
type IdR[ID comparable] struct{}
```
### IDR

```go
type IDR[ID comparable] struct{}
```
### IdW

```go
type IdW[ID comparable] struct{}
```
### IDW

```go
type IDW[ID comparable] struct{}
```
### IdR2W

```go
type IdR2W[ID comparable] struct{}
```
### IDR2W

```go
type IDR2W[ID comparable] struct{}
```
### Ordered
可排序类型
```go
type Ordered struct{}
```
### Number
数字类型
```go
type Number struct{}
```
### SignedNumber
有符号数字类型
```go
type SignedNumber struct{}
```
### Integer
整数类型
```go
type Integer struct{}
```
### Signed
有符号整数类型
```go
type Signed struct{}
```
### Unsigned
无符号整数类型
```go
type Unsigned struct{}
```
### UnsignedNumber
无符号数字类型
```go
type UnsignedNumber struct{}
```
### Float
浮点类型
```go
type Float struct{}
```
### Basic
基本类型
```go
type Basic struct{}
```
