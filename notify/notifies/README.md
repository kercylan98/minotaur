# Notifies

notifies 包含了内置通知内容的实现

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/notifies)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewFeiShu](#NewFeiShu)|创建飞书通知消息
|[FeiShuMessageWithText](#FeiShuMessageWithText)|飞书文本消息
|[FeiShuMessageWithRichText](#FeiShuMessageWithRichText)|飞书富文本消息
|[FeiShuMessageWithImage](#FeiShuMessageWithImage)|飞书图片消息
|[FeiShuMessageWithInteractive](#FeiShuMessageWithInteractive)|飞书卡片消息
|[FeiShuMessageWithShareChat](#FeiShuMessageWithShareChat)|飞书分享群名片
|[FeiShuMessageWithShareUser](#FeiShuMessageWithShareUser)|飞书分享个人名片
|[FeiShuMessageWithAudio](#FeiShuMessageWithAudio)|飞书语音消息
|[FeiShuMessageWithMedia](#FeiShuMessageWithMedia)|飞书视频消息
|[FeiShuMessageWithMediaAndCover](#FeiShuMessageWithMediaAndCover)|飞书带封面的视频消息
|[FeiShuMessageWithFile](#FeiShuMessageWithFile)|飞书文件消息
|[FeiShuMessageWithSticker](#FeiShuMessageWithSticker)|飞书表情包消息
|[NewFeiShuRichText](#NewFeiShuRichText)|创建一个飞书富文本


> 结构体定义

|结构体|描述
|:--|:--
|[FeiShu](#feishu)|飞书通知消息
|[FeiShuMessage](#feishumessage)|暂无描述...
|[FeiShuRichText](#feishurichtext)|飞书富文本结构
|[FeiShuRichTextContent](#feishurichtextcontent)|飞书富文本内容体

</details>


#### func NewFeiShu(message FeiShuMessage)  *FeiShu
<span id="NewFeiShu"></span>
> 创建飞书通知消息
***
#### func FeiShuMessageWithText(text string)  FeiShuMessage
<span id="FeiShuMessageWithText"></span>
> 飞书文本消息
>   - 支持通过换行符进行消息换行
>   - 支持通过 <at user_id="OpenID">名字</at> 进行@用户
>   - 支持通过 <at user_id="all">所有人</at> 进行@所有人（必须满足所在群开启@所有人功能。）
> 
> 支持加粗、斜体、下划线、删除线四种样式，可嵌套使用：
>   - 加粗: <b>文本示例</b>
>   - 斜体: <i>文本示例</i>
>   - 下划线 : <u>文本示例</u>
>   - 删除线: <s>文本示例</s>
> 
> 超链接使用说明
>   - 超链接的使用格式为[文本](链接)， 如[Feishu Open Platform](https://open.feishu.cn) 。
>   - 请确保链接是合法的，否则会以原始内容发送消息。
***
#### func FeiShuMessageWithRichText(richText *FeiShuRichText)  FeiShuMessage
<span id="FeiShuMessageWithRichText"></span>
> 飞书富文本消息
***
#### func FeiShuMessageWithImage(imageKey string)  FeiShuMessage
<span id="FeiShuMessageWithImage"></span>
> 飞书图片消息
>   - imageKey 可通过上传图片接口获取
***
#### func FeiShuMessageWithInteractive(json string)  FeiShuMessage
<span id="FeiShuMessageWithInteractive"></span>
> 飞书卡片消息
>   - json 表示卡片的 json 数据或者消息模板的 json 数据
***
#### func FeiShuMessageWithShareChat(chatId string)  FeiShuMessage
<span id="FeiShuMessageWithShareChat"></span>
> 飞书分享群名片
>   - chatId 群ID获取方式请参见群ID说明
> 
> 群ID说明：https://open.feishu.cn/document/server-docs/group/chat/chat-id-description
***
#### func FeiShuMessageWithShareUser(userId string)  FeiShuMessage
<span id="FeiShuMessageWithShareUser"></span>
> 飞书分享个人名片
>   - userId 表示用户的 OpenID 获取方式请参见了解更多：如何获取 Open ID
> 
> 如何获取 Open ID：https://open.feishu.cn/document/faq/trouble-shooting/how-to-obtain-openid
***
#### func FeiShuMessageWithAudio(fileKey string)  FeiShuMessage
<span id="FeiShuMessageWithAudio"></span>
> 飞书语音消息
>   - fileKey 语音文件Key，可通过上传文件接口获取
> 
> 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
***
#### func FeiShuMessageWithMedia(fileKey string)  FeiShuMessage
<span id="FeiShuMessageWithMedia"></span>
> 飞书视频消息
>   - fileKey 视频文件Key，可通过上传文件接口获取
> 
> 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
***
#### func FeiShuMessageWithMediaAndCover(fileKey string, imageKey string)  FeiShuMessage
<span id="FeiShuMessageWithMediaAndCover"></span>
> 飞书带封面的视频消息
>   - fileKey 视频文件Key，可通过上传文件接口获取
>   - imageKey 图片文件Key，可通过上传文件接口获取
> 
> 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
***
#### func FeiShuMessageWithFile(fileKey string)  FeiShuMessage
<span id="FeiShuMessageWithFile"></span>
> 飞书文件消息
>   - fileKey 文件Key，可通过上传文件接口获取
> 
> 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
***
#### func FeiShuMessageWithSticker(fileKey string)  FeiShuMessage
<span id="FeiShuMessageWithSticker"></span>
> 飞书表情包消息
>   - fileKey 表情包文件Key，目前仅支持发送机器人收到的表情包，可通过接收消息事件的推送获取表情包 file_key。
> 
> 接收消息事件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
***
#### func NewFeiShuRichText()  *FeiShuRichText
<span id="NewFeiShuRichText"></span>
> 创建一个飞书富文本
***
### FeiShu
飞书通知消息
```go
type FeiShu struct {
	Content any
	MsgType string
}
```
#### func (*FeiShu) Format()  string,  error
> 格式化通知内容
***
### FeiShuMessage

```go
type FeiShuMessage struct{}
```
### FeiShuRichText
飞书富文本结构
```go
type FeiShuRichText struct {
	content map[string]*FeiShuRichTextContent
}
```
#### func (*FeiShuRichText) Create(lang string, title string)  *FeiShuRichTextContent
> 创建一个特定语言和标题的富文本内容
***
### FeiShuRichTextContent
飞书富文本内容体
```go
type FeiShuRichTextContent struct {
	richText *FeiShuRichText
	Title    string
	Content  [][]map[string]any
}
```
#### func (*FeiShuRichTextContent) AddText(text string, styles ...string)  *FeiShuRichTextContent
> 添加文本
***
#### func (*FeiShuRichTextContent) AddUnescapeText(text string, styles ...string)  *FeiShuRichTextContent
> 添加 unescape 解码的文本
***
#### func (*FeiShuRichTextContent) AddLink(text string, href string, styles ...string)  *FeiShuRichTextContent
> 添加超链接文本
>   - 请确保链接地址的合法性，否则消息会发送失败
***
#### func (*FeiShuRichTextContent) AddAt(userId string, styles ...string)  *FeiShuRichTextContent
> 添加@的用户
>   - @单个用户时，userId 字段必须是有效值
>   - @所有人填"all"。
***
#### func (*FeiShuRichTextContent) AddAtWithUsername(userId string, username string, styles ...string)  *FeiShuRichTextContent
> 添加包含用户名的@用户
>   - @单个用户时，userId 字段必须是有效值
>   - @所有人填"all"。
***
#### func (*FeiShuRichTextContent) AddImg(imageKey string)  *FeiShuRichTextContent
> 添加图片
>   - imageKey 表示图片的唯一标识，可通过上传图片接口获取
***
#### func (*FeiShuRichTextContent) AddMedia(fileKey string)  *FeiShuRichTextContent
> 添加视频
>   - fileKey 表示视频文件的唯一标识，可通过上传文件接口获取
***
#### func (*FeiShuRichTextContent) AddMediaWithCover(fileKey string, imageKey string)  *FeiShuRichTextContent
> 添加包含封面的视频
>   - fileKey 表示视频文件的唯一标识，可通过上传文件接口获取
>   - imageKey 表示图片的唯一标识，可通过上传图片接口获取
***
#### func (*FeiShuRichTextContent) AddEmotion(emojiType string)  *FeiShuRichTextContent
> 添加表情
>   - emojiType 表示表情类型，部分可选值请参见表情文案。
> 
> 表情文案：https://open.feishu.cn/document/server-docs/im-v1/message-reaction/emojis-introduce
***
#### func (*FeiShuRichTextContent) Ok()  *FeiShuRichText
> 确认完成，将返回 FeiShuRichText 可继续创建多语言富文本
***
