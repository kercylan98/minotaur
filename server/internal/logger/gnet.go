package logger

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
)

type GNet struct {
}

func (slf *GNet) Debugf(format string, args ...interface{}) {
	log.Debug(fmt.Sprintf(format, args...))
}

func (slf *GNet) Infof(format string, args ...interface{}) {
	log.Info(fmt.Sprintf(format, args...))
}

func (slf *GNet) Warnf(format string, args ...interface{}) {
	log.Warn(fmt.Sprintf(format, args...))
}

func (slf *GNet) Errorf(format string, args ...interface{}) {
	log.Error(fmt.Sprintf(format, args...))
}

func (slf *GNet) Fatalf(format string, args ...interface{}) {
	log.Fatal(fmt.Sprintf(format, args...))
}
