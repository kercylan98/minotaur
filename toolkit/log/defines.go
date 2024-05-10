package log

import (
	"github.com/fatih/color"
	"log/slog"
)

type (
	Level   = slog.Level
	Handler = slog.Handler
	Logger  = slog.Logger
	Field   = slog.Attr
)

type (
	Attribute = color.Attribute
	Color     struct {
		*color.Color
		attrs []Attribute
	}
)

type AttrType int

const (
	AttrTypeTime AttrType = iota + 1
	AttrTypeCaller
	AttrTypeMessage
	AttrTypeField
	AttrTypeTrace
	AttrTypeError
)

const (
	levelNone  = slog.LevelDebug - 1
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

var levels = []Level{LevelDebug, LevelInfo, LevelWarn, LevelError}

func (c *Color) Add(value ...Attribute) *Color {
	c.Color.Add(value...)
	c.attrs = append(c.attrs, value...)
	return c
}

func (c *Color) clone() *Color {
	return NewColor(c.attrs...)
}

func Levels() []Level {
	return levels
}
