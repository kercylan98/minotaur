# Router

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
|[NewMultistage](#NewMultistage)|创建一个支持多级分类的路由器
|[WithRouteTrim](#WithRouteTrim)|路由修剪选项


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[MultistageBind](#struct_MultistageBind)|多级分类路由绑定函数
|`STRUCT`|[Multistage](#struct_Multistage)|支持多级分类的路由器
|`STRUCT`|[MultistageOption](#struct_MultistageOption)|路由器选项

</details>


***
## 详情信息
#### func NewMultistage\[HandleFunc any\](options ...MultistageOption[HandleFunc]) *Multistage[HandleFunc]
<span id="NewMultistage"></span>
> 创建一个支持多级分类的路由器

**示例代码：**

```go

func ExampleNewMultistage() {
	router.NewMultistage[func()]()
}

```

***
#### func WithRouteTrim\[HandleFunc any\](handle func (route any)  any) MultistageOption[HandleFunc]
<span id="WithRouteTrim"></span>
> 路由修剪选项
>   - 将在路由注册前对路由进行对应处理

***
<span id="struct_MultistageBind"></span>
### MultistageBind `STRUCT`
多级分类路由绑定函数
```go
type MultistageBind[HandleFunc any] func(HandleFunc)
```
#### func (MultistageBind) Bind(handleFunc HandleFunc)
> 将处理函数绑定到预设的路由中
***
<span id="struct_Multistage"></span>
### Multistage `STRUCT`
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
**示例代码：**

```go

func ExampleMultistage_Register() {
	r := router.NewMultistage[func()]()
	r.Register("System", "Network", "Ping")(func() {
	})
}

```

***
#### func (*Multistage) Route(route any, handleFunc HandleFunc)
> 为特定路由绑定处理函数，被绑定的处理函数将可以通过 Match 函数进行匹配
**示例代码：**

```go

func ExampleMultistage_Route() {
	r := router.NewMultistage[func()]()
	r.Route("ServerTime", func() {
	})
}

```

***
#### func (*Multistage) Match(routes ...any)  HandleFunc
> 匹配已绑定处理函数的路由，返回处理函数
>   - 如果未找到将会返回空指针
**示例代码：**

```go

func ExampleMultistage_Match() {
	r := router.NewMultistage[func()]()
	r.Route("ServerTime", func() {
	})
	r.Register("System", "Network", "Ping").Bind(func() {
	})
	r.Match("ServerTime")()
	r.Match("System", "Network", "Ping")()
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestMultistage_Match(t *testing.T) {
	r := router.NewMultistage[func()]()
	r.Sub("System").Route("Heartbeat", func() {
		fmt.Println("Heartbeat")
	})
	r.Route("ServerTime", func() {
		fmt.Println("ServerTime")
	})
	r.Register("System", "Network", "Ping")(func() {
		fmt.Println("Ping")
	})
	r.Register("System", "Network", "Echo").Bind(onEcho)
	r.Match("System", "Heartbeat")()
	r.Match("ServerTime")()
	r.Match("System", "Network", "Ping")()
	r.Match("System", "Network", "Echo")()
	fmt.Println(r.Match("None") == nil)
}

```


</details>


***
#### func (*Multistage) Sub(route any)  *Multistage[HandleFunc]
> 获取子路由器
**示例代码：**

```go

func ExampleMultistage_Sub() {
	r := router.NewMultistage[func()]()
	r.Sub("System").Route("Heartbeat", func() {
	})
}

```

***
<span id="struct_MultistageOption"></span>
### MultistageOption `STRUCT`
路由器选项
```go
type MultistageOption[HandleFunc any] func(multistage *Multistage[HandleFunc])
```
