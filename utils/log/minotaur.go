package log

import "go.uber.org/zap"

type Minotaur struct {
	*zap.Logger
	Sugared *zap.SugaredLogger
}
