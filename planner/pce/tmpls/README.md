# Tmpls

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
|[NewGolang](#NewGolang)|创建一个 Golang 配置导出模板
|[NewJSON](#NewJSON)|暂无描述...


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Golang](#struct_Golang)|配置导出模板
|`STRUCT`|[JSON](#struct_JSON)|暂无描述...

</details>


***
## 详情信息
#### func NewGolang(packageName string) *Golang
<span id="NewGolang"></span>
> 创建一个 Golang 配置导出模板

***
#### func NewJSON() *JSON
<span id="NewJSON"></span>

***
<span id="struct_Golang"></span>
### Golang `STRUCT`
配置导出模板
```go
type Golang struct {
	Package   string
	Templates []*pce.TmplStruct
}
```
<span id="struct_Golang_Render"></span>

#### func (*Golang) Render(templates ...*pce.TmplStruct) ( string,  error)

***
<span id="struct_Golang_GetVariable"></span>

#### func (*Golang) GetVariable(config *pce.TmplStruct)  string

***
<span id="struct_Golang_GetConfigName"></span>

#### func (*Golang) GetConfigName(config *pce.TmplStruct)  string

***
<span id="struct_Golang_HasIndex"></span>

#### func (*Golang) HasIndex(config *pce.TmplStruct)  bool

***
<span id="struct_JSON"></span>
### JSON `STRUCT`

```go
type JSON struct {
	jsonIter.API
}
```
<span id="struct_JSON_Render"></span>

#### func (*JSON) Render(data map[any]any) ( string,  error)

***
