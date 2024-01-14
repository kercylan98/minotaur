# Pce



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/pce)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewExporter](#NewExporter)|创建导出器
|[GetFieldGolangType](#GetFieldGolangType)|获取字段的 Golang 类型
|[GetFields](#GetFields)|获取所有内置支持的字段
|[NewLoader](#NewLoader)|创建加载器


> 结构体定义

|结构体|描述
|:--|:--
|[Config](#config)|配置解析接口
|[DataTmpl](#datatmpl)|数据导出模板
|[Exporter](#exporter)|导出器
|[Field](#field)|基本字段类型接口
|[Int](#int)|暂无描述...
|[Int8](#int8)|暂无描述...
|[Int16](#int16)|暂无描述...
|[Int32](#int32)|暂无描述...
|[Int64](#int64)|暂无描述...
|[Uint](#uint)|暂无描述...
|[Uint8](#uint8)|暂无描述...
|[Uint16](#uint16)|暂无描述...
|[Uint32](#uint32)|暂无描述...
|[Uint64](#uint64)|暂无描述...
|[Float32](#float32)|暂无描述...
|[Float64](#float64)|暂无描述...
|[String](#string)|暂无描述...
|[Bool](#bool)|暂无描述...
|[Byte](#byte)|暂无描述...
|[Rune](#rune)|暂无描述...
|[Complex64](#complex64)|暂无描述...
|[Complex128](#complex128)|暂无描述...
|[Uintptr](#uintptr)|暂无描述...
|[Double](#double)|暂无描述...
|[Float](#float)|暂无描述...
|[Long](#long)|暂无描述...
|[Short](#short)|暂无描述...
|[Char](#char)|暂无描述...
|[Number](#number)|暂无描述...
|[Integer](#integer)|暂无描述...
|[Boolean](#boolean)|暂无描述...
|[Loader](#loader)|配置加载器
|[DataInfo](#datainfo)|配置数据
|[DataField](#datafield)|配置数据字段
|[Tmpl](#tmpl)|配置结构模板接口
|[TmplField](#tmplfield)|模板字段
|[TmplStruct](#tmplstruct)|模板结构

</details>


#### func NewExporter()  *Exporter
<span id="NewExporter"></span>
> 创建导出器
***
#### func GetFieldGolangType(field Field)  string
<span id="GetFieldGolangType"></span>
> 获取字段的 Golang 类型
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
### Config
配置解析接口
  - 用于将配置文件解析为可供分析的数据结构
  - 可以在 cs 包中找到内置提供的实现及其模板，例如 cs.XlsxIndexConfig
```go
type Config struct{}
```
### DataTmpl
数据导出模板
```go
type DataTmpl struct{}
```
### Exporter
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
### Field
基本字段类型接口
```go
type Field struct{}
```
### Int

```go
type Int struct{}
```
#### func (Int) TypeName()  string
***
#### func (Int) Zero()  any
***
#### func (Int) Parse(value string)  any
***
### Int8

```go
type Int8 struct{}
```
#### func (Int8) TypeName()  string
***
#### func (Int8) Zero()  any
***
#### func (Int8) Parse(value string)  any
***
### Int16

```go
type Int16 struct{}
```
#### func (Int16) TypeName()  string
***
#### func (Int16) Zero()  any
***
#### func (Int16) Parse(value string)  any
***
### Int32

```go
type Int32 struct{}
```
#### func (Int32) TypeName()  string
***
#### func (Int32) Zero()  any
***
#### func (Int32) Parse(value string)  any
***
### Int64

```go
type Int64 struct{}
```
#### func (Int64) TypeName()  string
***
#### func (Int64) Zero()  any
***
#### func (Int64) Parse(value string)  any
***
### Uint

```go
type Uint struct{}
```
#### func (Uint) TypeName()  string
***
#### func (Uint) Zero()  any
***
#### func (Uint) Parse(value string)  any
***
### Uint8

```go
type Uint8 struct{}
```
#### func (Uint8) TypeName()  string
***
#### func (Uint8) Zero()  any
***
#### func (Uint8) Parse(value string)  any
***
### Uint16

```go
type Uint16 struct{}
```
#### func (Uint16) TypeName()  string
***
#### func (Uint16) Zero()  any
***
#### func (Uint16) Parse(value string)  any
***
### Uint32

```go
type Uint32 struct{}
```
#### func (Uint32) TypeName()  string
***
#### func (Uint32) Zero()  any
***
#### func (Uint32) Parse(value string)  any
***
### Uint64

```go
type Uint64 struct{}
```
#### func (Uint64) TypeName()  string
***
#### func (Uint64) Zero()  any
***
#### func (Uint64) Parse(value string)  any
***
### Float32

```go
type Float32 struct{}
```
#### func (Float32) TypeName()  string
***
#### func (Float32) Zero()  any
***
#### func (Float32) Parse(value string)  any
***
### Float64

```go
type Float64 struct{}
```
#### func (Float64) TypeName()  string
***
#### func (Float64) Zero()  any
***
#### func (Float64) Parse(value string)  any
***
### String

```go
type String struct{}
```
#### func (String) TypeName()  string
***
#### func (String) Zero()  any
***
#### func (String) Parse(value string)  any
***
### Bool

```go
type Bool struct{}
```
#### func (Bool) TypeName()  string
***
#### func (Bool) Zero()  any
***
#### func (Bool) Parse(value string)  any
***
### Byte

```go
type Byte struct{}
```
#### func (Byte) TypeName()  string
***
#### func (Byte) Zero()  any
***
#### func (Byte) Parse(value string)  any
***
### Rune

```go
type Rune struct{}
```
#### func (Rune) TypeName()  string
***
#### func (Rune) Zero()  any
***
#### func (Rune) Parse(value string)  any
***
### Complex64

```go
type Complex64 struct{}
```
#### func (Complex64) TypeName()  string
***
#### func (Complex64) Zero()  any
***
#### func (Complex64) Parse(value string)  any
***
### Complex128

```go
type Complex128 struct{}
```
#### func (Complex128) TypeName()  string
***
#### func (Complex128) Zero()  any
***
#### func (Complex128) Parse(value string)  any
***
### Uintptr

```go
type Uintptr struct{}
```
#### func (Uintptr) TypeName()  string
***
#### func (Uintptr) Zero()  any
***
#### func (Uintptr) Parse(value string)  any
***
### Double

```go
type Double struct{}
```
#### func (Double) TypeName()  string
***
#### func (Double) Zero()  any
***
#### func (Double) Parse(value string)  any
***
### Float

```go
type Float struct{}
```
#### func (Float) TypeName()  string
***
#### func (Float) Zero()  any
***
#### func (Float) Parse(value string)  any
***
### Long

```go
type Long struct{}
```
#### func (Long) TypeName()  string
***
#### func (Long) Zero()  any
***
#### func (Long) Parse(value string)  any
***
### Short

```go
type Short struct{}
```
#### func (Short) TypeName()  string
***
#### func (Short) Zero()  any
***
#### func (Short) Parse(value string)  any
***
### Char

```go
type Char struct{}
```
#### func (Char) TypeName()  string
***
#### func (Char) Zero()  any
***
#### func (Char) Parse(value string)  any
***
### Number

```go
type Number struct{}
```
#### func (Number) TypeName()  string
***
#### func (Number) Zero()  any
***
#### func (Number) Parse(value string)  any
***
### Integer

```go
type Integer struct{}
```
#### func (Integer) TypeName()  string
***
#### func (Integer) Zero()  any
***
#### func (Integer) Parse(value string)  any
***
### Boolean

```go
type Boolean struct{}
```
#### func (Boolean) TypeName()  string
***
#### func (Boolean) Zero()  any
***
#### func (Boolean) Parse(value string)  any
***
### Loader
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
### DataInfo
配置数据
```go
type DataInfo struct {
	DataField
	Value string
}
```
### DataField
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
### Tmpl
配置结构模板接口
```go
type Tmpl struct{}
```
### TmplField
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
### TmplStruct
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
