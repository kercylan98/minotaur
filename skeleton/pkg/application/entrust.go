package application

import "github.com/kercylan98/minotaur/skeleton/pkg/fiber"

type FiberEntrust interface {
	OnInitialize(ctx *Context, fiber *fiber.Server[*Context]) (err error)
}
