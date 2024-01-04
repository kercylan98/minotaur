package log

import (
	"log/slog"
)

const (
	// DebugLevel 调试级别日志通常非常庞大，并且通常在生产中被禁用
	DebugLevel = slog.LevelDebug
	// InfoLevel 是默认的日志记录优先级
	InfoLevel = slog.LevelInfo
	// WarnLevel 日志比信息更重要，但不需要单独的人工审核
	WarnLevel = slog.LevelWarn
	// ErrorLevel 日志具有高优先级。如果应用程序运行顺利，它不应该生成任何错误级别的日志
	ErrorLevel = slog.LevelError
	// PanicLevel 记录一条消息，然后出现恐慌
	PanicLevel slog.Level = 16
	// DPanicLevel 日志是特别重要的错误。在开发中，记录器在写入消息后会出现恐慌
	DPanicLevel slog.Level = 17
	// FatalLevel 记录一条消息，然后调用 os.Exit(1)
	FatalLevel slog.Level = 20
)
