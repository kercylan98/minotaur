package log

// LoggerProvider Logger 提供者
type LoggerProvider interface {
	// Provide 提供 Logger
	Provide() *Logger
}

// FunctionalLoggerProvider Logger 提供者
type FunctionalLoggerProvider func() *Logger

// Provide 提供 Logger
func (f FunctionalLoggerProvider) Provide() *Logger {
	return f()
}
