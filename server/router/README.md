# Router



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/router)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewMultistage](#NewMultistage)|创建一个支持多级分类的路由器
|[WithRouteTrim](#WithRouteTrim)|路由修剪选项


> 结构体定义

|结构体|描述
|:--|:--
|[MultistageBind](#multistagebind)|多级分类路由绑定函数
|[Multistage](#multistage)|支持多级分类的路由器
|[MultistageOption](#multistageoption)|路由器选项

</details>


#### func NewMultistage(options ...MultistageOption[HandleFunc])  *Multistage[HandleFunc]
<span id="NewMultistage"></span>
> 创建一个支持多级分类的路由器
***
#### func WithRouteTrim(handle func (route any)  any)  MultistageOption[HandleFunc]
<span id="WithRouteTrim"></span>
> 路由修剪选项
>   - 将在路由注册前对路由进行对应处理
***
### MultistageBind
多级分类路由绑定函数
```go
type MultistageBind[HandleFunc any] struct{}
```
#### func (MultistageBind) Bind(handleFunc HandleFunc)
> 将处理函数绑定到预设的路由中
***
### Multistage
支持多级分类的路由器
```go
type Multistage[HandleFunc any] struct {
	routes map[any]HandleFunc
	subs   map[any]*Multistage[HandleFunc]
	tag    any
	trim   func(route any) any
}
```
#### func (*Multistage) Register(routes ...any)  MultistageBind[HandleFunc]
> 注册路由是结合 Sub 和 Route 的快捷方式，用于一次性注册多级路由
>   - 该函数将返回一个注册函数，可通过调用其将路由绑定到特定处理函数，例如：router.Register("a", "b").Bind(onExec())
***
#### func (*Multistage) Route(route any, handleFunc HandleFunc)
> 为特定路由绑定处理函数，被绑定的处理函数将可以通过 Match 函数进行匹配
***
#### func (*Multistage) Match(routes ...any)  HandleFunc
> 匹配已绑定处理函数的路由，返回处理函数
>   - 如果未找到将会返回空指针
***
#### func (*Multistage) Sub(route any)  *Multistage[HandleFunc]
> 获取子路由器
***
### MultistageOption
路由器选项
```go
type MultistageOption[HandleFunc any] struct{}
```
