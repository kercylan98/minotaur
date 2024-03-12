package dimension

import (
	"github.com/kercylan98/minotaur/modular"
	"github.com/kercylan98/minotaur/modular/example/internal/dimension/core"
	"github.com/kercylan98/minotaur/modular/example/internal/dimension/dimensions/dimensions/visitors"
)

func New(roomId int64) error {
	visitorsDimension := new(visitors.Dimension)

	return modular.RunDimensions(&core.Room{
		RoomId: roomId,
		Events: &core.Events{},

		VisitorsExpose: visitorsDimension,
	},
		visitorsDimension,
	)
}
