package notifies

const (
	FeiShuMsgTypeText        = "text"        // 文本
	FeiShuMsgTypeRichText    = "post"        // 富文本
	FeiShuMsgTypeImage       = "image"       // 图片
	FeiShuMsgTypeInteractive = "interactive" // 消息卡片
	FeiShuMsgTypeShareChat   = "share_chat"  // 分享群名片
	FeiShuMsgTypeShareUser   = "share_user"  // 分享个人名片
	FeiShuMsgTypeAudio       = "audio"       // 音频
	FeiShuMsgTypeMedia       = "media"       // 视频
	FeiShuMsgTypeFile        = "file"        // 文件
	FeiShuMsgTypeSticker     = "sticker"     // 表情包
)

const (
	FeiShuStyleBold        = "bold"        // 加粗
	FeiShuStyleUnderline   = "underline"   // 下划线
	FeiShuStyleLineThrough = "lineThrough" // 删除线
	FeiShuStyleItalic      = "italic"      // 斜体
)

type FeiShuMessage func(feishu *FeiShu)

// FeiShuMessageWithText 飞书文本消息
//   - 支持通过换行符进行消息换行
//   - 支持通过 <at user_id="OpenID">名字</at> 进行@用户
//   - 支持通过 <at user_id="all">所有人</at> 进行@所有人（必须满足所在群开启@所有人功能。）
//
// 支持加粗、斜体、下划线、删除线四种样式，可嵌套使用：
//   - 加粗: <b>文本示例</b>
//   - 斜体: <i>文本示例</i>
//   - 下划线 : <u>文本示例</u>
//   - 删除线: <s>文本示例</s>
//
// 超链接使用说明
//   - 超链接的使用格式为[文本](链接)， 如[Feishu Open Platform](https://open.feishu.cn) 。
//   - 请确保链接是合法的，否则会以原始内容发送消息。
func FeiShuMessageWithText(text string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			Text string `json:"text"`
		}{text}
		feishu.MsgType = FeiShuMsgTypeText
	}
}

// FeiShuMessageWithRichText 飞书富文本消息
func FeiShuMessageWithRichText(richText *FeiShuRichText) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			Post any `json:"post,omitempty"`
		}{richText.content}
		feishu.MsgType = FeiShuMsgTypeRichText
	}
}

// FeiShuMessageWithImage 飞书图片消息
//   - imageKey 可通过上传图片接口获取
func FeiShuMessageWithImage(imageKey string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			ImageKey string `json:"imageKey"`
		}{imageKey}
		feishu.MsgType = FeiShuMsgTypeImage
	}
}

// FeiShuMessageWithInteractive 飞书卡片消息
//   - json 表示卡片的 json 数据或者消息模板的 json 数据
func FeiShuMessageWithInteractive(json string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = json
		feishu.MsgType = FeiShuMsgTypeInteractive
	}
}

// FeiShuMessageWithShareChat 飞书分享群名片
//   - chatId 群ID获取方式请参见群ID说明
//
// 群ID说明：https://open.feishu.cn/document/server-docs/group/chat/chat-id-description
func FeiShuMessageWithShareChat(chatId string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			ChatID string `json:"chat_id"`
		}{chatId}
		feishu.MsgType = FeiShuMsgTypeShareChat
	}
}

// FeiShuMessageWithShareUser 飞书分享个人名片
//   - userId 表示用户的 OpenID 获取方式请参见了解更多：如何获取 Open ID
//
// 如何获取 Open ID：https://open.feishu.cn/document/faq/trouble-shooting/how-to-obtain-openid
func FeiShuMessageWithShareUser(userId string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			UserID string `json:"user_id"`
		}{userId}
		feishu.MsgType = FeiShuMsgTypeShareUser
	}
}

// FeiShuMessageWithAudio 飞书语音消息
//   - fileKey 语音文件Key，可通过上传文件接口获取
//
// 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
func FeiShuMessageWithAudio(fileKey string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			FileKey string `json:"file_key"`
		}{fileKey}
		feishu.MsgType = FeiShuMsgTypeAudio
	}
}

// FeiShuMessageWithMedia 飞书视频消息
//   - fileKey 视频文件Key，可通过上传文件接口获取
//
// 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
func FeiShuMessageWithMedia(fileKey string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			FileKey string `json:"file_key"`
		}{fileKey}
		feishu.MsgType = FeiShuMsgTypeMedia
	}
}

// FeiShuMessageWithMediaAndCover 飞书带封面的视频消息
//   - fileKey 视频文件Key，可通过上传文件接口获取
//   - imageKey 图片文件Key，可通过上传文件接口获取
//
// 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
func FeiShuMessageWithMediaAndCover(fileKey, imageKey string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			FileKey  string `json:"file_key"`
			ImageKey string `json:"image_key"`
		}{fileKey, imageKey}
		feishu.MsgType = FeiShuMsgTypeMedia
	}
}

// FeiShuMessageWithFile 飞书文件消息
//   - fileKey 文件Key，可通过上传文件接口获取
//
// 上传文件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
func FeiShuMessageWithFile(fileKey string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			FileKey string `json:"file_key"`
		}{fileKey}
		feishu.MsgType = FeiShuMsgTypeFile
	}
}

// FeiShuMessageWithSticker 飞书表情包消息
//   - fileKey 表情包文件Key，目前仅支持发送机器人收到的表情包，可通过接收消息事件的推送获取表情包 file_key。
//
// 接收消息事件：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
func FeiShuMessageWithSticker(fileKey string) FeiShuMessage {
	return func(feishu *FeiShu) {
		feishu.Content = struct {
			FileKey string `json:"file_key"`
		}{fileKey}
		feishu.MsgType = FeiShuMsgTypeSticker
	}
}
