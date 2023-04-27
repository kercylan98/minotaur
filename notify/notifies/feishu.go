package notifies

import "encoding/json"

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

func NewFeiShu(msgType, receiveId, content string) *FeiShu {
	return &FeiShu{
		ReceiveId: receiveId,
		Content:   content,
		MsgType:   msgType,
	}
}

type FeiShu struct {
	ReceiveId string `json:"receive_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
}

func (slf *FeiShu) GetTitle() string {
	return ""
}

func (slf *FeiShu) GetContent() (string, error) {
	data, err := json.Marshal(slf)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
