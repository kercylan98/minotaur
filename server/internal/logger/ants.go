package logger

import "github.com/kercylan98/minotaur/utils/log"

type Ants struct {
}

func (slf *Ants) Printf(format string, args ...interface{}) {
	log.Warn(format, log.Any("args", args))
}
