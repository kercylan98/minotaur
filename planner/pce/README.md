# Pce

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
|[NewExporter](#NewExporter)|创建导出器
|[GetFieldGolangType](#GetFieldGolangType)|获取字段的 Golang 类型
|[GetFields](#GetFields)|获取所有内置支持的字段
|[NewLoader](#NewLoader)|创建加载器


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`INTERFACE`|[Config](#config)|配置解析接口
|`INTERFACE`|[DataTmpl](#datatmpl)|数据导出模板
|`STRUCT`|[Exporter](#exporter)|导出器
|`INTERFACE`|[Field](#field)|基本字段类型接口
|`STRUCT`|[Int](#int)|暂无描述...
|`STRUCT`|[Int8](#int8)|暂无描述...
|`STRUCT`|[Int16](#int16)|暂无描述...
|`STRUCT`|[Int32](#int32)|暂无描述...
|`STRUCT`|[Int64](#int64)|暂无描述...
|`STRUCT`|[Uint](#uint)|暂无描述...
|`STRUCT`|[Uint8](#uint8)|暂无描述...
|`STRUCT`|[Uint16](#uint16)|暂无描述...
|`STRUCT`|[Uint32](#uint32)|暂无描述...
|`STRUCT`|[Uint64](#uint64)|暂无描述...
|`STRUCT`|[Float32](#float32)|暂无描述...
|`STRUCT`|[Float64](#float64)|暂无描述...
|`STRUCT`|[String](#string)|暂无描述...
|`STRUCT`|[Bool](#bool)|暂无描述...
|`STRUCT`|[Byte](#byte)|暂无描述...
|`STRUCT`|[Rune](#rune)|暂无描述...
|`STRUCT`|[Complex64](#complex64)|暂无描述...
|`STRUCT`|[Complex128](#complex128)|暂无描述...
|`STRUCT`|[Uintptr](#uintptr)|暂无描述...
|`STRUCT`|[Double](#double)|暂无描述...
|`STRUCT`|[Float](#float)|暂无描述...
|`STRUCT`|[Long](#long)|暂无描述...
|`STRUCT`|[Short](#short)|暂无描述...
|`STRUCT`|[Char](#char)|暂无描述...
|`STRUCT`|[Number](#number)|暂无描述...
|`STRUCT`|[Integer](#integer)|暂无描述...
|`STRUCT`|[Boolean](#boolean)|暂无描述...
|`STRUCT`|[Loader](#loader)|配置加载器
|`STRUCT`|[DataInfo](#datainfo)|配置数据
|`STRUCT`|[DataField](#datafield)|配置数据字段
|`INTERFACE`|[Tmpl](#tmpl)|配置结构模板接口
|`STRUCT`|[TmplField](#tmplfield)|模板字段
|`STRUCT`|[TmplStruct](#tmplstruct)|模板结构

</details>


***
## 详情信息
#### func NewExporter()  *Exporter
<span id="NewExporter"></span>
> 创建导出器

***
#### func GetFieldGolangType(field Field)  string
<span id="GetFieldGolangType"></span>
> 获取字段的 Golang 类型

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestGetFieldGolangType(t *testing.T) {
	fmt.Println(pce.GetFieldGolangType(new(pce.String)))
}

```


</details>


***
#### func GetFields()  []Field
<span id="GetFields"></span>
> 获取所有内置支持的字段

***
#### func NewLoader(fields []Field)  *Loader
<span id="NewLoader"></span>
> 创建加载器
>   - 加载器被用于加载配置表的数据和结构信息

***
### Config `INTERFACE`
配置解析接口
  - 用于将配置文件解析为可供分析的数据结构
  - 可以在 cs 包中找到内置提供的实现及其模板，例如 cs.XlsxIndexConfig
```go
type Config interface {
	GetConfigName() string
	GetDisplayName() string
	GetDescription() string
	GetIndexCount() int
	GetFields() []DataField
	GetData() [][]DataInfo
}
```
### DataTmpl `INTERFACE`
数据导出模板
```go
type DataTmpl interface {
	Render(data map[any]any) (string, error)
}
```
### Exporter `STRUCT`
导出器
```go
type Exporter struct{}
```
#### func (*Exporter) ExportStruct(tmpl Tmpl, tmplStruct ...*TmplStruct)  []byte,  error
> 导出结构
***
#### func (*Exporter) ExportData(tmpl DataTmpl, data map[any]any)  []byte,  error
> 导出数据
***
### Field `INTERFACE`
基本字段类型接口
```go
type Field interface {
	TypeName() string
	Zero() any
	Parse(value string) any
}
```
### Int `STRUCT`

```go
type Int int
```
#### func (Int) TypeName()  string
***
#### func (Int) Zero()  any
***
#### func (Int) Parse(value string)  any
***
### Int8 `STRUCT`

```go
type Int8 int8
```
#### func (Int8) TypeName()  string
***
#### func (Int8) Zero()  any
***
#### func (Int8) Parse(value string)  any
***
### Int16 `STRUCT`

```go
type Int16 int16
```
#### func (Int16) TypeName()  string
***
#### func (Int16) Zero()  any
***
#### func (Int16) Parse(value string)  any
***
### Int32 `STRUCT`

```go
type Int32 int32
```
#### func (Int32) TypeName()  string
***
#### func (Int32) Zero()  any
***
#### func (Int32) Parse(value string)  any
***
### Int64 `STRUCT`

```go
type Int64 int64
```
#### func (Int64) TypeName()  string
***
#### func (Int64) Zero()  any
***
#### func (Int64) Parse(value string)  any
***
### Uint `STRUCT`

```go
type Uint uint
```
#### func (Uint) TypeName()  string
***
#### func (Uint) Zero()  any
***
#### func (Uint) Parse(value string)  any
***
### Uint8 `STRUCT`

```go
type Uint8 uint8
```
#### func (Uint8) TypeName()  string
***
#### func (Uint8) Zero()  any
***
#### func (Uint8) Parse(value string)  any
***
### Uint16 `STRUCT`

```go
type Uint16 uint16
```
#### func (Uint16) TypeName()  string
***
#### func (Uint16) Zero()  any
***
#### func (Uint16) Parse(value string)  any
***
### Uint32 `STRUCT`

```go
type Uint32 uint32
```
#### func (Uint32) TypeName()  string
***
#### func (Uint32) Zero()  any
***
#### func (Uint32) Parse(value string)  any
***
### Uint64 `STRUCT`

```go
type Uint64 uint64
```
#### func (Uint64) TypeName()  string
***
#### func (Uint64) Zero()  any
***
#### func (Uint64) Parse(value string)  any
***
### Float32 `STRUCT`

```go
type Float32 float32
```
#### func (Float32) TypeName()  string
***
#### func (Float32) Zero()  any
***
#### func (Float32) Parse(value string)  any
***
### Float64 `STRUCT`

```go
type Float64 float64
```
#### func (Float64) TypeName()  string
***
#### func (Float64) Zero()  any
***
#### func (Float64) Parse(value string)  any
***
### String `STRUCT`

```go
type String string
```
#### func (String) TypeName()  string
***
#### func (String) Zero()  any
***
#### func (String) Parse(value string)  any
***
### Bool `STRUCT`

```go
type Bool bool
```
#### func (Bool) TypeName()  string
***
#### func (Bool) Zero()  any
***
#### func (Bool) Parse(value string)  any
***
### Byte `STRUCT`

```go
type Byte byte
```
#### func (Byte) TypeName()  string
***
#### func (Byte) Zero()  any
***
#### func (Byte) Parse(value string)  any
***
### Rune `STRUCT`

```go
type Rune rune
```
#### func (Rune) TypeName()  string
***
#### func (Rune) Zero()  any
***
#### func (Rune) Parse(value string)  any
***
### Complex64 `STRUCT`

```go
type Complex64 complex64
```
#### func (Complex64) TypeName()  string
***
#### func (Complex64) Zero()  any
***
#### func (Complex64) Parse(value string)  any
***
### Complex128 `STRUCT`

```go
type Complex128 complex128
```
#### func (Complex128) TypeName()  string
***
#### func (Complex128) Zero()  any
***
#### func (Complex128) Parse(value string)  any
***
### Uintptr `STRUCT`

```go
type Uintptr uintptr
```
#### func (Uintptr) TypeName()  string
***
#### func (Uintptr) Zero()  any
***
#### func (Uintptr) Parse(value string)  any
***
### Double `STRUCT`

```go
type Double float64
```
#### func (Double) TypeName()  string
***
#### func (Double) Zero()  any
***
#### func (Double) Parse(value string)  any
***
### Float `STRUCT`

```go
type Float float32
```
#### func (Float) TypeName()  string
***
#### func (Float) Zero()  any
***
#### func (Float) Parse(value string)  any
***
### Long `STRUCT`

```go
type Long int64
```
#### func (Long) TypeName()  string
***
#### func (Long) Zero()  any
***
#### func (Long) Parse(value string)  any
***
### Short `STRUCT`

```go
type Short int16
```
#### func (Short) TypeName()  string
***
#### func (Short) Zero()  any
***
#### func (Short) Parse(value string)  any
***
### Char `STRUCT`

```go
type Char int8
```
#### func (Char) TypeName()  string
***
#### func (Char) Zero()  any
***
#### func (Char) Parse(value string)  any
***
### Number `STRUCT`

```go
type Number float64
```
#### func (Number) TypeName()  string
***
#### func (Number) Zero()  any
***
#### func (Number) Parse(value string)  any
***
### Integer `STRUCT`

```go
type Integer int64
```
#### func (Integer) TypeName()  string
***
#### func (Integer) Zero()  any
***
#### func (Integer) Parse(value string)  any
***
### Boolean `STRUCT`

```go
type Boolean bool
```
#### func (Boolean) TypeName()  string
***
#### func (Boolean) Zero()  any
***
#### func (Boolean) Parse(value string)  any
***
### Loader `STRUCT`
配置加载器
```go
type Loader struct {
	fields map[string]Field
}
```
#### func (*Loader) LoadStruct(config Config)  *TmplStruct
> 加载结构
***
#### func (*Loader) LoadData(config Config)  map[any]any
> 加载配置并得到配置数据
***
### DataInfo `STRUCT`
配置数据
```go
type DataInfo struct {
	DataField
	Value string
}
```
### DataField `STRUCT`
配置数据字段
```go
type DataField struct {
	Index      int
	Name       string
	Desc       string
	Type       string
	ExportType string
}
```
### Tmpl `INTERFACE`
配置结构模板接口
```go
type Tmpl interface {
	Render(templates ...*TmplStruct) (string, error)
}
```
### TmplField `STRUCT`
模板字段
```go
type TmplField struct {
	Name    string
	Desc    string
	Type    string
	Struct  *TmplStruct
	Index   int
	slice   bool
	isIndex bool
}
```
#### func (*TmplField) IsIndex()  bool
> 是否是索引字段
***
#### func (*TmplField) IsStruct()  bool
> 是否是结构类型
***
#### func (*TmplField) IsSlice()  bool
> 是否是切片类型
***
### TmplStruct `STRUCT`
模板结构
```go
type TmplStruct struct {
	Name       string
	Desc       string
	Fields     []*TmplField
	IndexCount int
}
```
#### func (*TmplStruct) AllChildren()  []*TmplStruct
> 获取所有子结构
***
