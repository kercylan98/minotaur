package vivid

import "github.com/kercylan98/minotaur/toolkit/log"

// LoggerProvider Logger 提供者
type LoggerProvider interface {
	// Provide 提供 Logger
	Provide() *log.Logger
}

// FunctionalLoggerProvider Logger 提供者
type FunctionalLoggerProvider func() *log.Logger

// Provide 提供 Logger
func (f FunctionalLoggerProvider) Provide() *log.Logger {
	return f()
}
