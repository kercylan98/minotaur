package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	serverMultipleMark = "Minotaur Multiple Server"
	serverMark         = "Minotaur Server"
)

const (
	DefaultAsyncPoolSize           = 256
	DefaultWebsocketReadDeadline   = 30 * time.Second
	DefaultPacketWarnSize          = 1024 * 1024 * 1 // 1MB
	DefaultDispatcherBufferSize    = 1024 * 16
	DefaultConnWriteBufferSize     = 1024 * 1
	DefaultConnHubBufferSize       = 1024 * 1
	DefaultLowMessageDuration      = 100 * time.Millisecond
	DefaultAsyncLowMessageDuration = time.Second
)

func DefaultWebsocketUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
