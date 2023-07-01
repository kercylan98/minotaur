package components_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/component"
	"github.com/kercylan98/minotaur/component/components"
	"sync"
	"testing"
	"time"
)

type MoveEntity struct {
	guid  int64
	x, y  float64
	speed float64
}

func (slf *MoveEntity) SetGuid(guid int64) {
}

func (slf *MoveEntity) GetGuid() int64 {
	return slf.guid
}

func (slf *MoveEntity) GetPosition() (x, y float64) {
	return slf.x, slf.y
}

func (slf *MoveEntity) SetPosition(x, y float64) {
	slf.x, slf.y = x, y
}

func (slf *MoveEntity) GetSpeed() float64 {
	return slf.speed
}

func NewEntity(guid int64, speed float64) *MoveEntity {
	return &MoveEntity{
		guid:  guid,
		speed: speed,
	}
}

func TestMoving2D_MoveTo(t *testing.T) {
	var wait sync.WaitGroup

	moving := components.NewMoving2D(components.WithMoving2DTimeUnit(time.Second))
	defer func() {
		moving.Release()
	}()

	moving.RegPosition2DChangeEvent(func(moving component.Moving2D, entity component.Moving2DEntity, oldX, oldY float64) {
		x, y := entity.GetPosition()
		fmt.Println(fmt.Sprintf("%d : %d | %f, %f > %f, %f", entity.GetGuid(), time.Now().UnixMilli(), oldX, oldY, x, y))
	})
	moving.RegPosition2DDestinationEvent(func(moving component.Moving2D, entity component.Moving2DEntity) {
		wait.Done()
	})

	for i := 0; i < 10; i++ {
		wait.Add(1)
		entity := NewEntity(int64(i)+1, float64(10+i))
		moving.MoveTo(entity, 50, 30)
	}

	wait.Wait()
}

func TestMoving2D_StopMove(t *testing.T) {
	var wait sync.WaitGroup

	moving := components.NewMoving2D(components.WithMoving2DTimeUnit(time.Second))
	defer func() {
		moving.Release()
	}()

	moving.RegPosition2DChangeEvent(func(moving component.Moving2D, entity component.Moving2DEntity, oldX, oldY float64) {
		x, y := entity.GetPosition()
		fmt.Println(fmt.Sprintf("%d : %d | %f, %f > %f, %f", entity.GetGuid(), time.Now().UnixMilli(), oldX, oldY, x, y))
	})
	moving.RegPosition2DDestinationEvent(func(moving component.Moving2D, entity component.Moving2DEntity) {
		fmt.Println(fmt.Sprintf("%d : %d | destination", entity.GetGuid(), time.Now().UnixMilli()))
		wait.Done()
	})
	moving.RegPosition2DStopMoveEvent(func(moving component.Moving2D, entity component.Moving2DEntity) {
		fmt.Println(fmt.Sprintf("%d : %d | stop", entity.GetGuid(), time.Now().UnixMilli()))
		wait.Done()
	})

	for i := 0; i < 10; i++ {
		wait.Add(1)
		entity := NewEntity(int64(i)+1, float64(10+i))
		moving.MoveTo(entity, 50, 30)
	}

	time.Sleep(time.Second * 1)

	for i := 0; i < 10; i++ {
		moving.StopMove(int64(i) + 1)
	}

	wait.Wait()
}
