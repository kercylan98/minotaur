package log

import (
	"fmt"
	"os"
)

func NewGNetLogger(logger *Logger) *GNetLogger {
	return &GNetLogger{Logger: logger}
}

type GNetLogger struct {
	Logger *Logger
}

func (l *GNetLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug("GNET", String("msg", fmt.Sprintf(format, args...)))
}

func (l *GNetLogger) Infof(format string, args ...interface{}) {
	l.Logger.Info("GNET", String("msg", fmt.Sprintf(format, args...)))
}

func (l *GNetLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warn("GNET", String("msg", fmt.Sprintf(format, args...)))
}

func (l *GNetLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Error("GNET", String("msg", fmt.Sprintf(format, args...)))
}

func (l *GNetLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Error("GNET", String("msg", fmt.Sprintf(format, args...)))
	os.Exit(1)
}
