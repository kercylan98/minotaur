package cluster

import (
	"errors"
	"github.com/kercylan98/minotaur/toolkit/log"
	"strings"
)

type logger struct {
	log *log.Logger
}

func (l *logger) Write(p []byte) (n int, err error) {
	// 2024/06/17 17:03:45 [INFO]
	// 2024/06/17 17:03:45 [DEBUG]
	var level string
	var message string

	str := string(p)
	levelEnd := strings.Index(str, "]")
	if levelEnd > 0 {
		level = str[21:levelEnd]
		message = str[levelEnd+2:]
	} else {
		return 0, nil
	}

	if message[len(message)-1] == '\n' {
		message = message[:len(message)-1]
	}

	switch level {
	case "INFO":
		l.log.Info("cluster", log.String("msg", message))
	case "DEBUG":
		l.log.Debug("cluster", log.String("msg", message))
	case "ERROR":
		l.log.Error("cluster", log.Err(errors.New(message)))
	case "WARN":
		l.log.Warn("cluster", log.String("msg", message))
	default:
	}

	return 0, nil
}
