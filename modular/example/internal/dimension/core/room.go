package core

import "github.com/kercylan98/minotaur/modular/example/internal/dimension/dimensions/exposes"

type Room struct {
	RoomId int64
	*Events
	exposes.VisitorsExpose
}
