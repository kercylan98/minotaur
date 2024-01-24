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
|`INTERFACE`|[Config](#struct_Config)|配置解析接口
|`INTERFACE`|[DataTmpl](#struct_DataTmpl)|数据导出模板
|`STRUCT`|[Exporter](#struct_Exporter)|导出器
|`INTERFACE`|[Field](#struct_Field)|基本字段类型接口
|`STRUCT`|[Int](#struct_Int)|暂无描述...
|`STRUCT`|[Int8](#struct_Int8)|暂无描述...
|`STRUCT`|[Int16](#struct_Int16)|暂无描述...
|`STRUCT`|[Int32](#struct_Int32)|暂无描述...
|`STRUCT`|[Int64](#struct_Int64)|暂无描述...
|`STRUCT`|[Uint](#struct_Uint)|暂无描述...
|`STRUCT`|[Uint8](#struct_Uint8)|暂无描述...
|`STRUCT`|[Uint16](#struct_Uint16)|暂无描述...
|`STRUCT`|[Uint32](#struct_Uint32)|暂无描述...
|`STRUCT`|[Uint64](#struct_Uint64)|暂无描述...
|`STRUCT`|[Float32](#struct_Float32)|暂无描述...
|`STRUCT`|[Float64](#struct_Float64)|暂无描述...
|`STRUCT`|[String](#struct_String)|暂无描述...
|`STRUCT`|[Bool](#struct_Bool)|暂无描述...
|`STRUCT`|[Byte](#struct_Byte)|暂无描述...
|`STRUCT`|[Rune](#struct_Rune)|暂无描述...
|`STRUCT`|[Complex64](#struct_Complex64)|暂无描述...
|`STRUCT`|[Complex128](#struct_Complex128)|暂无描述...
|`STRUCT`|[Uintptr](#struct_Uintptr)|暂无描述...
|`STRUCT`|[Double](#struct_Double)|暂无描述...
|`STRUCT`|[Float](#struct_Float)|暂无描述...
|`STRUCT`|[Long](#struct_Long)|暂无描述...
|`STRUCT`|[Short](#struct_Short)|暂无描述...
|`STRUCT`|[Char](#struct_Char)|暂无描述...
|`STRUCT`|[Number](#struct_Number)|暂无描述...
|`STRUCT`|[Integer](#struct_Integer)|暂无描述...
|`STRUCT`|[Boolean](#struct_Boolean)|暂无描述...
|`STRUCT`|[Loader](#struct_Loader)|配置加载器
|`STRUCT`|[DataInfo](#struct_DataInfo)|配置数据
|`STRUCT`|[DataField](#struct_DataField)|配置数据字段
|`INTERFACE`|[Tmpl](#struct_Tmpl)|配置结构模板接口
|`STRUCT`|[TmplField](#struct_TmplField)|模板字段
|`STRUCT`|[TmplStruct](#struct_TmplStruct)|模板结构

</details>


***
## 详情信息
#### func NewExporter() *Exporter
<span id="NewExporter"></span>
> 创建导出器

***
#### func GetFieldGolangType(field Field) string
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
#### func GetFields() []Field
<span id="GetFields"></span>
> 获取所有内置支持的字段

***
#### func NewLoader(fields []Field) *Loader
<span id="NewLoader"></span>
> 创建加载器
>   - 加载器被用于加载配置表的数据和结构信息

***
<span id="struct_Config"></span>
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
<span id="struct_DataTmpl"></span>
### DataTmpl `INTERFACE`
数据导出模板
```go
type DataTmpl interface {
	Render(data map[any]any) (string, error)
}
```
<span id="struct_Exporter"></span>
### Exporter `STRUCT`
导出器
```go
type Exporter struct{}
```
<span id="struct_Exporter_ExportStruct"></span>

#### func (*Exporter) ExportStruct(tmpl Tmpl, tmplStruct ...*TmplStruct) ( []byte,  error)
> 导出结构

***
<span id="struct_Exporter_ExportData"></span>

#### func (*Exporter) ExportData(tmpl DataTmpl, data map[any]any) ( []byte,  error)
> 导出数据

***
<span id="struct_Field"></span>
### Field `INTERFACE`
基本字段类型接口
```go
type Field interface {
	TypeName() string
	Zero() any
	Parse(value string) any
}
```
<span id="struct_Int"></span>
### Int `STRUCT`

```go
type Int int
```
<span id="struct_Int_TypeName"></span>

#### func (Int) TypeName()  string

***
<span id="struct_Int_Zero"></span>

#### func (Int) Zero()  any

***
<span id="struct_Int_Parse"></span>

#### func (Int) Parse(value string)  any

***
<span id="struct_Int8"></span>
### Int8 `STRUCT`

```go
type Int8 int8
```
<span id="struct_Int8_TypeName"></span>

#### func (Int8) TypeName()  string

***
<span id="struct_Int8_Zero"></span>

#### func (Int8) Zero()  any

***
<span id="struct_Int8_Parse"></span>

#### func (Int8) Parse(value string)  any

***
<span id="struct_Int16"></span>
### Int16 `STRUCT`

```go
type Int16 int16
```
<span id="struct_Int16_TypeName"></span>

#### func (Int16) TypeName()  string

***
<span id="struct_Int16_Zero"></span>

#### func (Int16) Zero()  any

***
<span id="struct_Int16_Parse"></span>

#### func (Int16) Parse(value string)  any

***
<span id="struct_Int32"></span>
### Int32 `STRUCT`

```go
type Int32 int32
```
<span id="struct_Int32_TypeName"></span>

#### func (Int32) TypeName()  string

***
<span id="struct_Int32_Zero"></span>

#### func (Int32) Zero()  any

***
<span id="struct_Int32_Parse"></span>

#### func (Int32) Parse(value string)  any

***
<span id="struct_Int64"></span>
### Int64 `STRUCT`

```go
type Int64 int64
```
<span id="struct_Int64_TypeName"></span>

#### func (Int64) TypeName()  string

***
<span id="struct_Int64_Zero"></span>

#### func (Int64) Zero()  any

***
<span id="struct_Int64_Parse"></span>

#### func (Int64) Parse(value string)  any

***
<span id="struct_Uint"></span>
### Uint `STRUCT`

```go
type Uint uint
```
<span id="struct_Uint_TypeName"></span>

#### func (Uint) TypeName()  string

***
<span id="struct_Uint_Zero"></span>

#### func (Uint) Zero()  any

***
<span id="struct_Uint_Parse"></span>

#### func (Uint) Parse(value string)  any

***
<span id="struct_Uint8"></span>
### Uint8 `STRUCT`

```go
type Uint8 uint8
```
<span id="struct_Uint8_TypeName"></span>

#### func (Uint8) TypeName()  string

***
<span id="struct_Uint8_Zero"></span>

#### func (Uint8) Zero()  any

***
<span id="struct_Uint8_Parse"></span>

#### func (Uint8) Parse(value string)  any

***
<span id="struct_Uint16"></span>
### Uint16 `STRUCT`

```go
type Uint16 uint16
```
<span id="struct_Uint16_TypeName"></span>

#### func (Uint16) TypeName()  string

***
<span id="struct_Uint16_Zero"></span>

#### func (Uint16) Zero()  any

***
<span id="struct_Uint16_Parse"></span>

#### func (Uint16) Parse(value string)  any

***
<span id="struct_Uint32"></span>
### Uint32 `STRUCT`

```go
type Uint32 uint32
```
<span id="struct_Uint32_TypeName"></span>

#### func (Uint32) TypeName()  string

***
<span id="struct_Uint32_Zero"></span>

#### func (Uint32) Zero()  any

***
<span id="struct_Uint32_Parse"></span>

#### func (Uint32) Parse(value string)  any

***
<span id="struct_Uint64"></span>
### Uint64 `STRUCT`

```go
type Uint64 uint64
```
<span id="struct_Uint64_TypeName"></span>

#### func (Uint64) TypeName()  string

***
<span id="struct_Uint64_Zero"></span>

#### func (Uint64) Zero()  any

***
<span id="struct_Uint64_Parse"></span>

#### func (Uint64) Parse(value string)  any

***
<span id="struct_Float32"></span>
### Float32 `STRUCT`

```go
type Float32 float32
```
<span id="struct_Float32_TypeName"></span>

#### func (Float32) TypeName()  string

***
<span id="struct_Float32_Zero"></span>

#### func (Float32) Zero()  any

***
<span id="struct_Float32_Parse"></span>

#### func (Float32) Parse(value string)  any

***
<span id="struct_Float64"></span>
### Float64 `STRUCT`

```go
type Float64 float64
```
<span id="struct_Float64_TypeName"></span>

#### func (Float64) TypeName()  string

***
<span id="struct_Float64_Zero"></span>

#### func (Float64) Zero()  any

***
<span id="struct_Float64_Parse"></span>

#### func (Float64) Parse(value string)  any

***
<span id="struct_String"></span>
### String `STRUCT`

```go
type String string
```
<span id="struct_String_TypeName"></span>

#### func (String) TypeName()  string

***
<span id="struct_String_Zero"></span>

#### func (String) Zero()  any

***
<span id="struct_String_Parse"></span>

#### func (String) Parse(value string)  any

***
<span id="struct_Bool"></span>
### Bool `STRUCT`

```go
type Bool bool
```
<span id="struct_Bool_TypeName"></span>

#### func (Bool) TypeName()  string

***
<span id="struct_Bool_Zero"></span>

#### func (Bool) Zero()  any

***
<span id="struct_Bool_Parse"></span>

#### func (Bool) Parse(value string)  any

***
<span id="struct_Byte"></span>
### Byte `STRUCT`

```go
type Byte byte
```
<span id="struct_Byte_TypeName"></span>

#### func (Byte) TypeName()  string

***
<span id="struct_Byte_Zero"></span>

#### func (Byte) Zero()  any

***
<span id="struct_Byte_Parse"></span>

#### func (Byte) Parse(value string)  any

***
<span id="struct_Rune"></span>
### Rune `STRUCT`

```go
type Rune rune
```
<span id="struct_Rune_TypeName"></span>

#### func (Rune) TypeName()  string

***
<span id="struct_Rune_Zero"></span>

#### func (Rune) Zero()  any

***
<span id="struct_Rune_Parse"></span>

#### func (Rune) Parse(value string)  any

***
<span id="struct_Complex64"></span>
### Complex64 `STRUCT`

```go
type Complex64 complex64
```
<span id="struct_Complex64_TypeName"></span>

#### func (Complex64) TypeName()  string

***
<span id="struct_Complex64_Zero"></span>

#### func (Complex64) Zero()  any

***
<span id="struct_Complex64_Parse"></span>

#### func (Complex64) Parse(value string)  any

***
<span id="struct_Complex128"></span>
### Complex128 `STRUCT`

```go
type Complex128 complex128
```
<span id="struct_Complex128_TypeName"></span>

#### func (Complex128) TypeName()  string

***
<span id="struct_Complex128_Zero"></span>

#### func (Complex128) Zero()  any

***
<span id="struct_Complex128_Parse"></span>

#### func (Complex128) Parse(value string)  any

***
<span id="struct_Uintptr"></span>
### Uintptr `STRUCT`

```go
type Uintptr uintptr
```
<span id="struct_Uintptr_TypeName"></span>

#### func (Uintptr) TypeName()  string

***
<span id="struct_Uintptr_Zero"></span>

#### func (Uintptr) Zero()  any

***
<span id="struct_Uintptr_Parse"></span>

#### func (Uintptr) Parse(value string)  any

***
<span id="struct_Double"></span>
### Double `STRUCT`

```go
type Double float64
```
<span id="struct_Double_TypeName"></span>

#### func (Double) TypeName()  string

***
<span id="struct_Double_Zero"></span>

#### func (Double) Zero()  any

***
<span id="struct_Double_Parse"></span>

#### func (Double) Parse(value string)  any

***
<span id="struct_Float"></span>
### Float `STRUCT`

```go
type Float float32
```
<span id="struct_Float_TypeName"></span>

#### func (Float) TypeName()  string

***
<span id="struct_Float_Zero"></span>

#### func (Float) Zero()  any

***
<span id="struct_Float_Parse"></span>

#### func (Float) Parse(value string)  any

***
<span id="struct_Long"></span>
### Long `STRUCT`

```go
type Long int64
```
<span id="struct_Long_TypeName"></span>

#### func (Long) TypeName()  string

***
<span id="struct_Long_Zero"></span>

#### func (Long) Zero()  any

***
<span id="struct_Long_Parse"></span>

#### func (Long) Parse(value string)  any

***
<span id="struct_Short"></span>
### Short `STRUCT`

```go
type Short int16
```
<span id="struct_Short_TypeName"></span>

#### func (Short) TypeName()  string

***
<span id="struct_Short_Zero"></span>

#### func (Short) Zero()  any

***
<span id="struct_Short_Parse"></span>

#### func (Short) Parse(value string)  any

***
<span id="struct_Char"></span>
### Char `STRUCT`

```go
type Char int8
```
<span id="struct_Char_TypeName"></span>

#### func (Char) TypeName()  string

***
<span id="struct_Char_Zero"></span>

#### func (Char) Zero()  any

***
<span id="struct_Char_Parse"></span>

#### func (Char) Parse(value string)  any

***
<span id="struct_Number"></span>
### Number `STRUCT`

```go
type Number float64
```
<span id="struct_Number_TypeName"></span>

#### func (Number) TypeName()  string

***
<span id="struct_Number_Zero"></span>

#### func (Number) Zero()  any

***
<span id="struct_Number_Parse"></span>

#### func (Number) Parse(value string)  any

***
<span id="struct_Integer"></span>
### Integer `STRUCT`

```go
type Integer int64
```
<span id="struct_Integer_TypeName"></span>

#### func (Integer) TypeName()  string

***
<span id="struct_Integer_Zero"></span>

#### func (Integer) Zero()  any

***
<span id="struct_Integer_Parse"></span>

#### func (Integer) Parse(value string)  any

***
<span id="struct_Boolean"></span>
### Boolean `STRUCT`

```go
type Boolean bool
```
<span id="struct_Boolean_TypeName"></span>

#### func (Boolean) TypeName()  string

***
<span id="struct_Boolean_Zero"></span>

#### func (Boolean) Zero()  any

***
<span id="struct_Boolean_Parse"></span>

#### func (Boolean) Parse(value string)  any

***
<span id="struct_Loader"></span>
### Loader `STRUCT`
配置加载器
```go
type Loader struct {
	fields map[string]Field
}
```
<span id="struct_Loader_LoadStruct"></span>

#### func (*Loader) LoadStruct(config Config)  *TmplStruct
> 加载结构

***
<span id="struct_Loader_LoadData"></span>

#### func (*Loader) LoadData(config Config)  map[any]any
> 加载配置并得到配置数据

***
<span id="struct_DataInfo"></span>
### DataInfo `STRUCT`
配置数据
```go
type DataInfo struct {
	DataField
	Value string
}
```
<span id="struct_DataField"></span>
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
<span id="struct_Tmpl"></span>
### Tmpl `INTERFACE`
配置结构模板接口
```go
type Tmpl interface {
	Render(templates ...*TmplStruct) (string, error)
}
```
<span id="struct_TmplField"></span>
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
<span id="struct_TmplField_IsIndex"></span>

#### func (*TmplField) IsIndex()  bool
> 是否是索引字段

***
<span id="struct_TmplField_IsStruct"></span>

#### func (*TmplField) IsStruct()  bool
> 是否是结构类型

***
<span id="struct_TmplField_IsSlice"></span>

#### func (*TmplField) IsSlice()  bool
> 是否是切片类型

***
<span id="struct_TmplStruct"></span>
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
<span id="struct_TmplStruct_AllChildren"></span>

#### func (*TmplStruct) AllChildren()  []*TmplStruct
> 获取所有子结构

***
