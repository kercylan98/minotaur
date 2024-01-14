# Astgo



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/astgo)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewPackage](#NewPackage)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[Comment](#comment)|暂无描述...
|[Field](#field)|暂无描述...
|[File](#file)|暂无描述...
|[Function](#function)|暂无描述...
|[Package](#package)|暂无描述...
|[Struct](#struct)|暂无描述...
|[Type](#type)|暂无描述...

</details>


#### func NewPackage(dir string)  *Package,  error
<span id="NewPackage"></span>
***
### Comment

```go
type Comment struct {
	Comments []string
	Clear    []string
}
```
### Field

```go
type Field struct {
	Anonymous bool
	Name      string
	Type      *Type
	Comments  *Comment
}
```
### File

```go
type File struct {
	af        *ast.File
	owner     *Package
	FilePath  string
	Structs   []*Struct
	Functions []*Function
	Comment   *Comment
}
```
#### func (*File) Package()  string
***
### Function

```go
type Function struct {
	Name        string
	Internal    bool
	Generic     []*Field
	Params      []*Field
	Results     []*Field
	Comments    *Comment
	Struct      *Field
	IsExample   bool
	IsTest      bool
	IsBenchmark bool
	Test        bool
}
```
### Package

```go
type Package struct {
	Dir   string
	Name  string
	Dirs  []string
	Files []*File
}
```
#### func (*Package) StructFunc(name string)  []*Function
***
#### func (*Package) PackageFunc()  []*Function
***
#### func (*Package) Structs()  []*Struct
***
#### func (*Package) FileComments()  *Comment
***
### Struct

```go
type Struct struct {
	Name     string
	Internal bool
	Comments *Comment
	Generic  []*Field
	Fields   []*Field
	Test     bool
}
```
### Type

```go
type Type struct {
	expr      ast.Expr
	Sign      string
	IsPointer bool
	Name      string
}
```
