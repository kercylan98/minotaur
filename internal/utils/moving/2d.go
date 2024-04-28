package moving

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
	"sync"
	"time"
)

// NewTwoDimensional 创建一个用于2D对象移动的实例(TwoDimensional)
func NewTwoDimensional[EID generic.Basic, PosType generic.SignedNumber](options ...TwoDimensionalOption[EID, PosType]) *TwoDimensional[EID, PosType] {
	moving2D := &TwoDimensional[EID, PosType]{
		entities: map[EID]*moving2DTarget[EID, PosType]{},
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
type TwoDimensional[EID generic.Basic, PosType generic.SignedNumber] struct {
	rw       sync.RWMutex
	entities map[EID]*moving2DTarget[EID, PosType]
	timeUnit float64
	idle     time.Duration
	interval time.Duration
	close    bool

	position2DChangeEventHandles      []Position2DChangeEventHandle[EID, PosType]
	position2DDestinationEventHandles []Position2DDestinationEventHandle[EID, PosType]
	position2DStopMoveEventHandles    []Position2DStopMoveEventHandle[EID, PosType]
}

// MoveTo 设置对象移动到特定位置
func (slf *TwoDimensional[EID, PosType]) MoveTo(entity TwoDimensionalEntity[EID, PosType], x, y PosType) {
	guid := entity.GetTwoDimensionalEntityID()
	current := time.Now().UnixMilli()
	slf.rw.Lock()
	defer slf.rw.Unlock()
	if slf.close {
		return
	}
	entityTarget, exist := slf.entities[guid]
	if !exist {
		entityTarget = &moving2DTarget[EID, PosType]{
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
func (slf *TwoDimensional[EID, PosType]) StopMove(id EID) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	entity, exist := slf.entities[id]
	if exist {
		slf.OnPosition2DStopMoveEvent(entity)
		delete(slf.entities, id)
	}
}

// RegPosition2DChangeEvent 在对象位置改变时将执行注册的事件处理函数
func (slf *TwoDimensional[EID, PosType]) RegPosition2DChangeEvent(handle Position2DChangeEventHandle[EID, PosType]) {
	slf.position2DChangeEventHandles = append(slf.position2DChangeEventHandles, handle)
}

func (slf *TwoDimensional[EID, PosType]) OnPosition2DChangeEvent(entity TwoDimensionalEntity[EID, PosType], oldX, oldY PosType) {
	for _, handle := range slf.position2DChangeEventHandles {
		handle(slf, entity, oldX, oldY)
	}
}

// RegPosition2DDestinationEvent 在对象到达终点时将执行被注册的事件处理函数
func (slf *TwoDimensional[EID, PosType]) RegPosition2DDestinationEvent(handle Position2DDestinationEventHandle[EID, PosType]) {
	slf.position2DDestinationEventHandles = append(slf.position2DDestinationEventHandles, handle)
}

func (slf *TwoDimensional[EID, PosType]) OnPosition2DDestinationEvent(entity TwoDimensionalEntity[EID, PosType]) {
	for _, handle := range slf.position2DDestinationEventHandles {
		handle(slf, entity)
	}
}

// RegPosition2DStopMoveEvent 在对象停止移动时将执行被注册的事件处理函数
func (slf *TwoDimensional[EID, PosType]) RegPosition2DStopMoveEvent(handle Position2DStopMoveEventHandle[EID, PosType]) {
	slf.position2DStopMoveEventHandles = append(slf.position2DStopMoveEventHandles, handle)
}

func (slf *TwoDimensional[EID, PosType]) OnPosition2DStopMoveEvent(entity TwoDimensionalEntity[EID, PosType]) {
	for _, handle := range slf.position2DStopMoveEventHandles {
		handle(slf, entity)
	}
}

type moving2DTarget[EID generic.Basic, PosType generic.SignedNumber] struct {
	TwoDimensionalEntity[EID, PosType]
	x, y         PosType
	lastMoveTime int64
}

// Release 释放对象移动对象所占用的资源
func (slf *TwoDimensional[EID, PosType]) Release() {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.close = true
}

func (slf *TwoDimensional[EID, PosType]) handle() {
	for {
		slf.rw.Lock()
		if slf.close {
			slf.rw.Unlock()
			return
		}
		for guid, entity := range slf.entities {
			entity := entity
			x, y := entity.GetPosition().GetXY()
			angle := geometry.CalcAngle(float64(x), float64(y), float64(entity.x), float64(entity.y))
			moveTime := time.Now().UnixMilli()
			interval := float64(moveTime - entity.lastMoveTime)
			if interval == 0 {
				continue
			}
			distance := geometry.CalcDistanceWithCoordinate(x, y, entity.x, entity.y)
			moveDistance := interval * (entity.GetSpeed() / (slf.timeUnit / 1000 / 1000))
			if moveDistance >= float64(distance) || (x == entity.x && y == entity.y) {
				entity.SetPosition(geometry.NewPoint(entity.x, entity.y))
				delete(slf.entities, guid)
				slf.OnPosition2DDestinationEvent(entity)
				continue
			} else {
				nx, ny := geometry.CalcNewCoordinate(float64(x), float64(y), angle, moveDistance)
				entity.SetPosition(geometry.NewPoint(PosType(nx), PosType(ny)))
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
