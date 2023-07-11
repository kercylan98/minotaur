package server

import "time"

const (
	DefaultMessageBufferSize     = 1024
	DefaultMessageChannelSize    = 1024 * 4096
	DefaultAsyncPoolSize         = 256
	DefaultWebsocketReadDeadline = 30 * time.Second
)
