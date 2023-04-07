package game

// PushFail 消息推送失败信息
type PushFail struct {
	Player *Player
	Data   []byte
	Err    error
}
