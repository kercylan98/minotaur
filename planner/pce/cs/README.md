# Cs



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/cs)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewXlsx](#NewXlsx)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[XlsxExportType](#xlsxexporttype)|暂无描述...
|[Xlsx](#xlsx)|内置的 Xlsx 配置

</details>


#### func NewXlsx(sheet *xlsx.Sheet, exportType XlsxExportType)  *Xlsx
<span id="NewXlsx"></span>
***
### XlsxExportType

```go
type XlsxExportType struct{}
```
### Xlsx
内置的 Xlsx 配置
```go
type Xlsx struct {
	sheet      *xlsx.Sheet
	exportType XlsxExportType
}
```
#### func (*Xlsx) GetConfigName()  string
***
#### func (*Xlsx) GetDisplayName()  string
***
#### func (*Xlsx) GetDescription()  string
***
#### func (*Xlsx) GetIndexCount()  int
***
#### func (*Xlsx) GetFields()  []pce.DataField
***
#### func (*Xlsx) GetData()  [][]pce.DataInfo
***
