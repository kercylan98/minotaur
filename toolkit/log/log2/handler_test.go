package log_test

import (
	"errors"
	log "github.com/kercylan98/minotaur/toolkit/log/log2"
	"log/slog"
	"os"
	"testing"
)

type Test struct {
	Name string
	Age  int
}

type TestList struct {
	Tag   string
	Tests []Test
}

func TestNewHandler(t *testing.T) {
	slog.Default().With(slog.String("User", "Jerry")).Info("Test")

	var h = log.NewHandler(os.Stdout, log.NewProdHandlerOptions())
	logger := slog.New(h)
	logger = logger.With(slog.String("fixedAttr", "test"))
	logger = logger.WithGroup("group")

	var fields = make([]any, 0)
	fields = append(fields, slog.String("stringKey", "stringValue"))
	fields = append(fields, slog.Int("intKey", 123))
	fields = append(fields, slog.Float64("floatKey", 123.456))
	fields = append(fields, slog.Bool("boolKey", true))
	fields = append(fields, slog.Any("anyKey", Test{Name: "Jerry", Age: 18}))
	fields = append(fields, slog.Any("anyListKey", TestList{
		Tag: "test",
		Tests: []Test{
			{Name: "Jerry", Age: 18},
			{Name: "Tom", Age: 20},
		},
	},
	))
	fields = append(fields, slog.Any("nilKey", nil))
	fields = append(fields, slog.Any("errKey", errors.New("test error")))

	for _, field := range fields {
		logger.Debug("log", field)
		logger.Info("log", field)
		logger.Warn("log", field)
		logger.Error("log", field)
	}

	logger.Error("testError", log.Err(errors.New("test error 1")), log.Err(errors.New("test error 2")), log.Err(errors.New("test error 3")))
}
