package components

import (
	"github.com/kercylan98/minotaur/vivid"
)

// BehaviorAutoExecutor 行为自动执行器，用于自动执行匹配的 Actor 行为
//   - 该组件仅需与 Actor 匿名组合，并在 Actor.OnPreStart 中调用 Init 方法即可
type BehaviorAutoExecutor struct {
	ctx vivid.ActorContext
}

// Init 初始化行为自动执行器
func (b *BehaviorAutoExecutor) Init(ctx vivid.ActorContext) *BehaviorAutoExecutor {
	b.ctx = ctx

	return b
}

func (b *BehaviorAutoExecutor) OnReceived(msg vivid.MessageContext) error {
	executor := vivid.MatchBehavior(b.ctx, msg)
	if executor == nil {
		return nil
	}

	return executor.Execute()
}
