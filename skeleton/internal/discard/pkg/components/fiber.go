package components

import (
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/fiber"
)

type FiberComponent interface {
	// RegisterFiberHandler 注册 Fiber 处理器，将在 fiber.App 启动前执行注册的处理器
	RegisterFiberHandler(handlers ...func(fiberApp *fiber.App))
}

type FiberWebSocket interface {
}
