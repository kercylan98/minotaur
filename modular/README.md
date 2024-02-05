# Modular

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
|[Run](#Run)|运行模块化应用程序
|[RegisterServices](#RegisterServices)|注册服务


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`INTERFACE`|[Block](#struct_Block)|标识模块化服务为阻塞进程的服务，当实现了 Service 且实现了 Block 接口时，模块化应用程序会在 Service.OnMount 阶段完成后执行 OnBlock 函数
|`INTERFACE`|[Service](#struct_Service)|模块化服务接口，所有的服务均需要实现该接口，在服务的生命周期内发生任何错误均应通过 panic 阻止服务继续运行

</details>


***
## 详情信息
#### func Run()
<span id="Run"></span>
> 运行模块化应用程序

***
#### func RegisterServices(s ...Service)
<span id="RegisterServices"></span>
> 注册服务

***
<span id="struct_Block"></span>
### Block `INTERFACE`
标识模块化服务为阻塞进程的服务，当实现了 Service 且实现了 Block 接口时，模块化应用程序会在 Service.OnMount 阶段完成后执行 OnBlock 函数

该接口适用于 Http 服务、WebSocket 服务等需要阻塞进程的服务。需要注意的是， OnBlock 的执行不能保证按照 Service 的注册顺序执行
```go
type Block interface {
	Service
	OnBlock()
}
```
<span id="struct_Service"></span>
### Service `INTERFACE`
模块化服务接口，所有的服务均需要实现该接口，在服务的生命周期内发生任何错误均应通过 panic 阻止服务继续运行
  - 生命周期示例： OnInit -> OnPreload -> OnMount

在 Golang 中，包与包之间互相引用会导致循环依赖，因此在模块化应用程序中，所有的服务均不应该直接引用其他服务。

服务应该在 OnInit 阶段将不依赖其他服务的内容初始化完成，并且如果服务需要暴露给其他服务调用，那么也应该在 OnInit 阶段完成对外暴露。
  - 暴露方式可参考 modular/example

在 OnPreload 阶段，服务应该完成对其依赖服务的依赖注入，最终在 OnMount 阶段完成对服务功能的定义、路由的声明等。
```go
type Service interface {
	OnInit()
	OnPreload()
	OnMount()
}
```
