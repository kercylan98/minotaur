package log

var logger Logger

func init() {
	logger = Default().Build()
}

// SetLogger 设置日志记录器
func SetLogger(l Logger) {
	if m, ok := l.(*Minotaur); ok && m != nil {
		_ = m.Sync()
		_ = m.Sugared.Sync()
	}
	logger = l
}

// Logger 适用于 Minotaur 的日志接口
type Logger interface {
	// Debug 在 DebugLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
	Debug(msg string, fields ...Field)
	// Info 在 InfoLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
	Info(msg string, fields ...Field)
	// Warn 在 WarnLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
	Warn(msg string, fields ...Field)
	// Error 在 ErrorLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
	Error(msg string, fields ...Field)
	// DPanic 在 DPanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
	//   - 如果记录器处于开发模式，它就会出现 panic（DPanic 的意思是“development panic”）。这对于捕获可恢复但不应该发生的错误很有用
	DPanic(msg string, fields ...Field)
	// Panic 在 PanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
	//   - 即使禁用了 PanicLevel 的日志记录，记录器也会出现 panic
	Panic(msg string, fields ...Field)
	// Fatal 在 FatalLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
	//   - 然后记录器调用 os.Exit(1)，即使 FatalLevel 的日志记录被禁用
	Fatal(msg string, fields ...Field)
}

// Debug 在 DebugLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Debug(msg string, fields ...Field) {
	logger.Debug(msg, fields...)
}

// Info 在 InfoLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}

// Warn 在 WarnLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Warn(msg string, fields ...Field) {
	logger.Warn(msg, fields...)
}

// Error 在 ErrorLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Error(msg string, fields ...Field) {
	logger.Error(msg, fields...)
}

// DPanic 在 DPanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 如果记录器处于开发模式，它就会出现 panic（DPanic 的意思是“development panic”）。这对于捕获可恢复但不应该发生的错误很有用
func DPanic(msg string, fields ...Field) {
	logger.DPanic(msg, fields...)
}

// Panic 在 PanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 即使禁用了 PanicLevel 的日志记录，记录器也会出现 panic
func Panic(msg string, fields ...Field) {
	logger.Panic(msg, fields...)
}

// Fatal 在 FatalLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 然后记录器调用 os.Exit(1)，即使 FatalLevel 的日志记录被禁用
func Fatal(msg string, fields ...Field) {
	logger.Fatal(msg, fields...)
}
