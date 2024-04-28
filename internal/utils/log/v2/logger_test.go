package log_test

import (
	"errors"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"testing"
)

func TestLogger(t *testing.T) {

	var msg = "TestLogger"
	var fields = make([]any, 0)

	fields = append(fields, log.String("Name", "Jerry"), log.Any("errhhha", errors.New("test error")), log.Err(errors.New("test error")))
	for _, level := range log.Levels() {
		log.Log(level, msg, fields...)
	}

}
