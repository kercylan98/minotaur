package log

import (
	"log/slog"
	"testing"
	"time"
)

func TestStack(t *testing.T) {

	var i int
	for {
		time.Sleep(time.Second)
		Debug("TestStack")
		Info("TestStack")
		Warn("TestStack")
		Error("TestStack")
		i++
		if i == 3 {
			Default().Logger.Handler().(*handler).opts.GerRuntimeHandler().ChangeLevel(slog.LevelInfo)
		}
	}
	//log.Panic("TestStack")
	//log.DPanic("TestStack")
	//log.Fatal("TestStack")
}
