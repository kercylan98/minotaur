package components

import (
	"github.com/kercylan98/minotaur/component"
	"github.com/kercylan98/minotaur/utils/geometry"
	"sync"
	"time"
)

// NewMoving2D 创建一个用于2D对象移动的实例(Moving2D)
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

// Moving2D 用于2D对象移动的数据结构
//   - 通过对象调用 MoveTo 方法后将开始执行该对象的移动
//   - 移动将在根据设置的每次移动间隔时间(WithMoving2DInterval)进行移动，当无对象移动需要移动时将会进入短暂的休眠
//   - 当对象移动速度永久为0时，将会导致永久无法完成的移动
type Moving2D struct {
	rw       sync.RWMutex
	entities map[int64]*moving2DTarget
	timeUnit float64
	idle     time.Duration
	interval time.Duration
	close    bool

	position2DChangeEventHandles      []component.Position2DChangeEventHandle
	position2DDestinationEventHandles []component.Position2DDestinationEventHandle
	position2DStopMoveEventHandles    []component.Position2DStopMoveEventHandle
}

// MoveTo 设置对象移动到特定位置
func (slf *Moving2D) MoveTo(entity component.Moving2DEntity, x float64, y float64) {
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

// StopMove 停止特定对象的移动
func (slf *Moving2D) StopMove(guid int64) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	entity, exist := slf.entities[guid]
	if exist {
		slf.OnPosition2DStopMoveEvent(entity)
		delete(slf.entities, guid)
	}
}

// RegPosition2DChangeEvent 在对象位置改变时将执行注册的事件处理函数
func (slf *Moving2D) RegPosition2DChangeEvent(handle component.Position2DChangeEventHandle) {
	slf.position2DChangeEventHandles = append(slf.position2DChangeEventHandles, handle)
}

func (slf *Moving2D) OnPosition2DChangeEvent(entity component.Moving2DEntity, oldX, oldY float64) {
	for _, handle := range slf.position2DChangeEventHandles {
		handle(slf, entity, oldX, oldY)
	}
}

// RegPosition2DDestinationEvent 在对象到达终点时将执行被注册的事件处理函数
func (slf *Moving2D) RegPosition2DDestinationEvent(handle component.Position2DDestinationEventHandle) {
	slf.position2DDestinationEventHandles = append(slf.position2DDestinationEventHandles, handle)
}

func (slf *Moving2D) OnPosition2DDestinationEvent(entity component.Moving2DEntity) {
	for _, handle := range slf.position2DDestinationEventHandles {
		handle(slf, entity)
	}
}

// RegPosition2DStopMoveEvent 在对象停止移动时将执行被注册的事件处理函数
func (slf *Moving2D) RegPosition2DStopMoveEvent(handle component.Position2DStopMoveEventHandle) {
	slf.position2DStopMoveEventHandles = append(slf.position2DStopMoveEventHandles, handle)
}

func (slf *Moving2D) OnPosition2DStopMoveEvent(entity component.Moving2DEntity) {
	for _, handle := range slf.position2DStopMoveEventHandles {
		handle(slf, entity)
	}
}

type moving2DTarget struct {
	component.Moving2DEntity
	x, y         float64
	lastMoveTime int64
}

// Release 释放对象移动对象所占用的资源
func (slf *Moving2D) Release() {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.close = true
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
			angle := geometry.CalcAngle(x, y, entity.x, entity.y)
			moveTime := time.Now().UnixMilli()
			interval := float64(moveTime - entity.lastMoveTime)
			if interval == 0 {
				continue
			}
			distance := geometry.CalcDistanceWithCoordinate(x, y, entity.x, entity.y)
			moveDistance := interval * (entity.GetSpeed() / (slf.timeUnit / 1000 / 1000))
			if moveDistance >= distance || (x == entity.x && y == entity.y) {
				entity.SetPosition(entity.x, entity.y)
				delete(slf.entities, guid)
				slf.OnPosition2DDestinationEvent(entity)
				continue
			} else {
				nx, ny := geometry.CalcNewCoordinate(x, y, angle, moveDistance)
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
