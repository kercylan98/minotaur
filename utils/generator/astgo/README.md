# Astgo

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
|[NewPackage](#NewPackage)|暂无描述...


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Comment](#struct_Comment)|暂无描述...
|`STRUCT`|[Field](#struct_Field)|暂无描述...
|`STRUCT`|[File](#struct_File)|暂无描述...
|`STRUCT`|[Function](#struct_Function)|暂无描述...
|`STRUCT`|[Package](#struct_Package)|暂无描述...
|`STRUCT`|[Struct](#struct_Struct)|暂无描述...
|`STRUCT`|[Type](#struct_Type)|暂无描述...

</details>


***
## 详情信息
#### func NewPackage(dir string) (*Package,  error)
<span id="NewPackage"></span>

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewPackage(t *testing.T) {
	p, err := astgo.NewPackage(`/Users/kercylan/Coding.localized/Go/minotaur/server`)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(super.MarshalIndentJSON(p, "", "  ")))
}

```


</details>


***
<span id="struct_Comment"></span>
### Comment `STRUCT`

```go
type Comment struct {
	Comments []string
	Clear    []string
}
```
<span id="struct_Field"></span>
### Field `STRUCT`

```go
type Field struct {
	Anonymous bool
	Name      string
	Type      *Type
	Comments  *Comment
}
```
<span id="struct_File"></span>
### File `STRUCT`

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
<span id="struct_Function"></span>
### Function `STRUCT`

```go
type Function struct {
	decl        *ast.FuncDecl
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
#### func (*Function) Code()  string
***
<span id="struct_Package"></span>
### Package `STRUCT`

```go
type Package struct {
	Dir       string
	Name      string
	Dirs      []string
	Files     []*File
	Functions map[string]*Function
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
#### func (*Package) GetUnitTest(f *Function)  *Function
***
#### func (*Package) GetExampleTest(f *Function)  *Function
***
#### func (*Package) GetBenchmarkTest(f *Function)  *Function
***
<span id="struct_Struct"></span>
### Struct `STRUCT`

```go
type Struct struct {
	Name      string
	Internal  bool
	Interface bool
	Comments  *Comment
	Generic   []*Field
	Fields    []*Field
	Type      *Type
	Test      bool
}
```
<span id="struct_Type"></span>
### Type `STRUCT`

```go
type Type struct {
	expr      ast.Expr
	Sign      string
	IsPointer bool
	Name      string
}
```
