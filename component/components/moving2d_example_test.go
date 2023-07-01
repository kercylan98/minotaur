package components_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/component"
	"github.com/kercylan98/minotaur/component/components"
	"sync"
	"time"
)

func ExampleNewMoving2D() {
	moving := components.NewMoving2D()
	defer func() {
		moving.Release()
	}()
	fmt.Println(moving != nil)

	// Output:
	// true
}

func ExampleMoving2D_MoveTo() {
	moving := components.NewMoving2D(components.WithMoving2DTimeUnit(time.Second))
	defer func() {
		moving.Release()
	}()

	var wait sync.WaitGroup
	moving.RegPosition2DDestinationEvent(func(moving component.Moving2D, entity component.Moving2DEntity) {
		fmt.Println("done")
		wait.Done()
	})

	wait.Add(1)
	entity := NewEntity(1, 100)
	moving.MoveTo(entity, 50, 30)

	wait.Wait()

	// Output:
	// done
}

func ExampleMoving2D_StopMove() {
	moving := components.NewMoving2D(components.WithMoving2DTimeUnit(time.Second))
	defer func() {
		moving.Release()
	}()

	var wait sync.WaitGroup
	moving.RegPosition2DChangeEvent(func(moving component.Moving2D, entity component.Moving2DEntity, oldX, oldY float64) {
		fmt.Println("move")
	})
	moving.RegPosition2DStopMoveEvent(func(moving component.Moving2D, entity component.Moving2DEntity) {
		fmt.Println("stop")
		wait.Done()
	})
	moving.RegPosition2DDestinationEvent(func(moving component.Moving2D, entity component.Moving2DEntity) {
		fmt.Println("done")
		wait.Done()
	})

	wait.Add(1)
	entity := NewEntity(1, 100)
	moving.MoveTo(entity, 50, 300)
	moving.StopMove(1)

	wait.Wait()

	// Output:
	// stop
}
