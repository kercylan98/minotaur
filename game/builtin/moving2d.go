package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/g2d"
	"sync"
	"time"
)

func NewMoving2D(options ...Moving2DOption) *Moving2D {
	moving2D := &Moving2D{
		entities: map[int64]*moving2DTarget{},
		timeUnit: float64(time.Millisecond),
		idle:     time.Millisecond * 100,
		interval: time.Millisecond * 100,
	}
	for _, option := range options {
		option(moving2D)
	}
	go moving2D.handle()
	return moving2D
}

type Moving2D struct {
	rw       sync.RWMutex
	entities map[int64]*moving2DTarget
	timeUnit float64
	idle     time.Duration
	interval time.Duration
	event    chan func()
	close    bool

	position2DChangeEventHandles      []game.Position2DChangeEventHandle
	position2DDestinationEventHandles []game.Position2DDestinationEventHandle
}

func (slf *Moving2D) MoveTo(entity game.Moving2DEntity, x float64, y float64) {
	guid := entity.GetGuid()
	current := time.Now().UnixMilli()
	slf.rw.Lock()
	defer slf.rw.Unlock()
	if slf.close {
		return
	}
	entityTarget, exist := slf.entities[guid]
	if !exist {
		entityTarget = &moving2DTarget{
			Moving2DEntity: entity,
			x:              x,
			y:              y,
			lastMoveTime:   current,
		}
		slf.entities[guid] = entityTarget
		return
	}
	entityTarget.x = x
	entityTarget.y = y
	entityTarget.lastMoveTime = current
}

func (slf *Moving2D) StopMove(guid int64) {
	slf.rw.Lock()
	delete(slf.entities, guid)
	slf.rw.Unlock()
}

func (slf *Moving2D) RegPosition2DChangeEvent(handle game.Position2DChangeEventHandle) {
	slf.position2DChangeEventHandles = append(slf.position2DChangeEventHandles, handle)
}

func (slf *Moving2D) OnPosition2DChangeEvent(entity game.Moving2DEntity, oldX, oldY float64) {
	for _, handle := range slf.position2DChangeEventHandles {
		handle(slf, entity, oldX, oldY)
	}
}

func (slf *Moving2D) RegPosition2DDestinationEvent(handle game.Position2DDestinationEventHandle) {
	slf.position2DDestinationEventHandles = append(slf.position2DDestinationEventHandles, handle)
}

func (slf *Moving2D) OnPosition2DDestinationEvent(entity game.Moving2DEntity) {
	for _, handle := range slf.position2DDestinationEventHandles {
		handle(slf, entity)
	}
}

type moving2DTarget struct {
	game.Moving2DEntity
	x, y         float64
	lastMoveTime int64
}

func (slf *Moving2D) Release() {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.close = true
	close(slf.event)
}

func (slf *Moving2D) handle() {
	for {
		slf.rw.Lock()
		if slf.close {
			slf.rw.Unlock()
			return
		}
		for guid, entity := range slf.entities {
			entity := entity
			x, y := entity.GetPosition()
			angle := g2d.CalcAngle(x, y, entity.x, entity.y)
			moveTime := time.Now().UnixMilli()
			interval := float64(moveTime - entity.lastMoveTime)
			if interval == 0 {
				continue
			}
			distance := g2d.CalcDistance(x, y, entity.x, entity.y)
			moveDistance := interval * (entity.GetSpeed() / (slf.timeUnit / 1000 / 1000))
			if moveDistance >= distance || (x == entity.x && y == entity.y) {
				entity.SetPosition(entity.x, entity.y)
				delete(slf.entities, guid)
				slf.OnPosition2DDestinationEvent(entity)
				continue
			} else {
				nx, ny := g2d.CalculateNewCoordinate(x, y, angle, moveDistance)
				entity.SetPosition(nx, ny)
				entity.lastMoveTime = moveTime
				slf.OnPosition2DChangeEvent(entity, x, y)
			}
		}

		time.Sleep(slf.interval)
		if len(slf.entities) == 0 {
			slf.rw.Unlock()
			time.Sleep(slf.idle)
		} else {
			slf.rw.Unlock()
		}
	}
}
