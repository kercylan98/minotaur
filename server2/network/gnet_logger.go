package network

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/toolkit/log"
)

type gNetLogger struct {
	server.Controller
}

func (l *gNetLogger) Debugf(format string, args ...interface{}) {
	l.GetServerLogger().Debug("gnet", log.String("message", fmt.Sprintf(format, args...)))
}

func (l *gNetLogger) Infof(format string, args ...interface{}) {
	l.GetServerLogger().Info("gnet", log.String("message", fmt.Sprintf(format, args...)))
}

func (l *gNetLogger) Warnf(format string, args ...interface{}) {
	l.GetServerLogger().Warn("gnet", log.String("message", fmt.Sprintf(format, args...)))
}

func (l *gNetLogger) Errorf(format string, args ...interface{}) {
	l.GetServerLogger().Error("gnet", log.String("message", fmt.Sprintf(format, args...)))
}

func (l *gNetLogger) Fatalf(format string, args ...interface{}) {
	l.GetServerLogger().Error("gnet fatal:", log.String("message", fmt.Sprintf(format, args...)))
}
