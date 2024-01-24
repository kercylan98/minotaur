# Cs

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
|[NewXlsx](#NewXlsx)|暂无描述...


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[XlsxExportType](#struct_XlsxExportType)|暂无描述...
|`STRUCT`|[Xlsx](#struct_Xlsx)|内置的 Xlsx 配置

</details>


***
## 详情信息
#### func NewXlsx(sheet *xlsx.Sheet, exportType XlsxExportType) *Xlsx
<span id="NewXlsx"></span>

***
<span id="struct_XlsxExportType"></span>
### XlsxExportType `STRUCT`

```go
type XlsxExportType int
```
<span id="struct_Xlsx"></span>
### Xlsx `STRUCT`
内置的 Xlsx 配置
```go
type Xlsx struct {
	sheet      *xlsx.Sheet
	exportType XlsxExportType
}
```
<span id="struct_Xlsx_GetConfigName"></span>

#### func (*Xlsx) GetConfigName()  string

***
<span id="struct_Xlsx_GetDisplayName"></span>

#### func (*Xlsx) GetDisplayName()  string

***
<span id="struct_Xlsx_GetDescription"></span>

#### func (*Xlsx) GetDescription()  string

***
<span id="struct_Xlsx_GetIndexCount"></span>

#### func (*Xlsx) GetIndexCount()  int

***
<span id="struct_Xlsx_GetFields"></span>

#### func (*Xlsx) GetFields()  []pce.DataField

***
<span id="struct_Xlsx_GetData"></span>

#### func (*Xlsx) GetData()  [][]pce.DataInfo

***
