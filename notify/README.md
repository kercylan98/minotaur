# Notify

notify 包含了对外部第三方通知的实现，如机器人消息等

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/notify)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewManager](#NewManager)|通过指定的 Sender 创建一个通知管理器， senders 包中提供了一些内置的 Sender


> 结构体定义

|结构体|描述
|:--|:--
|[Manager](#manager)|通知管理器，可用于将通知同时发送至多个渠道
|[Notify](#notify)|通用通知接口定义
|[Sender](#sender)|通知发送器接口声明

</details>


#### func NewManager(senders ...Sender)  *Manager
<span id="NewManager"></span>
> 通过指定的 Sender 创建一个通知管理器， senders 包中提供了一些内置的 Sender
***
### Manager
通知管理器，可用于将通知同时发送至多个渠道
```go
type Manager struct {
	senders       []Sender
	notifyChannel chan Notify
	closeChannel  chan struct{}
}
```
#### func (*Manager) PushNotify(notify Notify)
> 推送通知
***
#### func (*Manager) Release()
> 释放通知管理器
***
### Notify
通用通知接口定义
```go
type Notify struct{}
```
### Sender
通知发送器接口声明
```go
type Sender struct{}
```
