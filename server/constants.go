package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"time"
)

type (
	RunMode = log.RunMode
)

const (
	RunModeDev  RunMode = log.RunModeDev
	RunModeProd RunMode = log.RunModeProd
	RunModeTest RunMode = log.RunModeTest
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

const (
	contextKeyWST = "_wst" // WebSocket 消息类型
)
