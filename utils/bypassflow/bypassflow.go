package bypassflow

import (
	"context"
	"go.uber.org/zap"
	"minotaur/utils/hash"
	"minotaur/utils/log"
	"runtime/debug"
)

// BypassFlow 分流器
type BypassFlow struct {
	consistency *hash.Consistency
	processor   []chan func()
}

func New(ctx context.Context, nodeCount, nodeBuffer int) *BypassFlow {
	bypassFlow := &BypassFlow{
		consistency: &hash.Consistency{},
		processor:   make([]chan func(), nodeCount),
	}

	for i := 0; i < nodeCount; i++ {
		bypassFlow.consistency.AddNode(i)

		nodeChan := make(chan func(), nodeBuffer)
		bypassFlow.processor[i] = nodeChan
		go func() {
			for {
				select {
				case f := <-nodeChan:
					go func() {
						f()
						defer func() {
							if err := recover(); err != nil {
								log.Error("BypassFlow", zap.Any("error", err), zap.Any("stack\n", debug.Stack()))
							}
						}()
					}()
				case <-ctx.Done():
					close(nodeChan)
					return
				}
			}
		}()
	}
	return bypassFlow
}

func (slf *BypassFlow) getNode(item Item) int {
	return slf.consistency.PickNode(item)
}

func (slf *BypassFlow) Handle(item Item, handleFunc func()) {
	node := slf.getNode(item)
	slf.processor[node] <- handleFunc
}
