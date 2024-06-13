package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type (
	// ConnectionInitMessage 连接初始化消息
	ConnectionInitMessage struct {
		Conn   net.Conn   // 连接
		Writer ConnWriter // 连接写入器
	}

	// ConnectionBecomeReactPacketMessage 切换响应数据包消息行为
	ConnectionBecomeReactPacketMessage struct {
		Behavior vivid.Behavior // 行为
	}

	// ConnectionReactPacketMessage 连接响应数据包消息
	ConnectionReactPacketMessage struct {
		Packet Packet // 数据包
	}

	// ConnectionWritePacketMessage 连接写入数据包消息
	ConnectionWritePacketMessage struct {
		Packet Packet
	}
)
