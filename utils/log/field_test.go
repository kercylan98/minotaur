package log_test

import (
	"github.com/kercylan98/minotaur/utils/log"
	"testing"
)

func TestStack(t *testing.T) {

	log.Debug("TestStack")
	log.Info("TestStack")
	log.Warn("TestStack")
	log.Error("TestStack")
	//log.Panic("TestStack")
	//log.DPanic("TestStack")
	//log.Fatal("TestStack")
}
