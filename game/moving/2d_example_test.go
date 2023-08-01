package moving_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/moving"
	"sync"
	"time"
)

func ExampleNewTwoDimensional() {
	m := moving.NewTwoDimensional()
	defer func() {
		m.Release()
	}()
	fmt.Println(m != nil)

	// Output:
	// true
}

func ExampleTwoDimensional_MoveTo() {
	m := moving.NewTwoDimensional(moving.WithTwoDimensionalTimeUnit(time.Second))
	defer func() {
		m.Release()
	}()

	var wait sync.WaitGroup
	m.RegPosition2DDestinationEvent(func(moving *moving.TwoDimensional, entity moving.TwoDimensionalEntity) {
		fmt.Println("done")
		wait.Done()
	})

	wait.Add(1)
	entity := NewEntity(1, 100)
	m.MoveTo(entity, 50, 30)

	wait.Wait()

	// Output:
	// done
}

func ExampleTwoDimensional_StopMove() {
	m := moving.NewTwoDimensional(moving.WithTwoDimensionalTimeUnit(time.Second))
	defer func() {
		m.Release()
	}()

	var wait sync.WaitGroup
	m.RegPosition2DChangeEvent(func(moving *moving.TwoDimensional, entity moving.TwoDimensionalEntity, oldX, oldY float64) {
		fmt.Println("move")
	})
	m.RegPosition2DStopMoveEvent(func(moving *moving.TwoDimensional, entity moving.TwoDimensionalEntity) {
		fmt.Println("stop")
		wait.Done()
	})
	m.RegPosition2DDestinationEvent(func(moving *moving.TwoDimensional, entity moving.TwoDimensionalEntity) {
		fmt.Println("done")
		wait.Done()
	})

	wait.Add(1)
	entity := NewEntity(1, 100)
	m.MoveTo(entity, 50, 300)
	m.StopMove(1)

	wait.Wait()

	// Output:
	// stop
}
