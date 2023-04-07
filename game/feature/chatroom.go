package feature

// ChatRoom 聊天室接口定义
type ChatRoom interface {
	// JoinChatRoom 加入聊天室
	JoinChatRoom(chat PlayerChat)
	// LeaveChatRoom 离开聊天室
	LeaveChatRoom(chat PlayerChat)
}

// PlayerChat 玩家聊天接口定义
type PlayerChat interface {
	Player
	// SendChatMessage 发送聊天消息给该玩家
	SendChatMessage(message string)
}
