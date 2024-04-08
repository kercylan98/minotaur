package message

import (
	"context"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
)

type Message[P Producer, Q Queue] interface {
	OnInitialize(ctx context.Context)
	OnProcess()

	// GetProducer 获取消息生产者
	GetProducer() P

	queue.Message[Q]
}
