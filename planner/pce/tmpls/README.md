# Tmpls



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/tmpls)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewGolang](#NewGolang)|创建一个 Golang 配置导出模板
|[NewJSON](#NewJSON)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[Golang](#golang)|配置导出模板
|[JSON](#json)|暂无描述...

</details>


#### func NewGolang(packageName string)  *Golang
<span id="NewGolang"></span>
> 创建一个 Golang 配置导出模板
***
#### func NewJSON()  *JSON
<span id="NewJSON"></span>
***
### Golang
配置导出模板
```go
type Golang struct {
	Package   string
	Templates []*pce.TmplStruct
}
```
#### func (*Golang) Render(templates ...*pce.TmplStruct)  string,  error
***
#### func (*Golang) GetVariable(config *pce.TmplStruct)  string
***
#### func (*Golang) GetConfigName(config *pce.TmplStruct)  string
***
#### func (*Golang) HasIndex(config *pce.TmplStruct)  bool
***
### JSON

```go
type JSON struct {
	jsonIter.API
}
```
#### func (*JSON) Render(data map[any]any)  string,  error
***
