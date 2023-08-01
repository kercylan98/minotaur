package moving

import (
	"github.com/kercylan98/minotaur/utils/geometry"
	"sync"
	"time"
)

// NewTwoDimensional 创建一个用于2D对象移动的实例(TwoDimensional)
func NewTwoDimensional(options ...TwoDimensionalOption) *TwoDimensional {
	moving2D := &TwoDimensional{
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

// TwoDimensional 用于2D对象移动的数据结构
//   - 通过对象调用 MoveTo 方法后将开始执行该对象的移动
//   - 移动将在根据设置的每次移动间隔时间(WithTwoDimensionalInterval)进行移动，当无对象移动需要移动时将会进入短暂的休眠
//   - 当对象移动速度永久为0时，将会导致永久无法完成的移动
type TwoDimensional struct {
	rw       sync.RWMutex
	entities map[int64]*moving2DTarget
	timeUnit float64
	idle     time.Duration
	interval time.Duration
	close    bool

	position2DChangeEventHandles      []Position2DChangeEventHandle
	position2DDestinationEventHandles []Position2DDestinationEventHandle
	position2DStopMoveEventHandles    []Position2DStopMoveEventHandle
}

// MoveTo 设置对象移动到特定位置
func (slf *TwoDimensional) MoveTo(entity TwoDimensionalEntity, x float64, y float64) {
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
			TwoDimensionalEntity: entity,
			x:                    x,
			y:                    y,
			lastMoveTime:         current,
		}
		slf.entities[guid] = entityTarget
		return
	}
	entityTarget.x = x
	entityTarget.y = y
	entityTarget.lastMoveTime = current
}

// StopMove 停止特定对象的移动
func (slf *TwoDimensional) StopMove(guid int64) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	entity, exist := slf.entities[guid]
	if exist {
		slf.OnPosition2DStopMoveEvent(entity)
		delete(slf.entities, guid)
	}
}

// RegPosition2DChangeEvent 在对象位置改变时将执行注册的事件处理函数
func (slf *TwoDimensional) RegPosition2DChangeEvent(handle Position2DChangeEventHandle) {
	slf.position2DChangeEventHandles = append(slf.position2DChangeEventHandles, handle)
}

func (slf *TwoDimensional) OnPosition2DChangeEvent(entity TwoDimensionalEntity, oldX, oldY float64) {
	for _, handle := range slf.position2DChangeEventHandles {
		handle(slf, entity, oldX, oldY)
	}
}

// RegPosition2DDestinationEvent 在对象到达终点时将执行被注册的事件处理函数
func (slf *TwoDimensional) RegPosition2DDestinationEvent(handle Position2DDestinationEventHandle) {
	slf.position2DDestinationEventHandles = append(slf.position2DDestinationEventHandles, handle)
}

func (slf *TwoDimensional) OnPosition2DDestinationEvent(entity TwoDimensionalEntity) {
	for _, handle := range slf.position2DDestinationEventHandles {
		handle(slf, entity)
	}
}

// RegPosition2DStopMoveEvent 在对象停止移动时将执行被注册的事件处理函数
func (slf *TwoDimensional) RegPosition2DStopMoveEvent(handle Position2DStopMoveEventHandle) {
	slf.position2DStopMoveEventHandles = append(slf.position2DStopMoveEventHandles, handle)
}

func (slf *TwoDimensional) OnPosition2DStopMoveEvent(entity TwoDimensionalEntity) {
	for _, handle := range slf.position2DStopMoveEventHandles {
		handle(slf, entity)
	}
}

type moving2DTarget struct {
	TwoDimensionalEntity
	x, y         float64
	lastMoveTime int64
}

// Release 释放对象移动对象所占用的资源
func (slf *TwoDimensional) Release() {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.close = true
}

func (slf *TwoDimensional) handle() {
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
