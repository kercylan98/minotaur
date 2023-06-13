package notifies

import "encoding/json"

// NewFeiShu 创建飞书通知消息
func NewFeiShu(message FeiShuMessage) *FeiShu {
	feishu := &FeiShu{
		Content: map[string]any{},
	}
	message(feishu)
	return feishu
}

// FeiShu 飞书通知消息
type FeiShu struct {
	Content any    `json:"content"`
	MsgType string `json:"msg_type"`
}

// Format 格式化通知内容
func (slf *FeiShu) Format() (string, error) {
	data, err := json.Marshal(slf)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
