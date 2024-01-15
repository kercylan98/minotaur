# Senders

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

senders Package 包含了内置通知发送器的实现


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewFeiShu](#NewFeiShu)|根据特定的 webhook 地址创建飞书发送器


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[FeiShu](#feishu)|飞书发送器

</details>


***
## 详情信息
#### func NewFeiShu(webhook string) *FeiShu
<span id="NewFeiShu"></span>
> 根据特定的 webhook 地址创建飞书发送器

***
### FeiShu `STRUCT`
飞书发送器
```go
type FeiShu struct {
	client  *resty.Client
	webhook string
}
```
#### func (*FeiShu) Push(notify notify.Notify)  error
> 推送通知
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFeiShu_Push(t *testing.T) {
	fs := NewFeiShu("https://open.feishu.cn/open-apis/bot/v2/hook/d886f30f-814c-47b1-aeb0-b508da0f7f22")
	rt := notifies.NewFeiShu(notifies.FeiShuMessageWithRichText(notifies.NewFeiShuRichText().Create("zh_cn", "标题咯").AddText("哈哈哈").Ok()))
	if err := fs.Push(rt); err != nil {
		panic(err)
	}
}

```


</details>


***
