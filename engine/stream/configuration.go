package stream

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/behavior"
)

func newConfiguration() *Configuration {
	return &Configuration{}
}

// Configuration 流配置
type Configuration struct {
	performance behavior.Performance[vivid.ActorContext]
}

// WithPerformance 设置流的行为表现，你可以像控制 Actor 一样控制它
func (c *Configuration) WithPerformance(performance behavior.Performance[vivid.ActorContext]) {
	c.performance = performance
}
