package module

import (
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
)

type FiberModule interface {
	application.Module

	Fiber() *fiber.Server[*application.Context]
}
