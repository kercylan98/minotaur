package notifies

// NewFeiShuRichText 创建一个飞书富文本
func NewFeiShuRichText() *FeiShuRichText {
	return &FeiShuRichText{
		content: map[string]*FeiShuRichTextContent{},
	}
}

// FeiShuRichText 飞书富文本结构
type FeiShuRichText struct {
	content map[string]*FeiShuRichTextContent
}

// Create 创建一个特定语言和标题的富文本内容
func (slf *FeiShuRichText) Create(lang, title string) *FeiShuRichTextContent {
	content := &FeiShuRichTextContent{
		richText: slf,
		Title:    title,
		Content:  make([][]map[string]any, 1),
	}
	slf.content[lang] = content
	return content
}

// FeiShuRichTextContent 飞书富文本内容体
type FeiShuRichTextContent struct {
	richText *FeiShuRichText
	Title    string             `json:"title,omitempty"`
	Content  [][]map[string]any `json:"content,omitempty"`
}

// AddText 添加文本
func (slf *FeiShuRichTextContent) AddText(text string, styles ...string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":   "text",
		"text":  text,
		"style": styles,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddUnescapeText 添加 unescape 解码的文本
func (slf *FeiShuRichTextContent) AddUnescapeText(text string, styles ...string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":       "text",
		"text":      text,
		"un_escape": true,
		"style":     styles,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddLink 添加超链接文本
//   - 请确保链接地址的合法性，否则消息会发送失败
func (slf *FeiShuRichTextContent) AddLink(text, href string, styles ...string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":   "a",
		"text":  text,
		"href":  href,
		"style": styles,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddAt 添加@的用户
//   - @单个用户时，userId 字段必须是有效值
//   - @所有人填"all"。
func (slf *FeiShuRichTextContent) AddAt(userId string, styles ...string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":     "at",
		"user_id": userId,
		"style":   styles,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddAtWithUsername 添加包含用户名的@用户
//   - @单个用户时，userId 字段必须是有效值
//   - @所有人填"all"。
func (slf *FeiShuRichTextContent) AddAtWithUsername(userId, username string, styles ...string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":       "at",
		"user_id":   userId,
		"user_name": username,
		"style":     styles,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddImg 添加图片
//   - imageKey 表示图片的唯一标识，可通过上传图片接口获取
func (slf *FeiShuRichTextContent) AddImg(imageKey string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":       "img",
		"image_key": imageKey,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddMedia 添加视频
//   - fileKey 表示视频文件的唯一标识，可通过上传文件接口获取
func (slf *FeiShuRichTextContent) AddMedia(fileKey string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":      "media",
		"file_key": fileKey,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddMediaWithCover 添加包含封面的视频
//   - fileKey 表示视频文件的唯一标识，可通过上传文件接口获取
//   - imageKey 表示图片的唯一标识，可通过上传图片接口获取
func (slf *FeiShuRichTextContent) AddMediaWithCover(fileKey, imageKey string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":       "media",
		"file_key":  fileKey,
		"image_key": imageKey,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// AddEmotion 添加表情
//   - emojiType 表示表情类型，部分可选值请参见表情文案。
//
// 表情文案：https://open.feishu.cn/document/server-docs/im-v1/message-reaction/emojis-introduce
func (slf *FeiShuRichTextContent) AddEmotion(emojiType string) *FeiShuRichTextContent {
	content := map[string]any{
		"tag":        "emotion",
		"emoji_type": emojiType,
	}
	slf.Content[0] = append(slf.Content[0], content)
	return slf
}

// Ok 确认完成，将返回 FeiShuRichText 可继续创建多语言富文本
func (slf *FeiShuRichTextContent) Ok() *FeiShuRichText {
	return slf.richText
}
