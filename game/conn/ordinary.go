package conn

import (
	"minotaur/game"
)

func NewOrdinary() game.OnCreateConnHandleFunc {
	return func() game.Conn {
		return new(Ordinary)
	}
}

type Ordinary struct {
}

func (slf *Ordinary) GetConn() any {
	return slf
}
