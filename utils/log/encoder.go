package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Encoder struct {
	e     zapcore.Encoder
	cores []Core
	conf  *Config
}

func (slf *Encoder) Split(config *lumberjack.Logger, level LevelEnabler) *Encoder {
	slf.cores = append(slf.cores, zapcore.NewCore(slf.e,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(config)),
		level))
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
	options = append(options, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(append(slf.cores, core)...)
	}))
	l = l.WithOptions(options...)
	return &Minotaur{
		Logger:  l,
		Sugared: l.Sugar(),
	}
}
