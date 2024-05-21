package vivid

import "github.com/kercylan98/minotaur/vivid/vivids"

// RegisterBehavior 注册行为
func RegisterBehavior[T vivids.Message](ctx vivids.ActorContext, behavior vivids.ActorBehavior[T]) {
	ctx.RegisterBehavior((*T)(nil), behavior)
}
