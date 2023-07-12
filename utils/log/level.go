package log

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level
type LevelEnablerFunc = zap.LevelEnablerFunc

const (
	// DebugLevel 调试级别日志通常非常庞大，并且通常在生产中被禁用
	DebugLevel Level = zapcore.DebugLevel
	// InfoLevel 是默认的日志记录优先级
	InfoLevel Level = zapcore.InfoLevel
	// WarnLevel 日志比信息更重要，但不需要单独的人工审核
	WarnLevel Level = zapcore.WarnLevel
	// ErrorLevel 日志具有高优先级。如果应用程序运行顺利，它不应该生成任何错误级别的日志
	ErrorLevel Level = zapcore.ErrorLevel
	// DPanicLevel 日志是特别重要的错误。在开发中，记录器在写入消息后会出现恐慌
	DPanicLevel Level = zapcore.DPanicLevel
	// PanicLevel 记录一条消息，然后出现恐慌
	PanicLevel Level = zapcore.PanicLevel
	// FatalLevel 记录一条消息，然后调用 os.Exit(1)
	FatalLevel Level = zapcore.FatalLevel
)

var (
	levels                = []Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel, DPanicLevel, PanicLevel, FatalLevel}
	defaultLevelPartition = map[Level]func() LevelEnablerFunc{
		DebugLevel:  DebugLevelPartition,
		InfoLevel:   InfoLevelPartition,
		WarnLevel:   WarnLevelPartition,
		ErrorLevel:  ErrorLevelPartition,
		DPanicLevel: DPanicLevelPartition,
		PanicLevel:  PanicLevelPartition,
		FatalLevel:  FatalLevelPartition,
	}
)

// Levels 返回所有日志级别
func Levels() []Level {
	return levels
}

// MultiLevelPartition 返回一个 LevelEnablerFunc，该函数在指定的多个级别时返回 true
//   - 该函数被用于划分不同级别的日志输出
func MultiLevelPartition(levels ...Level) LevelEnablerFunc {
	var levelMap = hash.ToIterator(levels)
	return func(level zapcore.Level) bool {
		return hash.Exist(levelMap, level)
	}
}

// DebugLevelPartition 返回一个 LevelEnablerFunc，该函数在 DebugLevel 时返回 true
//   - 该函数被用于划分不同级别的日志输出
func DebugLevelPartition() LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == DebugLevel
	}
}

// InfoLevelPartition 返回一个 LevelEnablerFunc，该函数在 InfoLevel 时返回 true
//   - 该函数被用于划分不同级别的日志输出
func InfoLevelPartition() LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == InfoLevel
	}
}

// WarnLevelPartition 返回一个 LevelEnablerFunc，该函数在 WarnLevel 时返回 true
//   - 该函数被用于划分不同级别的日志输出
func WarnLevelPartition() LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == WarnLevel
	}
}

// ErrorLevelPartition 返回一个 LevelEnablerFunc，该函数在 ErrorLevel 时返回 true
//   - 该函数被用于划分不同级别的日志输出
func ErrorLevelPartition() LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == ErrorLevel
	}
}

// DPanicLevelPartition 返回一个 LevelEnablerFunc，该函数在 DPanicLevel 时返回 true
//   - 该函数被用于划分不同级别的日志输出
func DPanicLevelPartition() LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == DPanicLevel
	}
}

// PanicLevelPartition 返回一个 LevelEnablerFunc，该函数在 PanicLevel 时返回 true
//   - 该函数被用于划分不同级别的日志输出
func PanicLevelPartition() LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == PanicLevel
	}
}

// FatalLevelPartition 返回一个 LevelEnablerFunc，该函数在 FatalLevel 时返回 true
//   - 该函数被用于划分不同级别的日志输出
func FatalLevelPartition() LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == FatalLevel
	}
}
