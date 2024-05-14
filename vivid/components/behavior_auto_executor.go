package components

import (
	"github.com/kercylan98/minotaur/vivid/actor2"
)

// BehaviorAutoExecutor 行为自动执行器，用于自动执行匹配的 Actor 行为
//   - 该组件仅需与 Actor 匿名组合，并在 Actor.OnPreStart 中调用 Init 方法即可
type BehaviorAutoExecutor struct {
	system *actor2.ActorSystem
	ctx    *actor2.ActorContext
}

// Init 初始化行为自动执行器
func (b *BehaviorAutoExecutor) Init(ctx *actor2.ActorContext) *BehaviorAutoExecutor {
	b.system = ctx.GetSystem()
	b.ctx = ctx

	return b
}

func (b *BehaviorAutoExecutor) OnReceived(msg actor2.Message) error {
	executor := actor2.MatchBehavior(b.ctx, msg)
	if executor == nil {
		return nil
	}

	return executor.Execute()
}
