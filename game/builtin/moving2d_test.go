package builtin

import (
	"fmt"
	"github.com/kercylan98/minotaur/game"
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
	moving := NewMoving2D(WithMoving2DTimeUnit(time.Second))
	var wait sync.WaitGroup
	moving.RegPosition2DDestinationEvent(func(moving game.Moving2D, entity game.Moving2DEntity) {
		wait.Done()
	})
	var res []string
	moving.RegPosition2DChangeEvent(func(moving game.Moving2D, entity game.Moving2DEntity, oldX, oldY float64) {
		x, y := entity.GetPosition()
		res = append(res, fmt.Sprintf("%d : %d | %f, %f > %f, %f", entity.GetGuid(), time.Now().UnixMilli(), oldX, oldY, x, y))
	})
	//moving.RegPosition2DChangeEvent(func(moving game.Moving2D, entity game.Moving2DEntity, oldX, oldY float64) {
	//	x, y := entity.GetPosition()
	//	fmt.Println("Moving", entity.GetGuid(), oldX, oldY, x, y)
	//})
	for i := 0; i < 1; i++ {
		wait.Add(1)
		entity := NewEntity(int64(i)+1, float64(10+i))
		moving.MoveTo(entity, 50, 30)
	}

	wait.Wait()
	for _, re := range res {
		fmt.Println(re)
	}
}
