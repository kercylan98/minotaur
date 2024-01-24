# Notify

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

notify 包含了对外部第三方通知的实现，如机器人消息等


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewManager](#NewManager)|通过指定的 Sender 创建一个通知管理器， senders 包中提供了一些内置的 Sender


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Manager](#struct_Manager)|通知管理器，可用于将通知同时发送至多个渠道
|`INTERFACE`|[Notify](#struct_Notify)|通用通知接口定义
|`INTERFACE`|[Sender](#struct_Sender)|通知发送器接口声明

</details>


***
## 详情信息
#### func NewManager(senders ...Sender) *Manager
<span id="NewManager"></span>
> 通过指定的 Sender 创建一个通知管理器， senders 包中提供了一些内置的 Sender

***
<span id="struct_Manager"></span>
### Manager `STRUCT`
通知管理器，可用于将通知同时发送至多个渠道
```go
type Manager struct {
	senders       []Sender
	notifyChannel chan Notify
	closeChannel  chan struct{}
}
```
<span id="struct_Manager_PushNotify"></span>

#### func (*Manager) PushNotify(notify Notify)
> 推送通知

***
<span id="struct_Manager_Release"></span>

#### func (*Manager) Release()
> 释放通知管理器

***
<span id="struct_Notify"></span>
### Notify `INTERFACE`
通用通知接口定义
```go
type Notify interface {
	Format() (string, error)
}
```
<span id="struct_Sender"></span>
### Sender `INTERFACE`
通知发送器接口声明
```go
type Sender interface {
	Push(notify Notify) error
}
```
