package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Encoder struct {
	e     zapcore.Encoder
	cores []Core
	conf  *Config
}

func (slf *Encoder) Split(config *lumberjack.Logger) *Encoder {
	slf.cores = append(slf.cores, zapcore.NewCore(slf.e, zapcore.AddSync(config), zapcore.DebugLevel))
	return slf
}

func (slf *Encoder) AddCore(ws WriteSyncer, enab LevelEnabler) *Encoder {
	slf.cores = append(slf.cores, zapcore.NewCore(slf.e, ws, enab))
	return slf
}

func (slf *Encoder) Build(options ...LoggerOption) *Minotaur {
	l, err := slf.conf.Build()
	if err != nil {
		panic(err)
	}
	options = append([]LoggerOption{zap.AddCaller(), zap.AddCallerSkip(1)}, options...)
	l = l.WithOptions(options...)
	if len(slf.cores) == 0 {
		// stdout、stderr，不使用 lumberjack.Logger
		slf.cores = append(slf.cores, zapcore.NewCore(
			slf.e,
			zapcore.Lock(os.Stdout),
			zapcore.InfoLevel,
		))
		slf.cores = append(slf.cores, zapcore.NewCore(
			slf.e,
			zapcore.Lock(os.Stderr),
			zapcore.ErrorLevel,
		))
	}
	l = l.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(slf.cores...)
	}))
	return &Minotaur{
		Logger:  l,
		Sugared: l.Sugar(),
	}
}
