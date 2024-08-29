package stream

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/behavior"
)

func newConfiguration() *Configuration {
	return &Configuration{}
}

// Configuration 流配置
//
// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type Configuration struct {
	performance behavior.Performance[vivid.ActorContext]
}

// WithPerformance 设置流的行为表现，你可以像控制 Actor 一样控制它
func (c *Configuration) WithPerformance(performance behavior.Performance[vivid.ActorContext]) {
	c.performance = performance
}
