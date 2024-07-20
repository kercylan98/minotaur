package supervision

// Logger 监管记录器是一个处理函数，它将在发生意外时被调用，以记录用户定义的日志信息
type Logger interface {
	// Log 当发生事故时
	Log(record *AccidentRecord)
}

// FunctionalLogger 是一个函数类型的监管记录器
type FunctionalLogger func(record *AccidentRecord)

// Log 当发生事故时
func (f FunctionalLogger) Log(record *AccidentRecord) {
	f(record)
}
