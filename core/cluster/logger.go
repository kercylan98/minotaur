package cluster

import (
	"github.com/kercylan98/minotaur/core/vivid"
	mlog "github.com/kercylan98/minotaur/toolkit/log"
	"log"
	"strings"
)

func newMemberlistLogger(provider vivid.LoggerProvider) *log.Logger {
	return log.New(&loggerWriter{provider: provider}, "", 0)
}

type loggerWriter struct {
	provider vivid.LoggerProvider
}

func (w *loggerWriter) Write(p []byte) (n int, err error) {
	str := string(p)
	info := strings.SplitN(str, "memberlist: ", 2)[1]
	info = strings.TrimSpace(info)
	switch {
	case strings.Contains(str, "ERR"):
		w.provider().Error("cluster", mlog.String("status", "memberlist"), mlog.String("info", info))
	case strings.Contains(str, "WARN"):
		w.provider().Warn("cluster", mlog.String("status", "memberlist"), mlog.String("info", info))
	case strings.Contains(str, "DEBUG"):
		w.provider().Debug("cluster", mlog.String("status", "memberlist"), mlog.String("info", info))
	case strings.Contains(str, "INFO"):
		w.provider().Info("cluster", mlog.String("status", "memberlist"), mlog.String("info", info))
	default:
	}
	return len(p), nil
}
