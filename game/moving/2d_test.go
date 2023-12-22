package moving_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/moving"
	"github.com/kercylan98/minotaur/utils/geometry"
	"sync"
	"testing"
	"time"
)

type MoveEntity struct {
	guid  int64
	pos   geometry.Point[float64]
	speed float64
}

func (slf *MoveEntity) GetTwoDimensionalEntityID() int64 {
	return slf.guid
}

func (slf *MoveEntity) GetSpeed() float64 {
	return slf.speed
}

func (slf *MoveEntity) GetPosition() geometry.Point[float64] {
	return slf.pos
}

func (slf *MoveEntity) SetPosition(pos geometry.Point[float64]) {
	slf.pos = pos
}

func NewEntity(guid int64, speed float64) *MoveEntity {
	return &MoveEntity{
		guid:  guid,
		speed: speed,
	}
}

func TestNewTwoDimensional(t *testing.T) {
	m := moving.NewTwoDimensional[int64, float64]()
	defer func() {
		m.Release()
	}()
}

func TestTwoDimensional_StopMove(t *testing.T) {
	var wait sync.WaitGroup

	m := moving.NewTwoDimensional(moving.WithTwoDimensionalTimeUnit[int64, float64](time.Second))
	defer func() {
		m.Release()
	}()

	m.RegPosition2DChangeEvent(func(moving *moving.TwoDimensional[int64, float64], entity moving.TwoDimensionalEntity[int64, float64], oldX, oldY float64) {
		x, y := entity.GetPosition().GetXY()
		fmt.Println(fmt.Sprintf("%d : %d | %f, %f > %f, %f", entity.GetTwoDimensionalEntityID(), time.Now().UnixMilli(), oldX, oldY, x, y))
	})
	m.RegPosition2DDestinationEvent(func(moving *moving.TwoDimensional[int64, float64], entity moving.TwoDimensionalEntity[int64, float64]) {
		fmt.Println(fmt.Sprintf("%d : %d | destination", entity.GetTwoDimensionalEntityID(), time.Now().UnixMilli()))
		wait.Done()
	})
	m.RegPosition2DStopMoveEvent(func(moving *moving.TwoDimensional[int64, float64], entity moving.TwoDimensionalEntity[int64, float64]) {
		fmt.Println(fmt.Sprintf("%d : %d | stop", entity.GetTwoDimensionalEntityID(), time.Now().UnixMilli()))
		wait.Done()
	})

	for i := 0; i < 10; i++ {
		wait.Add(1)
		entity := NewEntity(int64(i)+1, float64(10+i))
		m.MoveTo(entity, 50, 30)
	}

	time.Sleep(time.Second * 1)

	for i := 0; i < 10; i++ {
		m.StopMove(int64(i) + 1)
	}

	wait.Wait()
}
