package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/g2d"
	"sync"
	"time"
)

func NewMoving2D() *Moving2D {
	moving2D := &Moving2D{
		entities: map[int64]*moving2DTarget{},
	}
	go moving2D.handle()
	return moving2D
}

type moving2DTarget struct {
	game.Moving2DEntity
	x, y         float64
	lastMoveTime int64
}

func (slf *Moving2D) MoveTo(entity game.Moving2DEntity, x float64, y float64) {
	slf.rw.Lock()
	slf.entities[entity.GetGuid()] = &moving2DTarget{
		Moving2DEntity: entity,
		x:              x,
		y:              y,
		lastMoveTime:   time.Now().UnixMilli(),
	}
	slf.rw.Unlock()
}

type Moving2D struct {
	rw       sync.RWMutex
	entities map[int64]*moving2DTarget
}

func (slf *Moving2D) handle() {
	for {
		slf.rw.Lock()
		for guid, entity := range slf.entities {
			x, y := entity.GetPosition()
			angle := g2d.CalcAngle(x, y, entity.x, entity.y)
			moveTime := time.Now().UnixMilli()
			interval := float64(moveTime - entity.lastMoveTime)
			distance := interval * entity.GetSpeed()
			nx, ny := g2d.CalculateNewCoordinate(x, y, angle, distance)
			if g2d.CalcDistance(nx, ny, entity.x, entity.y) <= distance {
				entity.SetPosition(entity.x, entity.y)
				delete(slf.entities, guid)
				return
			}
			entity.SetPosition(nx, ny)
			entity.lastMoveTime = moveTime
		}

		if len(slf.entities) == 0 {
			slf.rw.Unlock()
			time.Sleep(100 * time.Millisecond)
		} else {
			slf.rw.Unlock()
		}
	}
}
