package server

import (
	"time"
)

const (
	serverMultipleMark     = "Minotaur Multiple Server"
	serverMark             = "Minotaur Server"
	serverSystemDispatcher = "system" // 系统消息分发器
)

const (
	DefaultMessageBufferSize     = 1024
	DefaultAsyncPoolSize         = 256
	DefaultWebsocketReadDeadline = 30 * time.Second
	DefaultPacketWarnSize        = 1024 * 1024 * 1 // 1MB
)
