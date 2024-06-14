package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
	"time"
)

type (
	ConnPacketHandler    func(ctx vivid.MessageContext, conn Conn, packet ConnectionReactPacketMessage)
	ConnTerminateHandler func(ctx vivid.MessageContext, conn Conn, message vivid.OnTerminate)
)

type (
	// ConnectionInitMessage 连接初始化消息
	ConnectionInitMessage struct {
		Conn      net.Conn               // 连接
		Writer    ConnWriter             // 连接写入器
		ActorHook func(actor *ConnActor) // 连接 Actor 引用钩子
	}

	// ConnectionSetPacketHandlerMessage 切换响应数据包消息行为
	ConnectionSetPacketHandlerMessage struct {
		Handler ConnPacketHandler
	}

	// ConnectionSetTerminateHandlerMessage 设置连接终止处理器消息
	ConnectionSetTerminateHandlerMessage struct {
		Handler ConnTerminateHandler
	}

	// ConnectionReactPacketMessage 连接响应数据包消息
	ConnectionReactPacketMessage struct {
		Packet // 数据包
	}

	// ConnectionLoadModMessage 加载模组消息
	ConnectionLoadModMessage struct {
		Mods []vivid.ModInfo
	}

	// ConnectionUnloadModMessage 卸载模组消息
	ConnectionUnloadModMessage struct {
		Mods []vivid.ModInfo
	}

	// ConnectionApplyModMessage 应用模组消息
	ConnectionApplyModMessage struct {
	}

	// ConnectionSetZombieTimeoutMessage 设置僵尸连接超时时间消息
	ConnectionSetZombieTimeoutMessage struct {
		Timeout time.Duration
	}
)
