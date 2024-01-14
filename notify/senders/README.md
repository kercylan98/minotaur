# Senders

senders Package 包含了内置通知发送器的实现

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/senders)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewFeiShu](#NewFeiShu)|根据特定的 webhook 地址创建飞书发送器


> 结构体定义

|结构体|描述
|:--|:--
|[FeiShu](#feishu)|飞书发送器

</details>


#### func NewFeiShu(webhook string)  *FeiShu
<span id="NewFeiShu"></span>
> 根据特定的 webhook 地址创建飞书发送器
***
### FeiShu
飞书发送器
```go
type FeiShu struct {
	client  *resty.Client
	webhook string
}
```
#### func (*FeiShu) Push(notify notify.Notify)  error
> 推送通知
***
