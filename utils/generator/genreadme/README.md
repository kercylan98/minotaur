# Genreadme



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/genreadme)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[New](#New)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[Builder](#builder)|暂无描述...

</details>


#### func New(pkgDirPath string, output string)  *Builder,  error
<span id="New"></span>
***
### Builder

```go
type Builder struct {
	p *astgo.Package
	b *strings.Builder
	o string
}
```
#### func (*Builder) Generate()  error
***
